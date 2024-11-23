package golanggoroutine

import (
	"fmt"
	"testing"
	"time"
)


func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
	}()

	for i := range ticker.C {
		fmt.Println("Tick ke ", i)
	}

}

func TestTick(t *testing.T) {
	ticker := time.Tick(1 * time.Second)


	for i := range ticker {
		fmt.Println("Tick ke ", i)
	}

}