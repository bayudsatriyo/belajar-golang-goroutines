package golanggoroutine

import (
	"fmt"
	"sync"
	"testing"
)

var counter = 0

func OnlyOnce() {
	fmt.Println("Halo")
	counter++
}

func TestOnce(t *testing.T) {
	once := &sync.Once{}
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		// group add harusnya dipanggil sebelum gorountine
		group.Add(1)
		go func ()  {
			once.Do(OnlyOnce)
			group.Done()	
		}()
	}

	group.Wait()

	fmt.Println("Counter = ", counter)
}