package main

import (
	"database/sql"
	"errors"
	"log"
	"os"

	"github.com/kahnaisehC/blog_aggregator/internal/config"
	"github.com/kahnaisehC/blog_aggregator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

type command struct {
	name      string
	arguments []string
}

type commands struct {
	cmdMap map[string]func(s *state, cmd command) error
}

func (cmds *commands) run(s *state, cmd command) error {
	f, ok := cmds.cmdMap[cmd.name]
	if !ok {
		return errors.New(cmd.name + " is not a registered command")
	}

	return f(s, cmd)
}

func (cmds *commands) register(name string, f func(*state, command) error) {
	cmds.cmdMap[name] = f
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		panic(err)
	}

	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	gatorState := state{}
	gatorState.cfg = &cfg
	gatorState.db = database.New(db)

	cmds := commands{}
	cmds.cmdMap = make(map[string]func(s *state, cmd command) error)
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", handlerAddfeed)
	cmds.register("feeds", handlerFeeds)

	args := os.Args
	if len(args) < 2 {
		log.Println("not enough arguments")
		os.Exit(1)
	}
	cmd := command{}
	cmd.name = args[1]
	cmd.arguments = args[2:]

	if err = cmds.run(&gatorState, cmd); err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}
