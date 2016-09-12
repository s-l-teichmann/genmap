// THIS IS A MACHINE GENERATED FILE!
// BE CAREFUL WITH EDITING BY HAND.
//
// This is Free Software covered by the terms of the MIT license.
// See LICENSE file for details.
// (c) 2016 by Sascha L. Teichmann.
// See the full list of contributors in the CONTRIBUTORS file.

package genmap

type entryIntToInt struct {
	k    int
	v    int
	next *entryIntToInt
}

// MapIntToInt implements a hash map from int to int.
type MapIntToInt struct {
	mask     int
	slots    []*entryIntToInt
	used     int
	size     int
	max      int
	freelist *entryIntToInt
}

// NewMapIntToInt creates a new MapIntToInt with at least a size of size.
func NewMapIntToInt(size int) *MapIntToInt {
	shift := nextShiftPowerOfTwo(size)
	return &MapIntToInt{
		mask:  ^(^0 << shift),
		max:   maxFill(1 << shift),
		slots: make([]*entryIntToInt, 1<<shift),
	}
}

// Size returns the current size of the map.
func (h *MapIntToInt) Size() int {
	return h.size
}

// Get looks up a key k and returns its value. 0 if not found.
func (h *MapIntToInt) Get(k int) int {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v
		}
	}
	return 0
}

// Contains looks up a key k and returns true if it is found else false.
func (h *MapIntToInt) Contains(k int) bool {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return true
		}
	}
	return false
}

// Modify looks up a key k and calls function fn with a pointer to its value.
func (h *MapIntToInt) Modify(k int, fn func(v *int)) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			fn(&e.v)
			return
		}
	}
}

// Find looks up a key k and returns its value and true.
// 0 and false if not found.
func (h *MapIntToInt) Find(k int) (int, bool) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v, true
		}
	}
	return 0, false
}

// Visit calls a given function fn for every key/value pair in the map.
func (h *MapIntToInt) Visit(fn func(int, int)) {
	for _, e := range h.slots {
		for ; e != nil; e = e.next {
			fn(e.k, e.v)
		}
	}
}

// Add increments the value associated with key k by the value of v.
// It creates a new key/value pair k/v in the map if the key does not exist.
func (h *MapIntToInt) Add(k int, v int) {
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
func (h *MapIntToInt) Put(k int, v int) {
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
func (h *MapIntToInt) Remove(k int) {
	p := &h.slots[int(k)&h.mask]
	var parent *entryIntToInt
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
func (h *MapIntToInt) Clear() {
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
func (h *MapIntToInt) Inc(k int) {
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

func (h *MapIntToInt) rehash() {
	ns := len(h.slots) << 1
	nslots := make([]*entryIntToInt, ns)
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

func (h *MapIntToInt) free(entry *entryIntToInt) {
	entry.next = h.freelist
	h.freelist = entry
	h.size--
}

func (h *MapIntToInt) alloc(k int, v int) *entryIntToInt {
	if h.freelist == nil {
		entries := make([]entryIntToInt, 256)
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

type entryIntToInt8 struct {
	k    int
	v    int8
	next *entryIntToInt8
}

// MapIntToInt8 implements a hash map from int to int8.
type MapIntToInt8 struct {
	mask     int
	slots    []*entryIntToInt8
	used     int
	size     int
	max      int
	freelist *entryIntToInt8
}

// NewMapIntToInt8 creates a new MapIntToInt8 with at least a size of size.
func NewMapIntToInt8(size int) *MapIntToInt8 {
	shift := nextShiftPowerOfTwo(size)
	return &MapIntToInt8{
		mask:  ^(^0 << shift),
		max:   maxFill(1 << shift),
		slots: make([]*entryIntToInt8, 1<<shift),
	}
}

// Size returns the current size of the map.
func (h *MapIntToInt8) Size() int {
	return h.size
}

// Get looks up a key k and returns its value. 0 if not found.
func (h *MapIntToInt8) Get(k int) int8 {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v
		}
	}
	return 0
}

// Contains looks up a key k and returns true if it is found else false.
func (h *MapIntToInt8) Contains(k int) bool {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return true
		}
	}
	return false
}

// Modify looks up a key k and calls function fn with a pointer to its value.
func (h *MapIntToInt8) Modify(k int, fn func(v *int8)) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			fn(&e.v)
			return
		}
	}
}

// Find looks up a key k and returns its value and true.
// 0 and false if not found.
func (h *MapIntToInt8) Find(k int) (int8, bool) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v, true
		}
	}
	return 0, false
}

// Visit calls a given function fn for every key/value pair in the map.
func (h *MapIntToInt8) Visit(fn func(int, int8)) {
	for _, e := range h.slots {
		for ; e != nil; e = e.next {
			fn(e.k, e.v)
		}
	}
}

// Add increments the value associated with key k by the value of v.
// It creates a new key/value pair k/v in the map if the key does not exist.
func (h *MapIntToInt8) Add(k int, v int8) {
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
func (h *MapIntToInt8) Put(k int, v int8) {
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
func (h *MapIntToInt8) Remove(k int) {
	p := &h.slots[int(k)&h.mask]
	var parent *entryIntToInt8
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
func (h *MapIntToInt8) Clear() {
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
func (h *MapIntToInt8) Inc(k int) {
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

func (h *MapIntToInt8) rehash() {
	ns := len(h.slots) << 1
	nslots := make([]*entryIntToInt8, ns)
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

func (h *MapIntToInt8) free(entry *entryIntToInt8) {
	entry.next = h.freelist
	h.freelist = entry
	h.size--
}

func (h *MapIntToInt8) alloc(k int, v int8) *entryIntToInt8 {
	if h.freelist == nil {
		entries := make([]entryIntToInt8, 256)
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

type entryIntToInt16 struct {
	k    int
	v    int16
	next *entryIntToInt16
}

// MapIntToInt16 implements a hash map from int to int16.
type MapIntToInt16 struct {
	mask     int
	slots    []*entryIntToInt16
	used     int
	size     int
	max      int
	freelist *entryIntToInt16
}

// NewMapIntToInt16 creates a new MapIntToInt16 with at least a size of size.
func NewMapIntToInt16(size int) *MapIntToInt16 {
	shift := nextShiftPowerOfTwo(size)
	return &MapIntToInt16{
		mask:  ^(^0 << shift),
		max:   maxFill(1 << shift),
		slots: make([]*entryIntToInt16, 1<<shift),
	}
}

// Size returns the current size of the map.
func (h *MapIntToInt16) Size() int {
	return h.size
}

// Get looks up a key k and returns its value. 0 if not found.
func (h *MapIntToInt16) Get(k int) int16 {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v
		}
	}
	return 0
}

// Contains looks up a key k and returns true if it is found else false.
func (h *MapIntToInt16) Contains(k int) bool {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return true
		}
	}
	return false
}

// Modify looks up a key k and calls function fn with a pointer to its value.
func (h *MapIntToInt16) Modify(k int, fn func(v *int16)) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			fn(&e.v)
			return
		}
	}
}

// Find looks up a key k and returns its value and true.
// 0 and false if not found.
func (h *MapIntToInt16) Find(k int) (int16, bool) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v, true
		}
	}
	return 0, false
}

// Visit calls a given function fn for every key/value pair in the map.
func (h *MapIntToInt16) Visit(fn func(int, int16)) {
	for _, e := range h.slots {
		for ; e != nil; e = e.next {
			fn(e.k, e.v)
		}
	}
}

// Add increments the value associated with key k by the value of v.
// It creates a new key/value pair k/v in the map if the key does not exist.
func (h *MapIntToInt16) Add(k int, v int16) {
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
func (h *MapIntToInt16) Put(k int, v int16) {
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
func (h *MapIntToInt16) Remove(k int) {
	p := &h.slots[int(k)&h.mask]
	var parent *entryIntToInt16
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
func (h *MapIntToInt16) Clear() {
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
func (h *MapIntToInt16) Inc(k int) {
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

func (h *MapIntToInt16) rehash() {
	ns := len(h.slots) << 1
	nslots := make([]*entryIntToInt16, ns)
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

func (h *MapIntToInt16) free(entry *entryIntToInt16) {
	entry.next = h.freelist
	h.freelist = entry
	h.size--
}

func (h *MapIntToInt16) alloc(k int, v int16) *entryIntToInt16 {
	if h.freelist == nil {
		entries := make([]entryIntToInt16, 256)
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

type entryIntToInt32 struct {
	k    int
	v    int32
	next *entryIntToInt32
}

// MapIntToInt32 implements a hash map from int to int32.
type MapIntToInt32 struct {
	mask     int
	slots    []*entryIntToInt32
	used     int
	size     int
	max      int
	freelist *entryIntToInt32
}

// NewMapIntToInt32 creates a new MapIntToInt32 with at least a size of size.
func NewMapIntToInt32(size int) *MapIntToInt32 {
	shift := nextShiftPowerOfTwo(size)
	return &MapIntToInt32{
		mask:  ^(^0 << shift),
		max:   maxFill(1 << shift),
		slots: make([]*entryIntToInt32, 1<<shift),
	}
}

// Size returns the current size of the map.
func (h *MapIntToInt32) Size() int {
	return h.size
}

// Get looks up a key k and returns its value. 0 if not found.
func (h *MapIntToInt32) Get(k int) int32 {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v
		}
	}
	return 0
}

// Contains looks up a key k and returns true if it is found else false.
func (h *MapIntToInt32) Contains(k int) bool {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return true
		}
	}
	return false
}

// Modify looks up a key k and calls function fn with a pointer to its value.
func (h *MapIntToInt32) Modify(k int, fn func(v *int32)) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			fn(&e.v)
			return
		}
	}
}

// Find looks up a key k and returns its value and true.
// 0 and false if not found.
func (h *MapIntToInt32) Find(k int) (int32, bool) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v, true
		}
	}
	return 0, false
}

// Visit calls a given function fn for every key/value pair in the map.
func (h *MapIntToInt32) Visit(fn func(int, int32)) {
	for _, e := range h.slots {
		for ; e != nil; e = e.next {
			fn(e.k, e.v)
		}
	}
}

// Add increments the value associated with key k by the value of v.
// It creates a new key/value pair k/v in the map if the key does not exist.
func (h *MapIntToInt32) Add(k int, v int32) {
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
func (h *MapIntToInt32) Put(k int, v int32) {
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
func (h *MapIntToInt32) Remove(k int) {
	p := &h.slots[int(k)&h.mask]
	var parent *entryIntToInt32
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
func (h *MapIntToInt32) Clear() {
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
func (h *MapIntToInt32) Inc(k int) {
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

func (h *MapIntToInt32) rehash() {
	ns := len(h.slots) << 1
	nslots := make([]*entryIntToInt32, ns)
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

func (h *MapIntToInt32) free(entry *entryIntToInt32) {
	entry.next = h.freelist
	h.freelist = entry
	h.size--
}

func (h *MapIntToInt32) alloc(k int, v int32) *entryIntToInt32 {
	if h.freelist == nil {
		entries := make([]entryIntToInt32, 256)
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

type entryIntToInt64 struct {
	k    int
	v    int64
	next *entryIntToInt64
}

// MapIntToInt64 implements a hash map from int to int64.
type MapIntToInt64 struct {
	mask     int
	slots    []*entryIntToInt64
	used     int
	size     int
	max      int
	freelist *entryIntToInt64
}

// NewMapIntToInt64 creates a new MapIntToInt64 with at least a size of size.
func NewMapIntToInt64(size int) *MapIntToInt64 {
	shift := nextShiftPowerOfTwo(size)
	return &MapIntToInt64{
		mask:  ^(^0 << shift),
		max:   maxFill(1 << shift),
		slots: make([]*entryIntToInt64, 1<<shift),
	}
}

// Size returns the current size of the map.
func (h *MapIntToInt64) Size() int {
	return h.size
}

// Get looks up a key k and returns its value. 0 if not found.
func (h *MapIntToInt64) Get(k int) int64 {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v
		}
	}
	return 0
}

// Contains looks up a key k and returns true if it is found else false.
func (h *MapIntToInt64) Contains(k int) bool {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return true
		}
	}
	return false
}

// Modify looks up a key k and calls function fn with a pointer to its value.
func (h *MapIntToInt64) Modify(k int, fn func(v *int64)) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			fn(&e.v)
			return
		}
	}
}

// Find looks up a key k and returns its value and true.
// 0 and false if not found.
func (h *MapIntToInt64) Find(k int) (int64, bool) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v, true
		}
	}
	return 0, false
}

// Visit calls a given function fn for every key/value pair in the map.
func (h *MapIntToInt64) Visit(fn func(int, int64)) {
	for _, e := range h.slots {
		for ; e != nil; e = e.next {
			fn(e.k, e.v)
		}
	}
}

// Add increments the value associated with key k by the value of v.
// It creates a new key/value pair k/v in the map if the key does not exist.
func (h *MapIntToInt64) Add(k int, v int64) {
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
func (h *MapIntToInt64) Put(k int, v int64) {
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
func (h *MapIntToInt64) Remove(k int) {
	p := &h.slots[int(k)&h.mask]
	var parent *entryIntToInt64
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
func (h *MapIntToInt64) Clear() {
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
func (h *MapIntToInt64) Inc(k int) {
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

func (h *MapIntToInt64) rehash() {
	ns := len(h.slots) << 1
	nslots := make([]*entryIntToInt64, ns)
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

func (h *MapIntToInt64) free(entry *entryIntToInt64) {
	entry.next = h.freelist
	h.freelist = entry
	h.size--
}

func (h *MapIntToInt64) alloc(k int, v int64) *entryIntToInt64 {
	if h.freelist == nil {
		entries := make([]entryIntToInt64, 256)
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

type entryIntToUint struct {
	k    int
	v    uint
	next *entryIntToUint
}

// MapIntToUint implements a hash map from int to uint.
type MapIntToUint struct {
	mask     int
	slots    []*entryIntToUint
	used     int
	size     int
	max      int
	freelist *entryIntToUint
}

// NewMapIntToUint creates a new MapIntToUint with at least a size of size.
func NewMapIntToUint(size int) *MapIntToUint {
	shift := nextShiftPowerOfTwo(size)
	return &MapIntToUint{
		mask:  ^(^0 << shift),
		max:   maxFill(1 << shift),
		slots: make([]*entryIntToUint, 1<<shift),
	}
}

// Size returns the current size of the map.
func (h *MapIntToUint) Size() int {
	return h.size
}

// Get looks up a key k and returns its value. 0 if not found.
func (h *MapIntToUint) Get(k int) uint {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v
		}
	}
	return 0
}

// Contains looks up a key k and returns true if it is found else false.
func (h *MapIntToUint) Contains(k int) bool {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return true
		}
	}
	return false
}

// Modify looks up a key k and calls function fn with a pointer to its value.
func (h *MapIntToUint) Modify(k int, fn func(v *uint)) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			fn(&e.v)
			return
		}
	}
}

// Find looks up a key k and returns its value and true.
// 0 and false if not found.
func (h *MapIntToUint) Find(k int) (uint, bool) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v, true
		}
	}
	return 0, false
}

// Visit calls a given function fn for every key/value pair in the map.
func (h *MapIntToUint) Visit(fn func(int, uint)) {
	for _, e := range h.slots {
		for ; e != nil; e = e.next {
			fn(e.k, e.v)
		}
	}
}

// Add increments the value associated with key k by the value of v.
// It creates a new key/value pair k/v in the map if the key does not exist.
func (h *MapIntToUint) Add(k int, v uint) {
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
func (h *MapIntToUint) Put(k int, v uint) {
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
func (h *MapIntToUint) Remove(k int) {
	p := &h.slots[int(k)&h.mask]
	var parent *entryIntToUint
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
func (h *MapIntToUint) Clear() {
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
func (h *MapIntToUint) Inc(k int) {
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

func (h *MapIntToUint) rehash() {
	ns := len(h.slots) << 1
	nslots := make([]*entryIntToUint, ns)
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

func (h *MapIntToUint) free(entry *entryIntToUint) {
	entry.next = h.freelist
	h.freelist = entry
	h.size--
}

func (h *MapIntToUint) alloc(k int, v uint) *entryIntToUint {
	if h.freelist == nil {
		entries := make([]entryIntToUint, 256)
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

type entryIntToUint8 struct {
	k    int
	v    uint8
	next *entryIntToUint8
}

// MapIntToUint8 implements a hash map from int to uint8.
type MapIntToUint8 struct {
	mask     int
	slots    []*entryIntToUint8
	used     int
	size     int
	max      int
	freelist *entryIntToUint8
}

// NewMapIntToUint8 creates a new MapIntToUint8 with at least a size of size.
func NewMapIntToUint8(size int) *MapIntToUint8 {
	shift := nextShiftPowerOfTwo(size)
	return &MapIntToUint8{
		mask:  ^(^0 << shift),
		max:   maxFill(1 << shift),
		slots: make([]*entryIntToUint8, 1<<shift),
	}
}

// Size returns the current size of the map.
func (h *MapIntToUint8) Size() int {
	return h.size
}

// Get looks up a key k and returns its value. 0 if not found.
func (h *MapIntToUint8) Get(k int) uint8 {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v
		}
	}
	return 0
}

// Contains looks up a key k and returns true if it is found else false.
func (h *MapIntToUint8) Contains(k int) bool {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return true
		}
	}
	return false
}

// Modify looks up a key k and calls function fn with a pointer to its value.
func (h *MapIntToUint8) Modify(k int, fn func(v *uint8)) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			fn(&e.v)
			return
		}
	}
}

// Find looks up a key k and returns its value and true.
// 0 and false if not found.
func (h *MapIntToUint8) Find(k int) (uint8, bool) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v, true
		}
	}
	return 0, false
}

// Visit calls a given function fn for every key/value pair in the map.
func (h *MapIntToUint8) Visit(fn func(int, uint8)) {
	for _, e := range h.slots {
		for ; e != nil; e = e.next {
			fn(e.k, e.v)
		}
	}
}

// Add increments the value associated with key k by the value of v.
// It creates a new key/value pair k/v in the map if the key does not exist.
func (h *MapIntToUint8) Add(k int, v uint8) {
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
func (h *MapIntToUint8) Put(k int, v uint8) {
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
func (h *MapIntToUint8) Remove(k int) {
	p := &h.slots[int(k)&h.mask]
	var parent *entryIntToUint8
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
func (h *MapIntToUint8) Clear() {
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
func (h *MapIntToUint8) Inc(k int) {
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

func (h *MapIntToUint8) rehash() {
	ns := len(h.slots) << 1
	nslots := make([]*entryIntToUint8, ns)
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

func (h *MapIntToUint8) free(entry *entryIntToUint8) {
	entry.next = h.freelist
	h.freelist = entry
	h.size--
}

func (h *MapIntToUint8) alloc(k int, v uint8) *entryIntToUint8 {
	if h.freelist == nil {
		entries := make([]entryIntToUint8, 256)
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

type entryIntToUint16 struct {
	k    int
	v    uint16
	next *entryIntToUint16
}

// MapIntToUint16 implements a hash map from int to uint16.
type MapIntToUint16 struct {
	mask     int
	slots    []*entryIntToUint16
	used     int
	size     int
	max      int
	freelist *entryIntToUint16
}

// NewMapIntToUint16 creates a new MapIntToUint16 with at least a size of size.
func NewMapIntToUint16(size int) *MapIntToUint16 {
	shift := nextShiftPowerOfTwo(size)
	return &MapIntToUint16{
		mask:  ^(^0 << shift),
		max:   maxFill(1 << shift),
		slots: make([]*entryIntToUint16, 1<<shift),
	}
}

// Size returns the current size of the map.
func (h *MapIntToUint16) Size() int {
	return h.size
}

// Get looks up a key k and returns its value. 0 if not found.
func (h *MapIntToUint16) Get(k int) uint16 {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v
		}
	}
	return 0
}

// Contains looks up a key k and returns true if it is found else false.
func (h *MapIntToUint16) Contains(k int) bool {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return true
		}
	}
	return false
}

// Modify looks up a key k and calls function fn with a pointer to its value.
func (h *MapIntToUint16) Modify(k int, fn func(v *uint16)) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			fn(&e.v)
			return
		}
	}
}

// Find looks up a key k and returns its value and true.
// 0 and false if not found.
func (h *MapIntToUint16) Find(k int) (uint16, bool) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v, true
		}
	}
	return 0, false
}

// Visit calls a given function fn for every key/value pair in the map.
func (h *MapIntToUint16) Visit(fn func(int, uint16)) {
	for _, e := range h.slots {
		for ; e != nil; e = e.next {
			fn(e.k, e.v)
		}
	}
}

// Add increments the value associated with key k by the value of v.
// It creates a new key/value pair k/v in the map if the key does not exist.
func (h *MapIntToUint16) Add(k int, v uint16) {
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
func (h *MapIntToUint16) Put(k int, v uint16) {
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
func (h *MapIntToUint16) Remove(k int) {
	p := &h.slots[int(k)&h.mask]
	var parent *entryIntToUint16
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
func (h *MapIntToUint16) Clear() {
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
func (h *MapIntToUint16) Inc(k int) {
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

func (h *MapIntToUint16) rehash() {
	ns := len(h.slots) << 1
	nslots := make([]*entryIntToUint16, ns)
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

func (h *MapIntToUint16) free(entry *entryIntToUint16) {
	entry.next = h.freelist
	h.freelist = entry
	h.size--
}

func (h *MapIntToUint16) alloc(k int, v uint16) *entryIntToUint16 {
	if h.freelist == nil {
		entries := make([]entryIntToUint16, 256)
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

type entryIntToUint32 struct {
	k    int
	v    uint32
	next *entryIntToUint32
}

// MapIntToUint32 implements a hash map from int to uint32.
type MapIntToUint32 struct {
	mask     int
	slots    []*entryIntToUint32
	used     int
	size     int
	max      int
	freelist *entryIntToUint32
}

// NewMapIntToUint32 creates a new MapIntToUint32 with at least a size of size.
func NewMapIntToUint32(size int) *MapIntToUint32 {
	shift := nextShiftPowerOfTwo(size)
	return &MapIntToUint32{
		mask:  ^(^0 << shift),
		max:   maxFill(1 << shift),
		slots: make([]*entryIntToUint32, 1<<shift),
	}
}

// Size returns the current size of the map.
func (h *MapIntToUint32) Size() int {
	return h.size
}

// Get looks up a key k and returns its value. 0 if not found.
func (h *MapIntToUint32) Get(k int) uint32 {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v
		}
	}
	return 0
}

// Contains looks up a key k and returns true if it is found else false.
func (h *MapIntToUint32) Contains(k int) bool {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return true
		}
	}
	return false
}

// Modify looks up a key k and calls function fn with a pointer to its value.
func (h *MapIntToUint32) Modify(k int, fn func(v *uint32)) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			fn(&e.v)
			return
		}
	}
}

// Find looks up a key k and returns its value and true.
// 0 and false if not found.
func (h *MapIntToUint32) Find(k int) (uint32, bool) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v, true
		}
	}
	return 0, false
}

// Visit calls a given function fn for every key/value pair in the map.
func (h *MapIntToUint32) Visit(fn func(int, uint32)) {
	for _, e := range h.slots {
		for ; e != nil; e = e.next {
			fn(e.k, e.v)
		}
	}
}

// Add increments the value associated with key k by the value of v.
// It creates a new key/value pair k/v in the map if the key does not exist.
func (h *MapIntToUint32) Add(k int, v uint32) {
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
func (h *MapIntToUint32) Put(k int, v uint32) {
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
func (h *MapIntToUint32) Remove(k int) {
	p := &h.slots[int(k)&h.mask]
	var parent *entryIntToUint32
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
func (h *MapIntToUint32) Clear() {
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
func (h *MapIntToUint32) Inc(k int) {
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

func (h *MapIntToUint32) rehash() {
	ns := len(h.slots) << 1
	nslots := make([]*entryIntToUint32, ns)
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

func (h *MapIntToUint32) free(entry *entryIntToUint32) {
	entry.next = h.freelist
	h.freelist = entry
	h.size--
}

func (h *MapIntToUint32) alloc(k int, v uint32) *entryIntToUint32 {
	if h.freelist == nil {
		entries := make([]entryIntToUint32, 256)
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

type entryIntToUint64 struct {
	k    int
	v    uint64
	next *entryIntToUint64
}

// MapIntToUint64 implements a hash map from int to uint64.
type MapIntToUint64 struct {
	mask     int
	slots    []*entryIntToUint64
	used     int
	size     int
	max      int
	freelist *entryIntToUint64
}

// NewMapIntToUint64 creates a new MapIntToUint64 with at least a size of size.
func NewMapIntToUint64(size int) *MapIntToUint64 {
	shift := nextShiftPowerOfTwo(size)
	return &MapIntToUint64{
		mask:  ^(^0 << shift),
		max:   maxFill(1 << shift),
		slots: make([]*entryIntToUint64, 1<<shift),
	}
}

// Size returns the current size of the map.
func (h *MapIntToUint64) Size() int {
	return h.size
}

// Get looks up a key k and returns its value. 0 if not found.
func (h *MapIntToUint64) Get(k int) uint64 {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v
		}
	}
	return 0
}

// Contains looks up a key k and returns true if it is found else false.
func (h *MapIntToUint64) Contains(k int) bool {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return true
		}
	}
	return false
}

// Modify looks up a key k and calls function fn with a pointer to its value.
func (h *MapIntToUint64) Modify(k int, fn func(v *uint64)) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			fn(&e.v)
			return
		}
	}
}

// Find looks up a key k and returns its value and true.
// 0 and false if not found.
func (h *MapIntToUint64) Find(k int) (uint64, bool) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v, true
		}
	}
	return 0, false
}

// Visit calls a given function fn for every key/value pair in the map.
func (h *MapIntToUint64) Visit(fn func(int, uint64)) {
	for _, e := range h.slots {
		for ; e != nil; e = e.next {
			fn(e.k, e.v)
		}
	}
}

// Add increments the value associated with key k by the value of v.
// It creates a new key/value pair k/v in the map if the key does not exist.
func (h *MapIntToUint64) Add(k int, v uint64) {
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
func (h *MapIntToUint64) Put(k int, v uint64) {
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
func (h *MapIntToUint64) Remove(k int) {
	p := &h.slots[int(k)&h.mask]
	var parent *entryIntToUint64
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
func (h *MapIntToUint64) Clear() {
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
func (h *MapIntToUint64) Inc(k int) {
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

func (h *MapIntToUint64) rehash() {
	ns := len(h.slots) << 1
	nslots := make([]*entryIntToUint64, ns)
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

func (h *MapIntToUint64) free(entry *entryIntToUint64) {
	entry.next = h.freelist
	h.freelist = entry
	h.size--
}

func (h *MapIntToUint64) alloc(k int, v uint64) *entryIntToUint64 {
	if h.freelist == nil {
		entries := make([]entryIntToUint64, 256)
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
