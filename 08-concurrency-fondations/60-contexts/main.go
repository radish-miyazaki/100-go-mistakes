package main

import (
	"context"
	"time"

	"github.com/radish-miyazaki/100-go-mistakes/08-concurrency-fondations/60-contexts/flight"
)

type publisher interface {
	Publish(ctx context.Context, position flight.Position) error
}

type publishHandler struct {
	pub publisher
}

func (h publishHandler) publishPosition(position flight.Position) error {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()
	return h.pub.Publish(ctx, position)
}
