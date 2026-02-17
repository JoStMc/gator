package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/JoStMc/gator/internal/config"
	"github.com/JoStMc/gator/internal/database"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal("error reading config:", err)
	} 

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatal("error loading database:", err)
	} 

	currentState := state{db: database.New(db), cfg: &cfg} 

	cmds := commands{cmdMap: make(map[string]func(*state, command) error)}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("users", handlerList)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", handlerAddFeed)
	cmds.register("reset", handlerReset)

	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatal("insufficient number of arguments")
	} 
	cmd := command{name: args[0], args: args[1:]} 
	err = cmds.run(&currentState, cmd)
	if err != nil {
		log.Fatal("error running command:", err)
	} 
}


type state struct {
	db *database.Queries
    cfg  *config.Config
} 

