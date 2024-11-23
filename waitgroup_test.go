package golanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsynchrounous(group *sync.WaitGroup)  {
	// untuk menurunkan group
	defer group.Done()

	group.Add(1)

	fmt.Println("Hello")
	time.Sleep(2 * time.Second)
	fmt.Println("World")
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 1; i < 100; i++ {
		go RunAsynchrounous(group)
	}

	group.Wait()

	fmt.Println("Selesai")
}