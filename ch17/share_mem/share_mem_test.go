package share_mem

import (
	"sync"
	"testing"
	"time"
)

const num = 50

func TestCounter(t *testing.T) {
	counter := 0
	for i := 0; i < num; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

func TestCounterThreadSafe(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	for i := 0; i < num; i++ {
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
		time.Sleep(1 * time.Second)
		t.Logf("counter = %d", counter)
	}
}

func TestCounterWaitGroup(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < num; i++ {
		wg.Add(i)
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done()
		}()

	}
	wg.Wait()
	t.Logf("conter = %d", counter)
}
