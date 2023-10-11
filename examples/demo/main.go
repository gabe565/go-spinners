package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gabe565/go-spinners"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGTERM)
	defer cancel()

	demos := []spinner.Spinner{
		spinner.Dots,
		spinner.Dots12,
		spinner.Sand,
		spinner.Line,
		spinner.SimpleDotsScrolling,
		spinner.Point,
		spinner.Pong,
		spinner.BluePulse,
		spinner.Moon,
		spinner.Arrow3,
		spinner.BouncingBall,
		spinner.Aesthetic,
	}

	// Hide cursor
	fmt.Println("\x1B[?25l")
	defer func() {
		// Show cursor before exit
		fmt.Println("\x1B[?25h")
	}()
	for _, sp := range demos {
		select {
		case <-ctx.Done():
			break
		default:
			animate(ctx, sp)
		}
	}
}

func animate(ctx context.Context, sp spinner.Spinner) {
	var exitAfter bool
	go func() {
		<-time.After(2 * time.Second)
		exitAfter = true
	}()

	var frame int
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Print("\r\x1B[K  " + sp.Frames[frame] + " " + sp.Name)
			frame += 1
			frame %= len(sp.Frames)
			select {
			case <-ctx.Done():
				return
			case <-time.After(sp.Interval):
			}
		}

		if exitAfter && frame == 0 {
			return
		}
	}
}
