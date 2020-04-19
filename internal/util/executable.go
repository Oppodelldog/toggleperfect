package util

import (
	"os"
	"path/filepath"
)

func GetExecutableDir() string {
	execDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	return execDir
}
