package main

import (
	"fmt"
	"os"
)

func main() {
	var results = make(chan Result)

	c, err := GetConfig(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	d := NewDatabase(c)
	count := 0

	go func() {
		for r := range results {
			count++
			fmt.Println(r)
		}
	}()

	fmt.Println("Running scan...")

	if err := d.Scan(results); err != nil {
		fmt.Println(err)
		os.Exit(2)
		return
	}

	fmt.Printf("\nDetected %d items\n", count)
	if count != 0 {
		os.Exit(1)
	}
}
