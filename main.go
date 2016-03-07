package main

import (
	"fmt"
	"os"
)

func main() {
	var results = make(chan Result)

	c := GetConfig()
	d := NewDatabase(c)
	count := 0

	go func() {
		for r := range results {
			count++
			fmt.Println(r)
		}
	}()

	fmt.Println("Running scan...")

	err := d.Scan(results)

	fmt.Printf("\nDetected %d items\n", count)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
