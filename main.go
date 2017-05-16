package main

import (
	"log"

	"github.com/justinfinch/gauge/cmd"
)

func main() {
	if err := cmd.NewRootCommand().Execute(); err != nil {
		log.Fatal(err)
	}
}
