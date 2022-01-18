package helpers

import (
	"path"
	"path/filepath"
	"runtime"
)

func StringLike(pointer *string) string {
	non_pointer := ""
	if pointer != nil {
		non_pointer = *pointer
	}
	return non_pointer
}

func RootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}
