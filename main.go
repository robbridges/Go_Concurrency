package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
    whatWins()
    
}

func writeAndAppend() {
	var wg sync.WaitGroup
	wg.Add(2)
	example := []int{1,2,3,4,5,6}
	numChan := make(chan int)
	var res []int
	go func() {
		defer wg.Done()
		SendNumToChan(example, numChan)
	}()
	go func() {
		defer wg.Done()
		readFromChannel(&res, numChan)
	}()
	wg.Wait()
	fmt.Println(res)

}

func SendNumToChan(nums []int, c chan int) {
    for _, num := range nums {
        if num % 2 == 0 {
			c <- num
		}
    }
	// this is fine because we can close a channel even if all data hasn't been read from it. read from Channel will still read value from a closed channel
	close(c)
}

func readFromChannel(nums *[]int, c chan int) {

    for num := range c {
        *nums = append(*nums, num)
    }
	
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

func whatWins() {
    var wg sync.WaitGroup
    numChan := make(chan int, 30) // Increase the buffer size to prevent blocking
    nums := []int{}
    wg.Add(3)
    go func() {
        for i := 0; i < 10; i++ {
            numChan <- 1
            time.Sleep(time.Millisecond) // Add a small delay
        }
        wg.Done()
    }()
    go func() {
        for i := 0; i < 10; i++ {
            numChan <- 2
            time.Sleep(time.Millisecond) // Add a small delay
        }
        wg.Done()
    }()
    go func() {
        for i := 0; i < 10; i++ {
            numChan <- 3
            time.Sleep(time.Millisecond) // Add a small delay
        }
        wg.Done()
    }()
    go func() {
        wg.Wait() // Wait for the other goroutines to finish
        close(numChan) // Then close the channel
    }()
    for num := range numChan { // Read from the channel until it's closed
        nums = append(nums, num)
    }
    fmt.Println(nums)
}