package main

import "fmt"

func channelSampleOne() {
	ch := make(chan int)
	wg.Add(2)
	go func() {
		i := <-ch
		fmt.Println("Received value: ", i)
		wg.Done()
	}()

	go func() {
		i := 500
		ch <- i
		i = 1000 //DOES NOT AFFECT SENT DATA
		wg.Done()

	}()
	wg.Wait()
}

func main() {
	println("__CHANNEL__")
	channelSampleOne()
}
