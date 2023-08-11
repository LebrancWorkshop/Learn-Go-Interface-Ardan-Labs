package main

import (
	"fmt"
	"log"
	"lebrancconvas/go-interface/interfaces"
)

func main() {
	max, err := interfaces.Max([]int{8, 1, 9, 2, 3})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(max)

}

