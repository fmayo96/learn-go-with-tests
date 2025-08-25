package syncch

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("increasing the counter 100 times sequentially lives it at 100", func(t *testing.T) {
		counter := NewCounter()
		want := 100
		for range want {
			counter.Inc()
		}
		assertCounter(t, counter, want)
	})

	t.Run("increasing the counter 100 times concurrently lives it at 1000", func(t *testing.T) {
		counter := NewCounter()
		var wg sync.WaitGroup
		want := 1000
		wg.Add(want)
		for range want {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()
		assertCounter(t, counter, want)
	})
}

func assertCounter(t testing.TB, got *Counter, want int) {
	if got.Value() != want {
		t.Errorf("got %d, wanted %d", got.Value(), want)
	}
}
