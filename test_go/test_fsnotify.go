package main

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
)

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	defer watcher.Close()

	err = watcher.Add("test_fsnotify.go")
	if err != nil {
		panic(err)
	}

	for {
		select {
		case evt := <-watcher.Events:
			fmt.Printf("%#v\n", evt.Name)

		case err := <-watcher.Errors:
			fmt.Println(err)
		}
	}
}
