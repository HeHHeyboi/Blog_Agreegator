package main

import (
	"blog_agreegator/internal/config"
	"fmt"
	"os"
)

type state struct {
	config *config.Config
}
type command struct {
	name string
	arg  []string
}
type commands struct {
	handler map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.handler[name] = f
}
func (c *commands) run(s *state, cmd command) error {
	if _, ok := c.handler[cmd.name]; !ok {
		return fmt.Errorf("Command doesn't exist")
	}
	return c.handler[cmd.name](s, cmd)

}

func main() {
	config, err := config.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	s := state{&config}
	commands := commands{make(map[string]func(*state, command) error)}
	commands.register("login", func(s *state, c command) error {
		return handlerLogin(s, c)
	})
	if len(os.Args) < 2 {
		fmt.Println("not enough arguments")
		os.Exit(1)
	}
	command := command{os.Args[1], os.Args[2:]}
	if err = commands.run(&s, command); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arg) < 1 {
		return fmt.Errorf("Username is require")
	}

	name := cmd.arg[0]
	err := s.config.SetUser(name)
	if err != nil {
		return err
	}
	fmt.Println("User has been set")
	return nil
}
