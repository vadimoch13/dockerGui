package daemon

import (
	"log"
	"net"
	"os"
	"syscall"
	"os/signal"

	"../ui"
	"../global"
)

func Run() error {
	log.Printf("Starting, HTTP on: %s\n", global.CfgDaemon.ListenSpec)
	l, err := net.Listen("tcp", global.CfgDaemon.ListenSpec)
	if err != nil {
		log.Printf("Error creating listener: %v\n", err)
		return err
	}
	ui.Start(l)

	waitForSignal()
	return nil
}

func waitForSignal() {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	s := <-ch
	log.Printf("Got signal: %v, exiting.", s)
}