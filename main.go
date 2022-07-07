package main

import (
	"fmt"
	"os"
	"log"
)

func main() {
	fmt.Println("check, check, 1-2-3, test 1-2-3, check 1-2-1-2...")
        fmt.Println("Give it another check, 2-2-1-2, yeah yeah yeah")

	file, err := os.Open("/rand.txt")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Print(file)
}
