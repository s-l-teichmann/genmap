// THIS IS A MACHINE GENERATED FILE!
// BE CAREFUL WITH EDITING BY HAND.
//
// This is Free Software covered by the terms of the MIT license.
// See LICENSE file for details.
// (c) 2016 by Sascha L. Teichmann.
// See the full list of contributors in the CONTRIBUTORS file.

package genmap

import "testing"

func TestMapUint64ToIntSize(t *testing.T) {
	m := NewMapUint64ToInt(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint64(k), int(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint64(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint64ToIntContains(t *testing.T) {
	m := NewMapUint64ToInt(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), int(signedData[i]))
		if !m.Contains(uint64(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint64(k))
		if m.Contains(uint64(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint64ToIntGet(t *testing.T) {
	m := NewMapUint64ToInt(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), int(signedData[i]))
		if g := m.Get(uint64(k)); g != int(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapUint64ToIntClear(t *testing.T) {
	m := NewMapUint64ToInt(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), int(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint64(k), int(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUint64ToIntInc(t *testing.T) {
	m := NewMapUint64ToInt(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), int(signedData[i]))
	}
	for i, k := range unsignedData {
		m.Inc(uint64(k))
		m.Inc(uint64(k))
		m.Inc(uint64(k))
		if g := m.Get(uint64(k)); g != int(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int(signedData[i])+3)
		}
	}
	m = NewMapUint64ToInt(13)
	for _, k := range unsignedData {
		m.Inc(uint64(k))
		if g := m.Get(uint64(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapUint64ToIntAdd(t *testing.T) {
	m := NewMapUint64ToInt(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), int(signedData[i]))
	}
	for i, k := range unsignedData {
		m.Add(uint64(k), 3)
		if g := m.Get(uint64(k)); g != int(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int(signedData[i])+3)
		}
	}
	m = NewMapUint64ToInt(13)
	for _, k := range unsignedData {
		m.Add(uint64(k), 42)
		if g := m.Get(uint64(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapUint64ToIntVisit(t *testing.T) {
	m := NewMapUint64ToInt(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), int(signedData[i]))
	}

	n := make(map[uint64]int, len(unsignedData))

	for i, k := range unsignedData {
		n[uint64(k)] = int(signedData[i])
	}

	m.Visit(func(k uint64, v int) {
		g, ok := n[k]
		if !ok {
			t.Errorf("key %d not found.\n", k)
			return
		}
		if g != v {
			t.Errorf("key %d: got %d, want %d\n", k, v, g)
		}
		delete(n, k)
	})

	if len(n) > 0 {
		t.Errorf("Size is %d, want 0\n", len(n))
	}
}

func TestMapUint64ToInt8Size(t *testing.T) {
	m := NewMapUint64ToInt8(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint64(k), int8(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint64(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint64ToInt8Contains(t *testing.T) {
	m := NewMapUint64ToInt8(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), int8(signedData[i]))
		if !m.Contains(uint64(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint64(k))
		if m.Contains(uint64(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint64ToInt8Get(t *testing.T) {
	m := NewMapUint64ToInt8(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), int8(signedData[i]))
		if g := m.Get(uint64(k)); g != int8(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapUint64ToInt8Clear(t *testing.T) {
	m := NewMapUint64ToInt8(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), int8(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint64(k), int8(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUint64ToInt8Inc(t *testing.T) {
	m := NewMapUint64ToInt8(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), int8(signedData[i]))
	}
	for i, k := range unsignedData {
		m.Inc(uint64(k))
		m.Inc(uint64(k))
		m.Inc(uint64(k))
		if g := m.Get(uint64(k)); g != int8(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int8(signedData[i])+3)
		}
	}
	m = NewMapUint64ToInt8(13)
	for _, k := range unsignedData {
		m.Inc(uint64(k))
		if g := m.Get(uint64(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapUint64ToInt8Add(t *testing.T) {
	m := NewMapUint64ToInt8(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), int8(signedData[i]))
	}
	for i, k := range unsignedData {
		m.Add(uint64(k), 3)
		if g := m.Get(uint64(k)); g != int8(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int8(signedData[i])+3)
		}
	}
	m = NewMapUint64ToInt8(13)
	for _, k := range unsignedData {
		m.Add(uint64(k), 42)
		if g := m.Get(uint64(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapUint64ToInt8Visit(t *testing.T) {
	m := NewMapUint64ToInt8(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), int8(signedData[i]))
	}

	n := make(map[uint64]int8, len(unsignedData))

	for i, k := range unsignedData {
		n[uint64(k)] = int8(signedData[i])
	}

	m.Visit(func(k uint64, v int8) {
		g, ok := n[k]
		if !ok {
			t.Errorf("key %d not found.\n", k)
			return
		}
		if g != v {
			t.Errorf("key %d: got %d, want %d\n", k, v, g)
		}
		delete(n, k)
	})

	if len(n) > 0 {
		t.Errorf("Size is %d, want 0\n", len(n))
	}
}

func TestMapUint64ToInt16Size(t *testing.T) {
	m := NewMapUint64ToInt16(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint64(k), int16(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint64(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint64ToInt16Contains(t *testing.T) {
	m := NewMapUint64ToInt16(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), int16(signedData[i]))
		if !m.Contains(uint64(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint64(k))
		if m.Contains(uint64(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint64ToInt16Get(t *testing.T) {
	m := NewMapUint64ToInt16(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), int16(signedData[i]))
		if g := m.Get(uint64(k)); g != int16(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapUint64ToInt16Clear(t *testing.T) {
	m := NewMapUint64ToInt16(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), int16(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint64(k), int16(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUint64ToInt16Inc(t *testing.T) {
	m := NewMapUint64ToInt16(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), int16(signedData[i]))
	}
	for i, k := range unsignedData {
		m.Inc(uint64(k))
		m.Inc(uint64(k))
		m.Inc(uint64(k))
		if g := m.Get(uint64(k)); g != int16(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int16(signedData[i])+3)
		}
	}
	m = NewMapUint64ToInt16(13)
	for _, k := range unsignedData {
		m.Inc(uint64(k))
		if g := m.Get(uint64(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapUint64ToInt16Add(t *testing.T) {
	m := NewMapUint64ToInt16(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), int16(signedData[i]))
	}
	for i, k := range unsignedData {
		m.Add(uint64(k), 3)
		if g := m.Get(uint64(k)); g != int16(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int16(signedData[i])+3)
		}
	}
	m = NewMapUint64ToInt16(13)
	for _, k := range unsignedData {
		m.Add(uint64(k), 42)
		if g := m.Get(uint64(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapUint64ToInt16Visit(t *testing.T) {
	m := NewMapUint64ToInt16(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), int16(signedData[i]))
	}

	n := make(map[uint64]int16, len(unsignedData))

	for i, k := range unsignedData {
		n[uint64(k)] = int16(signedData[i])
	}

	m.Visit(func(k uint64, v int16) {
		g, ok := n[k]
		if !ok {
			t.Errorf("key %d not found.\n", k)
			return
		}
		if g != v {
			t.Errorf("key %d: got %d, want %d\n", k, v, g)
		}
		delete(n, k)
	})

	if len(n) > 0 {
		t.Errorf("Size is %d, want 0\n", len(n))
	}
}

func TestMapUint64ToInt32Size(t *testing.T) {
	m := NewMapUint64ToInt32(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint64(k), int32(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint64(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint64ToInt32Contains(t *testing.T) {
	m := NewMapUint64ToInt32(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), int32(signedData[i]))
		if !m.Contains(uint64(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint64(k))
		if m.Contains(uint64(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint64ToInt32Get(t *testing.T) {
	m := NewMapUint64ToInt32(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), int32(signedData[i]))
		if g := m.Get(uint64(k)); g != int32(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapUint64ToInt32Clear(t *testing.T) {
	m := NewMapUint64ToInt32(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), int32(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint64(k), int32(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUint64ToInt32Inc(t *testing.T) {
	m := NewMapUint64ToInt32(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), int32(signedData[i]))
	}
	for i, k := range unsignedData {
		m.Inc(uint64(k))
		m.Inc(uint64(k))
		m.Inc(uint64(k))
		if g := m.Get(uint64(k)); g != int32(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int32(signedData[i])+3)
		}
	}
	m = NewMapUint64ToInt32(13)
	for _, k := range unsignedData {
		m.Inc(uint64(k))
		if g := m.Get(uint64(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapUint64ToInt32Add(t *testing.T) {
	m := NewMapUint64ToInt32(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), int32(signedData[i]))
	}
	for i, k := range unsignedData {
		m.Add(uint64(k), 3)
		if g := m.Get(uint64(k)); g != int32(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int32(signedData[i])+3)
		}
	}
	m = NewMapUint64ToInt32(13)
	for _, k := range unsignedData {
		m.Add(uint64(k), 42)
		if g := m.Get(uint64(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapUint64ToInt32Visit(t *testing.T) {
	m := NewMapUint64ToInt32(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), int32(signedData[i]))
	}

	n := make(map[uint64]int32, len(unsignedData))

	for i, k := range unsignedData {
		n[uint64(k)] = int32(signedData[i])
	}

	m.Visit(func(k uint64, v int32) {
		g, ok := n[k]
		if !ok {
			t.Errorf("key %d not found.\n", k)
			return
		}
		if g != v {
			t.Errorf("key %d: got %d, want %d\n", k, v, g)
		}
		delete(n, k)
	})

	if len(n) > 0 {
		t.Errorf("Size is %d, want 0\n", len(n))
	}
}

func TestMapUint64ToInt64Size(t *testing.T) {
	m := NewMapUint64ToInt64(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint64(k), int64(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint64(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint64ToInt64Contains(t *testing.T) {
	m := NewMapUint64ToInt64(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), int64(signedData[i]))
		if !m.Contains(uint64(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint64(k))
		if m.Contains(uint64(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint64ToInt64Get(t *testing.T) {
	m := NewMapUint64ToInt64(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), int64(signedData[i]))
		if g := m.Get(uint64(k)); g != int64(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapUint64ToInt64Clear(t *testing.T) {
	m := NewMapUint64ToInt64(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), int64(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint64(k), int64(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUint64ToInt64Inc(t *testing.T) {
	m := NewMapUint64ToInt64(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), int64(signedData[i]))
	}
	for i, k := range unsignedData {
		m.Inc(uint64(k))
		m.Inc(uint64(k))
		m.Inc(uint64(k))
		if g := m.Get(uint64(k)); g != int64(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int64(signedData[i])+3)
		}
	}
	m = NewMapUint64ToInt64(13)
	for _, k := range unsignedData {
		m.Inc(uint64(k))
		if g := m.Get(uint64(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapUint64ToInt64Add(t *testing.T) {
	m := NewMapUint64ToInt64(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), int64(signedData[i]))
	}
	for i, k := range unsignedData {
		m.Add(uint64(k), 3)
		if g := m.Get(uint64(k)); g != int64(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int64(signedData[i])+3)
		}
	}
	m = NewMapUint64ToInt64(13)
	for _, k := range unsignedData {
		m.Add(uint64(k), 42)
		if g := m.Get(uint64(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapUint64ToInt64Visit(t *testing.T) {
	m := NewMapUint64ToInt64(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), int64(signedData[i]))
	}

	n := make(map[uint64]int64, len(unsignedData))

	for i, k := range unsignedData {
		n[uint64(k)] = int64(signedData[i])
	}

	m.Visit(func(k uint64, v int64) {
		g, ok := n[k]
		if !ok {
			t.Errorf("key %d not found.\n", k)
			return
		}
		if g != v {
			t.Errorf("key %d: got %d, want %d\n", k, v, g)
		}
		delete(n, k)
	})

	if len(n) > 0 {
		t.Errorf("Size is %d, want 0\n", len(n))
	}
}

func TestMapUint64ToUintSize(t *testing.T) {
	m := NewMapUint64ToUint(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint64(k), uint(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint64(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint64ToUintContains(t *testing.T) {
	m := NewMapUint64ToUint(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), uint(unsignedData[i]))
		if !m.Contains(uint64(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint64(k))
		if m.Contains(uint64(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint64ToUintGet(t *testing.T) {
	m := NewMapUint64ToUint(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), uint(unsignedData[i]))
		if g := m.Get(uint64(k)); g != uint(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapUint64ToUintClear(t *testing.T) {
	m := NewMapUint64ToUint(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), uint(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint64(k), uint(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUint64ToUintInc(t *testing.T) {
	m := NewMapUint64ToUint(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), uint(unsignedData[i]))
	}
	for i, k := range unsignedData {
		m.Inc(uint64(k))
		m.Inc(uint64(k))
		m.Inc(uint64(k))
		if g := m.Get(uint64(k)); g != uint(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint(unsignedData[i])+3)
		}
	}
	m = NewMapUint64ToUint(13)
	for _, k := range unsignedData {
		m.Inc(uint64(k))
		if g := m.Get(uint64(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapUint64ToUintAdd(t *testing.T) {
	m := NewMapUint64ToUint(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), uint(unsignedData[i]))
	}
	for i, k := range unsignedData {
		m.Add(uint64(k), 3)
		if g := m.Get(uint64(k)); g != uint(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint(unsignedData[i])+3)
		}
	}
	m = NewMapUint64ToUint(13)
	for _, k := range unsignedData {
		m.Add(uint64(k), 42)
		if g := m.Get(uint64(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapUint64ToUintVisit(t *testing.T) {
	m := NewMapUint64ToUint(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), uint(unsignedData[i]))
	}

	n := make(map[uint64]uint, len(unsignedData))

	for i, k := range unsignedData {
		n[uint64(k)] = uint(unsignedData[i])
	}

	m.Visit(func(k uint64, v uint) {
		g, ok := n[k]
		if !ok {
			t.Errorf("key %d not found.\n", k)
			return
		}
		if g != v {
			t.Errorf("key %d: got %d, want %d\n", k, v, g)
		}
		delete(n, k)
	})

	if len(n) > 0 {
		t.Errorf("Size is %d, want 0\n", len(n))
	}
}

func TestMapUint64ToUint8Size(t *testing.T) {
	m := NewMapUint64ToUint8(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint64(k), uint8(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint64(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint64ToUint8Contains(t *testing.T) {
	m := NewMapUint64ToUint8(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), uint8(unsignedData[i]))
		if !m.Contains(uint64(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint64(k))
		if m.Contains(uint64(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint64ToUint8Get(t *testing.T) {
	m := NewMapUint64ToUint8(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), uint8(unsignedData[i]))
		if g := m.Get(uint64(k)); g != uint8(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapUint64ToUint8Clear(t *testing.T) {
	m := NewMapUint64ToUint8(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), uint8(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint64(k), uint8(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUint64ToUint8Inc(t *testing.T) {
	m := NewMapUint64ToUint8(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), uint8(unsignedData[i]))
	}
	for i, k := range unsignedData {
		m.Inc(uint64(k))
		m.Inc(uint64(k))
		m.Inc(uint64(k))
		if g := m.Get(uint64(k)); g != uint8(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint8(unsignedData[i])+3)
		}
	}
	m = NewMapUint64ToUint8(13)
	for _, k := range unsignedData {
		m.Inc(uint64(k))
		if g := m.Get(uint64(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapUint64ToUint8Add(t *testing.T) {
	m := NewMapUint64ToUint8(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), uint8(unsignedData[i]))
	}
	for i, k := range unsignedData {
		m.Add(uint64(k), 3)
		if g := m.Get(uint64(k)); g != uint8(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint8(unsignedData[i])+3)
		}
	}
	m = NewMapUint64ToUint8(13)
	for _, k := range unsignedData {
		m.Add(uint64(k), 42)
		if g := m.Get(uint64(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapUint64ToUint8Visit(t *testing.T) {
	m := NewMapUint64ToUint8(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), uint8(unsignedData[i]))
	}

	n := make(map[uint64]uint8, len(unsignedData))

	for i, k := range unsignedData {
		n[uint64(k)] = uint8(unsignedData[i])
	}

	m.Visit(func(k uint64, v uint8) {
		g, ok := n[k]
		if !ok {
			t.Errorf("key %d not found.\n", k)
			return
		}
		if g != v {
			t.Errorf("key %d: got %d, want %d\n", k, v, g)
		}
		delete(n, k)
	})

	if len(n) > 0 {
		t.Errorf("Size is %d, want 0\n", len(n))
	}
}

func TestMapUint64ToUint16Size(t *testing.T) {
	m := NewMapUint64ToUint16(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint64(k), uint16(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint64(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint64ToUint16Contains(t *testing.T) {
	m := NewMapUint64ToUint16(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), uint16(unsignedData[i]))
		if !m.Contains(uint64(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint64(k))
		if m.Contains(uint64(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint64ToUint16Get(t *testing.T) {
	m := NewMapUint64ToUint16(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), uint16(unsignedData[i]))
		if g := m.Get(uint64(k)); g != uint16(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapUint64ToUint16Clear(t *testing.T) {
	m := NewMapUint64ToUint16(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), uint16(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint64(k), uint16(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUint64ToUint16Inc(t *testing.T) {
	m := NewMapUint64ToUint16(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), uint16(unsignedData[i]))
	}
	for i, k := range unsignedData {
		m.Inc(uint64(k))
		m.Inc(uint64(k))
		m.Inc(uint64(k))
		if g := m.Get(uint64(k)); g != uint16(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint16(unsignedData[i])+3)
		}
	}
	m = NewMapUint64ToUint16(13)
	for _, k := range unsignedData {
		m.Inc(uint64(k))
		if g := m.Get(uint64(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapUint64ToUint16Add(t *testing.T) {
	m := NewMapUint64ToUint16(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), uint16(unsignedData[i]))
	}
	for i, k := range unsignedData {
		m.Add(uint64(k), 3)
		if g := m.Get(uint64(k)); g != uint16(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint16(unsignedData[i])+3)
		}
	}
	m = NewMapUint64ToUint16(13)
	for _, k := range unsignedData {
		m.Add(uint64(k), 42)
		if g := m.Get(uint64(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapUint64ToUint16Visit(t *testing.T) {
	m := NewMapUint64ToUint16(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), uint16(unsignedData[i]))
	}

	n := make(map[uint64]uint16, len(unsignedData))

	for i, k := range unsignedData {
		n[uint64(k)] = uint16(unsignedData[i])
	}

	m.Visit(func(k uint64, v uint16) {
		g, ok := n[k]
		if !ok {
			t.Errorf("key %d not found.\n", k)
			return
		}
		if g != v {
			t.Errorf("key %d: got %d, want %d\n", k, v, g)
		}
		delete(n, k)
	})

	if len(n) > 0 {
		t.Errorf("Size is %d, want 0\n", len(n))
	}
}

func TestMapUint64ToUint32Size(t *testing.T) {
	m := NewMapUint64ToUint32(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint64(k), uint32(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint64(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint64ToUint32Contains(t *testing.T) {
	m := NewMapUint64ToUint32(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), uint32(unsignedData[i]))
		if !m.Contains(uint64(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint64(k))
		if m.Contains(uint64(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint64ToUint32Get(t *testing.T) {
	m := NewMapUint64ToUint32(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), uint32(unsignedData[i]))
		if g := m.Get(uint64(k)); g != uint32(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapUint64ToUint32Clear(t *testing.T) {
	m := NewMapUint64ToUint32(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), uint32(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint64(k), uint32(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUint64ToUint32Inc(t *testing.T) {
	m := NewMapUint64ToUint32(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), uint32(unsignedData[i]))
	}
	for i, k := range unsignedData {
		m.Inc(uint64(k))
		m.Inc(uint64(k))
		m.Inc(uint64(k))
		if g := m.Get(uint64(k)); g != uint32(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint32(unsignedData[i])+3)
		}
	}
	m = NewMapUint64ToUint32(13)
	for _, k := range unsignedData {
		m.Inc(uint64(k))
		if g := m.Get(uint64(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapUint64ToUint32Add(t *testing.T) {
	m := NewMapUint64ToUint32(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), uint32(unsignedData[i]))
	}
	for i, k := range unsignedData {
		m.Add(uint64(k), 3)
		if g := m.Get(uint64(k)); g != uint32(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint32(unsignedData[i])+3)
		}
	}
	m = NewMapUint64ToUint32(13)
	for _, k := range unsignedData {
		m.Add(uint64(k), 42)
		if g := m.Get(uint64(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapUint64ToUint32Visit(t *testing.T) {
	m := NewMapUint64ToUint32(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), uint32(unsignedData[i]))
	}

	n := make(map[uint64]uint32, len(unsignedData))

	for i, k := range unsignedData {
		n[uint64(k)] = uint32(unsignedData[i])
	}

	m.Visit(func(k uint64, v uint32) {
		g, ok := n[k]
		if !ok {
			t.Errorf("key %d not found.\n", k)
			return
		}
		if g != v {
			t.Errorf("key %d: got %d, want %d\n", k, v, g)
		}
		delete(n, k)
	})

	if len(n) > 0 {
		t.Errorf("Size is %d, want 0\n", len(n))
	}
}

func TestMapUint64ToUint64Size(t *testing.T) {
	m := NewMapUint64ToUint64(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint64(k), uint64(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint64(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint64ToUint64Contains(t *testing.T) {
	m := NewMapUint64ToUint64(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), uint64(unsignedData[i]))
		if !m.Contains(uint64(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint64(k))
		if m.Contains(uint64(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint64ToUint64Get(t *testing.T) {
	m := NewMapUint64ToUint64(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), uint64(unsignedData[i]))
		if g := m.Get(uint64(k)); g != uint64(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapUint64ToUint64Clear(t *testing.T) {
	m := NewMapUint64ToUint64(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), uint64(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint64(k), uint64(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUint64ToUint64Inc(t *testing.T) {
	m := NewMapUint64ToUint64(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), uint64(unsignedData[i]))
	}
	for i, k := range unsignedData {
		m.Inc(uint64(k))
		m.Inc(uint64(k))
		m.Inc(uint64(k))
		if g := m.Get(uint64(k)); g != uint64(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint64(unsignedData[i])+3)
		}
	}
	m = NewMapUint64ToUint64(13)
	for _, k := range unsignedData {
		m.Inc(uint64(k))
		if g := m.Get(uint64(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapUint64ToUint64Add(t *testing.T) {
	m := NewMapUint64ToUint64(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), uint64(unsignedData[i]))
	}
	for i, k := range unsignedData {
		m.Add(uint64(k), 3)
		if g := m.Get(uint64(k)); g != uint64(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint64(unsignedData[i])+3)
		}
	}
	m = NewMapUint64ToUint64(13)
	for _, k := range unsignedData {
		m.Add(uint64(k), 42)
		if g := m.Get(uint64(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapUint64ToUint64Visit(t *testing.T) {
	m := NewMapUint64ToUint64(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), uint64(unsignedData[i]))
	}

	n := make(map[uint64]uint64, len(unsignedData))

	for i, k := range unsignedData {
		n[uint64(k)] = uint64(unsignedData[i])
	}

	m.Visit(func(k uint64, v uint64) {
		g, ok := n[k]
		if !ok {
			t.Errorf("key %d not found.\n", k)
			return
		}
		if g != v {
			t.Errorf("key %d: got %d, want %d\n", k, v, g)
		}
		delete(n, k)
	})

	if len(n) > 0 {
		t.Errorf("Size is %d, want 0\n", len(n))
	}
}
