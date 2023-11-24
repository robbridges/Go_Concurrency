package channels
import(
	"sync"
)

func WriteAndAppend(nums []int) []int {
	var wg sync.WaitGroup
    wg.Add(2)
    numChan := make(chan int)
    var res []int
    go func(){
        defer wg.Done()
        sendNumToChan(nums, numChan)
    }()
    go func(){
        defer wg.Done()
        readFromChannel(&res, numChan)
    }()
    wg.Wait()
	return res

}

func sendNumToChan(nums []int, c chan int) {
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
		// best to make this a pointer, while we could just return the array that we create, it's good pointer practice
        *nums = append(*nums, num)
    }
	
}