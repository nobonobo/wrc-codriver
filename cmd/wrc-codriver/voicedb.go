package main

import (
	"flag"

	"github.com/aethiopicuschan/nanoda"
)

var (
	ActorID = 3
	Speed   = 1.5
	Pitch   = 0.0
	Offset  = 0.0
	Volume  = 1.5
)

func init() {
	flag.IntVar(&ActorID, "actor", ActorID, "actor id")
	flag.Float64Var(&Speed, "speed", Speed, "speed")
	flag.Float64Var(&Pitch, "pitch", Pitch, "pitch")
	flag.Float64Var(&Volume, "volume", Volume, "volume")
	flag.Float64Var(&Offset, "offset", Offset, "offset")
	flag.Parse()
}

type Dict map[string]struct {
	Text string
}

var (
	marks = Dict{
		"1-left":              {Text: "いち・レフト"},
		"1-right":             {Text: "いち・ライト"},
		"2-left":              {Text: "に・レフト"},
		"2-right":             {Text: "に・ライト"},
		"3-left":              {Text: "さん・レフト"},
		"3-right":             {Text: "さん・ライト"},
		"4-left":              {Text: "よん・レフト"},
		"4-right":             {Text: "よん・ライト"},
		"5-left":              {Text: "ご・レフト"},
		"5-right":             {Text: "ご・ライト"},
		"6-left":              {Text: "ろく・レフト"},
		"6-right":             {Text: "ろく・ライト"},
		"slight-left":         {Text: "スライトレフト"},
		"slight-right":        {Text: "スライトライト"},
		"square-left":         {Text: "スクウェアレフト"},
		"square-right":        {Text: "スクウェアライト"},
		"hp-left":             {Text: "ヘアピンレフト"},
		"hp-right":            {Text: "ヘアピンライト"},
		"acute-hp-left":       {Text: "アキュートヘアピンレフト"},
		"acute-hp-right":      {Text: "アキュートヘアピンライト"},
		"jump":                {Text: "ジャンプ"},
		"crest":               {Text: "クレスト"},
		"dip":                 {Text: "ディップ"},
		"bump":                {Text: "バンプ"},
		"bridge":              {Text: "ブリッジ"},
		"water-splash":        {Text: "ウォータースプラッシュ"},
		"left-entry-chicane":  {Text: "レフトエントリーシケイン"},
		"right-entry-chicane": {Text: "ライトエントリーシケイン"},
		"straight":            {Text: "ストレート"},
		"finish":              {Text: "フィニッシュ！"},
	}
	icons = Dict{
		"bridge":           {Text: "ブリッジ"},
		"dont-cut":         {Text: "ドントカット"},
		"cut":              {Text: "カット"},
		"caution":          {Text: "コーション"},
		"terrible-caution": {Text: "テリブル・コーション"},
		"narrow":           {Text: "ナロー"},
		"widen":            {Text: "ワイドゥン"},
		"tighten":          {Text: "タイトゥン"},
		"twisty":           {Text: "ツウィスティ"},
	}
	dists = Dict{
		"30":  {Text: "３０メートル"},
		"40":  {Text: "４０メートル"},
		"50":  {Text: "５０メートル"},
		"60":  {Text: "６０メートル"},
		"70":  {Text: "７０メートル"},
		"80":  {Text: "８０メートル"},
		"90":  {Text: "９０メートル"},
		"100": {Text: "１００メートル"},
		"110": {Text: "１１０メートル"},
		"120": {Text: "１２０メートル"},
		"140": {Text: "１４０メートル"},
		"160": {Text: "１６０メートル"},
		"170": {Text: "１７０メートル"},
		"180": {Text: "１８０メートル"},
		"190": {Text: "１９０メートル"},
		"200": {Text: "２００メートル"},
		"220": {Text: "２２０メートル"},
		"250": {Text: "２５０メートル"},
		"280": {Text: "２８０メートル"},
		"300": {Text: "３００メートル"},
		"350": {Text: "３５０メートル"},
		"370": {Text: "３７０メートル"},
	}
)

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

func Setup(s nanoda.Synthesizer, dict Dict) (map[string]nanoda.AudioQuery, error) {
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
