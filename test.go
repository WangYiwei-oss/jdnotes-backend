package main

import (
	"github.com/fsnotify/fsnotify"
	"log"
)

func main() {
	watcher, err := fsnotify.NewWatcher() //1. 先new
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close() //记得释放
	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
				}
				if event.Op&fsnotify.Write == fsnotify.Create {
					log.Println("create file:", event.Name)
				}
				if event.Op&fsnotify.Write == fsnotify.Remove {
					log.Println("remove file:", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add("D:\\test")
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
