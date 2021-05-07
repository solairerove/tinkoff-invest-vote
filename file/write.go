package file

import (
	"io/ioutil"
	"strings"
)

func WriteIntoFile(e *[]User) {
	dat, err := ioutil.ReadFile(Basepath + "/emails_new.txt")
	Check(err)

	emails := strings.Split(string(dat), "\n")

	users := []User{}
	for _, email := range emails {
		users = append(users, User{email, strings.Split(email, "@")[0]})
	}
}
