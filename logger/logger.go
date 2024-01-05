package logger

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/nobonobo/wrc-logger/easportswrc"
	"gocv.io/x/gocv"
)

var CameraDeviceID = 2

func init() {
	flag.IntVar(&CameraDeviceID, "camera", CameraDeviceID, "camera device id")
}

func Setup(ctx context.Context) func(*easportswrc.PacketEASportsWRC) error {
	var logFile *os.File
	var webcam *gocv.VideoCapture
	img := gocv.NewMat()
	logCloser := func() {}
	camCloser := func() {}
	closing := false
	timeout := time.AfterFunc(3*time.Second, func() {
		camCloser()
		if closing {
			logCloser()
			closing = false
		}
		webcam = nil
	})
	timeout.Stop()
	compute := gocv.NewMat()
	lastTime := float32(0)
	lastDetected := "unknown"
	lastDistance := float64(0)
	blockDistance := float64(0)
	cnt := 0
	return func(packet *easportswrc.PacketEASportsWRC) error {
		if webcam == nil {
			cam, err := gocv.VideoCaptureDevice(CameraDeviceID)
			if err != nil {
				log.Fatal("[log]", err)
			}
			camCloser = sync.OnceFunc(func() { cam.Close() })
			cam.Set(gocv.VideoCaptureFPS, 60)
			cam.Set(gocv.VideoCaptureFrameWidth, 1920)
			cam.Set(gocv.VideoCaptureFrameHeight, 1024)
			webcam = cam
			timeout.Reset(3 * time.Second)
		}
		timeout.Reset(3 * time.Second)
		logName := filepath.Join("log", fmt.Sprintf("%v.log", packet.StageLength))
		// 上書きを避けるために既存ファイルがある場合はファイル名にサフィックスを付与する
		for i := 1; ; i++ {
			_, err := os.Stat(logName)
			if err != nil {
				if os.IsNotExist(err) {
					break
				}
			}
			logName = filepath.Join("log", fmt.Sprintf("%v.log.%d", packet.StageLength, i))
		}
		// ステージの最初に戻ったらいったんファイルを閉じる
		if logFile != nil && lastTime > 0 && packet.StageCurrentTime == 0 {
			logCloser()
		}
		lastTime = packet.StageCurrentTime
		if logFile == nil {
			f, err := os.Create(logName)
			if err != nil {
				log.Fatal(err)
			}
			logFile = f
			logCloser = sync.OnceFunc(func() {
				logFile.Close()
				logFile = nil
				closing = false
				log.Println("[log]", logName, "closed")
			})
			blockDistance = 0
			closing = false
			log.Println("[log]", logName, "created")
		}
		if webcam != nil {
			webcam.Read(&img)
			if img.Empty() {
				return nil
			}
			cnt++
			if cnt%2 == 0 {
				return nil
			}
			mark := img.Region(markRect)
			save := mark.Clone()
			if !getMotionStop(mark.Clone()) {
				return nil
			}
			icon := mark.Region(iconRect)
			dist := mark.Region(distRect)
			markPreProcess(&mark)
			hash.Compute(mark, &compute)
			detect := "unknown"
			detectMin := 10000.0
			for k, v := range Marks {
				similar := hash.Compare(compute, v)
				if similar < 10 {
					if detectMin > similar {
						detectMin = similar
						detect = k
					}
				}
			}
			// 終盤じゃないfinishは誤判定
			if detect == "finish" && packet.StageLength-500 > packet.StageCurrentDistance {
				detect = "unknown"
			}
			// 判定不能だった場合、その画像を記録しておく
			if detect == "unknown" {
				gocv.IMWrite(fmt.Sprintf("mark/%v_unknown.png", packet.StageCurrentDistance), save)
				gocv.IMWrite(fmt.Sprintf("mark/%v_th.png", packet.StageCurrentDistance), mark)
				icon.Close()
				dist.Close()
				mark.Close()
				save.Close()
				return nil
			}
			// 最後に検知したものとの30メートル以内の重複検出
			sameDetected := detect == lastDetected && packet.StageCurrentDistance < lastDistance+30
			// シケインは100m以内の重複を除外
			if strings.HasSuffix(lastDetected, "chicane") {
				sameDetected = sameDetected || (detect == lastDetected && packet.StageCurrentDistance < lastDistance+100)
			}
			// finishの重複分は無視
			if lastDetected == "finish" && detect == "finish" {
				return nil
			}
			lastDetected = detect
			lastDistance = packet.StageCurrentDistance
			// 最後に検知したものとの30メートル以内の重複および
			// straightとfinish以外のブロック（指示なし）距離の検出を無視する
			normalSign := !(detect == "finish" || detect == "straight")
			blocking := normalSign && packet.StageCurrentDistance < blockDistance
			if sameDetected || blocking {
				return nil
			}
			distPreProcess(&dist)
			hash.Compute(dist, &compute)
			distDetected := 0
			distMin := 1000.0
			for k, v := range Dists {
				similar := hash.Compare(compute, v)
				if similar < 15 {
					if distMin > similar {
						distMin = similar
						distDetected, _ = strconv.Atoi(k)
					}
				}
			}
			if distDetected > 0 {
				// 直線の宣言距離x0.8m分検出ブロック、50m差し引いてマイナスなものはブロックしない
				blockDistance = packet.StageCurrentDistance + 0.8*float64(distDetected) - 50
				if blockDistance < packet.StageCurrentDistance {
					blockDistance = packet.StageCurrentDistance
				}
			}
			iconPreProcess(&icon)
			hash.Compute(icon, &compute)
			iconDetected := "unknown"
			iconMin := 1000.0
			for k, v := range Icons {
				similar := hash.Compare(compute, v)
				if similar < 15 {
					if iconMin > similar {
						iconMin = similar
						iconDetected = k
					}
				}
			}
			log.Printf("[log] %v/%v:%s,%s,%d",
				packet.StageCurrentDistance, packet.StageLength,
				detect, iconDetected, distDetected,
			)
			if logFile != nil {
				logFile.WriteString(fmt.Sprintf("%v,%s,%s,%d\n",
					packet.StageCurrentDistance,
					detect, iconDetected, distDetected,
				))
			}
			if packet.StageCurrentDistance >= packet.StageLength || detect == "finish" {
				closing = true
			}
			gocv.IMWrite(fmt.Sprintf("mark/%v_%s_%s_%d.png", packet.StageCurrentDistance, detect, iconDetected, distDetected), save)
			gocv.IMWrite(fmt.Sprintf("mark/%v_th.png", packet.StageCurrentDistance), mark)
			icon.Close()
			dist.Close()
			mark.Close()
			save.Close()
		}
		return nil
	}
}
