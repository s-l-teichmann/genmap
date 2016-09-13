// This is Free Software covered by the terms of the MIT license.
// See LICENSE file for details.
// (c) 2016 by Sascha L. Teichmann.
// See the full list of contributors in the CONTRIBUTORS file.

package main

type typ struct {
	Name   string
	Signed bool
}

var types = []typ{
	{"int", true},
	{"int8", true},
	{"int16", true},
	{"int32", true},
	{"int64", true},
	{"uint", false},
	{"uint8", false},
	{"uint16", false},
	{"uint32", false},
	{"uint64", false},
}
