package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"gocv.io/x/gocv"
)

func main() {
	conn, err := net.ListenPacket("udp", "127.0.0.1:20777")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	var logFile *os.File
	var webcam *gocv.VideoCapture
	img := gocv.NewMat()
	logCloser := func() {}
	camCloser := func() {}
	timeout := time.AfterFunc(3*time.Second, func() {
		camCloser()
		webcam = nil
	})
	timeout.Stop()
	b := make([]byte, 4096)
	compute := gocv.NewMat()
	cnt := 0
	lastDetected := "unknown"
	lastDistance := float64(0)
	for {
		n, _, err := conn.ReadFrom(b)
		if err != nil {
			log.Fatal(err)
		}
		if n != PacketEASportsWRCLength {
			continue
		}
		cnt++
		if cnt%2 == 0 {
			continue
		}
		if webcam == nil {
			cam, err := gocv.VideoCaptureDevice(2)
			if err != nil {
				log.Fatal(err)
			}
			camCloser = sync.OnceFunc(func() { cam.Close() })
			cam.Set(gocv.VideoCaptureFPS, 30)
			cam.Set(gocv.VideoCaptureFrameWidth, 1920)
			cam.Set(gocv.VideoCaptureFrameHeight, 1024)
			webcam = cam
			timeout.Reset(3 * time.Second)
		}
		var packet PacketEASportsWRC
		packet.UnmarshalBinary(b[:n])
		timeout.Reset(3 * time.Second)
		logName := filepath.Join("log", fmt.Sprintf("%v.log", packet.StageLength))
		if logFile != nil && packet.StageCurrentDistance == 0 {
			logCloser()
		}
		if logFile == nil {
			f, err := os.Create(logName)
			if err != nil {
				log.Fatal(err)
			}
			logFile = f
			logCloser = sync.OnceFunc(func() { logFile.Close(); logFile = nil })
			log.Println(logName, "created")
		}
		if webcam != nil {
			webcam.Read(&img)
			if img.Empty() {
				continue
			}
			mark := img.Region(markRect)
			save := mark.Clone()
			if !getMotionStop(mark) {
				continue
			}
			icon := mark.Region(iconRect)
			dist := mark.Region(distRect)
			markPreProcess(&mark)
			hash.Compute(mark, &compute)
			detect := "unknown"
			detectMin := 100.0
			for k, v := range Marks {
				similar := hash.Compare(compute, v)
				if similar < 10 {
					if detectMin > similar {
						detectMin = similar
						detect = k
					}
				}
			}
			if detect == "unknown" {
				//log.Println(packet.StageCurrentDistance, detect)
				gocv.IMWrite(fmt.Sprintf("mark/%v_unknown.png", packet.StageCurrentDistance), save)
				gocv.IMWrite(fmt.Sprintf("mark/%v_th.png", packet.StageCurrentDistance), mark)
				icon.Close()
				dist.Close()
				mark.Close()
				save.Close()
				continue
			}
			sameDetected := detect == lastDetected && packet.StageCurrentDistance < lastDistance+10
			lastDetected = detect
			lastDistance = packet.StageCurrentDistance
			if sameDetected {
				continue
			}
			distPreProcess(&dist)
			hash.Compute(dist, &compute)
			distDetected := 0
			distMin := 100.0
			for k, v := range Dists {
				similar := hash.Compare(compute, v)
				if similar < 15 {
					if distMin > similar {
						distMin = similar
						distDetected, _ = strconv.Atoi(k)
					}
				}
			}
			iconPreProcess(&icon)
			hash.Compute(icon, &compute)
			iconDetected := "unknown"
			iconMin := 100.0
			for k, v := range Icons {
				similar := hash.Compare(compute, v)
				if similar < 15 {
					if iconMin > similar {
						iconMin = similar
						iconDetected = k
					}
				}
			}
			log.Printf("%v/%v:%s,%s,%d",
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
				log.Println(logName, "closed")
				logCloser()
			}
			gocv.IMWrite(fmt.Sprintf("mark/%v_%s_%s_%d.png", packet.StageCurrentDistance, detect, iconDetected, distDetected), save)
			gocv.IMWrite(fmt.Sprintf("mark/%v_th.png", packet.StageCurrentDistance), mark)
			icon.Close()
			dist.Close()
			mark.Close()
			save.Close()
		}
	}
}
