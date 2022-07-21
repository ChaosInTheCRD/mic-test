package main

import (
	"fmt"
	"os"
	"log"
)

func main() {
	fmt.Println("check, check, 1-2-3, test 1-2-3, check 1-2-1-2...")
        fmt.Println("another check, 2-2-1-2, oi oi yeah yeah")

        fmt.Println("I am now officially attestagon'd")

	file, err := os.Open("/rand.txt")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Print(file)
}
