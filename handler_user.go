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

func handlerList(s *state, cmd command) error {
	users, err := s.db.ListUsers(context.Background())
	if err != nil {
		return fmt.Errorf("unable to list users: %v", err)
	} 

	if len(users) == 0 {
	    fmt.Println("No users found.")
		return nil
	} 

	for _, name := range users {
		if name == s.cfg.CurrentUserName {
		    fmt.Println("* ", name, "(current)")
			continue
		} 
		fmt.Println("* ", name)
	}

	return nil
} 


func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		u, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return err
		}
		return handler(s, cmd, u)
	} 
} 
