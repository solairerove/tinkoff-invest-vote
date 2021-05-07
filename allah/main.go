package main

import (
	"log"

	load "github.com/solairerove/tinkoff-invest-vote/file"
)

func main() {
	log.Println(load.LoadFromFile())
}
