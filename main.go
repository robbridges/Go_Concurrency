package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
    selectChanExample()
    
}

func writeAndAppend() {
	var wg sync.WaitGroup
    c := make(chan int)
    myInts := []int{1,2,3,4}
    res := []int{}
    wg.Add(2)
    go func(){
        SendNumToChan(myInts, c)
        wg.Done()
    }()
    go func(){
        res = readFromChannel(res, c)
        wg.Done()
    }()
    wg.Wait() // Wait for all goroutines to finish
    fmt.Println(res)
}

func SendNumToChan(nums []int, c chan int) {
    for _, num := range nums {
        c <- num
    }
	// this is fine because we can close a channel even if all data hasn't been read from it. read from Channel will still read value from a closed channel
	close(c)
}

func readFromChannel(nums []int, c chan int) []int {

    for num := range c {
        nums = append(nums, num)
    }
	return nums
}

func countConcurrently() {
	var wg sync.WaitGroup
	c := make(chan string)
	wg.Add(2)
	go func() {
		count("sheep", c)
		wg.Done()
	}()
	go func() {
		count("fish", c)
		wg.Done()
	}()
	go func() {
		wg.Wait()
		close(c)
	}()
	for value := range c {
		fmt.Println(value)
	}
}

func count (thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}
	
}

func selectChanExample() {
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		for {
			c1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()
	go func() {
		for {
			c2 <- "Every 2 seconds"
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		select {
		case msg1 := <- c1:
			fmt.Println(msg1)
		case msg2 := <- c2:
			fmt.Println(msg2)
		}
		
	}
}