package main

import (
	"errors"
	"fmt"

	"github.com/kahnaisehC/blog_aggregator/internal/config"
)

type state struct {
	Cfg *config.Config
}

type command struct {
	name      string
	arguments []string
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
	err = cfg.SetUser("ian")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", cfg)
}
