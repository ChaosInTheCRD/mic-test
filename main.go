package main

import (
	"fmt"
	"os"
	"log"
)

func main() {
	fmt.Println("check, check, 1-2-3, test 1-2-3, check 1-2-1-2...")

	file, err := os.Open("/rand.txt")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Print(file)
}
