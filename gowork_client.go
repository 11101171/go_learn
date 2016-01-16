package main

//****************************************************/
//Copyright(c) 2015 Tencent, all rights reserved
// File        : gowork_client.go
// Author      : ningzhong.zeng
// Revision    : 2015-12-12 15:23:34
// Description :
//****************************************************/

import (
	"fmt"

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
}
