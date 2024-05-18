package main

func main() {
	w := newWatcher()
	defer w.close()
}

func newWatcher() watcher {
	w := watcher{}
	go w.watch()
	return w
}

type watcher struct{}

func (w watcher) watch() {}

func (w watcher) close() {
	// リソースをクローズする
}
