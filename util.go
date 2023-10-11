package spinner

// ReverseFrames makes the spinner play backwards.
func ReverseFrames[T any](frames []T) []T {
	result := make([]T, len(frames))
	copy(result, frames)
	for i := 0; i < len(result)/2; i += 1 {
		result[i], result[len(frames)-1-i] = result[len(frames)-1-i], result[i]
	}
	return result
}

// BoomerangFrames makes the spinner loop back and forth.
func BoomerangFrames[T any](frames []T) []T {
	if len(frames) < 3 {
		return frames
	}
	result := make([]T, len(frames))
	copy(result, frames)
	for i := len(result) - 2; i > 0; i -= 1 {
		result = append(result, result[i])
	}
	return result
}
