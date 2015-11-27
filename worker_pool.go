package main

//****************************************************/
//Copyright(c) 2015 Tencent, all rights reserved
// File        : worker_pool.go
// Author      : ningzhong.zeng
// Revision    : 2015-11-24 14:43:45
// Description :
//****************************************************/

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * j
	}
}

func main() {
	fmt.Println("Start Main func()")

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 0; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)
	for a := 1; a <= 9; a++ {
		// <-results
		fmt.Println(<-results)
	}

}
