package main

import (
	"example/fem/data"
	"fmt"
)

func main() {
	max := data.NewInstructor("max", "maxiliano")
	fmt.Printf("%v", max.Id)
}
