package main

import (
	"flag"
	"log"

	"./daemon"
	"./global"
)

func processFlags(){
	// cfg := &global.ConfigDaemon{}

	flag.StringVar(&global.CfgDaemon.ListenSpec, "listen", "localhost:3000", "HTTP listen spec")
	// flag.StringVar(&cfg.UI.BasePath, "assets-path", "/home/user/GoglandProjects/docker-gui/Public", "Setting path to public files")
	flag.StringVar(&global.Auth.Login, "auth-login", "admin", "username")
	flag.StringVar(&global.Auth.Password, "auth-password", "admin", "password")
	flag.Parse()
	// return cfg
}

func main(){
	processFlags()

	if err := daemon.Run(); err != nil {
		log.Printf("Error in main(): %v", err)

	}
}