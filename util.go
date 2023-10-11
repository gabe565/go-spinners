package spinner

// ReverseFrames makes the spinner plays backwards
func ReverseFrames[T any](frames []T) []T {
	for i := 0; i < len(frames)/2; i += 1 {
		frames[i], frames[len(frames)-1-i] = frames[len(frames)-1-i], frames[i]
	}
	return frames
}

// BoomerangFrames makes the spinner loop back and forth
func BoomerangFrames[T any](frames []T) []T {
	if len(frames) < 3 {
		return frames
	}
	for i := len(frames) - 2; i > 0; i -= 1 {
		frames = append(frames, frames[i])
	}
	return frames
}
