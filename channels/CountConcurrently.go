package channels

import (
	"sync"
 	"time"
)

func CountConcurrently(delay time.Duration) []string {
	var wg sync.WaitGroup
	c := make(chan string)
	results := []string{}
	wg.Add(2)
	go func() {
		count("sheep", c, delay)
		wg.Done()
	}()
	go func() {
		count("fish", c, delay)
		wg.Done()
	}()
	go func() {
		wg.Wait()
		close(c)
	}()
	
	for value := range c {
		results = append(results, value)
	}
	return results
}

func count (thing string, c chan string, delay time.Duration) {
	for i := 1; i <= 5; i++ {
		c <- thing
		time.Sleep(time.Millisecond * delay)
	}
	
}