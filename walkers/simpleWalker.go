package walkers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
)

type SimpleWalker struct {
	MiddlePrefix string
	EndPredix    string
	Indentation  string
}

var b bytes.Buffer

func (sw SimpleWalker) Walk(path string, level int) (string, error) {
	if err := sw.recWalk(path, 0); err != nil {
		return "", err
	}
	return b.String(), nil
}

func (sw SimpleWalker) recWalk(path string, level int) error {
	fileInfo, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	if len(fileInfo) == 0 {
		return nil
	}

	for i, entry := range fileInfo {
		if i == len(fileInfo)-1 {
			b.WriteString(strings.Repeat(sw.Indentation, level) + sw.EndPredix)
		} else {
			b.WriteString(strings.Repeat(sw.Indentation, level) + sw.MiddlePrefix)
		}
		b.WriteString(entry.Name() + "\n")
		if entry.IsDir() {
			if err := sw.recWalk(fmt.Sprintf("%s/%s", path, entry.Name()), level+1); err != nil {
				return err
			}
		}
	}
	return nil
}
