package main

import (
	"blog_agreegator/internal/config"
	"blog_agreegator/internal/database"
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

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
	handler map[string]func(*state, context.Context, command) error
}

func (c *commands) register(name string, f func(*state, context.Context, command) error) {
	c.handler[name] = f
}
func (c *commands) run(s *state, cmd command, ctx context.Context) error {
	if _, ok := c.handler[cmd.name]; !ok {
		return fmt.Errorf("Command doesn't exist")
	}
	return c.handler[cmd.name](s, ctx, cmd)

}

const feed = "https://www.wagslane.dev/index.xml"

func main() {
	config, err := config.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	ctx := context.Background()

	db, err := sql.Open("postgres", config.Url)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	dbQueries := database.New(db)
	s := state{&config, dbQueries}
	commands := commands{make(map[string]func(*state, context.Context, command) error)}
	commands.register("login", func(s *state, ctx context.Context, c command) error {
		return handlerLogin(s, c, ctx)
	})
	commands.register("register", func(s *state, ctx context.Context, c command) error {
		return handlerRegister(s, c, ctx)
	})
	commands.register("reset", func(s *state, ctx context.Context, c command) error {
		return handlerReset(s, c, ctx)
	})
	commands.register("users", func(s *state, ctx context.Context, c command) error {
		return handlerUsers(s, c, ctx)
	})
	commands.register("agg", func(s *state, ctx context.Context, c command) error {
		if len(c.arg) < 1 {
			return fmt.Errorf("Please input time ex. 1s (1 second) 1m (1 minute)")
		}
		time_between_req, err := time.ParseDuration(c.arg[0])
		if err != nil {
			return fmt.Errorf("Please input correct format ex. 1s (1 second) 1m (1 minute)")
		}
		fmt.Println("Collect feeds every", c.arg[0])
		tick := time.NewTicker(time_between_req)
		for ; ; <-tick.C {
			_, err := scrapeFeeds(s, ctx)
			if err != nil {
				log.Print(err)
			}
		}
	})
	commands.register("addfeed", middlewareLogin(handlerAddfeed))
	commands.register("feeds", func(s *state, ctx context.Context, c command) error {
		return handlerFeeds(s, ctx)
	})
	commands.register("follow", middlewareLogin(handlerFollow))
	commands.register("following", middlewareLogin(handlerFollowing))
	commands.register("unfollow", middlewareLogin(handlerUnfollow))
	commands.register("browse", func(s *state, ctx context.Context, c command) error {
		return handlerBrowse(s, ctx, c)
	})
	if len(os.Args) < 2 {
		log.Fatalln("not enough arguments")
	}
	command := command{os.Args[1], os.Args[2:]}
	if err = commands.run(&s, command, ctx); err != nil {
		log.Fatalln(err)
	}
	os.Exit(0)
}
