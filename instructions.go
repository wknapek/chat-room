package main

const (
	Join  = "join"
	Name  = "name"
	MSG   = "msg"
	Rooms = "room_list"
	Quit  = "quit"
)

type instruction struct {
	inst   string
	client *client
	args   []string
}
