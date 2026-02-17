package main

import (
	"log"
	"os"

	"github.com/JoStMc/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal("error reading config: ", err)
	} 
	currentState := state{cfg: &cfg} 
	cmds := commands{cmdMap: make(map[string]func(*state, command) error)}
	cmds.register("login", handlerLogin)

	args := os.Args[1:]
	if len(args) < 2 {
		log.Fatal("insufficient number of arguments")
	} 
	cmd := command{name: args[0], args: args[1:]} 
	err = cmds.run(&currentState, cmd)
	if err != nil {
		log.Fatal("error running command: ", err)
	} 
}


type state struct {
    cfg  *config.Config
} 

