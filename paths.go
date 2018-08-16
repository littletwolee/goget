package main

import (
	"io/ioutil"
)

type paths struct {
	m           map[string]int64
	root, child string
}

func (p *paths) fetch() error {
	rd, err := ioutil.ReadDir(p.root + p.child)
	for _, fi := range rd {
		fName := fi.Name()
		if fName[:1] != "." {
			if fi.IsDir() {
				p.child += fName + "/"
				p.fetch()
			} else {
				p.m[p.child+fName] = fi.ModTime().UnixNano()
			}
		}
	}
	return err
}
