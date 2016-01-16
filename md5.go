package main

//****************************************************/
//Copyright(c) 2015 Tencent, all rights reserved
// File        : md5.go
// Author      : ningzhong.zeng
// Revision    : 2016-01-15 17:49:18
// Description :
//****************************************************/
import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte("test md5 encrypto"))
	cipherStr := md5Ctx.Sum(nil)
	fmt.Print(cipherStr)
	fmt.Print("\n")
	fmt.Print(hex.EncodeToString(cipherStr))
}
