package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	spinner "github.com/gabe565/go-spinners"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGTERM)
	defer cancel()

	// Hide cursor
	fmt.Print("\x1B[?25l")
	defer func() {
		// Show cursor before exit
		fmt.Println("\x1B[?25h")
	}()

	select {
	case <-ctx.Done():
		break
	default:
		animate(ctx, spinner.Random())
	}
}

func animate(ctx context.Context, sp spinner.Spinner) {
	var exitAfter bool
	go func() {
		<-time.After(4 * time.Second)
		exitAfter = true
	}()

	var frame int
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Print("\r\x1B[K" + sp.Frames[frame] + " " + sp.Name)
			frame++
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
