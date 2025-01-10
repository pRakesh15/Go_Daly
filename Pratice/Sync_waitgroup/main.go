package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("worker %d started \n", i)
	fmt.Printf("worker %d ended \n", i)

}

func main() {

	// fmt.Println("this is sync wait group")

	//crate a variable  to store the wait groups
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		//lets make a individual channel for goroutine  ....
		//add the wait group by 1
		//define this worker function is a wait group
		wg.Add(1)
		go worker(i, &wg)
		//there are  3 individual channels created for 1,2 and 3
		//according to the channel 3 wg also created

	}

	//then wait for the worker groups completion
	wg.Wait()
	//this a individual channel
	fmt.Println("end of the function")

	//both are running concurrently

	// this is how we manually delly a function or make a static wait
	// but its not the write way to do this bcz if we delly by  1 sec but the channel takes 2 sec to execute ...
	// // time.Sleep(1 * time.Second)
}
