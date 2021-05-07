package main

import (
	"log"

	load "github.com/solairerove/tinkoff-invest-vote/file"
)

func main() {
	for _, user := range load.ReadFromFile() {
		log.Println(user)
	}
}
