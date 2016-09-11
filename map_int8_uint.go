// THIS IS A MACHINE GENERATED FILE!
// BE CAREFUL WITH EDITING BY HAND.
//
// This is Free Software covered by the terms of the MIT license.
// See LICENSE file for details.
// (c) 2016 by the contributors.
// See CONTRIBUTORS for a full list.

package genmap

type entryInt8ToUint struct {
	k    int8
	v    uint
	next *entryInt8ToUint
}

// MapInt8ToUint implements a hash map from int8 to uint.
type MapInt8ToUint struct {
	mask     int
	slots    []*entryInt8ToUint
	used     int
	size     int
	max      int
	freelist *entryInt8ToUint
}

// NewMapInt8ToUint creates a new MapInt8ToUint with at least a size of size.
func NewMapInt8ToUint(size int) *MapInt8ToUint {
	shift := nextShiftPowerOfTwo(size)
	return &MapInt8ToUint{
		mask:  ^(^0 << shift),
		max:   maxFill(1 << shift),
		slots: make([]*entryInt8ToUint, 1<<shift),
	}
}

// Get looks up a key k returns its value. 0 if not found.
func (h *MapInt8ToUint) Get(k int8) uint {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v
		}
	}
	return 0
}

// Visit calls a given function fn for every key/value pair in the map.
func (h *MapInt8ToUint) Visit(fn func(int8, uint)) {
	for _, e := range h.slots {
		for ; e != nil; e = e.next {
			fn(e.k, e.v)
		}
	}
}

// Add increments the value associated with key k by the value of v.
// It creates a new key/value pair k/v in the map if the key does not exist.
func (h *MapInt8ToUint) Add(k int8, v uint) {
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
func (h *MapInt8ToUint) Put(k int8, v uint) {
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
func (h *MapInt8ToUint) Remove(k int8) {
	p := &h.slots[int(k)&h.mask]
	var parent *entryInt8ToUint
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
func (h *MapInt8ToUint) Inc(k int8) {
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

func (h *MapInt8ToUint) rehash() {
	ns := len(h.slots) << 1
	nslots := make([]*entryInt8ToUint, ns)
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

func (h *MapInt8ToUint) free(entry *entryInt8ToUint) {
	entry.next = h.freelist
	h.freelist = entry
	h.size--
}

func (h *MapInt8ToUint) alloc(k int8, v uint) *entryInt8ToUint {
	if h.freelist == nil {
		entries := make([]entryInt8ToUint, 256)
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
