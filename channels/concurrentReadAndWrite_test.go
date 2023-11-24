package channels

import (
	"reflect"
	"sync"
	"testing"
)

func TestConcurrentReadAndWrite(t *testing.T) {
	ints := []int{2,3,4,5,7,8}
	want := []int{2,4,8}
	got := WriteAndAppend(ints) 
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Wrong values were returned: got %v, want %v", got, want)
	}
}

func TestSendNumToChan(t *testing.T) {
	ints := []int{12, 14, 15, 16}
	want := []int{12, 14, 16}
	numChan := make(chan int)

	go sendNumToChan(ints, numChan)

	i :=0
	for  num := range numChan {
		if num != want[i] {
		t.Errorf("Wrong numbers were written to chan")
		}
		i++
	}
}

func TestReadNumsFromChan(t *testing.T) {
	var wg sync.WaitGroup
	ints := []int{2,3,4,6,8}
	res := []int{}
	numChan := make(chan int)
	want := []int{2,4,6,8}
	wg.Add(2)
	go func(){
		defer wg.Done()
		sendNumToChan(ints, numChan)
	}()
	i :=0
	for  num := range numChan {
		if num != want[i] {
			t.Errorf("Wrong numbers were written to chan")
		}
		i++
	}

	go func(){
		defer wg.Done()
		readFromChannel(&res, numChan)
	}()
	
	wg.Wait()
	
	if reflect.DeepEqual(want, res) {
		t.Errorf("The wrong values were written to the channel")
	}

}