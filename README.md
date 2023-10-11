# go-spinners

70+ spinners for use in terminal Go apps.

This repo is autogenerated from [sindresorhus/cli-spinners](https://github.com/sindresorhus/cli-spinners).

Intended to be used with [schollz/progressbar](https://github.com/schollz/progressbar), but the spinners will work with any tool that accepts a `[]string`.

## Installation

```shell
go get github.com/gabe565/go-spinners
```

## Usage

See [`spinners.go`](./spinners.go) for a list of available spinners, or see one of the [preview pages](#preview).

### Examples

<details>
  <summary>Simple</summary>

```go
package main

import (
	"fmt"

	"github.com/gabe565/go-spinners"
)

func main() {
	fmt.Println(spinner.Dots)
}
```
</details>

<details>
  <summary>Progress bar</summary>

```go
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
```
</details>

## Preview

See all [spinners at once](https://jsfiddle.net/sindresorhus/2eLtsbey/embedded/result/) or [one at the time](https://asciinema.org/a/95348?size=big).
