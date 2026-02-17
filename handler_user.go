package main

import "fmt"

func handlerLogin(s *state, cmd command) error {
	if cmd.args == nil {
		return fmt.Errorf("no args passed")
	} 
	err := s.cfg.SetUser(cmd.args[0])
	if err != nil {
	    return err
	} 
	fmt.Println("User has been set to", cmd.args[0])
	return nil
} 
