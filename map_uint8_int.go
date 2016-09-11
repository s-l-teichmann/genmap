// THIS IS A MACHINE GENERATED FILE!
// BE CAREFUL WITH EDITING BY HAND.
//
// This is Free Software covered by the terms of the MIT license.
// See LICENSE file for details.
// (c) 2016 by the contributors.
// See CONTRIBUTORS for a full list.

package genmap

type entryUint8ToInt struct {
	k    uint8
	v    int
	next *entryUint8ToInt
}

// MapUint8ToInt implements a hash map from uint8 to int.
type MapUint8ToInt struct {
	mask     int
	slots    []*entryUint8ToInt
	used     int
	size     int
	max      int
	freelist *entryUint8ToInt
}

// NewMapUint8ToInt creates a new MapUint8ToInt with at least a size of size.
func NewMapUint8ToInt(size int) *MapUint8ToInt {
	shift := nextShiftPowerOfTwo(size)
	return &MapUint8ToInt{
		mask:  ^(^0 << shift),
		max:   maxFill(1 << shift),
		slots: make([]*entryUint8ToInt, 1<<shift),
	}
}

// Get looks up a key k returns its value. 0 if not found.
func (h *MapUint8ToInt) Get(k uint8) int {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v
		}
	}
	return 0
}

// Visit calls a given function fn for every key/value pair in the map.
func (h *MapUint8ToInt) Visit(fn func(uint8, int)) {
	for _, e := range h.slots {
		for ; e != nil; e = e.next {
			fn(e.k, e.v)
		}
	}
}

// Add increments the value associated with key k by the value of v.
// It creates a new key/value pair k/v in the map if the key does not exist.
func (h *MapUint8ToInt) Add(k uint8, v int) {
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
func (h *MapUint8ToInt) Put(k uint8, v int) {
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
func (h *MapUint8ToInt) Remove(k uint8) {
	p := &h.slots[int(k)&h.mask]
	var parent *entryUint8ToInt
	for e := *p; e != nil; e = e.next {
		if e.k == k {
			if parent == nil { // head
				if e.next == nil { // last in chain
					*p = nil
					h.used--
				} else {
					*p = e.next
				}
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
func (h *MapUint8ToInt) Inc(k uint8) {
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

func (h *MapUint8ToInt) rehash() {
	ns := len(h.slots) << 1
	nslots := make([]*entryUint8ToInt, ns)
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

func (h *MapUint8ToInt) free(entry *entryUint8ToInt) {
	entry.next = h.freelist
	h.freelist = entry
	h.size--
}

func (h *MapUint8ToInt) alloc(k uint8, v int) *entryUint8ToInt {
	if h.freelist == nil {
		entries := make([]entryUint8ToInt, 256)
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
