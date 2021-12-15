package common

import (
	"github.com/WangYiwei-oss/jdframe/src/jdft"
	"github.com/fsnotify/fsnotify"
	"log"
	"sync"
)

var fdnOnce sync.Once

type fdNotify struct {
	watcher *fsnotify.Watcher
	stopch  chan struct{} //用于停止监控
}

var FdNotify *fdNotify

func GetFdNotify() *fdNotify {
	fdnOnce.Do(func() {
		wtc, err := fsnotify.NewWatcher()
		if err != nil {
			jdft.GetLogger("Global").Error("failed to init fdnotify")
		}
		FdNotify = &fdNotify{
			watcher: wtc,
			stopch:  make(chan struct{}),
		}
	})
	return FdNotify
}

func (f *fdNotify) Mount(filepaths ...string) *fdNotify {
	for _, filepath := range filepaths {
		err := f.watcher.Add(filepath)
		if err != nil {
			jdft.GetLogger("Global").Warnf("failed to add %s to fdnotify", filepath)
		}
	}
	return f
}

func (f *fdNotify) Stop() {
	f.stopch <- struct{}{}
	err := f.watcher.Close()
	if err != nil {
		jdft.GetLogger("Global").Error("failed to close fdnotify")
	}
}

func (f *fdNotify) Start() {
	go func() {
		for {
			select {
			case event, ok := <-f.watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Create == fsnotify.Create { //创建
					log.Println("create file:", event.Name)
				}
				if event.Op&fsnotify.Write == fsnotify.Write { //写入
					log.Println("modified file:", event.Name)
				}
				if event.Op&fsnotify.Remove == fsnotify.Remove { //移除
					log.Println("remove file:", event.Name)
				}
				if event.Op&fsnotify.Rename == fsnotify.Rename { //移除
					log.Println("rename file:", event.Name)
				}
				if event.Op&fsnotify.Chmod == fsnotify.Chmod { //移除
					log.Println("chmod file:", event.Name)
				}
			case err, ok := <-f.watcher.Errors:
				if !ok {
					return
				}
				jdft.GetLogger("Global").Warn("fdnotify err:", err)
			}
		}
	}()
}
