package main

import (
	"blog_agreegator/internal/config"
	"blog_agreegator/internal/database"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type state struct {
	config *config.Config
	db     *database.Queries
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
	db, err := sql.Open("postgres", config.Url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	dbQueries := database.New(db)
	s := state{&config, dbQueries}
	commands := commands{make(map[string]func(*state, command) error)}
	commands.register("login", func(s *state, c command) error {
		return handlerLogin(s, c)
	})
	commands.register("register", func(s *state, c command) error {
		return handlerRegister(s, c)
	})
	commands.register("reset", func(s *state, c command) error {
		return handlerReset(s, c)
	})
	commands.register("users", func(s *state, c command) error {
		return handlerUsers(s, c)
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
