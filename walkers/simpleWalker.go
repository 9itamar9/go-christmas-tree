package walkers

import (
	"os"
	"path/filepath"
	"strings"
)

type SimpleWalker struct {
	filepath.WalkFunc
	MiddlePrefix string
	EndPredix    string
	Indentation  string
}

func (sw SimpleWalker) Walk(path string, level int) string {
	var files []string
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	return strings.Join(files, ",\n")
}
