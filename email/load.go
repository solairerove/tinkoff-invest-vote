package email

import (
	"io/ioutil"
	"strings"
)

type User struct {
	email, password string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func LoadFromFile() []User {
	dat, err := ioutil.ReadFile("emails.txt")
	check(err)
	// fmt.Println(string(dat))

	emails := strings.Split(string(dat), "\n")
	// fmt.Println(emails)

	users := []User{}
	for _, email := range emails {
		users = append(users, User{email, strings.Split(email, "@")[0]})
	}

	return users
}
