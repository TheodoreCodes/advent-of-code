package util

import (
	"path/filepath"
	"runtime"
)

func Dirname() string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		panic("getting calling function")
	}

	return filepath.Dir(filename)
}
