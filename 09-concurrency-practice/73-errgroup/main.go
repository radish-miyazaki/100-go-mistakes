package main

import (
	"context"
	"sync"

	"golang.org/x/sync/errgroup"
)

type (
	Circle struct{}
	Result struct{}
)

func handler1(ctx context.Context, circles []Circle) ([]Result, error) {
	results := make([]Result, len(circles))

	wg := sync.WaitGroup{}
	wg.Add(len(results))

	for i, circle := range circles {
		go func() {
			defer wg.Done()

			result, err := foo(ctx, circle)
			if err != nil {
				// ?
			}
			results[i] = result
		}()
	}
	wg.Wait()

	return results, nil
}

func handler2(ctx context.Context, circles []Circle) ([]Result, error) {
	results := make([]Result, len(circles))

	g, ctx := errgroup.WithContext(ctx)

	for i, circle := range circles {
		// エラーを処理して結果を集計するロジックを新しいゴルーチンで行うために Go を呼び出す
		g.Go(func() error {
			result, err := foo(ctx, circle)
			if err != nil {
				return err
			}
			results[i] = result
			return nil
		})
	}

	// すべてのゴルーチンを待つために Wait を呼び出す
	if err := g.Wait(); err != nil {
		return nil, err
	}
	return results, nil
}

func foo(context.Context, Circle) (Result, error) {
	return Result{}, nil
}
