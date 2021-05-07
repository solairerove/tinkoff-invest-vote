package file

import (
	"io/ioutil"
	"strings"
)

func ReadFromFile() []User {
	dat, err := ioutil.ReadFile(file)
	Check(err)

	emails := strings.Split(string(dat), "\n")

	users := []User{}
	for _, email := range emails {
		users = append(users, User{email, strings.Split(email, "@")[0]})
	}

	return users
}
