// This is Free Software covered by the terms of the MIT license.
// See LICENSE file for details.
// (c) 2016 by Sascha L. Teichmann.
// See the full list of contributors in the CONTRIBUTORS file.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

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

var headerMapsTmplTxt = `// THIS IS A MACHINE GENERATED FILE!
// BE CAREFUL WITH EDITING BY HAND.
//
// This is Free Software covered by the terms of the MIT license.
// See LICENSE file for details.
// (c) 2016 by Sascha L. Teichmann.
// See the full list of contributors in the CONTRIBUTORS file.

package genmap
`

var headerTestsTmplTxt = `// THIS IS A MACHINE GENERATED FILE!
// BE CAREFUL WITH EDITING BY HAND.
//
// This is Free Software covered by the terms of the MIT license.
// See LICENSE file for details.
// (c) 2016 by Sascha L. Teichmann.
// See the full list of contributors in the CONTRIBUTORS file.

package genmap

import "testing"
`

var mapsTmplTxt = `
{{     $FROM  := .From.Name }}{{/*
*/}}{{ $TO    := .To.Name }}{{/*
*/}}{{ $F     := $FROM | title }}{{/*
*/}}{{ $T     := $TO | title }}{{/*
*/}}{{ $TYPE  := printf "Map%sTo%s" $F $T }}{{/*
*/}}{{ $ENTRY := printf "entry%sTo%s" $F $T }}{{/*
*/}}type {{ $ENTRY }} struct {
	k    {{ $FROM }}
	v    {{ $TO }}
	next *{{ $ENTRY }}
}

// {{ $TYPE }} implements a hash map from {{ $FROM }} to {{ $TO }}.
type {{ $TYPE }} struct {
	mask     int
	slots    []*{{ $ENTRY }}
	used     int
	size     int
	max      int
	freelist *{{ $ENTRY }}
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

// Size returns the current size of the map.
func (h *{{ $TYPE }}) Size() int {
	return h.size
}

// Get looks up a key k and returns its value. 0 if not found.
func (h *{{ $TYPE }}) Get(k {{ $FROM }}) {{ $TO }} {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v
		}
	}
	return 0
}

// Contains looks up a key k and returns true if it is found else false.
func (h *{{ $TYPE }}) Contains(k {{ $FROM }}) bool {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return true
		}
	}
	return false
}

// Modify looks up a key k and calls function fn with a pointer to its value.
func (h *{{ $TYPE }}) Modify(k {{ $FROM }}, fn func(v *{{ $TO }})) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			fn(&e.v)
			return
		}
	}
}

// Find looks up a key k and returns its value and true.
// 0 and false if not found.
func (h *{{ $TYPE }}) Find(k {{ $FROM }}) ({{ $TO }}, bool) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v, true
		}
	}
	return 0, false
}

// Visit calls a given function fn for every key/value pair in the map.
func (h *{{ $TYPE }}) Visit(fn func({{ $FROM }}, {{ $TO }})) {
	for _, e := range h.slots {
		for ; e != nil; e = e.next {
			fn(e.k, e.v)
		}
	}
}

// Add increments the value associated with key k by the value of v.
// It creates a new key/value pair k/v in the map if the key does not exist.
func (h *{{ $TYPE }}) Add(k {{ $FROM }}, v {{ $TO }}) {
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

// Put sets the value associated with key k to the value of v.
// It creates a new key/value pair k/v in the map if the key does not exist.
func (h *{{ $TYPE }}) Put(k {{ $FROM }}, v {{ $TO }}) {
	p := &h.slots[int(k)&h.mask]
	for e := *p; e != nil; e = e.next {
		if e.k == k {
			e.v = v
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

// Remove removes the key/value pair associated with key k from this map..
func (h *{{ $TYPE }}) Remove(k {{ $FROM }}) {
	p := &h.slots[int(k)&h.mask]
	var parent *{{ $ENTRY }}
	for e := *p; e != nil; e = e.next {
		if e.k == k {
			if parent == nil { // head
				if e.next == nil { // last in chain
					h.used--
				}
				*p = e.next
			} else {
				parent.next = e.next
			}
			h.free(e)
			return
		}
		parent = e
	}
}

// Clear removes all elements from the map.
func (h *{{ $TYPE }}) Clear() {
	for i, e := range h.slots {
		if e != nil {
			for e != nil {
				n := e.next
				h.free(e)
				e = n
			}
			h.slots[i] = nil
		}
	}
	h.used = 0
}

// Inc increments a value associated with key k by one.
// A new entry is created with value 1 if the key
// does not exist.
func (h *{{ $TYPE }}) Inc(k {{ $FROM }}) {
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

func (h *{{ $TYPE }}) free(entry *{{ $ENTRY }}) {
	entry.next = h.freelist
	h.freelist = entry
	h.size--
}

func (h *{{ $TYPE }}) alloc(k {{ $FROM }}, v {{ $TO }}) *{{ $ENTRY }} {
	if h.freelist == nil {
		entries := make([]{{ $ENTRY }}, 256)
		for i := 0; i < 256-1; i++ {
			entries[i].next = &entries[i+1]
		}
		h.freelist = &entries[0]
	}
	h.size++
	x := h.freelist
	h.freelist = x.next
	x.next = nil
	x.k = k
	x.v = v
	return x
}
`

var testsTmplText = `
{{     $FROM  := .From.Name }}{{/*
*/}}{{ $TO    := .To.Name }}{{/*
*/}}{{ $F     := $FROM | title }}{{/*
*/}}{{ $T     := $TO | title }}{{/*
*/}}{{ $TYPE  := printf "Map%sTo%s" $F $T }}{{/*
*/}}{{ $KEYS  := or (and .From.Signed "signedData") "unsignedData" }}{{/*
*/}}{{ $VALS  := or (and .To.Signed "signedData") "unsignedData" }}{{/*
*/}}func Test{{ $TYPE }}Size(t *testing.T) {
	m := New{{ $TYPE }}(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range {{ $KEYS }} {
		m.Put({{ $FROM }}(k), {{ $TO }}({{ $VALS }}[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range {{ $KEYS }} {
		m.Remove({{ $FROM }}(k))
		if want := len({{ $KEYS }}) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func Test{{ $TYPE }}Contains(t *testing.T) {
	m := New{{ $TYPE }}(13)
	for i, k := range {{ $KEYS }} {
		m.Put({{ $FROM }}(k), {{ $TO }}({{ $VALS }}[i]))
		if !m.Contains({{ $FROM }}(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range {{ $KEYS }} {
		m.Remove({{ $FROM }}(k))
		if m.Contains({{ $FROM }}(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func Test{{ $TYPE }}Get(t *testing.T) {
	m := New{{ $TYPE }}(13)
	for i, k := range {{ $KEYS }} {
		m.Put({{ $FROM }}(k), {{ $TO }}({{ $VALS }}[i]))
		if g := m.Get({{ $FROM }}(k)); g != {{ $TO }}({{ $VALS }}[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, {{ $VALS }}[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func Test{{ $TYPE }}Remove(t *testing.T) {
	m := New{{ $TYPE }}(13)
	// Used to trigger code path when all keys are
	// hashed to same slot.
	mask := ^(^0 << nextShiftPowerOfTwo(13))
	for i, k := range {{ $KEYS }} {
		if int(k)&mask == 0 {
			m.Put({{ $FROM }}(k), {{ $TO }}({{ $VALS }}[i]))
		}
	}
	s := m.Size()
	for _, k := range {{ $KEYS }} {
		if int(k)&mask == 0 {
			m.Remove({{ $FROM }}(k))
			s--
			if l := m.Size(); l != s {
				t.Errorf("key: %d size %d, want %d\n", k, l, s)
			}
		}
	}

}

func Test{{ $TYPE }}Clear(t *testing.T) {
	m := New{{ $TYPE }}(13)
	for i, k := range {{ $KEYS }} {
		m.Put({{ $FROM }}(k), {{ $TO }}({{ $VALS }}[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range {{ $KEYS }} {
		m.Put({{ $FROM }}(k), {{ $TO }}({{ $VALS }}[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func Test{{ $TYPE }}Inc(t *testing.T) {
	m := New{{ $TYPE }}(13)
	for i, k := range {{ $KEYS }} {
		m.Put({{ $FROM }}(k), {{ $TO }}({{ $VALS }}[i]))
	}
	for i, k := range {{ $KEYS }} {
		m.Inc({{ $FROM }}(k))
		m.Inc({{ $FROM }}(k))
		m.Inc({{ $FROM }}(k))
		if g := m.Get({{ $FROM }}(k)); g != {{ $TO }}({{ $VALS }}[i])+3 {
			t.Errorf("got %d, want %d\n", g, {{ $TO }}({{ $VALS }}[i])+3)
		}
	}
	m = New{{ $TYPE }}(13)
	for _, k := range {{ $KEYS }} {
		m.Inc({{ $FROM }}(k))
		if g := m.Get({{ $FROM }}(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func Test{{ $TYPE }}Add(t *testing.T) {
	m := New{{ $TYPE }}(13)
	for i, k := range {{ $KEYS }} {
		m.Put({{ $FROM }}(k), {{ $TO }}({{ $VALS }}[i]))
	}
	for i, k := range {{ $KEYS }} {
		m.Add({{ $FROM }}(k), 3)
		if g := m.Get({{ $FROM }}(k)); g != {{ $TO }}({{ $VALS }}[i])+3 {
			t.Errorf("got %d, want %d\n", g, {{ $TO }}({{ $VALS }}[i])+3)
		}
	}
	m = New{{ $TYPE }}(13)
	for _, k := range {{ $KEYS }} {
		m.Add({{ $FROM }}(k), 42)
		if g := m.Get({{ $FROM }}(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func Test{{ $TYPE }}Modify(t *testing.T) {
	m := New{{ $TYPE }}(13)
	for i, k := range {{ $KEYS }} {
		m.Put({{ $FROM }}(k), {{ $TO }}({{ $VALS }}[i]))
	}
	for _, k := range {{ $KEYS }} {
		m.Modify({{ $FROM }}(k), func(v *{{ $TO }}) {
			*v += 3
		})
	}
	for i, k := range {{ $KEYS }} {
		if g := m.Get({{ $FROM }}(k)); g != {{ $TO }}({{ $VALS }}[i])+3 {
			t.Errorf("got %d, want %d\n", g, {{ $TO }}({{ $VALS }}[i])+3)
		}
	}
}

func Test{{ $TYPE }}Find(t *testing.T) {
	m := New{{ $TYPE }}(13)

	low := {{ $KEYS }}[:len({{ $KEYS }})/2]
	hi := {{ $KEYS }}[len({{ $KEYS }})/2:]

	for i, k := range low {
		m.Put({{ $FROM }}(k), {{ $TO }}({{ $VALS }}[i]))
	}
	for i, k := range low {
		v, ok := m.Find({{ $FROM }}(k))
		if !ok {
			t.Errorf("missing value for key %d\n", k)
			continue
		}
		if want := {{ $TO }}({{ $VALS }}[i]); v != want {
			t.Errorf("key %d: got %d, want %d\n", k, v, want)
		}
	}
	for _, k := range hi {
		if v, ok := m.Find({{ $FROM }}(k)); ok {
			t.Errorf("key %d: got %d, want nothing\n", k, v)
		}
	}
}

func Test{{ $TYPE }}Put(t *testing.T) {
	m := New{{ $TYPE }}(13)
	for i, k := range {{ $KEYS }} {
		m.Put({{ $FROM }}(k), {{ $TO }}({{ $VALS }}[i]))
	}
	for _, k := range {{ $KEYS }} {
		m.Put({{ $FROM }}(k), m.Get({{ $FROM }}(k))+3)
	}
	for i, k := range {{ $KEYS }} {
		if g := m.Get({{ $FROM }}(k)); g != {{ $TO }}({{ $VALS }}[i])+3 {
			t.Errorf("got %d, want %d\n", g, {{ $TO }}({{ $VALS }}[i])+3)
		}
	}
}

func Test{{ $TYPE }}Visit(t *testing.T) {
	m := New{{ $TYPE }}(13)
	for i, k := range {{ $KEYS }} {
		m.Put({{ $FROM }}(k), {{ $TO }}({{ $VALS }}[i]))
	}

	n := make(map[{{ $FROM }}]{{ $TO }}, len({{ $KEYS }}))

	for i, k := range {{ $KEYS }} {
		n[{{ $FROM }}(k)] = {{ $TO }}({{ $VALS }}[i])
	}

	m.Visit(func(k {{ $FROM }}, v {{ $TO }}) {
		g, ok := n[k]
		if !ok {
			t.Errorf("key %d not found.\n", k)
			return
		}
		if g != v {
			t.Errorf("key %d: got %d, want %d\n", k, v, g)
		}
		delete(n, k)
	})

	if len(n) > 0 {
		t.Errorf("Size is %d, want 0\n", len(n))
	}
}
`

var funcMap = template.FuncMap{
	"title": strings.Title,
}

func compileTemplate(name, txt string) *template.Template {
	return template.Must(template.New(name).Funcs(funcMap).Parse(txt))
}

var (
	headerMapsTmpl  = compileTemplate("hmaps", headerMapsTmplTxt)
	headerTestsTmpl = compileTemplate("htests", headerTestsTmplTxt)
	mapsTmpl        = compileTemplate("maps", mapsTmplTxt)
	testsTmpl       = compileTemplate("tests", testsTmplText)
)

type parameters struct {
	From *typ
	To   *typ
}

func generateMaps(t *typ) error {
	filename := fmt.Sprintf("map_%s.go", t.Name)
	log.Printf("Generating %s...\n", filename)

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	out := bufio.NewWriter(f)

	if err = headerMapsTmpl.Execute(out, nil); err != nil {
		log.Printf("templating header %s failed: %v\n", filename, err)
		f.Close()
		return err
	}

	p := parameters{From: t}

	for j := range types {
		p.To = &types[j]
		if err = mapsTmpl.Execute(out, &p); err != nil {
			f.Close()
			return err
		}
	}

	if err = out.Flush(); err != nil {
		f.Close()
		return err
	}
	return f.Close()
}

func generateTests(t *typ) error {
	filename := fmt.Sprintf("map_%s_test.go", t.Name)
	log.Printf("Generating %s...\n", filename)

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	out := bufio.NewWriter(f)

	if err = headerTestsTmpl.Execute(out, nil); err != nil {
		f.Close()
		return err
	}

	p := parameters{From: t}

	for j := range types {
		p.To = &types[j]
		if err = testsTmpl.Execute(out, &p); err != nil {
			f.Close()
			return err
		}
	}

	if err = out.Flush(); err != nil {
		f.Close()
		return err
	}
	return f.Close()
}

func main() {
	for i := range types {
		t := &types[i]
		if err := generateMaps(t); err != nil {
			log.Printf("generating maps failed: %v\n", err)
			continue
		}
		if err := generateTests(t); err != nil {
			log.Printf("generating tests failed: %v\n", err)
		}
	}
}
