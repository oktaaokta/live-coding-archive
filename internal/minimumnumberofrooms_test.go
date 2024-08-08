package internal

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MinimumNumberOfRooms(t *testing.T) {
	type args struct {
		Timeslots [][]int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "When timeslot is sequential return 1",
			args: args{
				Timeslots: [][]int{
					{
						0, 30,
					},
					{
						31, 35,
					},
					{
						40, 50,
					}},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMinimumNumberOfRooms(tt.args.Timeslots)
			if !assert.Equal(t, tt.want, got, tt.name) {
				fmt.Printf("Error when asserting. Want: %v, got: %v", tt.want, got)
			}
		})
	}
}
