package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	serv := newServer()
	//go serv.run()
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("unable stert chat room: %s", err.Error())
	}
	defer listener.Close()
	log.Printf("start chat room on port 8888")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	serv.start(listener)
	<-stop
	serv.stop()
}
