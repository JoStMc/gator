package main

import "fmt"

type command struct {
    name string
	args []string
} 

type commands struct {
	cmdMap map[string]func(*state, command) error
} 

func (c *commands) register(name string, f func(*state, command) error) {
	c.cmdMap[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	fnct := c.cmdMap[cmd.name]
	if fnct == nil {
		return fmt.Errorf("command not found: %s", cmd.name)
	} 
	if err := fnct(s, cmd); err != nil {
	    return err
	} 
	return nil
} 
