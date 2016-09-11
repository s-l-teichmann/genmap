// THIS IS A MACHINE GENERATED FILE!
// BE CAREFUL WITH EDITING BY HAND.
//
// This is Free Software covered by the terms of the MIT license.
// See LICENSE file for details.
// (c) 2016 by the contributors.
// See CONTRIBUTORS for a full list.

package genmap

type entryInt64ToInt64 struct {
	k    int64
	v    int64
	next *entryInt64ToInt64
}

// MapInt64ToInt64 implements a hash map from int64 to int64.
type MapInt64ToInt64 struct {
	mask  int
	slots []*entryInt64ToInt64
	used  int
	size  int
	max   int
	free  []entryInt64ToInt64
}

// NewMapInt64ToInt64 creates a new MapInt64ToInt64 with at least a size of size.
func NewMapInt64ToInt64(size int) *MapInt64ToInt64 {
	shift := nextShiftPowerOfTwo(size)
	return &MapInt64ToInt64{
		mask:  ^(^0 << shift),
		max:   maxFill(1 << shift),
		slots: make([]*entryInt64ToInt64, 1<<shift),
	}
}

// Get looks up a key k returns its value. 0 if not found.
func (h *MapInt64ToInt64) Get(k int64) int64 {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v
		}
	}
	return 0
}

// Visit calls a given function fn for every key/value pair in the map.
func (h *MapInt64ToInt64) Visit(fn func(int64, int64)) {
	for _, e := range h.slots {
		for ; e != nil; e = e.next {
			fn(e.k, e.v)
		}
	}
}

// Add increments the value associated with key k by the value of v.
// It creates a new key/value pair k/v in the map if the key does not exist.
func (h *MapInt64ToInt64) Add(k int64, v int64) {
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
func (h *MapInt64ToInt64) Inc(k int64) {
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

func (h *MapInt64ToInt64) rehash() {
	ns := len(h.slots) << 1
	nslots := make([]*entryInt64ToInt64, ns)
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

func (h *MapInt64ToInt64) alloc(k int64, v int64) *entryInt64ToInt64 {
	if len(h.free) == 0 {
		h.free = make([]entryInt64ToInt64, 256)
	}
	h.size++
	x := &h.free[0]
	x.k = k
	x.v = v
	h.free = h.free[1:]
	return x
}
