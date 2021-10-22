package main

//FloorRequestButton is a button on the pannel at the lobby to request any floor
type FloorRequestButton struct {
	ID        int
	status    string
	floor     int
	direction string
}

func NewFloorRequestButton(_id int, _floor int, _direction string) *FloorRequestButton {
	return &FloorRequestButton{ID: _id, floor: _floor, direction: _direction, status: "off"}
}
