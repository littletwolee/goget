package main

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/littletwolee/commons/logger"
)

type monitor struct {
	build *b
	cmd   *cmd
}

func newMonitor() *monitor {
	return &monitor{
		build: &b{},
		cmd:   newCmd(),
	}
}

func (m *monitor) run() {
	fmt.Println(1)
	m.cmd.stop()
	fmt.Println(2)
	m.build.run()
	fmt.Println(3)
	go m.cmd.run()
	fmt.Println(4)
}
func (m *monitor) monitoring() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		logger.Console().Error(err)
	}
	defer watcher.Close()
	done := make(chan bool)

	go func() {
		for {
			select {
			case _ = <-watcher.Events:
				m.run()
			case err := <-watcher.Errors:
				logger.Console().Error(err)
			}
		}
	}()

	if err = watcher.Add(m.cmd.p); err != nil {
		logger.Console().Error(err)
	}
	<-done
}
