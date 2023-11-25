package main

import (
	// "concurrency/channels"
	"concurrency/channels"
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
    res := channels.WhoWins()
    fmt.Println(res)
}

func whatWins() {
    // multiple go routines writing to the same channel needs to be closed in a third go routine
    var wg sync.WaitGroup
    numChan := make(chan int) // Increase the buffer size to prevent blocking
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



