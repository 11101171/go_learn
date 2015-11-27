package main

//****************************************************/
//Copyright(c) 2015 Tencent, all rights reserved
// File        : mutex.go
// Author      : ningzhong.zeng
// Revision    : 2015-11-24 18:01:27
// Description :
//****************************************************/

import (
	"fmt"
	"sync"
	// "sync/atomic"
	"time"
)

var total_tickets int32 = 10

func main() {
	fmt.Println("Start Main func()")
	var mutex = &sync.Mutex{}

	for r := 0; r < 10; r++ {
		go func() {
			var i int32 = 1
			for i <= 5 {
				mutex.Lock()
				total_tickets++
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(time.Second)

	fmt.Println("ticket ", total_tickets)
}
