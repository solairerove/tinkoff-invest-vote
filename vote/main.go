package main

import (
	"log"

	"github.com/gocolly/colly"
	load "github.com/solairerove/tinkoff-invest-vote/file"
)

func main() {

	for _, user := range load.ReadFromFile() {
		email := user.Email
		password := user.Password

		// create a new collector
		c := colly.NewCollector()

		// login
		log.Println("trying to login")
		err := c.Post("https://invest-terminal.useresponse.com/login",
			map[string]string{"email": email, "password": password, "submit": "Войти"})

		if err != nil {
			log.Fatal(err)
		}

		// vote add
		log.Println("trying to add vote")
		err = c.Post("https://invest-terminal.useresponse.com/topic/1921/vote/add/1",
			map[string]string{"format": "json", "full": "true"})

		if err != nil {
			log.Fatal(err)
		}

		// attach callbacks after vote
		c.OnResponse(func(r *colly.Response) {
			log.Println("response received", r.StatusCode)
		})
	}
}
