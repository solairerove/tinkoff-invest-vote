package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("hello from read")

	dat, err := ioutil.ReadFile("emails.txt")
	check(err)
	// fmt.Println(string(dat))

	emails := strings.Split(string(dat), "\n")
	// fmt.Println(emails)

	for _, email := range emails {
		fmt.Println("email - " + email)
		fmt.Println("password - " + strings.Split(email, "@")[0])
	}
}
