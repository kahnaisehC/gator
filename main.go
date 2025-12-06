package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/kahnaisehc/blog_aggregator/internal/config"
	"github.com/kahnaisehc/blog_aggregator/internal/database"
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

func handlerRegister(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return errors.New("the register command expects a name as an argument")
	}
	name := cmd.arguments[0]

	userParams := database.CreateUserParams{
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		ID:        uuid.New(),
	}

	_, err := s.db.CreateUser(context.Background(), userParams)
	if err != nil {
		return err
	}
	s.cfg.SetUser(name)
	fmt.Println(name + " Has been registered successfully!!")

	return nil
}

func handlerReset(s *state, cmd command) error {
	err := s.db.ResetUsers(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, u := range users {
		fmt.Print(u.Name)
		if u.Name == s.cfg.CurrentUserName {
			fmt.Print(" (current)")
		}
		fmt.Println()
	}
	return nil
}

func handlerAgg(s *state, cmd command) error {
	feedURL := "https://www.wagslane.dev/index.xml"
	feed, err := fetchFeed(context.Background(), feedURL)
	if err != nil {
		return err
	}
	fmt.Println(feed)
	return nil
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
