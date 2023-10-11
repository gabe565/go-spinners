package spinner

import "time"

// Spinner defines a command-line spinner.
type Spinner struct {
	// Name is the spinner name
	Name string
	// Frames is a list of frames for a spinner.
	Frames []string
	// Interval is the recommended interval for a spinner.
	Interval time.Duration
}
