package main

import (
	"os"

	"gocv.io/x/gocv"
)

func main() {
	filename := os.Args[1]
	img := gocv.IMRead(filename, gocv.IMReadColor)

	grayimg := gocv.NewMat()
	defer grayimg.Close()

	gocv.CvtColor(img, &grayimg, gocv.ColorBGRToGray)
	gocv.MedianBlur(grayimg, &grayimg, 5)

	maskedges := gocv.NewMat()
	defer maskedges.Close()

	gocv.AdaptiveThreshold(grayimg, &maskedges, 255, gocv.AdaptiveThresholdMean, gocv.ThresholdBinary, 9, 9)

	imgcolor := gocv.NewMat()
	defer imgcolor.Close()

	gocv.BilateralFilter(img, &imgcolor, 9, 10, 10)

	imgsticker := gocv.NewMat()
	defer imgsticker.Close()
	gocv.BitwiseAndWithMask(imgcolor, imgcolor, &imgsticker, maskedges)

	gocv.IMWrite("output_go.png", imgsticker)

}
