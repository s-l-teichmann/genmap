// THIS IS A MACHINE GENERATED FILE!
// BE CAREFUL WITH EDITING BY HAND.
//
// This is Free Software covered by the terms of the MIT license.
// See LICENSE file for details.
// (c) 2016 by the contributors.
// See CONTRIBUTORS for a full list.

package genmap

type entryUint32ToUint struct {
	k    uint32
	v    uint
	next *entryUint32ToUint
}

// MapUint32ToUint implements a hash map from uint32 to uint.
type MapUint32ToUint struct {
	mask  int
	slots []*entryUint32ToUint
	used  int
	size  int
	max   int
	free  []entryUint32ToUint
}

// NewMapUint32ToUint creates a new MapUint32ToUint with at least a size of size.
func NewMapUint32ToUint(size int) *MapUint32ToUint {
	shift := nextShiftPowerOfTwo(size)
	return &MapUint32ToUint{
		mask:  ^(^0 << shift),
		max:   maxFill(1 << shift),
		slots: make([]*entryUint32ToUint, 1<<shift),
	}
}

// Get looks up a key k returns its value. 0 if not found.
func (h *MapUint32ToUint) Get(k uint32) uint {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v
		}
	}
	return 0
}

// Visit calls a given function fn for every key/value pair in the map.
func (h *MapUint32ToUint) Visit(fn func(uint32, uint)) {
	for _, e := range h.slots {
		for ; e != nil; e = e.next {
			fn(e.k, e.v)
		}
	}
}

// Add increments the value associated with key k by the value of v.
// It creates a new key/value pair k/v in the map if the key does not exist.
func (h *MapUint32ToUint) Add(k uint32, v uint) {
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
func (h *MapUint32ToUint) Inc(k uint32) {
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

func (h *MapUint32ToUint) rehash() {
	ns := len(h.slots) << 1
	nslots := make([]*entryUint32ToUint, ns)
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

func (h *MapUint32ToUint) alloc(k uint32, v uint) *entryUint32ToUint {
	if len(h.free) == 0 {
		h.free = make([]entryUint32ToUint, 256)
	}
	h.size++
	x := &h.free[0]
	x.k = k
	x.v = v
	h.free = h.free[1:]
	return x
}
