package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/kahnaisehC/blog_aggregator/internal/config"
)

type state struct {
	Cfg *config.Config
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

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return errors.New("the login command expects a username as an arguments")
	}
	username := cmd.arguments[0]

	if err := s.Cfg.SetUser(username); err != nil {
		return err
	}

	fmt.Println("The username has been changed successfully!")

	return nil
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		panic(err)
	}
	gatorState := state{}
	gatorState.Cfg = &cfg

	cmds := commands{}
	cmds.cmdMap = make(map[string]func(s *state, cmd command) error)
	cmds.register("login", handlerLogin)

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
