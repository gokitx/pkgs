package limiter

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestLimiter_Allow(t *testing.T) {
	l := New(3e3)
	for i := 0; i < 1e4; i++ {
		l.Allow()
		fmt.Println("left", l.Left())
		go func(num int) {
			fmt.Println("i am", num)
			time.Sleep(time.Millisecond * time.Duration(1000+rand.Intn(2000)))
			l.Done()
		}(i)
	}
	l.Wait()
	t.Log(l.Left(), l.Count())
}
