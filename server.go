package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

type server struct {
	rooms        map[string]*room
	instructions chan instruction
	stopServer   bool
}

func newServer() *server {
	return &server{
		rooms:        make(map[string]*room),
		instructions: make(chan instruction),
	}
}

func (s *server) run() {
	for ins := range s.instructions {
		switch ins.inst {
		case "join":
			s.join(ins.client, ins.args)
		case "name":
			s.name(ins.client, ins.args)
		case "msg":
			s.msg(ins.client, ins.args)
		case "room_room":
			s.list(ins.client, ins.args)
		case "quit":
			s.quit(ins.client, ins.args)
		default:
		}
	}
}
func (s *server) start(listener net.Listener) {
	go s.run()
	go func() {
		for {
			con, errAcc := listener.Accept()
			if errAcc != nil {
				log.Printf("unable accept connection: %s", errAcc.Error())
				continue
			}
			go s.addClient(con)
			if s.stopServer {
				for _, r := range s.rooms {
					r.stopAll()
				}
				return
			}
		}
	}()
}

func (s *server) stop() {
	log.Println("got interruption signal")
	go func() {
		for _, r := range s.rooms {
			r.broadcastAll("server will be shutdown")
		}
	}()
	time.Sleep(5 * time.Second)
}

func (s *server) addClient(netCon net.Conn) {
	log.Printf("new client added:%s", netCon.RemoteAddr().String())
	cli := &client{
		con:          netCon,
		name:         netCon.RemoteAddr().String(),
		room:         nil,
		instructions: s.instructions,
	}
	cli.readInstructions()
}

func (s *server) name(cli *client, args []string) {
	cli.name = args[1]
	cli.msg(fmt.Sprintf("name set to: %s", cli.name))
}

func (s *server) join(cli *client, args []string) {
	r, exist := s.rooms[args[1]]
	if !exist {
		r = &room{
			name:    args[1],
			members: make(map[net.Addr]*client),
		}
		s.rooms[args[1]] = r
	}
	r.members[cli.con.RemoteAddr()] = cli

	s.quitCurrentRoom(cli)
	cli.room = r
	r.broadcast(cli, fmt.Sprintf("%s joined to room", cli.name))
	cli.msg(fmt.Sprintf("you joined to room: %s", r.name))
}

func (s *server) msg(cli *client, args []string) {
	if cli.room == nil {
		cli.err(errors.New("you must join to room first"))
		return
	}
	cli.room.broadcast(cli, cli.name+": "+strings.Join(args[1:len(args)], ""))
}
func (s *server) list(cli *client, args []string) {
	var rooms []string
	for name := range s.rooms {
		rooms = append(rooms, name)
	}
	cli.msg(fmt.Sprintf("room list are%v", rooms))
}

func (s *server) quit(cli *client, args []string) {
	log.Printf("client: %s left ", cli.con.RemoteAddr().String())
	s.quitCurrentRoom(cli)
	cli.msg("you left")
	cli.con.Close()
}

func (s *server) quitCurrentRoom(cli *client) {
	if cli.room != nil {
		oldRoom := s.rooms[cli.room.name]
		delete(s.rooms[cli.room.name].members, cli.con.RemoteAddr())
		oldRoom.broadcast(cli, fmt.Sprintf("%s has left room", cli.name))
	}
}
