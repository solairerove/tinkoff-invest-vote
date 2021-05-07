package file

import (
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)
	Basepath   = filepath.Dir(b)
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
