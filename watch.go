package main

// import (
// 	"github.com/go-fsnotify/fsnotify"
// 	"github.com/littletwolee/commons"
// )

// func Watch() {
// 	output, _ := Build()
// 	if output != "" {
// 		Get(output)
// 	}
// 	watcher, err := fsnotify.NewWatcher()
// 	if err != nil {
// 		commons.GetLogger().OutErr(err)
// 	}
// 	defer watcher.Close()
// 	done := make(chan bool)
// 	go func() {
// 		for {
// 			select {
// 			case _ = <-watcher.Events:
// 				output, _ := Build()
// 				if output != "" {
// 					Get(output)
// 				}
// 			case err := <-watcher.Errors:
// 				commons.GetLogger().OutErr(err)
// 			}
// 		}
// 	}()

// 	err = watcher.Add("slog")
// 	if err != nil {
// 		commons.GetLogger().OutErr(err)
// 	}
// 	<-done

// 	if err != nil {
// 		commons.GetLogger().OutErr(err)
// 	}
// }
