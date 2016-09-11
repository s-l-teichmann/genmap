// THIS IS A MACHINE GENERATED FILE!
// BE CAREFUL WITH EDITING BY HAND.
//
// This is Free Software covered by the terms of the MIT license.
// See LICENSE file for details.
// (c) 2016 by the contributors.
// See CONTRIBUTORS for a full list.

package genmap

type entryInt64ToInt32 struct {
	k    int64
	v    int32
	next *entryInt64ToInt32
}

// MapInt64ToInt32 implements a hash map from int64 to int32.
type MapInt64ToInt32 struct {
	mask  int
	slots []*entryInt64ToInt32
	used  int
	size  int
	max   int
	free  []entryInt64ToInt32
}

// NewMapInt64ToInt32 creates a new MapInt64ToInt32 with at least a size of size.
func NewMapInt64ToInt32(size int) *MapInt64ToInt32 {
	shift := nextShiftPowerOfTwo(size)
	return &MapInt64ToInt32{
		mask:  ^(^0 << shift),
		max:   maxFill(1 << shift),
		slots: make([]*entryInt64ToInt32, 1<<shift),
	}
}

// Get looks up a key k returns its value. 0 if not found.
func (h *MapInt64ToInt32) Get(k int64) int32 {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v
		}
	}
	return 0
}

// Visit calls a given function fn for every key/value pair in the map.
func (h *MapInt64ToInt32) Visit(fn func(int64, int32)) {
	for _, e := range h.slots {
		for ; e != nil; e = e.next {
			fn(e.k, e.v)
		}
	}
}

// Add increments the value associated with key k by the value of v.
// It creates a new key/value pair k/v in the map if the key does not exist.
func (h *MapInt64ToInt32) Add(k int64, v int32) {
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
func (h *MapInt64ToInt32) Inc(k int64) {
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

func (h *MapInt64ToInt32) rehash() {
	ns := len(h.slots) << 1
	nslots := make([]*entryInt64ToInt32, ns)
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

func (h *MapInt64ToInt32) alloc(k int64, v int32) *entryInt64ToInt32 {
	if len(h.free) == 0 {
		h.free = make([]entryInt64ToInt32, 256)
	}
	h.size++
	x := &h.free[0]
	x.k = k
	x.v = v
	h.free = h.free[1:]
	return x
}
