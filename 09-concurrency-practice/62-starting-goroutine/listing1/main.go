package main

func main() {
	newWatch()
}

func newWatch() {
	w := watcher{}
	go w.watch()
}

type watcher struct{}

func (w watcher) watch() {}
