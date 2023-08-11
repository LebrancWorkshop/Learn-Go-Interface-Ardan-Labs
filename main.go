package main

import (
	"fmt"
	"log"
	"lebrancconvas/go-interface/interfaces"
)

func main() {
	max, err := interfaces.MaxInts([]int{1, 3, 2, 5, 7, 2})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(max)

}

