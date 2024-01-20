package codriver

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/aethiopicuschan/nanoda"
	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/nobonobo/wrc-codriver/easportswrc"
)

func speech(ctx *oto.Context, s nanoda.Synthesizer, q nanoda.AudioQuery) error {
	w, err := s.Synthesis(q, nanoda.StyleId(ActorID))
	if err != nil {
		return err
	}
	defer w.Close()
	decoded, err := wav.DecodeWithoutResampling(w)
	if err != nil {
		return err
	}
	p := ctx.NewPlayer(decoded)
	p.Play()
	for p.IsPlaying() {
		time.Sleep(20 * time.Millisecond)
	}
	return nil
}

type Info struct {
	Position float64
	Mark     string
	Icon     string
	Dist     string
}

var (
	prevDist float64
	prevInfo *Info
)

func nextInfo(scanner *bufio.Scanner, d float64) (*Info, error) {
	empty := new(Info)
	if prevInfo != nil {
		if prevDist >= d {
			return empty, nil
		}
		res := prevInfo
		prevInfo = nil
		prevDist = d
		return res, nil
	}
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, err
		}
		return nil, nil
	}
	line := scanner.Text()
	log.Println("[snd]", d, "line", line)
	fields := strings.Split(line, ",")
	if len(fields) != 4 {
		return nil, fmt.Errorf("invalid line: %s", line)
	}
	next, err := strconv.ParseFloat(fields[0], 64)
	if err != nil {
		return nil, fmt.Errorf("invalid line: %s", line)
	}
	res := &Info{
		Position: next,
		Mark:     fields[1],
		Icon:     fields[2],
		Dist:     fields[3],
	}
	if next+100 < d {
		log.Println("[snd]", "skip", res)
		return empty, nil
	}
	if next < d {
		return res, nil
	}
	prevInfo = res
	prevDist = next
	return empty, nil
}

func Setup(ctx context.Context) func(*easportswrc.PacketEASportsWRC) error {
	v, err := nanoda.NewVoicevox(
		"voicevox_core/voicevox_core.dll",
		"voicevox_core/open_jtalk_dic_utf_8-1.11",
		"voicevox_core/model")
	if err != nil {
		log.Fatal(err)
	}
	ctxOto, _, err := oto.NewContext(&oto.NewContextOptions{
		SampleRate:   48000,
		ChannelCount: 1,
		Format:       oto.FormatSignedInt16LE,
	})
	if err != nil {
		log.Fatal(err)
	}
	s, err := v.NewSynthesizer()
	if err != nil {
		log.Fatal(err)
	}
	if err := s.LoadModelsFromStyleId(nanoda.StyleId(ActorID)); err != nil {
		log.Fatal(err)
	}
	qMarks, err := Init(s, Dict.Marks)
	if err != nil {
		log.Fatal(err)
	}
	qIcons, err := Init(s, Dict.Icons)
	if err != nil {
		log.Fatal(err)
	}
	qDists, err := Init(s, Dict.Dists)
	if err != nil {
		log.Fatal(err)
	}
	speechCh := make(chan nanoda.AudioQuery, 10)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case q := <-speechCh:
				speech(ctxOto, s, q)
			}
		}
	}()
	var logFile io.ReadCloser
	var scanner *bufio.Scanner
	logCloser := func() {}
	lastTime := float32(0)
	completed := ""
	return func(packet *easportswrc.PacketEASportsWRC) error {
		logName := filepath.Join("log", fmt.Sprintf("%v.log", packet.StageLength))
		v, ok := easportswrc.Stages[packet.StageLength]
		if ok {
			loc := easportswrc.Locations[v.Location]
			dir := fmt.Sprintf("%02d.%s", v.Location+1, easportswrc.LocationKeys[v.Location])
			name := fmt.Sprintf("%02d.%s", v.Stage+1, loc.Stages[v.Stage])
			logName = filepath.Join("pacenotes", dir, name+".log")
		}
		if lastTime > 0 && packet.StageCurrentTime == 0 {
			if logFile != nil {
				logCloser()
			}
			completed = "" // リスタートしたら読み込み済み解除
			log.Println("[snd]", logName, "restart")
			lastTime = packet.StageCurrentTime
		}
		if logFile == nil && completed != logName {
			f, err := os.Open(logName)
			if err != nil {
				completed = logName // 読み込み済み
				lastTime = 1
				log.Println("[snd]", logName, "not found")
				return err
			}
			logFile = f
			logCloser = sync.OnceFunc(func() {
				logFile.Close()
				logFile = nil
				scanner = nil
				prevDist = 0
				prevInfo = nil
			})
			scanner = bufio.NewScanner(logFile)
			log.Println("[snd]", logName, "opend", easportswrc.GetStage(packet.StageLength))
		}
		if scanner == nil {
			return nil
		}
		if packet.StageCurrentTime == 0 {
			return nil
		}
		lastTime = packet.StageCurrentTime
		info, err := nextInfo(scanner, packet.StageCurrentDistance+2*float64(packet.VehicleSpeed)+Offset)
		if err != nil {
			log.Println("[snd]", err)
			logCloser()
			completed = logName // 読み込み済み
			log.Println("[snd]", logName, "completed")
			return nil
		}
		if info == nil {
			// 読み込み進捗が無い場合、scannerを再構築して次の機会を待つ（追記あれば再開する）
			scanner = bufio.NewScanner(logFile)
			return nil
		}
		if info.Position == 0 {
			// 次の情報まで未到達
			return nil
		}
		words := []nanoda.AudioQuery{}
		//log.Println("[snd]", packet.StageCurrentDistance, info)
		qm, ok := qMarks[info.Mark]
		if ok {
			words = append(words, qm)
		} else {
			if info.Mark != "unknown" {
				q, err := makeAudioQuery(s, info.Mark)
				if err != nil {
					log.Println("[snd]", err)
				} else {
					words = append(words, q)
				}
			}
		}
		qi, ok := qIcons[info.Icon]
		if ok {
			words = append(words, qi)
		} else {
			if info.Icon != "unknown" {
				q, err := makeAudioQuery(s, info.Icon)
				if err != nil {
					log.Println("[snd]", err)
				} else {
					words = append(words, q)
				}
			}
		}
		qd, ok := qDists[info.Dist]
		if ok {
			words = append(words, qd)
		} else {
			if info.Dist != "0" {
				q, err := makeAudioQuery(s, info.Dist)
				if err != nil {
					log.Println("[snd]", err)
				} else {
					words = append(words, q)
				}
			}
		}
		if info.Mark == "finish" {
			logCloser()
			completed = logName // 読み込み済み
			log.Println("[snd]", logName, "completed")
		}
		for _, w := range words {
			speechCh <- w
		}
		return nil
	}
}
