package file

import (
	"io/ioutil"
	"strings"
)

func WriteIntoFile(emails []string) {
	err := ioutil.WriteFile(file, []byte(strings.Join(emails, "\n")), 0644)
	Check(err)
}
