package main

import (
	"fmt"
	"testing"
)

func Testfetch(t *testing.T) {
	p := &paths{m: make(map[string]int64)}
	path, err := getPath()
	if err != nil {
		t.Fatal(err)
	}
	p.root = path
	if err := p.fetch(); err != nil {
		t.Fatal(err)
	}
	if len(p.m) == 0 {
		t.Fatal(fmt.Errorf("map is empty!"))
	}
}
