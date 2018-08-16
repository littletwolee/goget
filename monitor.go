package main

import (
	"github.com/fsnotify/fsnotify"
	"github.com/littletwolee/commons/logger"
)

type monitor struct {
	build   *b
	cmd     *cmd
	c       chan bool
	watcher *fsnotify.Watcher
	path    string
}

func newMonitor(path string) *monitor {
	m := &monitor{
		build: &b{},
		cmd:   newCmd(),
		c:     make(chan bool),
	}
	w, err := fsnotify.NewWatcher()
	if err != nil {
		logger.Console().Error(err)
		return nil
	}
	m.watcher = w
	m.path = path
	return m
}

// func (m *monitor) run() {
// 	//m.cmd.stop()
// 	m.build.run()
// 	go func() {
// 		if err := m.cmd.run(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()
// }

func (m *monitor) run() {
	// watcher, err := fsnotify.NewWatcher()
	// if err != nil {
	// 	logger.Console().Error(err)
	// }
	// defer watcher.Close()
	done := make(chan bool)

	go func() {
		for {
			select {
			case event := <-m.watcher.Events:
				switch event.Op {
				case fsnotify.Write, fsnotify.Create, fsnotify.Remove:
					m.run()
				}
			case err := <-m.watcher.Errors:
				logger.Console().Error(err)
			}
		}
	}()

	if err := m.watcher.Add(m.path); err != nil {
		logger.Console().Error(err)
	}
	<-done
}
