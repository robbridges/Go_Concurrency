package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
    countConcurrently()
    
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