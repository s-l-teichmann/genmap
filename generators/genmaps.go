package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

var types = []string{
	"int", "int8", "int16", "int32", "int64",
	"uint", "uint8", "uint16", "uint32", "uint64",
}

var mapTmplTxt = `// THIS IS A MACHINE GENERATED FILE!
// BE CAREFUL WITH EDITING BY HAND.
//
// This is Free Software covered by the terms of the MIT license.
// See LICENSE file for details.
// (c) 2016 by the contributors.
// See CONTRIBUTORS for a full list.

package genmap

{{ $F := .From | title -}}
{{- $T := .To | title -}}
{{- $TYPE :=  printf "Map%sTo%s" $F $T -}}
{{- $ENTRY :=  printf "entry%sTo%s" $F $T -}}

type {{ $ENTRY }} struct {
	k    {{.From }}
	v    {{ .To }}
	next *{{ $ENTRY }}
}

// {{ $TYPE }} implements a hash map from {{ .From }} to {{ .To }}.
type {{ $TYPE }} struct {
	mask  int
	slots []*{{ $ENTRY }}
	used  int
	size  int
	max   int
	free  []{{ $ENTRY }}
}

// New{{ $TYPE }} creates a new {{ $TYPE }} with at least a size of size.
func New{{ $TYPE }}(size int) *{{ $TYPE }} {
	shift := nextShiftPowerOfTwo(size)
	return &{{ $TYPE }}{
		mask:  ^(^0 << shift),
		max:   maxFill(1 << shift),
		slots: make([]*{{ $ENTRY }}, 1<<shift),
	}
}

// Get looks up a key k returns its value. 0 if not found.
func (h *{{ $TYPE }}) Get(k {{ .From }}) {{ .To }} {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v
		}
	}
	return 0
}

// Visit calls a given function fn for every key/value pair in the map.
func (h *{{ $TYPE }}) Visit(fn func({{ .From }}, {{ .To }})) {
	for _, e := range h.slots {
		for ; e != nil; e = e.next {
			fn(e.k, e.v)
		}
	}
}

// Add increments the value associated with key k by the value of v.
// It creates a new key/value pair k/v in the map if the key does not exist.
func (h *{{ $TYPE }}) Add(k {{ .From }}, v {{ .To }}) {
	p := &h.slots[int(k)&h.mask]
	for e := *p; e != nil; e = e.next {
		if e.k == k {
			e.v += v
			return
		}
	}
	n := h.alloc(k, v)
	if *p != nil {
		n.next = *p
		*p = n
	} else {
		*p = n
		h.used++
		if h.used > h.max {
			h.rehash()
		}
	}
}

// Inc increments a value associated with key k by one.
// A new entry is created with value 1 if the key
// does not exist.
func (h *{{ $TYPE }}) Inc(k {{ .From }}) {
	p := &h.slots[int(k)&h.mask]
	for e := *p; e != nil; e = e.next {
		if e.k == k {
			e.v++
			return
		}
	}
	n := h.alloc(k, 1)
	if *p != nil {
		n.next = *p
		*p = n
	} else {
		*p = n
		h.used++
		if h.used > h.max {
			h.rehash()
		}
	}
}

func (h *{{ $TYPE }}) rehash() {
	ns := len(h.slots) << 1
	nslots := make([]*{{ $ENTRY }}, ns)
	nmask := (h.mask << 1) | 1
	h.mask = nmask
	nu := 0
	for i, e := range h.slots {
		if e == nil {
			continue
		}
		for e != nil {
			n := e.next
			p := &nslots[int(e.k)&nmask]
			if *p == nil {
				nu++
			}
			e.next = *p
			*p = e
			e = n
		}
		h.slots[i] = nil
	}
	h.used = nu
	h.slots = nslots
	h.max = maxFill(ns)
}

func (h *{{ $TYPE }}) alloc(k {{ .From }}, v {{ .To }}) *{{ $ENTRY }} {
	if len(h.free) == 0 {
		h.free = make([]{{ $ENTRY }}, 256)
	}
	h.size++
	x := &h.free[0]
	x.k = k
	x.v = v
	h.free = h.free[1:]
	return x
}
`

var funcMap = template.FuncMap{
	"title": strings.Title,
}

var mapTmpl = template.Must(template.New("maps").Funcs(funcMap).Parse(mapTmplTxt))

func matrix(fn func(from, to string)) {
	for _, a := range types {
		for _, b := range types {
			fn(a, b)
		}
	}
}

func main() {
	matrix(func(from, to string) {
		filename := fmt.Sprintf("map_%s_%s.go", from, to)
		log.Printf("Generating %s...\n", filename)
		f, err := os.Create(filename)
		if err != nil {
			log.Printf("creating %s failed: %v\n", filename, err)
			return
		}
		out := bufio.NewWriter(f)
		var parameters = struct {
			From string
			To   string
		}{
			From: from,
			To:   to,
		}
		if err = mapTmpl.Execute(out, &parameters); err != nil {
			log.Printf("templating %s failed: %v\n", filename, err)
		}
		if err = out.Flush(); err != nil {
			log.Printf("flushing %s failed: %v\n", filename, err)
		}
		if err = f.Close(); err != nil {
			log.Printf("closing %s failed: %v\n", filename, err)
		}
	})
}
