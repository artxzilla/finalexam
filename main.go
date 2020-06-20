package main

import (
	"github.com/artxzilla/finalexam/customer"
)

func main() {
	r := customer.SetupRouter()
	r.Run(":2019")
}
