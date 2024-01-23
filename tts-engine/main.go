package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aethiopicuschan/nanoda"
	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
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
		time.Sleep(10 * time.Millisecond)
	}
	return nil
}

type Info struct {
	Position float64
	Words    []string
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

func main() {
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
	qDicts, err := Init(s, Dict)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		for _, w := range strings.Fields(scanner.Text()) {
			qm, ok := qDicts[w]
			if !ok {
				q, err := makeAudioQuery(s, w)
				if err != nil {
					log.Println(err)
					continue
				}
				qm = q
			}
			speech(ctxOto, s, qm)
		}
	}
}
