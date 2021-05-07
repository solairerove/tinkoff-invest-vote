package file

import (
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
	file       = basepath + "/emails_new.txt"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
