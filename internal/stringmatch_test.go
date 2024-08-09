package internal

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ContainsWord(t *testing.T) {
	type args struct {
		Input, Contains string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Case * at end of word and contains word",
			args: args{
				Input:    "makan",
				Contains: "*an",
			},
			want: true,
		},
		{
			name: "Case * at start of word and contains word",
			args: args{
				Input:    "makan",
				Contains: "*ak",
			},
			want: true,
		},
		{
			name: "Case * at both end of word and contains word",
			args: args{
				Input:    "makan",
				Contains: "*ak*",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ContainsWord(tt.args.Input, tt.args.Contains)
			if !assert.Equal(t, tt.want, got, tt.name) {
				fmt.Printf("Error when asserting. Want: %v, got: %v", tt.want, got)
			}
		})
	}
}
