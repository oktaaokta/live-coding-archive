package internal

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LargestNumber(t *testing.T) {
	type args struct {
		Numbers []int
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Case 1",
			args: args{
				Numbers: []int{
					54, 546, 548, 60,
				},
			},
			want: "6054854654",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LargestNumber(tt.args.Numbers)
			if !assert.Equal(t, tt.want, got, tt.name) {
				fmt.Printf("Error when asserting. Want: %v, got: %v", tt.want, got)
			}
		})
	}
}
