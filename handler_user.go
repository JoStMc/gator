package main

import (
	"context"
	"fmt"
	"time"

	"github.com/JoStMc/gator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if cmd.args == nil {
		return fmt.Errorf("no args passed")
	} 

	_, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("couldn't find user: %v", err)
	} 

	err = s.cfg.SetUser(cmd.args[0])
	if err != nil {
	    return err
	} 

	fmt.Println("User has been set to", cmd.args[0])
	return nil
} 

func handlerRegister(s *state, cmd command) error {
	ctx := context.Background()
    if cmd.args == nil {
        return fmt.Errorf("no args passed")
    } 

	_, err := s.db.GetUser(ctx, cmd.args[0])
	if err == nil {
		return fmt.Errorf("user already exists")
	} 

	currentTime := time.Now()
	params := database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		Name: cmd.args[0],
	} 

	_, err = s.db.CreateUser(ctx, params)
	if err != nil {
		return fmt.Errorf("unable to create user: %v", err)
	} 
	err = s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return fmt.Errorf("unable to set user: %v", err)
	} 
	fmt.Println("User has been created!")
	return nil
} 
