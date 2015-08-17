package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "nobody", "The name of the person")
	street := flag.String("street", "No street", "The name of the street")

	flag.Parse()
	fmt.Printf("I'm %s from the %s\n", *name, *street)

}
