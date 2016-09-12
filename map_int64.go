// THIS IS A MACHINE GENERATED FILE!
// BE CAREFUL WITH EDITING BY HAND.
//
// This is Free Software covered by the terms of the MIT license.
// See LICENSE file for details.
// (c) 2016 by Sascha L. Teichmann.
// See the full list of contributors in the CONTRIBUTORS file.

package genmap

type entryInt64ToInt struct {
	k    int64
	v    int
	next *entryInt64ToInt
}

// MapInt64ToInt implements a hash map from int64 to int.
type MapInt64ToInt struct {
	mask     int
	slots    []*entryInt64ToInt
	used     int
	size     int
	max      int
	freelist *entryInt64ToInt
}

// NewMapInt64ToInt creates a new MapInt64ToInt with at least a size of size.
func NewMapInt64ToInt(size int) *MapInt64ToInt {
	shift := nextShiftPowerOfTwo(size)
	return &MapInt64ToInt{
		mask:  ^(^0 << shift),
		max:   maxFill(1 << shift),
		slots: make([]*entryInt64ToInt, 1<<shift),
	}
}

// Size returns the current size of the map.
func (h *MapInt64ToInt) Size() int {
	return h.size
}

// Get looks up a key k and returns its value. 0 if not found.
func (h *MapInt64ToInt) Get(k int64) int {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v
		}
	}
	return 0
}

// Contains looks up a key k and returns true if it is found else false.
func (h *MapInt64ToInt) Contains(k int64) bool {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return true
		}
	}
	return false
}

// Modify looks up a key k and calls function fn with a pointer to its value.
func (h *MapInt64ToInt) Modify(k int64, fn func(v *int)) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			fn(&e.v)
			return
		}
	}
}

// Find looks up a key k and returns its value and true.
// 0 and false if not found.
func (h *MapInt64ToInt) Find(k int64) (int, bool) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v, true
		}
	}
	return 0, false
}

// Visit calls a given function fn for every key/value pair in the map.
func (h *MapInt64ToInt) Visit(fn func(int64, int)) {
	for _, e := range h.slots {
		for ; e != nil; e = e.next {
			fn(e.k, e.v)
		}
	}
}

// Add increments the value associated with key k by the value of v.
// It creates a new key/value pair k/v in the map if the key does not exist.
func (h *MapInt64ToInt) Add(k int64, v int) {
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
func (h *MapInt64ToInt) Put(k int64, v int) {
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
func (h *MapInt64ToInt) Remove(k int64) {
	p := &h.slots[int(k)&h.mask]
	var parent *entryInt64ToInt
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
func (h *MapInt64ToInt) Clear() {
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
func (h *MapInt64ToInt) Inc(k int64) {
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

func (h *MapInt64ToInt) rehash() {
	ns := len(h.slots) << 1
	nslots := make([]*entryInt64ToInt, ns)
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

func (h *MapInt64ToInt) free(entry *entryInt64ToInt) {
	entry.next = h.freelist
	h.freelist = entry
	h.size--
}

func (h *MapInt64ToInt) alloc(k int64, v int) *entryInt64ToInt {
	if h.freelist == nil {
		entries := make([]entryInt64ToInt, 256)
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

type entryInt64ToInt8 struct {
	k    int64
	v    int8
	next *entryInt64ToInt8
}

// MapInt64ToInt8 implements a hash map from int64 to int8.
type MapInt64ToInt8 struct {
	mask     int
	slots    []*entryInt64ToInt8
	used     int
	size     int
	max      int
	freelist *entryInt64ToInt8
}

// NewMapInt64ToInt8 creates a new MapInt64ToInt8 with at least a size of size.
func NewMapInt64ToInt8(size int) *MapInt64ToInt8 {
	shift := nextShiftPowerOfTwo(size)
	return &MapInt64ToInt8{
		mask:  ^(^0 << shift),
		max:   maxFill(1 << shift),
		slots: make([]*entryInt64ToInt8, 1<<shift),
	}
}

// Size returns the current size of the map.
func (h *MapInt64ToInt8) Size() int {
	return h.size
}

// Get looks up a key k and returns its value. 0 if not found.
func (h *MapInt64ToInt8) Get(k int64) int8 {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v
		}
	}
	return 0
}

// Contains looks up a key k and returns true if it is found else false.
func (h *MapInt64ToInt8) Contains(k int64) bool {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return true
		}
	}
	return false
}

// Modify looks up a key k and calls function fn with a pointer to its value.
func (h *MapInt64ToInt8) Modify(k int64, fn func(v *int8)) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			fn(&e.v)
			return
		}
	}
}

// Find looks up a key k and returns its value and true.
// 0 and false if not found.
func (h *MapInt64ToInt8) Find(k int64) (int8, bool) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v, true
		}
	}
	return 0, false
}

// Visit calls a given function fn for every key/value pair in the map.
func (h *MapInt64ToInt8) Visit(fn func(int64, int8)) {
	for _, e := range h.slots {
		for ; e != nil; e = e.next {
			fn(e.k, e.v)
		}
	}
}

// Add increments the value associated with key k by the value of v.
// It creates a new key/value pair k/v in the map if the key does not exist.
func (h *MapInt64ToInt8) Add(k int64, v int8) {
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
func (h *MapInt64ToInt8) Put(k int64, v int8) {
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
func (h *MapInt64ToInt8) Remove(k int64) {
	p := &h.slots[int(k)&h.mask]
	var parent *entryInt64ToInt8
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
func (h *MapInt64ToInt8) Clear() {
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
func (h *MapInt64ToInt8) Inc(k int64) {
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

func (h *MapInt64ToInt8) rehash() {
	ns := len(h.slots) << 1
	nslots := make([]*entryInt64ToInt8, ns)
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

func (h *MapInt64ToInt8) free(entry *entryInt64ToInt8) {
	entry.next = h.freelist
	h.freelist = entry
	h.size--
}

func (h *MapInt64ToInt8) alloc(k int64, v int8) *entryInt64ToInt8 {
	if h.freelist == nil {
		entries := make([]entryInt64ToInt8, 256)
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

type entryInt64ToInt16 struct {
	k    int64
	v    int16
	next *entryInt64ToInt16
}

// MapInt64ToInt16 implements a hash map from int64 to int16.
type MapInt64ToInt16 struct {
	mask     int
	slots    []*entryInt64ToInt16
	used     int
	size     int
	max      int
	freelist *entryInt64ToInt16
}

// NewMapInt64ToInt16 creates a new MapInt64ToInt16 with at least a size of size.
func NewMapInt64ToInt16(size int) *MapInt64ToInt16 {
	shift := nextShiftPowerOfTwo(size)
	return &MapInt64ToInt16{
		mask:  ^(^0 << shift),
		max:   maxFill(1 << shift),
		slots: make([]*entryInt64ToInt16, 1<<shift),
	}
}

// Size returns the current size of the map.
func (h *MapInt64ToInt16) Size() int {
	return h.size
}

// Get looks up a key k and returns its value. 0 if not found.
func (h *MapInt64ToInt16) Get(k int64) int16 {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v
		}
	}
	return 0
}

// Contains looks up a key k and returns true if it is found else false.
func (h *MapInt64ToInt16) Contains(k int64) bool {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return true
		}
	}
	return false
}

// Modify looks up a key k and calls function fn with a pointer to its value.
func (h *MapInt64ToInt16) Modify(k int64, fn func(v *int16)) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			fn(&e.v)
			return
		}
	}
}

// Find looks up a key k and returns its value and true.
// 0 and false if not found.
func (h *MapInt64ToInt16) Find(k int64) (int16, bool) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v, true
		}
	}
	return 0, false
}

// Visit calls a given function fn for every key/value pair in the map.
func (h *MapInt64ToInt16) Visit(fn func(int64, int16)) {
	for _, e := range h.slots {
		for ; e != nil; e = e.next {
			fn(e.k, e.v)
		}
	}
}

// Add increments the value associated with key k by the value of v.
// It creates a new key/value pair k/v in the map if the key does not exist.
func (h *MapInt64ToInt16) Add(k int64, v int16) {
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
func (h *MapInt64ToInt16) Put(k int64, v int16) {
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
func (h *MapInt64ToInt16) Remove(k int64) {
	p := &h.slots[int(k)&h.mask]
	var parent *entryInt64ToInt16
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
func (h *MapInt64ToInt16) Clear() {
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
func (h *MapInt64ToInt16) Inc(k int64) {
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

func (h *MapInt64ToInt16) rehash() {
	ns := len(h.slots) << 1
	nslots := make([]*entryInt64ToInt16, ns)
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

func (h *MapInt64ToInt16) free(entry *entryInt64ToInt16) {
	entry.next = h.freelist
	h.freelist = entry
	h.size--
}

func (h *MapInt64ToInt16) alloc(k int64, v int16) *entryInt64ToInt16 {
	if h.freelist == nil {
		entries := make([]entryInt64ToInt16, 256)
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

type entryInt64ToInt32 struct {
	k    int64
	v    int32
	next *entryInt64ToInt32
}

// MapInt64ToInt32 implements a hash map from int64 to int32.
type MapInt64ToInt32 struct {
	mask     int
	slots    []*entryInt64ToInt32
	used     int
	size     int
	max      int
	freelist *entryInt64ToInt32
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

// Size returns the current size of the map.
func (h *MapInt64ToInt32) Size() int {
	return h.size
}

// Get looks up a key k and returns its value. 0 if not found.
func (h *MapInt64ToInt32) Get(k int64) int32 {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v
		}
	}
	return 0
}

// Contains looks up a key k and returns true if it is found else false.
func (h *MapInt64ToInt32) Contains(k int64) bool {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return true
		}
	}
	return false
}

// Modify looks up a key k and calls function fn with a pointer to its value.
func (h *MapInt64ToInt32) Modify(k int64, fn func(v *int32)) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			fn(&e.v)
			return
		}
	}
}

// Find looks up a key k and returns its value and true.
// 0 and false if not found.
func (h *MapInt64ToInt32) Find(k int64) (int32, bool) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v, true
		}
	}
	return 0, false
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

// Put sets the value associated with key k to the value of v.
// It creates a new key/value pair k/v in the map if the key does not exist.
func (h *MapInt64ToInt32) Put(k int64, v int32) {
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
func (h *MapInt64ToInt32) Remove(k int64) {
	p := &h.slots[int(k)&h.mask]
	var parent *entryInt64ToInt32
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
func (h *MapInt64ToInt32) Clear() {
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

func (h *MapInt64ToInt32) free(entry *entryInt64ToInt32) {
	entry.next = h.freelist
	h.freelist = entry
	h.size--
}

func (h *MapInt64ToInt32) alloc(k int64, v int32) *entryInt64ToInt32 {
	if h.freelist == nil {
		entries := make([]entryInt64ToInt32, 256)
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

type entryInt64ToInt64 struct {
	k    int64
	v    int64
	next *entryInt64ToInt64
}

// MapInt64ToInt64 implements a hash map from int64 to int64.
type MapInt64ToInt64 struct {
	mask     int
	slots    []*entryInt64ToInt64
	used     int
	size     int
	max      int
	freelist *entryInt64ToInt64
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

// Size returns the current size of the map.
func (h *MapInt64ToInt64) Size() int {
	return h.size
}

// Get looks up a key k and returns its value. 0 if not found.
func (h *MapInt64ToInt64) Get(k int64) int64 {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v
		}
	}
	return 0
}

// Contains looks up a key k and returns true if it is found else false.
func (h *MapInt64ToInt64) Contains(k int64) bool {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return true
		}
	}
	return false
}

// Modify looks up a key k and calls function fn with a pointer to its value.
func (h *MapInt64ToInt64) Modify(k int64, fn func(v *int64)) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			fn(&e.v)
			return
		}
	}
}

// Find looks up a key k and returns its value and true.
// 0 and false if not found.
func (h *MapInt64ToInt64) Find(k int64) (int64, bool) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v, true
		}
	}
	return 0, false
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

// Put sets the value associated with key k to the value of v.
// It creates a new key/value pair k/v in the map if the key does not exist.
func (h *MapInt64ToInt64) Put(k int64, v int64) {
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
func (h *MapInt64ToInt64) Remove(k int64) {
	p := &h.slots[int(k)&h.mask]
	var parent *entryInt64ToInt64
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
func (h *MapInt64ToInt64) Clear() {
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

func (h *MapInt64ToInt64) free(entry *entryInt64ToInt64) {
	entry.next = h.freelist
	h.freelist = entry
	h.size--
}

func (h *MapInt64ToInt64) alloc(k int64, v int64) *entryInt64ToInt64 {
	if h.freelist == nil {
		entries := make([]entryInt64ToInt64, 256)
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

type entryInt64ToUint struct {
	k    int64
	v    uint
	next *entryInt64ToUint
}

// MapInt64ToUint implements a hash map from int64 to uint.
type MapInt64ToUint struct {
	mask     int
	slots    []*entryInt64ToUint
	used     int
	size     int
	max      int
	freelist *entryInt64ToUint
}

// NewMapInt64ToUint creates a new MapInt64ToUint with at least a size of size.
func NewMapInt64ToUint(size int) *MapInt64ToUint {
	shift := nextShiftPowerOfTwo(size)
	return &MapInt64ToUint{
		mask:  ^(^0 << shift),
		max:   maxFill(1 << shift),
		slots: make([]*entryInt64ToUint, 1<<shift),
	}
}

// Size returns the current size of the map.
func (h *MapInt64ToUint) Size() int {
	return h.size
}

// Get looks up a key k and returns its value. 0 if not found.
func (h *MapInt64ToUint) Get(k int64) uint {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v
		}
	}
	return 0
}

// Contains looks up a key k and returns true if it is found else false.
func (h *MapInt64ToUint) Contains(k int64) bool {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return true
		}
	}
	return false
}

// Modify looks up a key k and calls function fn with a pointer to its value.
func (h *MapInt64ToUint) Modify(k int64, fn func(v *uint)) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			fn(&e.v)
			return
		}
	}
}

// Find looks up a key k and returns its value and true.
// 0 and false if not found.
func (h *MapInt64ToUint) Find(k int64) (uint, bool) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v, true
		}
	}
	return 0, false
}

// Visit calls a given function fn for every key/value pair in the map.
func (h *MapInt64ToUint) Visit(fn func(int64, uint)) {
	for _, e := range h.slots {
		for ; e != nil; e = e.next {
			fn(e.k, e.v)
		}
	}
}

// Add increments the value associated with key k by the value of v.
// It creates a new key/value pair k/v in the map if the key does not exist.
func (h *MapInt64ToUint) Add(k int64, v uint) {
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
func (h *MapInt64ToUint) Put(k int64, v uint) {
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
func (h *MapInt64ToUint) Remove(k int64) {
	p := &h.slots[int(k)&h.mask]
	var parent *entryInt64ToUint
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
func (h *MapInt64ToUint) Clear() {
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
func (h *MapInt64ToUint) Inc(k int64) {
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

func (h *MapInt64ToUint) rehash() {
	ns := len(h.slots) << 1
	nslots := make([]*entryInt64ToUint, ns)
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

func (h *MapInt64ToUint) free(entry *entryInt64ToUint) {
	entry.next = h.freelist
	h.freelist = entry
	h.size--
}

func (h *MapInt64ToUint) alloc(k int64, v uint) *entryInt64ToUint {
	if h.freelist == nil {
		entries := make([]entryInt64ToUint, 256)
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

type entryInt64ToUint8 struct {
	k    int64
	v    uint8
	next *entryInt64ToUint8
}

// MapInt64ToUint8 implements a hash map from int64 to uint8.
type MapInt64ToUint8 struct {
	mask     int
	slots    []*entryInt64ToUint8
	used     int
	size     int
	max      int
	freelist *entryInt64ToUint8
}

// NewMapInt64ToUint8 creates a new MapInt64ToUint8 with at least a size of size.
func NewMapInt64ToUint8(size int) *MapInt64ToUint8 {
	shift := nextShiftPowerOfTwo(size)
	return &MapInt64ToUint8{
		mask:  ^(^0 << shift),
		max:   maxFill(1 << shift),
		slots: make([]*entryInt64ToUint8, 1<<shift),
	}
}

// Size returns the current size of the map.
func (h *MapInt64ToUint8) Size() int {
	return h.size
}

// Get looks up a key k and returns its value. 0 if not found.
func (h *MapInt64ToUint8) Get(k int64) uint8 {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v
		}
	}
	return 0
}

// Contains looks up a key k and returns true if it is found else false.
func (h *MapInt64ToUint8) Contains(k int64) bool {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return true
		}
	}
	return false
}

// Modify looks up a key k and calls function fn with a pointer to its value.
func (h *MapInt64ToUint8) Modify(k int64, fn func(v *uint8)) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			fn(&e.v)
			return
		}
	}
}

// Find looks up a key k and returns its value and true.
// 0 and false if not found.
func (h *MapInt64ToUint8) Find(k int64) (uint8, bool) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v, true
		}
	}
	return 0, false
}

// Visit calls a given function fn for every key/value pair in the map.
func (h *MapInt64ToUint8) Visit(fn func(int64, uint8)) {
	for _, e := range h.slots {
		for ; e != nil; e = e.next {
			fn(e.k, e.v)
		}
	}
}

// Add increments the value associated with key k by the value of v.
// It creates a new key/value pair k/v in the map if the key does not exist.
func (h *MapInt64ToUint8) Add(k int64, v uint8) {
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
func (h *MapInt64ToUint8) Put(k int64, v uint8) {
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
func (h *MapInt64ToUint8) Remove(k int64) {
	p := &h.slots[int(k)&h.mask]
	var parent *entryInt64ToUint8
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
func (h *MapInt64ToUint8) Clear() {
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
func (h *MapInt64ToUint8) Inc(k int64) {
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

func (h *MapInt64ToUint8) rehash() {
	ns := len(h.slots) << 1
	nslots := make([]*entryInt64ToUint8, ns)
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

func (h *MapInt64ToUint8) free(entry *entryInt64ToUint8) {
	entry.next = h.freelist
	h.freelist = entry
	h.size--
}

func (h *MapInt64ToUint8) alloc(k int64, v uint8) *entryInt64ToUint8 {
	if h.freelist == nil {
		entries := make([]entryInt64ToUint8, 256)
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

type entryInt64ToUint16 struct {
	k    int64
	v    uint16
	next *entryInt64ToUint16
}

// MapInt64ToUint16 implements a hash map from int64 to uint16.
type MapInt64ToUint16 struct {
	mask     int
	slots    []*entryInt64ToUint16
	used     int
	size     int
	max      int
	freelist *entryInt64ToUint16
}

// NewMapInt64ToUint16 creates a new MapInt64ToUint16 with at least a size of size.
func NewMapInt64ToUint16(size int) *MapInt64ToUint16 {
	shift := nextShiftPowerOfTwo(size)
	return &MapInt64ToUint16{
		mask:  ^(^0 << shift),
		max:   maxFill(1 << shift),
		slots: make([]*entryInt64ToUint16, 1<<shift),
	}
}

// Size returns the current size of the map.
func (h *MapInt64ToUint16) Size() int {
	return h.size
}

// Get looks up a key k and returns its value. 0 if not found.
func (h *MapInt64ToUint16) Get(k int64) uint16 {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v
		}
	}
	return 0
}

// Contains looks up a key k and returns true if it is found else false.
func (h *MapInt64ToUint16) Contains(k int64) bool {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return true
		}
	}
	return false
}

// Modify looks up a key k and calls function fn with a pointer to its value.
func (h *MapInt64ToUint16) Modify(k int64, fn func(v *uint16)) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			fn(&e.v)
			return
		}
	}
}

// Find looks up a key k and returns its value and true.
// 0 and false if not found.
func (h *MapInt64ToUint16) Find(k int64) (uint16, bool) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v, true
		}
	}
	return 0, false
}

// Visit calls a given function fn for every key/value pair in the map.
func (h *MapInt64ToUint16) Visit(fn func(int64, uint16)) {
	for _, e := range h.slots {
		for ; e != nil; e = e.next {
			fn(e.k, e.v)
		}
	}
}

// Add increments the value associated with key k by the value of v.
// It creates a new key/value pair k/v in the map if the key does not exist.
func (h *MapInt64ToUint16) Add(k int64, v uint16) {
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
func (h *MapInt64ToUint16) Put(k int64, v uint16) {
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
func (h *MapInt64ToUint16) Remove(k int64) {
	p := &h.slots[int(k)&h.mask]
	var parent *entryInt64ToUint16
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
func (h *MapInt64ToUint16) Clear() {
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
func (h *MapInt64ToUint16) Inc(k int64) {
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

func (h *MapInt64ToUint16) rehash() {
	ns := len(h.slots) << 1
	nslots := make([]*entryInt64ToUint16, ns)
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

func (h *MapInt64ToUint16) free(entry *entryInt64ToUint16) {
	entry.next = h.freelist
	h.freelist = entry
	h.size--
}

func (h *MapInt64ToUint16) alloc(k int64, v uint16) *entryInt64ToUint16 {
	if h.freelist == nil {
		entries := make([]entryInt64ToUint16, 256)
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

type entryInt64ToUint32 struct {
	k    int64
	v    uint32
	next *entryInt64ToUint32
}

// MapInt64ToUint32 implements a hash map from int64 to uint32.
type MapInt64ToUint32 struct {
	mask     int
	slots    []*entryInt64ToUint32
	used     int
	size     int
	max      int
	freelist *entryInt64ToUint32
}

// NewMapInt64ToUint32 creates a new MapInt64ToUint32 with at least a size of size.
func NewMapInt64ToUint32(size int) *MapInt64ToUint32 {
	shift := nextShiftPowerOfTwo(size)
	return &MapInt64ToUint32{
		mask:  ^(^0 << shift),
		max:   maxFill(1 << shift),
		slots: make([]*entryInt64ToUint32, 1<<shift),
	}
}

// Size returns the current size of the map.
func (h *MapInt64ToUint32) Size() int {
	return h.size
}

// Get looks up a key k and returns its value. 0 if not found.
func (h *MapInt64ToUint32) Get(k int64) uint32 {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v
		}
	}
	return 0
}

// Contains looks up a key k and returns true if it is found else false.
func (h *MapInt64ToUint32) Contains(k int64) bool {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return true
		}
	}
	return false
}

// Modify looks up a key k and calls function fn with a pointer to its value.
func (h *MapInt64ToUint32) Modify(k int64, fn func(v *uint32)) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			fn(&e.v)
			return
		}
	}
}

// Find looks up a key k and returns its value and true.
// 0 and false if not found.
func (h *MapInt64ToUint32) Find(k int64) (uint32, bool) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v, true
		}
	}
	return 0, false
}

// Visit calls a given function fn for every key/value pair in the map.
func (h *MapInt64ToUint32) Visit(fn func(int64, uint32)) {
	for _, e := range h.slots {
		for ; e != nil; e = e.next {
			fn(e.k, e.v)
		}
	}
}

// Add increments the value associated with key k by the value of v.
// It creates a new key/value pair k/v in the map if the key does not exist.
func (h *MapInt64ToUint32) Add(k int64, v uint32) {
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
func (h *MapInt64ToUint32) Put(k int64, v uint32) {
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
func (h *MapInt64ToUint32) Remove(k int64) {
	p := &h.slots[int(k)&h.mask]
	var parent *entryInt64ToUint32
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
func (h *MapInt64ToUint32) Clear() {
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
func (h *MapInt64ToUint32) Inc(k int64) {
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

func (h *MapInt64ToUint32) rehash() {
	ns := len(h.slots) << 1
	nslots := make([]*entryInt64ToUint32, ns)
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

func (h *MapInt64ToUint32) free(entry *entryInt64ToUint32) {
	entry.next = h.freelist
	h.freelist = entry
	h.size--
}

func (h *MapInt64ToUint32) alloc(k int64, v uint32) *entryInt64ToUint32 {
	if h.freelist == nil {
		entries := make([]entryInt64ToUint32, 256)
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

type entryInt64ToUint64 struct {
	k    int64
	v    uint64
	next *entryInt64ToUint64
}

// MapInt64ToUint64 implements a hash map from int64 to uint64.
type MapInt64ToUint64 struct {
	mask     int
	slots    []*entryInt64ToUint64
	used     int
	size     int
	max      int
	freelist *entryInt64ToUint64
}

// NewMapInt64ToUint64 creates a new MapInt64ToUint64 with at least a size of size.
func NewMapInt64ToUint64(size int) *MapInt64ToUint64 {
	shift := nextShiftPowerOfTwo(size)
	return &MapInt64ToUint64{
		mask:  ^(^0 << shift),
		max:   maxFill(1 << shift),
		slots: make([]*entryInt64ToUint64, 1<<shift),
	}
}

// Size returns the current size of the map.
func (h *MapInt64ToUint64) Size() int {
	return h.size
}

// Get looks up a key k and returns its value. 0 if not found.
func (h *MapInt64ToUint64) Get(k int64) uint64 {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v
		}
	}
	return 0
}

// Contains looks up a key k and returns true if it is found else false.
func (h *MapInt64ToUint64) Contains(k int64) bool {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return true
		}
	}
	return false
}

// Modify looks up a key k and calls function fn with a pointer to its value.
func (h *MapInt64ToUint64) Modify(k int64, fn func(v *uint64)) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			fn(&e.v)
			return
		}
	}
}

// Find looks up a key k and returns its value and true.
// 0 and false if not found.
func (h *MapInt64ToUint64) Find(k int64) (uint64, bool) {
	for e := h.slots[int(k)&h.mask]; e != nil; e = e.next {
		if e.k == k {
			return e.v, true
		}
	}
	return 0, false
}

// Visit calls a given function fn for every key/value pair in the map.
func (h *MapInt64ToUint64) Visit(fn func(int64, uint64)) {
	for _, e := range h.slots {
		for ; e != nil; e = e.next {
			fn(e.k, e.v)
		}
	}
}

// Add increments the value associated with key k by the value of v.
// It creates a new key/value pair k/v in the map if the key does not exist.
func (h *MapInt64ToUint64) Add(k int64, v uint64) {
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
func (h *MapInt64ToUint64) Put(k int64, v uint64) {
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
func (h *MapInt64ToUint64) Remove(k int64) {
	p := &h.slots[int(k)&h.mask]
	var parent *entryInt64ToUint64
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
func (h *MapInt64ToUint64) Clear() {
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
func (h *MapInt64ToUint64) Inc(k int64) {
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

func (h *MapInt64ToUint64) rehash() {
	ns := len(h.slots) << 1
	nslots := make([]*entryInt64ToUint64, ns)
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

func (h *MapInt64ToUint64) free(entry *entryInt64ToUint64) {
	entry.next = h.freelist
	h.freelist = entry
	h.size--
}

func (h *MapInt64ToUint64) alloc(k int64, v uint64) *entryInt64ToUint64 {
	if h.freelist == nil {
		entries := make([]entryInt64ToUint64, 256)
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
