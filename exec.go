package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("./dir", "ls")
	buf, err := cmd.Output()
	fmt.Println(err)
	fmt.Println(buf)
}
