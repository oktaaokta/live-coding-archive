package internal

import "fmt"

func MinimumNumberOfRooms() {
	// Input consisting of start and end times [[s1,e1],[s2,e2],...].
	// find the minimum number of rooms required.
	input := [][]int{
		{
			0, 30,
		},
		{
			31, 35,
		},
		{
			40, 50,
		},
	}

	minimumRooms := findMinimumNumberOfRooms(input)
	fmt.Println("Minimum number of rooms needed:", minimumRooms)
}

type Rooms struct {
	StartTime int
	EndTime   int
}

func findMinimumNumberOfRooms(input [][]int) int {
	var listOfRooms []Rooms
	var roomSet bool
	listOfRooms = append(listOfRooms, Rooms{StartTime: input[0][0], EndTime: input[0][1]})
	for i := 1; i < len(input); i++ {
		for j := range listOfRooms {
			if listOfRooms[j].EndTime < input[i][0] {
				listOfRooms[j] = Rooms{
					StartTime: input[i][0],
					EndTime:   input[i][1],
				}
				roomSet = true
				break
			}
		}
		if !roomSet {
			listOfRooms = append(listOfRooms, Rooms{StartTime: input[i][0], EndTime: input[i][1]})
		}
	}

	return len(listOfRooms)
}
