package pool

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	handle := func(arg interface{}) interface{} {
		a, ok := arg.(int)
		if !ok {
			return nil
		}
		time.Sleep(time.Second)
		return a * 2
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	in := make(chan interface{}, 1e1)
	go func() {
		defer close(in)

		for i := 0; i <= 20; i++ {
			in <- i
		}
	}()

	for r := range Run(ctx, in, handle) {
		res, ok := r.(int)
		if !ok {
			continue
		}
		fmt.Println(res)
	}
}
