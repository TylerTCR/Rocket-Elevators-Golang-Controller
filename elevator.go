package main

import (
	"fmt"
	"sort"
)

type Elevator struct {
	ID                    string
	status                string
	currentFloor          int
	direction             string
	door                  Door
	floorRequestsList     []int
	completedRequestsList []int
}

func NewElevator(_elevatorID string) *Elevator {
	return &Elevator{
		ID:                    _elevatorID,
		status:                "idle",
		currentFloor:          1,
		direction:             "",
		door:                  *NewDoor(1),
		floorRequestsList:     []int{},
		completedRequestsList: []int{},
	}
}

func (elev *Elevator) move() {
	// While the elevator's floor request list is not empty
	if len(elev.floorRequestsList) != 0 {
		destination := elev.floorRequestsList[0]
		// Loop until all floors in the list have been gone through
		for i := 0; i < len(elev.floorRequestsList); i++ {
			// If the elevator's current floor is lower than the destination
			if elev.currentFloor < destination {
				elev.direction = "up"
				elev.sortFloorList(elev.floorRequestsList)
				if elev.door.status == "opened" {
					elev.door.status = "closed"
				}
				elev.status = "moving"

				// Move the elevator until it reaches the floor
				for j := elev.currentFloor; j < destination; j++ {
					elev.currentFloor++
					fmt.Printf("Floor: %d", elev.currentFloor)
				}
				// If the elevator's current floor is higher than the destination
			} else if elev.currentFloor > destination {
				elev.direction = "down"
				elev.sortFloorList(elev.floorRequestsList)
				if elev.door.status == "opened" {
					elev.door.status = "closed"
				}
				elev.status = "moving"

				// Move the elevator until it reaches the floor
				for j := elev.currentFloor; j > destination; j++ {
					elev.currentFloor--
					fmt.Printf("Floor: %d", elev.currentFloor)
				}
			}
			// Set elevator's status to stopped
			elev.status = "stopped"
			// Open the doors
			elev.door.status = "opened"
			// Add the destination to completed request list
			elev.completedRequestsList = append(elev.completedRequestsList, destination)
			// Remove the first element in the request list, keeping the order of the list
			r := 0 // Element to remove
			copy(elev.floorRequestsList[r:], elev.floorRequestsList[r+1:])
			elev.floorRequestsList[len(elev.floorRequestsList)-1] = 0
			elev.floorRequestsList = elev.floorRequestsList[:len(elev.floorRequestsList)-1]
		}
	}
}

// Sort the request list in either asceding or descending order
func (e *Elevator) sortFloorList(requestList []int) {
	if e.direction == "up" {
		sort.Ints(requestList)
	} else {
		sort.Ints(requestList)
		sort.Sort(sort.Reverse(sort.IntSlice(requestList)))
	}

}

// Add a new request to the floor request list, then set the direction
func (e *Elevator) addNewRequest(requestedFloor int) {
	// If the request is not already in the list, add it in
	if !(isElementExist(e.floorRequestsList, requestedFloor)) {
		e.floorRequestsList = append(e.floorRequestsList, requestedFloor)
	}

	if e.currentFloor < requestedFloor {
		e.direction = "up"
	} else if e.currentFloor > requestedFloor {
		e.direction = "down"
	}
}

func isElementExist(requestList []int, checkNum int) bool {
	for _, v := range requestList {
		if v == checkNum {
			return true
		}
	}

	return false
}
