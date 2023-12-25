package main

import (
	"image"
	"image/color"
	"log"
	"math"
	"path/filepath"

	"gocv.io/x/gocv"
	"gocv.io/x/gocv/contrib"
)

const center = 1920 / 2

var (
	Marks    = map[string]gocv.Mat{}
	Dists    = map[string]gocv.Mat{}
	Icons    = map[string]gocv.Mat{}
	markRect = image.Rectangle{
		Min: image.Point{X: center - 72, Y: 112},
		Max: image.Point{X: center + 72, Y: 112 + 144},
	}
	iconRect = image.Rectangle{
		Min: image.Point{X: 104, Y: 4},
		Max: image.Point{X: 143, Y: 43},
	}
	distRect = image.Rectangle{
		Min: image.Point{X: 65, Y: 94},
		Max: image.Point{X: 138, Y: 132},
	}
	hash = contrib.PHash{}
)

func init() {
	names := []string{
		"1-left",
		"1-right",
		"2-left",
		"2-right",
		"3-left",
		"3-right",
		"4-left",
		"4-right",
		"5-left",
		"5-right",
		"6-left",
		"6-right",
		"slight-left",
		"slight-right",
		"square-left",
		"square-right",
		"hp-left",
		"acute-hp-left",
		"hp-right",
		"acute-hp-right",
		"jump",
		"crest",
		"dip",
		"bump",
		"bridge",
		"water-splash",
		"left-entry-chicane",
		"right-entry-chicane",
		"straight",
		"finish",
		"sample",
	}
	nums := []string{
		"30",
		"40",
		"50",
		"60",
		"70",
		"80",
		"90",
		"100",
		"110",
		"120",
		"140",
		"160",
		"170",
		"180",
		"190",
		"200",
		"220",
		"250",
		"280",
		"300",
		"350",
		"370",
	}
	icons := []string{
		"bridge",
		"dont-cut",
		"cut",
		"caution",
		"terrible-caution",
		"narrow",
		"widen",
		"twisty",
	}
	for _, n := range names {
		log.Print("load:", n)
		img := gocv.IMRead(filepath.Join("assets", "sign", n+".png"), gocv.IMReadColor)
		markPreProcess(&img)
		gocv.IMWrite(filepath.Join("assets", "sign", n+"_th.png"), img)
		compute := gocv.NewMat()
		hash.Compute(img, &compute)
		img.Close()
		if compute.Empty() {
			log.Print("empty")
			continue
		}
		if n != "sample" {
			Marks[n] = compute
		}
	}
	for _, n := range nums {
		log.Print("load:", n)
		img := gocv.IMRead(filepath.Join("assets", "distance", n+".png"), gocv.IMReadColor)
		dist := img.Region(distRect)
		distPreProcess(&dist)
		gocv.IMWrite(filepath.Join("assets", "distance", n+"_th.png"), dist)
		compute := gocv.NewMat()
		hash.Compute(dist, &compute)
		if compute.Empty() {
			log.Print("empty")
			continue
		}
		Dists[n] = compute
	}
	for _, n := range icons {
		log.Print("load:", n)
		img := gocv.IMRead(filepath.Join("assets", "icon", n+".png"), gocv.IMReadColor)
		sub := img.Region(iconRect)
		iconPreProcess(&sub)
		gocv.IMWrite(filepath.Join("assets", "icon", n+"_th.png"), sub)
		compute := gocv.NewMat()
		hash.Compute(sub, &compute)
		if compute.Empty() {
			log.Print("empty")
			continue
		}
		Icons[n] = compute
	}
}

var (
	innerRect = image.Rectangle{
		Min: image.Point{X: 22, Y: 22},
		Max: image.Point{X: 144 - 24, Y: 144 - 24},
	}
	acc      = gocv.NewMat()
	frame    = gocv.NewMat()
	empty    = gocv.NewMat()
	histgram = gocv.NewMat()
	cntStop  = 0
)

func getMotion(mark gocv.Mat) int32 {
	inner := mark.Region(innerRect)
	gocv.DetailEnhance(inner, &inner, 20, 0.5)
	gocv.CvtColor(inner, &inner, gocv.ColorBGRToGray)
	gocv.GaussianBlur(inner, &inner, image.Point{X: 5, Y: 5}, 0, 0, gocv.BorderDefault)
	gocv.AdaptiveThreshold(inner, &inner, 255, gocv.AdaptiveThresholdGaussian, gocv.ThresholdBinary, 13, 20)
	if acc.Empty() {
		inner.ConvertTo(&acc, gocv.MatTypeCV32FC1)
		return 0
	}
	gocv.AccumulatedWeighted(inner, &acc, 0.3)
	gocv.ConvertScaleAbs(acc, &frame, 1.0, 0)
	gocv.AbsDiff(frame, inner, &frame)
	gocv.GaussianBlur(frame, &frame, image.Point{X: 5, Y: 5}, 0, 0, gocv.BorderDefault)
	gocv.CalcHist([]gocv.Mat{frame}, []int{0}, empty, &histgram, []int{2}, []float64{0, 255}, false)
	return histgram.GetIntAt(0, 1)
}

func getMotionStop(mark gocv.Mat) bool {
	stop := false
	m := getMotion(mark)
	if m > 100 {
		if cntStop > 0 {
			cntStop--
		}
	} else {
		if cntStop < 3 {
			cntStop++
		}
		if cntStop == 2 {
			stop = true
		}
	}
	return stop
}

var sigmoid_cvt = map[float64]gocv.Mat{}

func sigmoid(x, alpha float64) float64 {
	return 1 / (1 + math.Exp(-alpha*x))
}

func sigmoidConversion(src gocv.Mat, dst *gocv.Mat, alpha float64) {
	cvt, ok := sigmoid_cvt[alpha]
	if !ok {
		cvt = gocv.NewMatWithSize(256, 1, gocv.MatTypeCV8U)
		for r := 0; r < 256; r++ {
			v := 255.0 * sigmoid(5*float64(r-127)/127, alpha)
			if v > 255 {
				v = 255
			}
			if v < 0 {
				v = 0
			}
			cvt.SetUCharAt(r, 0, uint8(v))
		}
		sigmoid_cvt[alpha] = cvt
	}
	gocv.LUT(src, cvt, dst)
}

var mask = gocv.NewMat()

func markPreProcess(img *gocv.Mat) {
	if iconMask.Empty() {
		mask = gocv.IMRead(filepath.Join("assets", "mask", "mask.png"), gocv.IMReadColor)
		gocv.CvtColor(mask, &mask, gocv.ColorBGRToGray)
	}
	sigmoidConversion(*img, img, 7)
	gocv.DetailEnhance(*img, img, 20, 0.1)
	gocv.CvtColor(*img, img, gocv.ColorBGRToGray)
	gocv.GaussianBlur(*img, img, image.Point{X: 5, Y: 5}, 0, 0, gocv.BorderDefault)
	// gocv.Normalize(*img, img, 0, 255, gocv.NormMinMax)
	gocv.AdaptiveThreshold(*img, img, 255, gocv.AdaptiveThresholdGaussian, gocv.ThresholdBinary, 13, 10)
	gocv.Add(*img, mask, img)
	gocv.Rectangle(img, iconRect, color.RGBA{255, 255, 255, 0}, -1)
	gocv.Rectangle(img, distRect, color.RGBA{255, 255, 255, 0}, -1)
}

var distMask = gocv.NewMat()

func distPreProcess(img *gocv.Mat) {
	sigmoidConversion(*img, img, 5)
	gocv.DetailEnhance(*img, img, 20, 0.1)
	//gocv.CvtColor(*img, img, gocv.ColorBGRToGray)
	//gocv.Normalize(*img, img, 0, 255, gocv.NormMinMax)
	//gocv.AdaptiveThreshold(*img, img, 255, gocv.AdaptiveThresholdGaussian, gocv.ThresholdBinary, 11, 15)
	if distMask.Empty() {
		distMask = gocv.IMRead(filepath.Join("assets", "mask", "num-mask.png"), gocv.IMReadColor)
		//gocv.CvtColor(distMask, &distMask, gocv.ColorBGRToGray)
	}
	gocv.Add(*img, distMask, img)
}

var iconMask = gocv.NewMat()

func iconPreProcess(img *gocv.Mat) {
	sigmoidConversion(*img, img, 5)
	gocv.DetailEnhance(*img, img, 20, 0.1)
	//gocv.CvtColor(*img, img, gocv.ColorBGRToGray)
	//gocv.Normalize(*img, img, 0, 255, gocv.NormMinMax)
	//gocv.AdaptiveThreshold(*img, img, 255, gocv.AdaptiveThresholdGaussian, gocv.ThresholdBinary, 9, 15)
	if iconMask.Empty() {
		iconMask = gocv.IMRead(filepath.Join("assets", "mask", "icon-mask.png"), gocv.IMReadColor)
		//gocv.CvtColor(iconMask, &iconMask, gocv.ColorBGRToGray)
	}
	gocv.Add(*img, iconMask, img)
}
