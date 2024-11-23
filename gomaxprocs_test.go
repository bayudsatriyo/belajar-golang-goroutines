package golanggoroutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)



func TestGetGomaxprocs(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func ()  {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	 totalCpu := runtime.NumCPU()
	 fmt.Println("Total CPU = ", totalCpu)

	//  untuk mengubah thread gunakan sintaks
	// runtime.GOMAXPROCS(10)

	 totalThread := runtime.GOMAXPROCS(-1)
	 fmt.Println("Total Thread = ", totalThread)

	 totalGoroutine := runtime.NumGoroutine()
	 fmt.Println("Total Goroutine = ", totalGoroutine)
	 time.Sleep(5 * time.Second)
	 
	 group.Wait()
}