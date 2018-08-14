package main

import "fmt"

func main() {
	//go Watch()
	//run()
	outChan := make(chan bool)
	go run(outChan)
	<-outChan
	fmt.Printf("build success!")
}
