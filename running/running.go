package main

import (
	"flag"
	"fmt"
)

func main() {
	d := *flag.String("run", "sam", "an string")
	fmt.Println(d)
}
