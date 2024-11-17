package golanggoroutine

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Ups")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100; i++ {
		// eksekusi goroutine tidak berurutan karna dilakukan secara ascy dan paralel
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}
