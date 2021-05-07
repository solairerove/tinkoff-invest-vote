package main

import (
	"log"

	"github.com/gocolly/colly"
	load "github.com/solairerove/tinkoff-invest-vote/file"
)

func main() {

	// create a new collector
	c := colly.NewCollector()

	// registrate
	// log.Println("trying to registrate")
	// err := c.Post("https://invest-terminal.useresponse.com/registration",
	// 	map[string]string{
	// 		"email":                 email,
	// 		"full_name":             "morris",
	// 		"password":              password,
	// 		"password_confirmation": password,
	// 		"submit":                "Регистрация"})

	// if err != nil {
	// 	log.Fatal(err)
	// }

	for _, user := range load.LoadFromFile() {
		email := user.Email
		password := user.Password

		// login
		log.Println("trying to login")
		err := c.Post("https://invest-terminal.useresponse.com/login", map[string]string{"email": email, "password": password, "submit": "Войти"})
		if err != nil {
			log.Fatal(err)
		}

		// vote add
		log.Println("trying to add vote")
		err = c.Post("https://invest-terminal.useresponse.com/topic/1921/vote/add/1", map[string]string{"format": "json", "full": "true"})
		if err != nil {
			log.Fatal(err)
		}

		// attach callbacks after vote
		c.OnResponse(func(r *colly.Response) {
			log.Println("response received", r.StatusCode)
		})
	}
}
