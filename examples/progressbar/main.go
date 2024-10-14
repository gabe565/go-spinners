package main

import (
	"fmt"
	"os"
	"time"

	spinner "gabe565.com/spinners"
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
	for range 1000 {
		_ = bar.Add(1)
		time.Sleep(5 * time.Millisecond)
	}
}
