package internal

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findTargetOrClosest(t *testing.T) {
	type args struct {
		Numbers []int
		Target  int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case Target exists within numbers, then return target",
			args: args{
				Numbers: []int{
					2, 3, 5, 6,
				},
				Target: 5,
			},
			want: 5,
		},
		{
			name: "Case Target doesn't exist within numbers but exists smaller than target, then return the number smaller than target which exists.",
			args: args{
				Numbers: []int{
					2, 3, 5, 6,
				},
				Target: 4,
			},
			want: 3,
		},
		{
			name: "Case Target doesn't exist within numbers, and is out of bounds, return -1 (Leftmost)",
			args: args{
				Numbers: []int{
					2, 3, 5, 6,
				},
				Target: 1,
			},
			want: -1,
		},
		{
			name: "Case Target doesn't exist within numbers, and is out of bounds, return -1 (Rightmost)",
			args: args{
				Numbers: []int{
					2, 3, 5, 6,
				},
				Target: 1,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findTargetOrClosest(tt.args.Numbers, tt.args.Target)
			if !assert.Equal(t, tt.want, got, tt.name) {
				fmt.Printf("Error when asserting. Want: %v, got: %v", tt.want, got)
			}
		})
	}
}
