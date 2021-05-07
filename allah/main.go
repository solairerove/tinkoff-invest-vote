package main

import (
	"log"

	file "github.com/solairerove/tinkoff-invest-vote/file"
)

func main() {
	for _, user := range file.ReadFromFile() {
		log.Println(user)
	}

	emails := []string{"cikecac613@goqoez.com", "morris.1@mentornkc.com"}
	file.WriteIntoFile(emails)
}
