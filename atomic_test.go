package golanggoroutine

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var x int64 = 0
	var group = sync.WaitGroup{}

	for i := 1; i <= 1000; i++ {
		group.Add(1)
		go func() {
			defer group.Done()
			for j := 1; j <= 100; j++ {
				// jika menggunakan atomic maka tiak perlu lock dan unlock karna sudah otomatis, hanya bisa tipe data primitif saja
				atomic.AddInt64(&x, 1)
			}
		}()
	}

	group.Wait()
	fmt.Println("Counter = ", x)
}