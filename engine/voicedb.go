package engine

import (
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/aethiopicuschan/nanoda"
)

var (
	ActorID           = 3
	Pitch             = 0.0
	Intnation         = 1.0
	Speed             = 1.5
	Volume            = 1.8
	Pause             = 0.25
	PrePhonemeLength  = 0.0
	PostPhonemeLength = 0.0
)

func init() {
	flag.IntVar(&ActorID, "actor", ActorID, "actor id")
	flag.Float64Var(&Pitch, "pitch", Pitch, "pitch")
	flag.Float64Var(&Intnation, "intnation", Volume, "intnation")
	flag.Float64Var(&Speed, "speed", Speed, "base speed")
	flag.Float64Var(&Volume, "volume", Volume, "volume magnification")
	flag.Float64Var(&Pause, "pause", Pause, "pause magnification")
	flag.Float64Var(&PrePhonemeLength, "pre-phoneme", PrePhonemeLength, "pre-phoneme-length")
	flag.Float64Var(&PostPhonemeLength, "post-phoneme", PostPhonemeLength, "post-phoneme-length")
}

type AQ struct {
	Text      string  `json:"text"`
	Speed     float64 `json:"speed"`
	Pitch     float64 `json:"pitch"`
	Intnation float64 `json:"intnation"`
	Volume    float64 `json:"volume"`
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
	q.IntonationScale = Intnation
	q.PitchScale = Pitch
	q.SpeedScale = Speed
	q.VolumeScale = Volume
	q.PrePhonemeLength = PrePhonemeLength
	q.PostPhonemeLength = PostPhonemeLength
	for _, p := range q.AccentPhrases {
		if p.PauseMora != nil {
			p.PauseMora.VowelLength *= Pause
		}
	}
	return q, nil
}

func Init(s nanoda.Synthesizer, dict map[string]AQ) (map[string]nanoda.AudioQuery, error) {
	res := map[string]nanoda.AudioQuery{}
	for k, v := range dict {
		q, err := makeAudioQuery(s, v.Text)
		if err != nil {
			return nil, err
		}
		if v.Intnation != 0.0 {
			q.IntonationScale = v.Intnation
		}
		if v.Pitch != 0.0 {
			q.PitchScale = v.Pitch
		}
		if v.Speed != 0.0 {
			q.SpeedScale *= v.Speed
		}
		if v.Volume != 0.0 {
			q.VolumeScale *= v.Volume
		}
		res[k] = q
	}
	return res, nil
}
