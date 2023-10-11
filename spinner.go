package spinner

import "time"

type Spinner struct {
	Frames   []string
	Interval time.Duration
}
