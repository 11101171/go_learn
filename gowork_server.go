package main

//****************************************************/
//Copyright(c) 2015 Tencent, all rights reserved
// File        : gowork.go
// Author      : ningzhong.zeng
// Revision    : 2015-12-12 14:11:58
// Description :
//****************************************************/

import (
	"fmt"
	"time"

	"github.com/benmanns/goworker"
)

func init() {
	goworker.Register("Hello", helloWorker)
}

func helloWorker(queue string, args ...interface{}) error {
	fmt.Println("Hello, world!")
	return nil
}

func main() {
	fmt.Println("Start Main func()")
	go func() {
		for {
			if err := goworker.Work(); err != nil {
				fmt.Println("Error", err)
			}
		}
	}()

	time.Sleep(5 * 60)
}
