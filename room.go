package main

import "net"

type room struct {
	name    string
	members map[net.Addr]*client
}

func (r *room) broadcast(cli *client, msg string) {
	for addr, member := range r.members {
		if addr != cli.con.RemoteAddr() {
			member.msg(msg)
		}
	}
}

func (r *room) broadcastAll(msg string) {
	for _, cli := range r.members {
		r.broadcast(cli, msg)
	}
}

func (r *room) stopAll() {
	for _, cli := range r.members {
		cli.stopClient = true
	}
}
