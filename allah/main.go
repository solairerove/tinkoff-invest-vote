package main

import (
	"log"

	load "github.com/solairerove/tinkoff-invest-vote/email"
)

func main() {
	log.Println(load.LoadFromFile())
}
