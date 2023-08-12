package main

import (
	"fmt"
	"log"
	"lebrancconvas/go-interface/interfaces"
)

type integer int

func main() {
	max, err := interfaces.Max([]int{8, 1, 9, 2, 3})
	if err != nil {
		log.Fatal(err)
	}

	max2, err := interfaces.Max([]integer{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(max)
	fmt.Println(max2)

}

