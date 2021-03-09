package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

type Config struct {
	a []int
}

func TestTrace(t *testing.T) {
	var v atomic.Value

	// writer
	go func() {
		var i int
		for {
			i++
			cfg := &Config{
				a: []int{i, i + 1, i + 2, i + 3, i + 4, i + 5},
			}
			v.Store(cfg)
		}
	}()

	// reader
	var wg sync.WaitGroup
	for n := 0; n < 4; n++ {
		wg.Add(1)
		go func() {
			for n := 0; n < 100; n++ {
				cfg := v.Load()
				t.Logf("%#v\n", cfg)
			}
			wg.Done()
		}()
	}

	wg.Wait()
}
