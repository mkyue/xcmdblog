package mdnotify

import (
	"github.com/fsnotify/fsnotify"
	"log"
	"path/filepath"
)

type Mdnotify struct {
	path string
}

type WatchFunc interface {
	Create(name string)
	Write(name string)
	Remove(name string)
	Rename(name string)
}

func New(path string) (*Mdnotify, error) {
	s, e := filepath.Abs(path)
	if e != nil {
		return nil, e
	}
	mdn := Mdnotify{s}

	return &mdn, nil
}

func (m *Mdnotify) Watch(wf WatchFunc) {
	watcher, e := fsnotify.NewWatcher()
	if e != nil {
		log.Fatal(e)
	}
	defer watcher.Close()

	done := make(chan bool)

	go func() {
		for {
			select {
			case events, ok := <-watcher.Events:
				if !ok {
					return
				}

				switch events.Op.String() {
				case "CREATE":
					wf.Create(events.Name)
				case "WRITE":
					wf.Write(events.Name)
				case "REMOVE":
					wf.Remove(events.Name)
				case "RENAME":
					wf.Rename(events.Name)
				default:
					log.Printf("op:%s,  name:%s\n", events.Op.String(), events.Name)
				}

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err := watcher.Add(m.path)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("watch dir :" + m.path)
	<-done
}
