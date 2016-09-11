/*
Package genmap provides specialized hashmaps for integer types.

Installation

	go get github.com/s-l-teichmann/genmap

Example program using a uint64 to int map:

	package main

	import (
		"fmt"
		"github.com/s-l-teichmann/genmap"
	)

	func main() {
		// Create a new hash map.
		h := genmap.NewHashUint64ToInt(16)

		// Creates new key/value pairs 42 -> 10 and 47 -> 11.
		h.Add(42, 10)
		h.Add(47, 11)

		// Increases the value of key 47 to 111.
		h.Add(47, 100)

		// Increments value 42 to 43 of key 42.
		h.Inc(42)

		// Fetch value of key 42 from map.
		fmt.Printf("value: %d\n", h.Get(42))

		// Prints the number of key/value pairs in the map.
		fmt.Printf("size: %d\n", h.Size())

		// Prints all key/value pairs in map in undefined order.
		h.Visit(func(k uint64, v int) {
			fmt.Printf("key: %d value: %d\n", k, v)
		})
	}

For a more complete set of methods see the index.

*/
package genmap
