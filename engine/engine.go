package engine

import (
	"context"
	"time"

	"github.com/aethiopicuschan/nanoda"
	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

func playback(ctxOto *oto.Context, s nanoda.Synthesizer, q nanoda.AudioQuery) error {
	w, err := s.Synthesis(q, nanoda.StyleId(ActorID))
	if err != nil {
		return err
	}
	defer w.Close()
	decoded, err := wav.DecodeWithoutResampling(w)
	if err != nil {
		return err
	}
	p := ctxOto.NewPlayer(decoded)
	defer p.Close()
	p.Play()
	for p.IsPlaying() {
		time.Sleep(10 * time.Millisecond)
	}
	return nil
}

func StartEngine(ctx context.Context, in <-chan string) error {
	v, err := nanoda.NewVoicevox(
		"voicevox_core/voicevox_core.dll",
		"voicevox_core/open_jtalk_dic_utf_8-1.11",
		"voicevox_core/model")
	if err != nil {
		return err
	}
	ctxOto, _, err := oto.NewContext(&oto.NewContextOptions{
		SampleRate:   48000,
		ChannelCount: 1,
		Format:       oto.FormatSignedInt16LE,
	})
	if err != nil {
		return err
	}
	s, err := v.NewSynthesizer()
	if err != nil {
		return err
	}
	if err := s.LoadModelsFromStyleId(nanoda.StyleId(ActorID)); err != nil {
		return err
	}
	qDicts, err := Init(s, Dict)
	if err != nil {
		return err
	}
	for {
		select {
		case <-ctx.Done():
			return nil
		case v := <-in:
			qm, ok := qDicts[v]
			if !ok {
				q, err := makeAudioQuery(s, v)
				if err != nil {
					return err
				}
				qm = q
			}
			return playback(ctxOto, s, qm)
		}
	}
}
