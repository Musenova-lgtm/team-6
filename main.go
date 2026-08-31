package main

import (
	"fmt"
	"os"

	"./cli"
	"./game"
)

func main() {
	config, err := cli.Parse(os.Args[1:])

	if err != nil {
		fmt.Println("Error:", err)
		fmt.Print(cli.Usage())
		os.Exit(1)
	}

	if config.Help {
		fmt.Print(cli.Usage())
		return
	}

	game.Run(config)
}
