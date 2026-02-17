package main

import (
	"fmt"
	"log"

	"github.com/JoStMc/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal("error reading config: ", err)
	} 
	fmt.Println(cfg)

	err = cfg.SetUser("me")
	if err != nil {
		log.Fatal("error setting user: ", err)
	} 

	cfg, err = config.Read()
	if err != nil {
		log.Fatal("error reading config: ", err)
	} 
	fmt.Println(cfg)
}
