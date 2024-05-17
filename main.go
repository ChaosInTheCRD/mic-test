package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("bloo bla blee blum")
	fmt.Println("check, check, 1-2-3, test 1-2-3, check 1-2-1-2..")
	fmt.Println("blah blah blah")
	for {
		fmt.Println("one more time.")
		fmt.Println("celebrate and dance so free.")
		fmt.Println("one more time.")
		fmt.Println("one more time.")
		fmt.Println("celebrate and dance so free.")
		time.Sleep(2 * time.Second)
	}
}
