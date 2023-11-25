package channels

import "fmt"

func WhoWins() string {
    numChan := make(chan int)
    stringChan := make(chan string)
    stringComplete := make(chan bool)
    numComplete := make(chan bool)
    nums := []int{1,2,3,4}
    var winner string
    words := []string{"a", "b", "c", "d"}

    go func() {
        for _, num := range nums {
            numChan <- num
        }
        numComplete <- true
    }()

    go func() {
        for _, word := range words {
            stringChan <- word
        }
        stringComplete <- true
    }()

    done := false
    for !done  {
        select {
        case chanNum := <- numChan:
            fmt.Println(chanNum)
        case chanString := <- stringChan:
            fmt.Println(chanString)
        case <- numComplete:
            winner = "numbers"
            done = true
        case <- stringComplete:
            winner = "strings"
            done = true
        }
    }
    return winner
}