package main

import "context"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	newWatch(ctx)
}

func newWatch(ctx context.Context) {
	w := watcher{}
	go w.watch(ctx)
}

type watcher struct{}

func (w watcher) watch(ctx context.Context) {}
