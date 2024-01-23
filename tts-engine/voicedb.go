package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/aethiopicuschan/nanoda"
)

var (
	ActorID = 3
	Speed   = 1.5
	Pitch   = 0.0
	Volume  = 1.8
)

func init() {
	flag.IntVar(&ActorID, "actor", ActorID, "actor id")
	flag.Float64Var(&Speed, "speed", Speed, "speed")
	flag.Float64Var(&Pitch, "pitch", Pitch, "pitch")
	flag.Float64Var(&Volume, "volume", Volume, "volume")
}

type AQ struct {
	Text string `json:"text"`
}

var (
	Dict map[string]AQ
)

func init() {
	fp, err := os.Open("base.json")
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()
	if err := json.NewDecoder(fp).Decode(&Dict); err != nil {
		log.Fatal(err)
	}
}

func makeAudioQuery(s nanoda.Synthesizer, text string) (nanoda.AudioQuery, error) {
	q, err := s.CreateAudioQuery(text, nanoda.StyleId(ActorID))
	if err != nil {
		return nanoda.AudioQuery{}, err
	}
	q.SpeedScale = Speed
	q.PitchScale = Pitch
	q.VolumeScale = Volume
	return q, nil
}

func Init(s nanoda.Synthesizer, dict map[string]AQ) (map[string]nanoda.AudioQuery, error) {
	res := map[string]nanoda.AudioQuery{}
	for k, v := range dict {
		q, err := makeAudioQuery(s, v.Text)
		if err != nil {
			return nil, err
		}
		res[k] = q
	}
	return res, nil
}
