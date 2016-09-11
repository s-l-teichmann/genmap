# genmap

Specialized hash maps for integer types in Go.

The builtin maps of Go are good but sometimes you need some extra
performance.

This library offers specialized hash maps from *A* to *B*
with *A*, *B* from all of  
`int`, `int8`, `int16`, `int32`, `int64`,  
`uint`, `uint8`, `uint16`, `uint32` and `uint64`.

The maps are named camel cased `MapAToB`.  
So a map from e.g. `int8` to `uint32` is called `MapInt8ToUint32`.

## Installation

You need at least a Go 1.6 environment.

    $ go get github.com/s-l-teichmann/genmap

Find the API documentation here: [![GoDoc](https://godoc.org/github.com/s-l-teichmann/genmap?status.svg)](https://godoc.org/github.com/s-l-teichmann/genmap)

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

For a more complete set of methods see the API index.

## For contributors

The maps are generated from a template in `generators/genmaps.go`.  
Re-generate the maps with

    $ go generate

For pull request ensure that the files are generated in a way
that a `go fmt` call would not change them.

## License
This is Free Software covered by the terms of the [MIT license](LICENSE).  
(c) 2016 by Sascha L. Teichmann
