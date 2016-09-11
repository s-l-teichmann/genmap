// THIS IS A MACHINE GENERATED FILE!
// BE CAREFUL WITH EDITING BY HAND.
//
// This is Free Software covered by the terms of the MIT license.
// See LICENSE file for details.
// (c) 2016 by the contributors.
// See CONTRIBUTORS for a full list.

package genmap

type entryUint8ToInt8 struct {
	k    uint8
	v    int8
	next *entryUint8ToInt8
}

// MapUint8ToInt8 implements a hash map from uint8 to int8.
type MapUint8ToInt8 struct {
	mask  int
	slots []*entryUint8ToInt8
	used  int
	size  int
	max   int
	free  []entryUint8ToInt8
}

// NewMapUint8ToInt8 creates a new MapUint8ToInt8 with at least a size of size.
func NewMapUint8ToInt8(size int) *MapUint8ToInt8 {
	shift := nextShiftPowerOfTwo(size)
	return &MapUint8ToInt8{
		mask:  ^(^0 << shift),
		max:   maxFill(1 << shift),
		slots: make([]*entryUint8ToInt8, 1<<shift),
	}
}

// Get looks up a key k returns its value. 0 if not found.
func (h *MapUint8ToInt8) Get(k uint8) int8 {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v
		}
	}
	return 0
}

// Visit calls a given function fn for every key/value pair in the map.
func (h *MapUint8ToInt8) Visit(fn func(uint8, int8)) {
	for _, e := range h.slots {
		for ; e != nil; e = e.next {
			fn(e.k, e.v)
		}
	}
}

// Add increments the value associated with key k by the value of v.
// It creates a new key/value pair k/v in the map if the key does not exist.
func (h *MapUint8ToInt8) Add(k uint8, v int8) {
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
func (h *MapUint8ToInt8) Inc(k uint8) {
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

func (h *MapUint8ToInt8) rehash() {
	ns := len(h.slots) << 1
	nslots := make([]*entryUint8ToInt8, ns)
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

func (h *MapUint8ToInt8) alloc(k uint8, v int8) *entryUint8ToInt8 {
	if len(h.free) == 0 {
		h.free = make([]entryUint8ToInt8, 256)
	}
	h.size++
	x := &h.free[0]
	x.k = k
	x.v = v
	h.free = h.free[1:]
	return x
}
