package main

import "fmt"

func main() {
	ch := make(chan string)
	sum := 1
	for sum < 5 {
		go shower(ch, "sum"+sum)
		sum++
	}
	for i := 0; i < 10; i++ {
		name <- ch
		fmt.Println("read-%d,name-%s", i, name)
	}
}
func shower(c chan string, name string) {
	for {
		ch <- name
		fmt.Printf("ch<-name:[%s]", name)
	}
}
