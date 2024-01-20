package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
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
	if prevInfo != nil {
		if prevDist >= d {
			return nil, nil
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
		return nil, io.EOF
	}
	line := scanner.Text()
	log.Println(d, "line", line)
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
		log.Println("skip", res)
		return nil, nil
	}
	if next < d {
		return res, nil
	}
	prevInfo = res
	prevDist = next
	return nil, nil
}

func main() {
	v, err := nanoda.NewVoicevox(
		"voicevox_core/voicevox_core.dll",
		"voicevox_core/open_jtalk_dic_utf_8-1.11",
		"voicevox_core/model")
	if err != nil {
		log.Fatal(err)
	}
	ctx, _, err := oto.NewContext(&oto.NewContextOptions{
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
	qMarks, err := Setup(s, Dict.Marks)
	if err != nil {
		log.Fatal(err)
	}
	qIcons, err := Setup(s, Dict.Icons)
	if err != nil {
		log.Fatal(err)
	}
	qDists, err := Setup(s, Dict.Dists)
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.ListenPacket("udp", "127.0.0.1:20777")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	var logFile io.ReadCloser
	var scanner *bufio.Scanner
	logCloser := func() {}
	lastTime := float32(0)
	b := make([]byte, 4096)
	completed := ""
	for {
		n, _, err := conn.ReadFrom(b)
		if err != nil {
			log.Fatal(err)
		}
		if n != easportswrc.PacketEASportsWRCLength {
			continue
		}
		var packet easportswrc.PacketEASportsWRC
		packet.UnmarshalBinary(b[:n])
		logName := filepath.Join("log", fmt.Sprintf("%v.log", packet.StageLength))
		if logFile != nil && lastTime > 0 && packet.StageCurrentTime == 0 {
			logCloser()
			completed = "" // リスタートしたら読み込み済み解除
			log.Println(logName, "completed")
		}
		lastTime = packet.StageCurrentTime
		if logFile == nil && completed != logName {
			f, err := os.Open(logName)
			if err != nil {
				log.Print(err)
				continue
			}
			logFile = f
			logCloser = sync.OnceFunc(func() { logFile.Close(); logFile = nil; scanner = nil })
			scanner = bufio.NewScanner(logFile)
			log.Println(logName, "opend", easportswrc.GetStage(packet.StageLength))
		}
		if scanner == nil {
			continue
		}
		if packet.StageCurrentDistance < 1 {
			continue
		}
		info, err := nextInfo(scanner, packet.StageCurrentDistance+Offset)
		if err != nil {
			log.Println(err)
			logCloser()
			completed = logName // 読み込み済み
			log.Println(logName, "completed")
			continue
		}
		words := []nanoda.AudioQuery{}
		if info != nil {
			log.Println(packet.StageCurrentDistance, info)
			qm, ok := qMarks[info.Mark]
			if ok {
				words = append(words, qm)
			} else {
				if info.Mark != "unknown" {
					q, err := makeAudioQuery(s, info.Mark)
					if err != nil {
						log.Println(err)
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
						log.Println(err)
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
						log.Println(err)
					} else {
						words = append(words, q)
					}
				}
			}
			if info.Mark == "finish" {
				logCloser()
				completed = logName // 読み込み済み
				log.Println(logName, "completed")
			}
		}
		for _, w := range words {
			speech(ctx, s, w)
		}
	}
}
