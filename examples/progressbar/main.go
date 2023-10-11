package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gabe565/go-spinners"
	"github.com/schollz/progressbar/v3"
)

func main() {
	bar := progressbar.NewOptions(-1,
		// Set spinner frames
		progressbar.OptionSpinnerCustom(spinner.Dots.Frames),
		progressbar.OptionSetDescription("downloading"),
		progressbar.OptionSetWriter(os.Stderr),
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetWidth(10),
		progressbar.OptionThrottle(65*time.Millisecond),
		progressbar.OptionShowCount(),
		progressbar.OptionOnCompletion(func() {
			fmt.Fprint(os.Stderr, "\n")
		}),
		progressbar.OptionFullWidth(),
		progressbar.OptionSetRenderBlankState(true),
	)
	for i := 0; i < 1000; i++ {
		bar.Add(1)
		time.Sleep(5 * time.Millisecond)
	}
}
