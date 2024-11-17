package golanggoroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)

	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Halo"
		fmt.Println("Selesai mengirim data ke channel")
	}()

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Halo"
	fmt.Println("Selesai mengirim data ke channel")
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)

	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Halo"
	fmt.Println("Selesai mengirim data ke channel")
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)

	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	data := <-channel

	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

func TestBufferChannel(t *testing.T) {
	channel := make(chan string, 1)
	defer close(channel)

	go func() {
		channel <- "Halo"
		channel <- "Kakak"
		channel <- "bayu"

		fmt.Println("Selesai mengirim data ke channel")
	}()

	go func() {
		// jika buffer channel sudah penuh maka data akan diambil secara asyntonic
		// jika buffer penerima harnya 1 sedangkan channel diisi oleh 3 maka akan deadlock karna antrian buffer hanya menerima 1
		fmt.Println(<-channel)

		// tambahkan penerima lagi minimal 1
	}()

	time.Sleep(5 * time.Second)
}

func TestForRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println(data)
	}

	fmt.Println("Selesai")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Menerima data dari channel 1", data)
			counter++

		case data := <-channel2:
			fmt.Println("Menerima data dari channel 2", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}
}

func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Menerima data dari channel 1", data)
			counter++

		case data := <-channel2:
			fmt.Println("Menerima data dari channel 2", data)
			counter++
		default:
			fmt.Println("menunggu data")
		}

		if counter == 2 {
			break
		}
	}
}
