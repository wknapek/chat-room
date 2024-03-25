package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

type client struct {
	con          net.Conn
	name         string
	room         *room
	stopClient   bool
	instructions chan<- instruction
}

func (c *client) readInstructions() {
	for {
		msg, err := bufio.NewReader(c.con).ReadString('\n')
		if err != nil {
			return
		}
		msg = strings.Trim(msg, "\r\n")
		args := strings.Split(msg, " ")
		inst := strings.TrimSpace(args[0])
		switch inst {
		case "/join":
			c.instructions <- instruction{
				inst:   Join,
				client: c,
				args:   args,
			}
		case "/name":
			c.instructions <- instruction{
				inst:   Name,
				client: c,
				args:   args,
			}
		case "/msg":
			c.instructions <- instruction{
				inst:   MSG,
				client: c,
				args:   args,
			}
		case "/room_list":
			c.instructions <- instruction{
				inst:   Rooms,
				client: c,
				args:   args,
			}
		case "/quit":
			c.instructions <- instruction{
				inst:   Quit,
				client: c,
				args:   args,
			}
		default:
			c.err(fmt.Errorf("unknown instruction %s", inst))
		}
		if c.stopClient {
			time.Sleep(5 * time.Second)
			return
		}
	}
}

func (c *client) err(err error) {
	c.con.Write([]byte("ERR " + err.Error() + "\n"))
}

func (c *client) msg(msg string) {
	tmp := msg[:128]
	c.con.Write([]byte("> " + tmp + "\n"))
}
