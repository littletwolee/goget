package main

import (
	"fmt"
)

func main() {
	// m := newMonitor()
	// m.run()
	// m.monitoring()
	p := make(paths)
	path, err := getPath()
	if err != nil {
		fmt.Println(err)
	}
	if err := p.fetch(path); err != nil {
		fmt.Println(err)
	}
	for k, v := range p {
		fmt.Printf("%s : %d \n", k, v)
	}
}
