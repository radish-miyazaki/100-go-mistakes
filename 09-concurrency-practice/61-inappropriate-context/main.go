package main

import (
	"context"
	"net/http"
	"time"
)

func handler1(w http.ResponseWriter, r *http.Request) {
	resp, err := doSomeTask(r.Context(), r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	go func() {
		err := publish(r.Context(), resp)
		_ = err
	}()

	writeResponse(resp)
}

func handler2(w http.ResponseWriter, r *http.Request) {
	resp, err := doSomeTask(r.Context(), r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	go func() {
		err := publish(context.WithoutCancel(r.Context()), resp)
		_ = err
	}()

	writeResponse(resp)
}

type detach struct {
	ctx context.Context
}

func (d detach) Deadline() (time.Time, bool) {
	return time.Time{}, false
}

func (d detach) Done() <-chan struct{} {
	return nil
}

func (d detach) Err() error {
	return nil
}

func (d detach) Value(key any) any {
	return d.ctx.Value(key)
}

func doSomeTask(ctx context.Context, r *http.Request) (string, error) {
	return "", nil
}

func publish(context.Context, string) error {
	return nil
}

func writeResponse(string) {}
