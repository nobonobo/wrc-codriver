package codriver

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"github.com/nobonobo/wrc-codriver/easportswrc"
)

type Info struct {
	Position float64
	Words    []string
}

var (
	ActorID           = 3
	Speed             = 1.8
	Pitch             = 0.0
	Volume            = 1.8
	Offset            = 5.0
	PrePhonemeLength  = 0.05
	PostPhonemeLength = 0.05
)

func init() {
	flag.IntVar(&ActorID, "actor", ActorID, "actor id")
	flag.Float64Var(&Speed, "speed", Speed, "speed")
	flag.Float64Var(&Pitch, "pitch", Pitch, "pitch")
	flag.Float64Var(&Volume, "volume", Volume, "volume")
	flag.Float64Var(&Offset, "offset", Offset, "offset [-50..50]")
	flag.Float64Var(&PrePhonemeLength, "pre-phoneme", PrePhonemeLength, "pre-phoneme-length")
	flag.Float64Var(&PostPhonemeLength, "post-phoneme", PostPhonemeLength, "post-phoneme-length")
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
	if len(fields) < 2 {
		return nil, fmt.Errorf("invalid line: %s", line)
	}
	next, err := strconv.ParseFloat(strings.TrimSpace(fields[0]), 64)
	if err != nil {
		return nil, fmt.Errorf("invalid line: %s", line)
	}
	res := &Info{
		Position: next,
		Words:    fields[1:],
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

func startEngine(ctx context.Context, in <-chan []string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	cmd := exec.CommandContext(ctx, ".\\tts-engine.exe",
		"-actor", fmt.Sprintf("%d", ActorID),
		"-speed", fmt.Sprintf("%f", Speed),
		"-pitch", fmt.Sprintf("%f", Pitch),
		"-volume", fmt.Sprintf("%f", Volume),
		"-pre-phoneme", fmt.Sprintf("%f", PrePhonemeLength),
		"-post-phoneme", fmt.Sprintf("%f", PostPhonemeLength),
	)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return fmt.Errorf("tts-engine pipe failed: %w", err)
	}
	defer stdin.Close()
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("tts-engine start failed: %w", err)
	}
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case words := <-in:
			fmt.Fprintln(stdin, strings.Join(words, " "))
		}
	}
}

func Setup(ctx context.Context) func(*easportswrc.PacketEASportsWRC) error {
	speechCh := make(chan []string, 10)
	var logFile io.ReadCloser
	var scanner *bufio.Scanner
	logCloser := func() {}
	lastTime := float32(0)
	completed := ""
	go func() {
		for {
			func() {
				c, cancel := context.WithCancel(context.Background())
				defer cancel()
				select {
				case <-ctx.Done():
					return
				default:
				}
				if err := startEngine(c, speechCh); err != nil {
					log.Println("[snd]", err)
				}
			}()
		}
	}()
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
		//log.Println("[snd]", packet.StageCurrentDistance, info)
		req := []string{}
		for _, w := range info.Words {
			w = strings.TrimSpace(w)
			if w == "unknown" || w == "0" {
				continue
			}
			req = append(req, w)
			if w == "finish" {
				logCloser()
				completed = logName // 読み込み済み
				log.Println("[snd]", logName, "completed")
			}
		}
		speechCh <- req
		return nil
	}
}
