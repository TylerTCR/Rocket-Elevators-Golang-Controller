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
	for len(elev.floorRequestsList) != 0 {
		// Creating a temp array to sort later; this array is used to move the elevator
		tempRequestList := []int{len(elev.floorRequestsList)}
		// Copying the elevator's floor request list so the original list is untouched when sorted
		copy(tempRequestList, elev.floorRequestsList)
		destination := tempRequestList[0]
		
		// If the elevator's current floor is lower than the destination
		if elev.direction == "up" {
			elev.sortFloorList(tempRequestList)
			if elev.door.status == "opened" {
				elev.door.status = "closed"
			}
			elev.status = "moving"

			// Move the elevator until it reaches the floor
			for elev.currentFloor < destination {
				elev.currentFloor++
				fmt.Printf("Floor: %d\n", elev.currentFloor)
			}
			// If the elevator's current floor is higher than the destination
		} else if elev.direction == "down" {
			elev.sortFloorList(tempRequestList)
			if elev.door.status == "opened" {
				elev.door.status = "closed"
			}
			elev.status = "moving"

			// Move the elevator until it reaches the floor
			for elev.currentFloor > destination {
				elev.currentFloor--
				fmt.Printf("Floor: %d\n", elev.currentFloor)
			}
		}
		// Set elevator's status to stopped
		elev.status = "stopped"
		// Open the doors
		elev.door.status = "opened"
		// Add the destination to completed request list
		elev.completedRequestsList = append(elev.completedRequestsList, destination)
		// Remove the first element in the request list, keeping the order of the list
		elev.floorRequestsList = elev.floorRequestsList[1:]
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
