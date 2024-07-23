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

// Reverse returns a new Spinner that plays backwards.
func (s Spinner) Reverse() Spinner {
	s.Frames = reverseFrames(s.Frames)
	return s
}

// ReverseFrames makes a slice of frames play backwards.
func reverseFrames[T any](frames []T) []T {
	result := make([]T, len(frames))
	for i := range len(result) {
		result[i] = frames[len(frames)-1-i]
	}
	return result
}

// Boomerang returns a new Spinner that loops back and forth.
func (s Spinner) Boomerang() Spinner {
	s.Frames = boomerangFrames(s.Frames)
	return s
}

// BoomerangFrames makes a slice of frames loop back and forth.
func boomerangFrames[T any](frames []T) []T {
	if len(frames) < 3 {
		return frames
	}
	result := make([]T, len(frames), 2*len(frames)-2)
	copy(result, frames)
	for i := len(result) - 2; i > 0; i-- {
		result = append(result, result[i])
	}
	return result
}
