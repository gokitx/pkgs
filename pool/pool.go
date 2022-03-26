package pool

import (
	"context"
	"log"
	"runtime"
	"sync"

	"golang.org/x/sync/semaphore"
)

var (
	sem = semaphore.NewWeighted(int64(runtime.GOMAXPROCS(0)))
)

type Handler func(interface{}) interface{}

func Run(ctx context.Context, in chan interface{}, handler Handler) (out chan interface{}) {
	var wg sync.WaitGroup
	out = make(chan interface{}, 1e2)

	go func() {
		defer func() {
			wg.Wait()
			close(out)
		}()

	Loop:
		for {
			select {
			case <-ctx.Done():
				break Loop
			case task, ok := <-in:
				if !ok {
					break Loop
				}
				if err := sem.Acquire(ctx, 1); err != nil {
					log.Printf("[pool] failed to acquire semaphore: %v \n", err)
					break Loop
				}

				wg.Add(1)
				go func(arg interface{}) {
					defer func() {
						sem.Release(1)
						wg.Done()
					}()

					out <- handler(arg)
				}(task)
			}
		}
	}()

	return
}
