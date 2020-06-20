package main

import (
	"fmt"

	"github.com/artxzilla/finalexam/customer"
)

func main() {
	fmt.Println("hello go")

	r := customer.SetupRouter()
	r.Run(":2019")
}
