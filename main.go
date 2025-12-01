package main

import (
	"fmt"

	"github.com/kahnaisehC/blog_aggregator/internal/config"
)

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
