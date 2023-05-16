package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("check, check, 1-2-3, test 1-2-3, check 1-2-1-2..")
	for {
		fmt.Println("Time to test all three!")
		time.Sleep(2 * time.Second)
	}
}
