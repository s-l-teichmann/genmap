// THIS IS A MACHINE GENERATED FILE!
// BE CAREFUL WITH EDITING BY HAND.
//
// This is Free Software covered by the terms of the MIT license.
// See LICENSE file for details.
// (c) 2016 by Sascha L. Teichmann.
// See the full list of contributors in the CONTRIBUTORS file.

package genmap

type entryInt8ToInt struct {
	k    int8
	v    int
	next *entryInt8ToInt
}

// MapInt8ToInt implements a hash map from int8 to int.
type MapInt8ToInt struct {
	mask     int
	slots    []*entryInt8ToInt
	used     int
	size     int
	max      int
	freelist *entryInt8ToInt
}

// NewMapInt8ToInt creates a new MapInt8ToInt with at least a size of size.
func NewMapInt8ToInt(size int) *MapInt8ToInt {
	shift := nextShiftPowerOfTwo(size)
	return &MapInt8ToInt{
		mask:  ^(^0 << shift),
		max:   maxFill(1 << shift),
		slots: make([]*entryInt8ToInt, 1<<shift),
	}
}

// Size returns the current size of the map.
func (h *MapInt8ToInt) Size() int {
	return h.size
}

// Get looks up a key k and returns its value. 0 if not found.
func (h *MapInt8ToInt) Get(k int8) int {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v
		}
	}
	return 0
}

// Contains looks up a key k and returns true if it is found else false.
func (h *MapInt8ToInt) Contains(k int8) bool {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return true
		}
	}
	return false
}

// Modify looks up a key k and calls function fn with a pointer to its value.
func (h *MapInt8ToInt) Modify(k int8, fn func(v *int)) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			fn(&e.v)
			return
		}
	}
}

// Find looks up a key k and returns its value and true.
// 0 and false if not found.
func (h *MapInt8ToInt) Find(k int8) (int, bool) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v, true
		}
	}
	return 0, false
}

// Visit calls a given function fn for every key/value pair in the map.
func (h *MapInt8ToInt) Visit(fn func(int8, int)) {
	for _, e := range h.slots {
		for ; e != nil; e = e.next {
			fn(e.k, e.v)
		}
	}
}

// Add increments the value associated with key k by the value of v.
// It creates a new key/value pair k/v in the map if the key does not exist.
func (h *MapInt8ToInt) Add(k int8, v int) {
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
func (h *MapInt8ToInt) Put(k int8, v int) {
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
func (h *MapInt8ToInt) Remove(k int8) {
	p := &h.slots[int(k)&h.mask]
	var parent *entryInt8ToInt
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
func (h *MapInt8ToInt) Clear() {
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
func (h *MapInt8ToInt) Inc(k int8) {
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

func (h *MapInt8ToInt) rehash() {
	ns := len(h.slots) << 1
	nslots := make([]*entryInt8ToInt, ns)
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

func (h *MapInt8ToInt) free(entry *entryInt8ToInt) {
	entry.next = h.freelist
	h.freelist = entry
	h.size--
}

func (h *MapInt8ToInt) alloc(k int8, v int) *entryInt8ToInt {
	if h.freelist == nil {
		entries := make([]entryInt8ToInt, 256)
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

type entryInt8ToInt8 struct {
	k    int8
	v    int8
	next *entryInt8ToInt8
}

// MapInt8ToInt8 implements a hash map from int8 to int8.
type MapInt8ToInt8 struct {
	mask     int
	slots    []*entryInt8ToInt8
	used     int
	size     int
	max      int
	freelist *entryInt8ToInt8
}

// NewMapInt8ToInt8 creates a new MapInt8ToInt8 with at least a size of size.
func NewMapInt8ToInt8(size int) *MapInt8ToInt8 {
	shift := nextShiftPowerOfTwo(size)
	return &MapInt8ToInt8{
		mask:  ^(^0 << shift),
		max:   maxFill(1 << shift),
		slots: make([]*entryInt8ToInt8, 1<<shift),
	}
}

// Size returns the current size of the map.
func (h *MapInt8ToInt8) Size() int {
	return h.size
}

// Get looks up a key k and returns its value. 0 if not found.
func (h *MapInt8ToInt8) Get(k int8) int8 {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v
		}
	}
	return 0
}

// Contains looks up a key k and returns true if it is found else false.
func (h *MapInt8ToInt8) Contains(k int8) bool {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return true
		}
	}
	return false
}

// Modify looks up a key k and calls function fn with a pointer to its value.
func (h *MapInt8ToInt8) Modify(k int8, fn func(v *int8)) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			fn(&e.v)
			return
		}
	}
}

// Find looks up a key k and returns its value and true.
// 0 and false if not found.
func (h *MapInt8ToInt8) Find(k int8) (int8, bool) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v, true
		}
	}
	return 0, false
}

// Visit calls a given function fn for every key/value pair in the map.
func (h *MapInt8ToInt8) Visit(fn func(int8, int8)) {
	for _, e := range h.slots {
		for ; e != nil; e = e.next {
			fn(e.k, e.v)
		}
	}
}

// Add increments the value associated with key k by the value of v.
// It creates a new key/value pair k/v in the map if the key does not exist.
func (h *MapInt8ToInt8) Add(k int8, v int8) {
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
func (h *MapInt8ToInt8) Put(k int8, v int8) {
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
func (h *MapInt8ToInt8) Remove(k int8) {
	p := &h.slots[int(k)&h.mask]
	var parent *entryInt8ToInt8
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
func (h *MapInt8ToInt8) Clear() {
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
func (h *MapInt8ToInt8) Inc(k int8) {
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

func (h *MapInt8ToInt8) rehash() {
	ns := len(h.slots) << 1
	nslots := make([]*entryInt8ToInt8, ns)
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

func (h *MapInt8ToInt8) free(entry *entryInt8ToInt8) {
	entry.next = h.freelist
	h.freelist = entry
	h.size--
}

func (h *MapInt8ToInt8) alloc(k int8, v int8) *entryInt8ToInt8 {
	if h.freelist == nil {
		entries := make([]entryInt8ToInt8, 256)
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

type entryInt8ToInt16 struct {
	k    int8
	v    int16
	next *entryInt8ToInt16
}

// MapInt8ToInt16 implements a hash map from int8 to int16.
type MapInt8ToInt16 struct {
	mask     int
	slots    []*entryInt8ToInt16
	used     int
	size     int
	max      int
	freelist *entryInt8ToInt16
}

// NewMapInt8ToInt16 creates a new MapInt8ToInt16 with at least a size of size.
func NewMapInt8ToInt16(size int) *MapInt8ToInt16 {
	shift := nextShiftPowerOfTwo(size)
	return &MapInt8ToInt16{
		mask:  ^(^0 << shift),
		max:   maxFill(1 << shift),
		slots: make([]*entryInt8ToInt16, 1<<shift),
	}
}

// Size returns the current size of the map.
func (h *MapInt8ToInt16) Size() int {
	return h.size
}

// Get looks up a key k and returns its value. 0 if not found.
func (h *MapInt8ToInt16) Get(k int8) int16 {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v
		}
	}
	return 0
}

// Contains looks up a key k and returns true if it is found else false.
func (h *MapInt8ToInt16) Contains(k int8) bool {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return true
		}
	}
	return false
}

// Modify looks up a key k and calls function fn with a pointer to its value.
func (h *MapInt8ToInt16) Modify(k int8, fn func(v *int16)) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			fn(&e.v)
			return
		}
	}
}

// Find looks up a key k and returns its value and true.
// 0 and false if not found.
func (h *MapInt8ToInt16) Find(k int8) (int16, bool) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v, true
		}
	}
	return 0, false
}

// Visit calls a given function fn for every key/value pair in the map.
func (h *MapInt8ToInt16) Visit(fn func(int8, int16)) {
	for _, e := range h.slots {
		for ; e != nil; e = e.next {
			fn(e.k, e.v)
		}
	}
}

// Add increments the value associated with key k by the value of v.
// It creates a new key/value pair k/v in the map if the key does not exist.
func (h *MapInt8ToInt16) Add(k int8, v int16) {
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
func (h *MapInt8ToInt16) Put(k int8, v int16) {
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
func (h *MapInt8ToInt16) Remove(k int8) {
	p := &h.slots[int(k)&h.mask]
	var parent *entryInt8ToInt16
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
func (h *MapInt8ToInt16) Clear() {
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
func (h *MapInt8ToInt16) Inc(k int8) {
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

func (h *MapInt8ToInt16) rehash() {
	ns := len(h.slots) << 1
	nslots := make([]*entryInt8ToInt16, ns)
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

func (h *MapInt8ToInt16) free(entry *entryInt8ToInt16) {
	entry.next = h.freelist
	h.freelist = entry
	h.size--
}

func (h *MapInt8ToInt16) alloc(k int8, v int16) *entryInt8ToInt16 {
	if h.freelist == nil {
		entries := make([]entryInt8ToInt16, 256)
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

type entryInt8ToInt32 struct {
	k    int8
	v    int32
	next *entryInt8ToInt32
}

// MapInt8ToInt32 implements a hash map from int8 to int32.
type MapInt8ToInt32 struct {
	mask     int
	slots    []*entryInt8ToInt32
	used     int
	size     int
	max      int
	freelist *entryInt8ToInt32
}

// NewMapInt8ToInt32 creates a new MapInt8ToInt32 with at least a size of size.
func NewMapInt8ToInt32(size int) *MapInt8ToInt32 {
	shift := nextShiftPowerOfTwo(size)
	return &MapInt8ToInt32{
		mask:  ^(^0 << shift),
		max:   maxFill(1 << shift),
		slots: make([]*entryInt8ToInt32, 1<<shift),
	}
}

// Size returns the current size of the map.
func (h *MapInt8ToInt32) Size() int {
	return h.size
}

// Get looks up a key k and returns its value. 0 if not found.
func (h *MapInt8ToInt32) Get(k int8) int32 {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v
		}
	}
	return 0
}

// Contains looks up a key k and returns true if it is found else false.
func (h *MapInt8ToInt32) Contains(k int8) bool {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return true
		}
	}
	return false
}

// Modify looks up a key k and calls function fn with a pointer to its value.
func (h *MapInt8ToInt32) Modify(k int8, fn func(v *int32)) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			fn(&e.v)
			return
		}
	}
}

// Find looks up a key k and returns its value and true.
// 0 and false if not found.
func (h *MapInt8ToInt32) Find(k int8) (int32, bool) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v, true
		}
	}
	return 0, false
}

// Visit calls a given function fn for every key/value pair in the map.
func (h *MapInt8ToInt32) Visit(fn func(int8, int32)) {
	for _, e := range h.slots {
		for ; e != nil; e = e.next {
			fn(e.k, e.v)
		}
	}
}

// Add increments the value associated with key k by the value of v.
// It creates a new key/value pair k/v in the map if the key does not exist.
func (h *MapInt8ToInt32) Add(k int8, v int32) {
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
func (h *MapInt8ToInt32) Put(k int8, v int32) {
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
func (h *MapInt8ToInt32) Remove(k int8) {
	p := &h.slots[int(k)&h.mask]
	var parent *entryInt8ToInt32
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
func (h *MapInt8ToInt32) Clear() {
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
func (h *MapInt8ToInt32) Inc(k int8) {
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

func (h *MapInt8ToInt32) rehash() {
	ns := len(h.slots) << 1
	nslots := make([]*entryInt8ToInt32, ns)
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

func (h *MapInt8ToInt32) free(entry *entryInt8ToInt32) {
	entry.next = h.freelist
	h.freelist = entry
	h.size--
}

func (h *MapInt8ToInt32) alloc(k int8, v int32) *entryInt8ToInt32 {
	if h.freelist == nil {
		entries := make([]entryInt8ToInt32, 256)
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

type entryInt8ToInt64 struct {
	k    int8
	v    int64
	next *entryInt8ToInt64
}

// MapInt8ToInt64 implements a hash map from int8 to int64.
type MapInt8ToInt64 struct {
	mask     int
	slots    []*entryInt8ToInt64
	used     int
	size     int
	max      int
	freelist *entryInt8ToInt64
}

// NewMapInt8ToInt64 creates a new MapInt8ToInt64 with at least a size of size.
func NewMapInt8ToInt64(size int) *MapInt8ToInt64 {
	shift := nextShiftPowerOfTwo(size)
	return &MapInt8ToInt64{
		mask:  ^(^0 << shift),
		max:   maxFill(1 << shift),
		slots: make([]*entryInt8ToInt64, 1<<shift),
	}
}

// Size returns the current size of the map.
func (h *MapInt8ToInt64) Size() int {
	return h.size
}

// Get looks up a key k and returns its value. 0 if not found.
func (h *MapInt8ToInt64) Get(k int8) int64 {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v
		}
	}
	return 0
}

// Contains looks up a key k and returns true if it is found else false.
func (h *MapInt8ToInt64) Contains(k int8) bool {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return true
		}
	}
	return false
}

// Modify looks up a key k and calls function fn with a pointer to its value.
func (h *MapInt8ToInt64) Modify(k int8, fn func(v *int64)) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			fn(&e.v)
			return
		}
	}
}

// Find looks up a key k and returns its value and true.
// 0 and false if not found.
func (h *MapInt8ToInt64) Find(k int8) (int64, bool) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v, true
		}
	}
	return 0, false
}

// Visit calls a given function fn for every key/value pair in the map.
func (h *MapInt8ToInt64) Visit(fn func(int8, int64)) {
	for _, e := range h.slots {
		for ; e != nil; e = e.next {
			fn(e.k, e.v)
		}
	}
}

// Add increments the value associated with key k by the value of v.
// It creates a new key/value pair k/v in the map if the key does not exist.
func (h *MapInt8ToInt64) Add(k int8, v int64) {
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
func (h *MapInt8ToInt64) Put(k int8, v int64) {
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
func (h *MapInt8ToInt64) Remove(k int8) {
	p := &h.slots[int(k)&h.mask]
	var parent *entryInt8ToInt64
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
func (h *MapInt8ToInt64) Clear() {
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
func (h *MapInt8ToInt64) Inc(k int8) {
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

func (h *MapInt8ToInt64) rehash() {
	ns := len(h.slots) << 1
	nslots := make([]*entryInt8ToInt64, ns)
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

func (h *MapInt8ToInt64) free(entry *entryInt8ToInt64) {
	entry.next = h.freelist
	h.freelist = entry
	h.size--
}

func (h *MapInt8ToInt64) alloc(k int8, v int64) *entryInt8ToInt64 {
	if h.freelist == nil {
		entries := make([]entryInt8ToInt64, 256)
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

// Size returns the current size of the map.
func (h *MapInt8ToUint) Size() int {
	return h.size
}

// Get looks up a key k and returns its value. 0 if not found.
func (h *MapInt8ToUint) Get(k int8) uint {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v
		}
	}
	return 0
}

// Contains looks up a key k and returns true if it is found else false.
func (h *MapInt8ToUint) Contains(k int8) bool {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return true
		}
	}
	return false
}

// Modify looks up a key k and calls function fn with a pointer to its value.
func (h *MapInt8ToUint) Modify(k int8, fn func(v *uint)) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			fn(&e.v)
			return
		}
	}
}

// Find looks up a key k and returns its value and true.
// 0 and false if not found.
func (h *MapInt8ToUint) Find(k int8) (uint, bool) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v, true
		}
	}
	return 0, false
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
func (h *MapInt8ToUint) Clear() {
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
	h.freelist = x.next
	x.next = nil
	x.k = k
	x.v = v
	return x
}

type entryInt8ToUint8 struct {
	k    int8
	v    uint8
	next *entryInt8ToUint8
}

// MapInt8ToUint8 implements a hash map from int8 to uint8.
type MapInt8ToUint8 struct {
	mask     int
	slots    []*entryInt8ToUint8
	used     int
	size     int
	max      int
	freelist *entryInt8ToUint8
}

// NewMapInt8ToUint8 creates a new MapInt8ToUint8 with at least a size of size.
func NewMapInt8ToUint8(size int) *MapInt8ToUint8 {
	shift := nextShiftPowerOfTwo(size)
	return &MapInt8ToUint8{
		mask:  ^(^0 << shift),
		max:   maxFill(1 << shift),
		slots: make([]*entryInt8ToUint8, 1<<shift),
	}
}

// Size returns the current size of the map.
func (h *MapInt8ToUint8) Size() int {
	return h.size
}

// Get looks up a key k and returns its value. 0 if not found.
func (h *MapInt8ToUint8) Get(k int8) uint8 {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v
		}
	}
	return 0
}

// Contains looks up a key k and returns true if it is found else false.
func (h *MapInt8ToUint8) Contains(k int8) bool {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return true
		}
	}
	return false
}

// Modify looks up a key k and calls function fn with a pointer to its value.
func (h *MapInt8ToUint8) Modify(k int8, fn func(v *uint8)) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			fn(&e.v)
			return
		}
	}
}

// Find looks up a key k and returns its value and true.
// 0 and false if not found.
func (h *MapInt8ToUint8) Find(k int8) (uint8, bool) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v, true
		}
	}
	return 0, false
}

// Visit calls a given function fn for every key/value pair in the map.
func (h *MapInt8ToUint8) Visit(fn func(int8, uint8)) {
	for _, e := range h.slots {
		for ; e != nil; e = e.next {
			fn(e.k, e.v)
		}
	}
}

// Add increments the value associated with key k by the value of v.
// It creates a new key/value pair k/v in the map if the key does not exist.
func (h *MapInt8ToUint8) Add(k int8, v uint8) {
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
func (h *MapInt8ToUint8) Put(k int8, v uint8) {
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
func (h *MapInt8ToUint8) Remove(k int8) {
	p := &h.slots[int(k)&h.mask]
	var parent *entryInt8ToUint8
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
func (h *MapInt8ToUint8) Clear() {
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
func (h *MapInt8ToUint8) Inc(k int8) {
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

func (h *MapInt8ToUint8) rehash() {
	ns := len(h.slots) << 1
	nslots := make([]*entryInt8ToUint8, ns)
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

func (h *MapInt8ToUint8) free(entry *entryInt8ToUint8) {
	entry.next = h.freelist
	h.freelist = entry
	h.size--
}

func (h *MapInt8ToUint8) alloc(k int8, v uint8) *entryInt8ToUint8 {
	if h.freelist == nil {
		entries := make([]entryInt8ToUint8, 256)
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

type entryInt8ToUint16 struct {
	k    int8
	v    uint16
	next *entryInt8ToUint16
}

// MapInt8ToUint16 implements a hash map from int8 to uint16.
type MapInt8ToUint16 struct {
	mask     int
	slots    []*entryInt8ToUint16
	used     int
	size     int
	max      int
	freelist *entryInt8ToUint16
}

// NewMapInt8ToUint16 creates a new MapInt8ToUint16 with at least a size of size.
func NewMapInt8ToUint16(size int) *MapInt8ToUint16 {
	shift := nextShiftPowerOfTwo(size)
	return &MapInt8ToUint16{
		mask:  ^(^0 << shift),
		max:   maxFill(1 << shift),
		slots: make([]*entryInt8ToUint16, 1<<shift),
	}
}

// Size returns the current size of the map.
func (h *MapInt8ToUint16) Size() int {
	return h.size
}

// Get looks up a key k and returns its value. 0 if not found.
func (h *MapInt8ToUint16) Get(k int8) uint16 {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v
		}
	}
	return 0
}

// Contains looks up a key k and returns true if it is found else false.
func (h *MapInt8ToUint16) Contains(k int8) bool {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return true
		}
	}
	return false
}

// Modify looks up a key k and calls function fn with a pointer to its value.
func (h *MapInt8ToUint16) Modify(k int8, fn func(v *uint16)) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			fn(&e.v)
			return
		}
	}
}

// Find looks up a key k and returns its value and true.
// 0 and false if not found.
func (h *MapInt8ToUint16) Find(k int8) (uint16, bool) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v, true
		}
	}
	return 0, false
}

// Visit calls a given function fn for every key/value pair in the map.
func (h *MapInt8ToUint16) Visit(fn func(int8, uint16)) {
	for _, e := range h.slots {
		for ; e != nil; e = e.next {
			fn(e.k, e.v)
		}
	}
}

// Add increments the value associated with key k by the value of v.
// It creates a new key/value pair k/v in the map if the key does not exist.
func (h *MapInt8ToUint16) Add(k int8, v uint16) {
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
func (h *MapInt8ToUint16) Put(k int8, v uint16) {
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
func (h *MapInt8ToUint16) Remove(k int8) {
	p := &h.slots[int(k)&h.mask]
	var parent *entryInt8ToUint16
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
func (h *MapInt8ToUint16) Clear() {
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
func (h *MapInt8ToUint16) Inc(k int8) {
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

func (h *MapInt8ToUint16) rehash() {
	ns := len(h.slots) << 1
	nslots := make([]*entryInt8ToUint16, ns)
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

func (h *MapInt8ToUint16) free(entry *entryInt8ToUint16) {
	entry.next = h.freelist
	h.freelist = entry
	h.size--
}

func (h *MapInt8ToUint16) alloc(k int8, v uint16) *entryInt8ToUint16 {
	if h.freelist == nil {
		entries := make([]entryInt8ToUint16, 256)
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

type entryInt8ToUint32 struct {
	k    int8
	v    uint32
	next *entryInt8ToUint32
}

// MapInt8ToUint32 implements a hash map from int8 to uint32.
type MapInt8ToUint32 struct {
	mask     int
	slots    []*entryInt8ToUint32
	used     int
	size     int
	max      int
	freelist *entryInt8ToUint32
}

// NewMapInt8ToUint32 creates a new MapInt8ToUint32 with at least a size of size.
func NewMapInt8ToUint32(size int) *MapInt8ToUint32 {
	shift := nextShiftPowerOfTwo(size)
	return &MapInt8ToUint32{
		mask:  ^(^0 << shift),
		max:   maxFill(1 << shift),
		slots: make([]*entryInt8ToUint32, 1<<shift),
	}
}

// Size returns the current size of the map.
func (h *MapInt8ToUint32) Size() int {
	return h.size
}

// Get looks up a key k and returns its value. 0 if not found.
func (h *MapInt8ToUint32) Get(k int8) uint32 {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v
		}
	}
	return 0
}

// Contains looks up a key k and returns true if it is found else false.
func (h *MapInt8ToUint32) Contains(k int8) bool {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return true
		}
	}
	return false
}

// Modify looks up a key k and calls function fn with a pointer to its value.
func (h *MapInt8ToUint32) Modify(k int8, fn func(v *uint32)) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			fn(&e.v)
			return
		}
	}
}

// Find looks up a key k and returns its value and true.
// 0 and false if not found.
func (h *MapInt8ToUint32) Find(k int8) (uint32, bool) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v, true
		}
	}
	return 0, false
}

// Visit calls a given function fn for every key/value pair in the map.
func (h *MapInt8ToUint32) Visit(fn func(int8, uint32)) {
	for _, e := range h.slots {
		for ; e != nil; e = e.next {
			fn(e.k, e.v)
		}
	}
}

// Add increments the value associated with key k by the value of v.
// It creates a new key/value pair k/v in the map if the key does not exist.
func (h *MapInt8ToUint32) Add(k int8, v uint32) {
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
func (h *MapInt8ToUint32) Put(k int8, v uint32) {
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
func (h *MapInt8ToUint32) Remove(k int8) {
	p := &h.slots[int(k)&h.mask]
	var parent *entryInt8ToUint32
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
func (h *MapInt8ToUint32) Clear() {
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
func (h *MapInt8ToUint32) Inc(k int8) {
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

func (h *MapInt8ToUint32) rehash() {
	ns := len(h.slots) << 1
	nslots := make([]*entryInt8ToUint32, ns)
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

func (h *MapInt8ToUint32) free(entry *entryInt8ToUint32) {
	entry.next = h.freelist
	h.freelist = entry
	h.size--
}

func (h *MapInt8ToUint32) alloc(k int8, v uint32) *entryInt8ToUint32 {
	if h.freelist == nil {
		entries := make([]entryInt8ToUint32, 256)
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

type entryInt8ToUint64 struct {
	k    int8
	v    uint64
	next *entryInt8ToUint64
}

// MapInt8ToUint64 implements a hash map from int8 to uint64.
type MapInt8ToUint64 struct {
	mask     int
	slots    []*entryInt8ToUint64
	used     int
	size     int
	max      int
	freelist *entryInt8ToUint64
}

// NewMapInt8ToUint64 creates a new MapInt8ToUint64 with at least a size of size.
func NewMapInt8ToUint64(size int) *MapInt8ToUint64 {
	shift := nextShiftPowerOfTwo(size)
	return &MapInt8ToUint64{
		mask:  ^(^0 << shift),
		max:   maxFill(1 << shift),
		slots: make([]*entryInt8ToUint64, 1<<shift),
	}
}

// Size returns the current size of the map.
func (h *MapInt8ToUint64) Size() int {
	return h.size
}

// Get looks up a key k and returns its value. 0 if not found.
func (h *MapInt8ToUint64) Get(k int8) uint64 {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v
		}
	}
	return 0
}

// Contains looks up a key k and returns true if it is found else false.
func (h *MapInt8ToUint64) Contains(k int8) bool {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return true
		}
	}
	return false
}

// Modify looks up a key k and calls function fn with a pointer to its value.
func (h *MapInt8ToUint64) Modify(k int8, fn func(v *uint64)) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			fn(&e.v)
			return
		}
	}
}

// Find looks up a key k and returns its value and true.
// 0 and false if not found.
func (h *MapInt8ToUint64) Find(k int8) (uint64, bool) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v, true
		}
	}
	return 0, false
}

// Visit calls a given function fn for every key/value pair in the map.
func (h *MapInt8ToUint64) Visit(fn func(int8, uint64)) {
	for _, e := range h.slots {
		for ; e != nil; e = e.next {
			fn(e.k, e.v)
		}
	}
}

// Add increments the value associated with key k by the value of v.
// It creates a new key/value pair k/v in the map if the key does not exist.
func (h *MapInt8ToUint64) Add(k int8, v uint64) {
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
func (h *MapInt8ToUint64) Put(k int8, v uint64) {
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
func (h *MapInt8ToUint64) Remove(k int8) {
	p := &h.slots[int(k)&h.mask]
	var parent *entryInt8ToUint64
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
func (h *MapInt8ToUint64) Clear() {
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
func (h *MapInt8ToUint64) Inc(k int8) {
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

func (h *MapInt8ToUint64) rehash() {
	ns := len(h.slots) << 1
	nslots := make([]*entryInt8ToUint64, ns)
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

func (h *MapInt8ToUint64) free(entry *entryInt8ToUint64) {
	entry.next = h.freelist
	h.freelist = entry
	h.size--
}

func (h *MapInt8ToUint64) alloc(k int8, v uint64) *entryInt8ToUint64 {
	if h.freelist == nil {
		entries := make([]entryInt8ToUint64, 256)
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
