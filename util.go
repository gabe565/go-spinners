package spinner

import "math/rand"

// Random returns a random spinner.
func Random() Spinner {
	return Slice[rand.Intn(len(Slice))]
}
