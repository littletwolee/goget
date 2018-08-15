package main

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/littletwolee/commons/logger"
)

type monitor struct {
	build *b
	cmd   *cmd
	c     chan bool
}

func newMonitor() *monitor {
	return &monitor{
		build: &b{},
		cmd:   newCmd(),
		c:     make(chan bool),
	}
}

func (m *monitor) run() {
	//m.cmd.stop()
	m.build.run()
	go func() {
		if err := m.cmd.run(); err != nil {
			fmt.Println(err)
		}
	}()
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
			case event := <-watcher.Events:
				switch event.Op {
				case fsnotify.Write, fsnotify.Create, fsnotify.Remove:
					m.run()
				}
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
