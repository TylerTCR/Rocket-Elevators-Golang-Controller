package main

import (
	"math"
	"strconv"
)

type Column struct {
	ID            string
	status        string
	servedFloors  []int
	isBasement    bool
	elevatorsList []Elevator
	callButtons   []CallButton
}

func NewColumn(_id string, _amountOfElevators int, _servedFloors []int, _isBasement bool) *Column {
	return &Column{
		ID:            _id,
		status:        "online",
		servedFloors:  _servedFloors,
		isBasement:    _isBasement,
		callButtons:   createCallButtons(_servedFloors, _isBasement),
		elevatorsList: createElevators(_servedFloors, _amountOfElevators),
	}
}

func createCallButtons(servedFloors []int, isBasement bool) []CallButton {
	buttonId := 1
	callButtonList := []CallButton{}

	if isBasement {
		buttonFloor := -1

		for i := 0; i < len(servedFloors); i++ {
			callButton := NewCallButton(buttonId, buttonFloor, "up")
			callButtonList = append(callButtonList, *callButton)
			buttonFloor--
			buttonId++
		}
	} else {
		buttonFloor := 1

		for i := 0; i < len(servedFloors); i++ {
			callButton := NewCallButton(buttonId, buttonFloor, "down")
			callButtonList = append(callButtonList, *callButton)
			buttonFloor++
			buttonId++
		}
	}

	return callButtonList
}

func createElevators(servedFloors []int, amountOfElevators int) []Elevator {
	elevatorID := 1
	elevList := []Elevator{}
	for i := 0; i < amountOfElevators; i++ {
		elevator := NewElevator(strconv.Itoa(elevatorID))
		elevList = append(elevList, *elevator)
		elevatorID++
	}

	return elevList
}

//Simulate when a user press a button on a floor to go back to the first floor
func (c *Column) requestElevator(_requestedFloor int, _direction string) *Elevator {

	// Find the elevator to pick up the person
	elevator := c.findElevator(_requestedFloor, _direction)
	elevator.addNewRequest(_requestedFloor)
	elevator.move()

	// Would then go back to lobby
	elevator.addNewRequest(1)
	elevator.move()

	return elevator
}

/* Find the best elevator, prioritizing ones already in motion, heading the same way of the user wants to go,
and closest to the floor where the user is on. */
func (c *Column) findElevator(requestedFloor int, direction string) *Elevator {
	bestElevator := &c.elevatorsList[0]
	bestScore := 100
	referenceGap := 100000

	// If requestedFloor is the lobby
	if requestedFloor == 1 {
		for i:=0 ; i < len(c.elevatorsList); i++ {
			// Elevator is stopped at the lobby and already has requests
			if c.elevatorsList[i].currentFloor == 1 && c.elevatorsList[i].status == "stopped" {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(1, &c.elevatorsList[i], bestScore, referenceGap, bestElevator, requestedFloor)
				// Elevator is idle at the lobby, has no requests
			} else if c.elevatorsList[i].currentFloor == 1 && c.elevatorsList[i].status == "idle" {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(2, &c.elevatorsList[i], bestScore, referenceGap, bestElevator, requestedFloor)
				// Elevator is lower than the user and coming up
			} else if c.elevatorsList[i].currentFloor < 1 && c.elevatorsList[i].direction == "up" {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(2, &c.elevatorsList[i], bestScore, referenceGap, bestElevator, requestedFloor)
				// Elevator is above the user and coming down
			} else if c.elevatorsList[i].currentFloor > 1 && c.elevatorsList[i].direction == "down" {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(3, &c.elevatorsList[i], bestScore, referenceGap, bestElevator, requestedFloor)
				// Elevator is not at the lobby, but is idle and has no requests
			} else if c.elevatorsList[i].status == "idle" {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(4, &c.elevatorsList[i], bestScore, referenceGap, bestElevator, requestedFloor)
				// Elevator is not available, but could take the request if there's nothing better
			} else {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(5, &c.elevatorsList[i], bestScore, referenceGap, bestElevator, requestedFloor)
			}

		} // End for loop
	// If requested floor is not the lobby...
	} else {
		for i:=0 ; i < len(c.elevatorsList); i++ {
			// Elevator is stopped at the same level as user, about to go to lobby
			if c.elevatorsList[i].currentFloor == requestedFloor && c.elevatorsList[i].status == "stopped" && c.elevatorsList[i].direction == direction {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(1, &c.elevatorsList[i], bestScore, referenceGap, bestElevator, requestedFloor)
				// Elevator is lower than user, and going up towards the lobby
			} else if c.elevatorsList[i].currentFloor < requestedFloor && c.elevatorsList[i].direction == "up" && c.elevatorsList[i].direction == direction {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(2, &c.elevatorsList[i], bestScore, referenceGap, bestElevator, requestedFloor)
				// Elevator is above user and going down towards the lobby
			} else if c.elevatorsList[i].currentFloor > requestedFloor && c.elevatorsList[i].direction == "down" && c.elevatorsList[i].direction == direction {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(2, &c.elevatorsList[i], bestScore, referenceGap, bestElevator, requestedFloor)
				// Elevator is idle
			} else if c.elevatorsList[i].status == "idle" {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(4, &c.elevatorsList[i], bestScore, referenceGap, bestElevator, requestedFloor)
				// Elevator is not available but can still take the request
			} else {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(5, &c.elevatorsList[i], bestScore, referenceGap, bestElevator, requestedFloor)
			}
		}
	}

	return bestElevator
}

func (c *Column) checkIfElevatorIsBetter(scoreToCheck int, newElevator *Elevator, bestScore int, referenceGap int, bestElevator *Elevator, floor int) (*Elevator, int, int) {
	if scoreToCheck < bestScore {
		bestScore = scoreToCheck
		bestElevator = newElevator
		referenceGap = int(math.Abs(float64(newElevator.currentFloor) - float64(floor)))
	} else if bestScore == scoreToCheck {
		var gap int = int(math.Abs(float64(newElevator.currentFloor) - float64(floor)))
		if referenceGap > gap {
			bestElevator = newElevator
			referenceGap = gap
		}
	}

	return bestElevator, bestScore, referenceGap
}
