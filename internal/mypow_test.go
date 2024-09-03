package internal

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Mypow(t *testing.T) {
	type args struct {
		Base, Power int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case base 2 power 10 then 1024",
			args: args{
				Base:  2,
				Power: 10,
			},
			want: 1024,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Mypow(tt.args.Base, tt.args.Power)
			if !assert.Equal(t, tt.want, got, tt.name) {
				fmt.Printf("Error when asserting. Want: %v, got: %v", tt.want, got)
			}
		})
	}
}
