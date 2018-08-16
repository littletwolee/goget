package main

import (
	"io/ioutil"
	"strings"
)

type paths map[string]int64

func (p paths) fetch(root string, children ...string) error {
	child := func() string {
		if len(children) > 0 {
			return strings.Join(children, "/") + "/"
		} else {
			return ""
		}
	}
	childStr := child()
	rd, err := ioutil.ReadDir(root + childStr)
	for _, fi := range rd {
		if fi.Name()[:1] != "." {
			if fi.IsDir() {
				p.fetch(root, append(children, "/"+fi.Name())...)
			} else {
				p[childStr+fi.Name()] = fi.ModTime().UnixNano()
			}
		}
	}
	return err
}
