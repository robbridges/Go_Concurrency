package main

import (
	"fmt"
	"sync"
	"time"
)


func main() {

    // c1 := make(chan string)
    // c2 := make(chan string)

    // go func() {
    //     time.Sleep(1 * time.Second)
    //     c1 <- "one"
    // }()
    // go func() {
    //     time.Sleep(2 * time.Second)
    //     c2 <- "two"
    // }()

    // for i := 0; i < 2; i++ {
    //     select {
    //     case msg1 := <-c1:
    //         fmt.Println("received", msg1)
    //     case msg2 := <-c2:
    //         fmt.Println("received", msg2)
    //     }
    // }
   channelSelect()
}

func writeAndAppend() {
	var wg sync.WaitGroup
    wg.Add(2)
    nums := []int{1,2,3,4,5,6}
    numChan := make(chan int)
    var res []int
    go func(){
        defer wg.Done()
        SendNumToChan(nums, numChan)
    }()
    go func(){
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
    numChan := make(chan int, 20) // Increase the buffer size to prevent blocking
    nums := []int{}
    wg.Add(2)
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
        wg.Wait() // Wait for the other goroutines to finish
        close(numChan) // Then close the channel
    }()
    for num := range numChan { // Read from the channel until it's closed
        nums = append(nums, num)
    }
    fmt.Println(nums)
}

func appendToSlices(slice1, slice2 []int) []int {
	var newSlice []int
	newSlice = append(slice1, slice2...)
	return newSlice
}

func channelSelect() {
    numChan := make(chan int)
    stringChan := make(chan string)
    quit := make(chan bool)
    nums := []int{1,2,3,4}
    words := []string{"yam", "yam", "yo"}

    go func() {
        for _, num := range nums {
            numChan <- num
        }
        quit <- true
    }()

    go func() {
        for _, word := range words {
            stringChan <- word
        }
        quit <- true
    }()

    completed := 0
    for completed < 2 {
        select {
        case chanNum := <- numChan:
            fmt.Println(chanNum)
        case chanString := <- stringChan:
            fmt.Println(chanString)
        case <- quit:
            completed++
        }
    }
}


