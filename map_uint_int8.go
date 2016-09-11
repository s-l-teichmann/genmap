// THIS IS A MACHINE GENERATED FILE!
// BE CAREFUL WITH EDITING BY HAND.
//
// This is Free Software covered by the terms of the MIT license.
// See LICENSE file for details.
// (c) 2016 by the contributors.
// See CONTRIBUTORS for a full list.

package genmap

type entryUintToInt8 struct {
	k    uint
	v    int8
	next *entryUintToInt8
}

// MapUintToInt8 implements a hash map from uint to int8.
type MapUintToInt8 struct {
	mask     int
	slots    []*entryUintToInt8
	used     int
	size     int
	max      int
	freelist *entryUintToInt8
}

// NewMapUintToInt8 creates a new MapUintToInt8 with at least a size of size.
func NewMapUintToInt8(size int) *MapUintToInt8 {
	shift := nextShiftPowerOfTwo(size)
	return &MapUintToInt8{
		mask:  ^(^0 << shift),
		max:   maxFill(1 << shift),
		slots: make([]*entryUintToInt8, 1<<shift),
	}
}

// Size returns the current size of the map.
func (h *MapUintToInt8) Size() int {
	return h.size
}

// Get looks up a key k and returns its value. 0 if not found.
func (h *MapUintToInt8) Get(k uint) int8 {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v
		}
	}
	return 0
}

// Contains looks up a key k and returns true if it is found else false.
func (h *MapUintToInt8) Contains(k uint) bool {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return true
		}
	}
	return false
}

// Find looks up a key k and returns its value and true.
// 0 and false if not found.
func (h *MapUintToInt8) Find(k uint) (int8, bool) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v, true
		}
	}
	return 0, false
}

// Visit calls a given function fn for every key/value pair in the map.
func (h *MapUintToInt8) Visit(fn func(uint, int8)) {
	for _, e := range h.slots {
		for ; e != nil; e = e.next {
			fn(e.k, e.v)
		}
	}
}

// Add increments the value associated with key k by the value of v.
// It creates a new key/value pair k/v in the map if the key does not exist.
func (h *MapUintToInt8) Add(k uint, v int8) {
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
func (h *MapUintToInt8) Put(k uint, v int8) {
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
func (h *MapUintToInt8) Remove(k uint) {
	p := &h.slots[int(k)&h.mask]
	var parent *entryUintToInt8
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

// Inc increments a value associated with key k by one.
// A new entry is created with value 1 if the key
// does not exist.
func (h *MapUintToInt8) Inc(k uint) {
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

func (h *MapUintToInt8) rehash() {
	ns := len(h.slots) << 1
	nslots := make([]*entryUintToInt8, ns)
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

func (h *MapUintToInt8) free(entry *entryUintToInt8) {
	entry.next = h.freelist
	h.freelist = entry
	h.size--
}

func (h *MapUintToInt8) alloc(k uint, v int8) *entryUintToInt8 {
	if h.freelist == nil {
		entries := make([]entryUintToInt8, 256)
		for i := 0; i < 256-1; i++ {
			entries[i].next = &entries[i+1]
		}
		h.freelist = &entries[0]
	}
	h.size++
	x := h.freelist
	x.k = k
	x.v = v
	h.freelist = h.freelist.next
	return x
}
