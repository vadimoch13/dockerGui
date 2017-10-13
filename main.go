package main

import (
	"flag"
	"log"

	"./daemon"
	"./global"
)

func processFlags(){
	flag.StringVar(&global.CfgDaemon.ListenSpec, "listen", "localhost:3000", "HTTP listen spec")
	flag.StringVar(&global.Auth.Login, "auth-login", "admin", "username")
	flag.StringVar(&global.Auth.Password, "auth-password", "admin", "password")
	flag.Parse()
}

func main(){
	processFlags()

	if err := daemon.Run(); err != nil {
		log.Printf("Error in main(): %v", err)

	}
}