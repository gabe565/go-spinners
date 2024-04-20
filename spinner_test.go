package spinner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseFrames(t *testing.T) {
	t.Parallel()
	type args[T any] struct {
		frames []T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[string]{
		{"zero frames", args[string]{frames: []string{}}, []string{}},
		{"one frame", args[string]{frames: []string{"a"}}, []string{"a"}},
		{"two frames", args[string]{frames: []string{"a", "b"}}, []string{"b", "a"}},
		{"three frames", args[string]{frames: []string{"a", "b", "c"}}, []string{"c", "b", "a"}},
		{"four frames", args[string]{frames: []string{"a", "b", "c", "d"}}, []string{"d", "c", "b", "a"}},
		{"five frames", args[string]{frames: []string{"a", "b", "c", "d", "e"}}, []string{"e", "d", "c", "b", "a"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, reverseFrames(tt.args.frames))
		})
	}
}

func Test_boomerangFrames(t *testing.T) {
	t.Parallel()
	type args[T any] struct {
		frames []T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[string]{
		{"zero frames", args[string]{frames: []string{}}, []string{}},
		{"one frame", args[string]{frames: []string{"a"}}, []string{"a"}},
		{"two frames", args[string]{frames: []string{"a", "b"}}, []string{"a", "b"}},
		{"three frames", args[string]{frames: []string{"a", "b", "c"}}, []string{"a", "b", "c", "b"}},
		{"four frames", args[string]{frames: []string{"a", "b", "c", "d"}}, []string{"a", "b", "c", "d", "c", "b"}},
		{"five frames", args[string]{frames: []string{"a", "b", "c", "d", "e"}}, []string{"a", "b", "c", "d", "e", "d", "c", "b"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, boomerangFrames(tt.args.frames))
		})
	}
}
