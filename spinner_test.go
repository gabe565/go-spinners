package spinner

import (
	"reflect"
	"testing"
)

func Test_reverseFrames(t *testing.T) {
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
			if got := reverseFrames(tt.args.frames); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseFrames() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_boomerangFrames(t *testing.T) {
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
			if got := boomerangFrames(tt.args.frames); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BoomerangFrames() = %v, want %v", got, tt.want)
			}
		})
	}
}
