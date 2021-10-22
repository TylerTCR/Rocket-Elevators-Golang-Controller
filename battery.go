package main

import (
	"math"
	"strconv"
)

type Battery struct {
	ID                        int
	amountOfColumns           int
	amountOfFloors            int
	amountOfBasements         int
	amountOfElevatorPerColumn int
	status                    string
	columnsList               []Column
	floorRequestButtonList    []FloorRequestButton
}

func NewBattery(_id, _amountOfColumns, _amountOfFloors, _amountOfBasements, _amountOfElevatorPerColumn int) *Battery {
	return &Battery{
		ID:                        _id,
		status:                    "online",
		amountOfColumns:           _amountOfColumns,
		amountOfFloors:            _amountOfFloors,
		amountOfBasements:         _amountOfBasements,
		amountOfElevatorPerColumn: _amountOfElevatorPerColumn,
		columnsList:               createColumns(_amountOfColumns, _amountOfFloors, _amountOfBasements, _amountOfElevatorPerColumn),
		floorRequestButtonList:    createFloorRequestButtons(_amountOfFloors, _amountOfBasements),
	}
}

// Find the best column based on the requested floor
func (b *Battery) findBestColumn(_requestedFloor int) *Column {
	for _, column := range b.columnsList {
		for i := 0; i < len(column.servedFloors); i++ {
			if column.servedFloors[i] == _requestedFloor {
				return &column
			}
		}
	}
	return nil
}

//Simulate when a user press a button at the lobby
func (b *Battery) assignElevator(_requestedFloor int, _direction string) (*Column, *Elevator) {
	// Determine the chosen column
	chosenColumn := b.findBestColumn(_requestedFloor)
	// Determine the chosen elevator within the column
	chosenElevator := chosenColumn.findElevator(1, _direction)

	// Add the Lobby to the elevator's request list
	chosenElevator.addNewRequest(1)
	// Now move the elevator to the lobby
	chosenElevator.move()

	// Add the requested floor to the elevator's request list
	chosenElevator.addNewRequest(_requestedFloor)
	// Now move the elevator to the requested floor
	chosenElevator.move()

	return chosenColumn, chosenElevator
}

// Create columns, both underground and above ground
func createColumns(amountOfColumns, amountOfFloors, amountOfBasements, amountOfElevatorsPerColumn int) []Column {
	columnsList := []Column{}
	colAmount := amountOfColumns
	colID := 1
	// Above-ground variables
	temp := float64(amountOfFloors) / (float64(amountOfColumns) - 1)
	amountOfFloorsPerColumn := math.Ceil(temp)
	floor := 1
	// Below-ground variables
	isBasementDone := false
	basementServedFloors := make([]int, 0)
	basementFloor := -1

	// If there are any basements, first create basement column
	if amountOfBasements > 0 && isBasementDone == false {
		// Get the amount of floors the basement column would serve
		for bi := 0; bi < amountOfBasements; bi++ {
			basementServedFloors = append(basementServedFloors, basementFloor)
			basementFloor--
		}

		basementColumn := NewColumn(strconv.Itoa(colID), amountOfElevatorsPerColumn, basementServedFloors, true)
		columnsList = append(columnsList, *basementColumn)
		isBasementDone = true
		colID++
		colAmount--
	}

	// For each column and each above-ground floor within the column, add a floor to the servedFloors list
	for i := 0; i < colAmount; i++ {
		servedFloors := make([]int, 0)
		// Loop until gone through the amount of floors for that column to get servedFloors
		for j := 0; j < int(amountOfFloorsPerColumn); j++ {
			if floor <= amountOfFloors {
				servedFloors = append(servedFloors, floor)
				floor++
			}
		}

		// Create a column then add it to the list of columns
		column := NewColumn(strconv.Itoa(colID), amountOfElevatorsPerColumn, servedFloors, false)
		columnsList = append(columnsList, *column)
	}
	return columnsList
}

func createFloorRequestButtons(amountOfFloors, amountOfBasements int) []FloorRequestButton {
	buttonsList := []FloorRequestButton{}
	buttonFloor := 1
	floorRequestButtonID := 1
	basementButtonFloor := -1
	basementRequestButtonID := 1

	// For each basement, create a floor request button
	for bi := 0; bi < amountOfBasements; bi++ {
		basementFloorRequestButton := NewFloorRequestButton(basementRequestButtonID, basementButtonFloor, "down")
		buttonsList = append(buttonsList, *basementFloorRequestButton)
		basementButtonFloor--
		basementRequestButtonID++
	}

	// For each above ground floor, create a floor request button
	for i := 0; i < amountOfFloors; i++ {
		floorRequestButton := NewFloorRequestButton(floorRequestButtonID, buttonFloor, "up")
		buttonsList = append(buttonsList, *floorRequestButton)
		buttonFloor++
		floorRequestButtonID++
	}

	return buttonsList
}
