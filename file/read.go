package file

import (
	"io/ioutil"
	"path/filepath"
	"runtime"
	"strings"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFromFile() []User {
	dat, err := ioutil.ReadFile(basepath + "/emails_new.txt")
	check(err)

	emails := strings.Split(string(dat), "\n")

	users := []User{}
	for _, email := range emails {
		users = append(users, User{email, strings.Split(email, "@")[0]})
	}

	return users
}
