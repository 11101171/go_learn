package main

//****************************************************/
//Copyright(c) 2015 Tencent, all rights reserved
// File        : pty.go
// Author      : ningzhong.zeng
// Revision    : 2015-12-25 19:06:24
// Description :
//****************************************************/

import (
	"io"
	"os"
	"os/exec"

	"github.com/kr/pty"
)

func main() {
	c := exec.Command("ssh", "souche@120.26.78.245")
	f, err := pty.Start(c)
	if err != nil {
		panic(err)
	}

	go func() {
		f.Write([]byte("pwd!Z@X@souche2013\n"))
		f.Write([]byte("cd /home\n"))
		f.Write([]byte("ls -a\n"))
		f.Write([]byte{4}) // EOT
	}()
	io.Copy(os.Stdout, f)
}
