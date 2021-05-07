package main

import (
	"log"

	"github.com/gocolly/colly"
	load "github.com/solairerove/tinkoff-invest-vote/file"
)

func main() {
	log.Println("trying to registrate new emails")

	// create a new collector
	c := colly.NewCollector()

	for _, user := range load.ReadFromFile() {
		email := user.Email
		password := user.Password

		err := c.Post("https://invest-terminal.useresponse.com/registration",
			map[string]string{
				"email":                 email,
				"full_name":             "morris",
				"password":              password,
				"password_confirmation": password,
				"submit":                "Регистрация"})

		if err != nil {
			log.Fatal(err)
		}

		// attach callbacks after vote
		c.OnResponse(func(r *colly.Response) {
			log.Println("response received", r.StatusCode)
		})
	}
}
