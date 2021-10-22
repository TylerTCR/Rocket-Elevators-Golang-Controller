package main

type Door struct {
	ID     int
	status string
}

func NewDoor(_id int) *Door {
	return &Door{ID: _id, status: "closed"}
}
