// THIS IS A MACHINE GENERATED FILE!
// BE CAREFUL WITH EDITING BY HAND.
//
// This is Free Software covered by the terms of the MIT license.
// See LICENSE file for details.
// (c) 2016 by Sascha L. Teichmann.
// See the full list of contributors in the CONTRIBUTORS file.

package genmap

import "testing"

func TestMapUintToIntSize(t *testing.T) {
	m := NewMapUintToInt(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint(k), int(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUintToIntContains(t *testing.T) {
	m := NewMapUintToInt(13)
	for i, k := range unsignedData {
		m.Put(uint(k), int(signedData[i]))
		if !m.Contains(uint(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint(k))
		if m.Contains(uint(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUintToIntGet(t *testing.T) {
	m := NewMapUintToInt(13)
	for i, k := range unsignedData {
		m.Put(uint(k), int(signedData[i]))
		if g := m.Get(uint(k)); g != int(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapUintToIntClear(t *testing.T) {
	m := NewMapUintToInt(13)
	for i, k := range unsignedData {
		m.Put(uint(k), int(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint(k), int(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUintToIntInc(t *testing.T) {
	m := NewMapUintToInt(13)
	for i, k := range unsignedData {
		m.Put(uint(k), int(signedData[i]))
	}
	for i, k := range unsignedData {
		m.Inc(uint(k))
		m.Inc(uint(k))
		m.Inc(uint(k))
		if g := m.Get(uint(k)); g != int(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int(signedData[i])+3)
		}
	}
	m = NewMapUintToInt(13)
	for _, k := range unsignedData {
		m.Inc(uint(k))
		if g := m.Get(uint(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapUintToIntAdd(t *testing.T) {
	m := NewMapUintToInt(13)
	for i, k := range unsignedData {
		m.Put(uint(k), int(signedData[i]))
	}
	for i, k := range unsignedData {
		m.Add(uint(k), 3)
		if g := m.Get(uint(k)); g != int(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int(signedData[i])+3)
		}
	}
	m = NewMapUintToInt(13)
	for _, k := range unsignedData {
		m.Add(uint(k), 42)
		if g := m.Get(uint(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapUintToIntModify(t *testing.T) {
	m := NewMapUintToInt(13)
	for i, k := range unsignedData {
		m.Put(uint(k), int(signedData[i]))
	}
	for _, k := range unsignedData {
		m.Modify(uint(k), func(v *int) {
			*v += 3
		})
	}
	for i, k := range unsignedData {
		if g := m.Get(uint(k)); g != int(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int(signedData[i])+3)
		}
	}
}

func TestMapUintToIntPut(t *testing.T) {
	m := NewMapUintToInt(13)
	for i, k := range unsignedData {
		m.Put(uint(k), int(signedData[i]))
	}
	for _, k := range unsignedData {
		m.Put(uint(k), m.Get(uint(k))+3)
	}
	for i, k := range unsignedData {
		if g := m.Get(uint(k)); g != int(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int(signedData[i])+3)
		}
	}
}

func TestMapUintToIntVisit(t *testing.T) {
	m := NewMapUintToInt(13)
	for i, k := range unsignedData {
		m.Put(uint(k), int(signedData[i]))
	}

	n := make(map[uint]int, len(unsignedData))

	for i, k := range unsignedData {
		n[uint(k)] = int(signedData[i])
	}

	m.Visit(func(k uint, v int) {
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

func TestMapUintToInt8Size(t *testing.T) {
	m := NewMapUintToInt8(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint(k), int8(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUintToInt8Contains(t *testing.T) {
	m := NewMapUintToInt8(13)
	for i, k := range unsignedData {
		m.Put(uint(k), int8(signedData[i]))
		if !m.Contains(uint(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint(k))
		if m.Contains(uint(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUintToInt8Get(t *testing.T) {
	m := NewMapUintToInt8(13)
	for i, k := range unsignedData {
		m.Put(uint(k), int8(signedData[i]))
		if g := m.Get(uint(k)); g != int8(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapUintToInt8Clear(t *testing.T) {
	m := NewMapUintToInt8(13)
	for i, k := range unsignedData {
		m.Put(uint(k), int8(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint(k), int8(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUintToInt8Inc(t *testing.T) {
	m := NewMapUintToInt8(13)
	for i, k := range unsignedData {
		m.Put(uint(k), int8(signedData[i]))
	}
	for i, k := range unsignedData {
		m.Inc(uint(k))
		m.Inc(uint(k))
		m.Inc(uint(k))
		if g := m.Get(uint(k)); g != int8(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int8(signedData[i])+3)
		}
	}
	m = NewMapUintToInt8(13)
	for _, k := range unsignedData {
		m.Inc(uint(k))
		if g := m.Get(uint(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapUintToInt8Add(t *testing.T) {
	m := NewMapUintToInt8(13)
	for i, k := range unsignedData {
		m.Put(uint(k), int8(signedData[i]))
	}
	for i, k := range unsignedData {
		m.Add(uint(k), 3)
		if g := m.Get(uint(k)); g != int8(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int8(signedData[i])+3)
		}
	}
	m = NewMapUintToInt8(13)
	for _, k := range unsignedData {
		m.Add(uint(k), 42)
		if g := m.Get(uint(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapUintToInt8Modify(t *testing.T) {
	m := NewMapUintToInt8(13)
	for i, k := range unsignedData {
		m.Put(uint(k), int8(signedData[i]))
	}
	for _, k := range unsignedData {
		m.Modify(uint(k), func(v *int8) {
			*v += 3
		})
	}
	for i, k := range unsignedData {
		if g := m.Get(uint(k)); g != int8(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int8(signedData[i])+3)
		}
	}
}

func TestMapUintToInt8Put(t *testing.T) {
	m := NewMapUintToInt8(13)
	for i, k := range unsignedData {
		m.Put(uint(k), int8(signedData[i]))
	}
	for _, k := range unsignedData {
		m.Put(uint(k), m.Get(uint(k))+3)
	}
	for i, k := range unsignedData {
		if g := m.Get(uint(k)); g != int8(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int8(signedData[i])+3)
		}
	}
}

func TestMapUintToInt8Visit(t *testing.T) {
	m := NewMapUintToInt8(13)
	for i, k := range unsignedData {
		m.Put(uint(k), int8(signedData[i]))
	}

	n := make(map[uint]int8, len(unsignedData))

	for i, k := range unsignedData {
		n[uint(k)] = int8(signedData[i])
	}

	m.Visit(func(k uint, v int8) {
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

func TestMapUintToInt16Size(t *testing.T) {
	m := NewMapUintToInt16(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint(k), int16(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUintToInt16Contains(t *testing.T) {
	m := NewMapUintToInt16(13)
	for i, k := range unsignedData {
		m.Put(uint(k), int16(signedData[i]))
		if !m.Contains(uint(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint(k))
		if m.Contains(uint(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUintToInt16Get(t *testing.T) {
	m := NewMapUintToInt16(13)
	for i, k := range unsignedData {
		m.Put(uint(k), int16(signedData[i]))
		if g := m.Get(uint(k)); g != int16(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapUintToInt16Clear(t *testing.T) {
	m := NewMapUintToInt16(13)
	for i, k := range unsignedData {
		m.Put(uint(k), int16(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint(k), int16(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUintToInt16Inc(t *testing.T) {
	m := NewMapUintToInt16(13)
	for i, k := range unsignedData {
		m.Put(uint(k), int16(signedData[i]))
	}
	for i, k := range unsignedData {
		m.Inc(uint(k))
		m.Inc(uint(k))
		m.Inc(uint(k))
		if g := m.Get(uint(k)); g != int16(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int16(signedData[i])+3)
		}
	}
	m = NewMapUintToInt16(13)
	for _, k := range unsignedData {
		m.Inc(uint(k))
		if g := m.Get(uint(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapUintToInt16Add(t *testing.T) {
	m := NewMapUintToInt16(13)
	for i, k := range unsignedData {
		m.Put(uint(k), int16(signedData[i]))
	}
	for i, k := range unsignedData {
		m.Add(uint(k), 3)
		if g := m.Get(uint(k)); g != int16(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int16(signedData[i])+3)
		}
	}
	m = NewMapUintToInt16(13)
	for _, k := range unsignedData {
		m.Add(uint(k), 42)
		if g := m.Get(uint(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapUintToInt16Modify(t *testing.T) {
	m := NewMapUintToInt16(13)
	for i, k := range unsignedData {
		m.Put(uint(k), int16(signedData[i]))
	}
	for _, k := range unsignedData {
		m.Modify(uint(k), func(v *int16) {
			*v += 3
		})
	}
	for i, k := range unsignedData {
		if g := m.Get(uint(k)); g != int16(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int16(signedData[i])+3)
		}
	}
}

func TestMapUintToInt16Put(t *testing.T) {
	m := NewMapUintToInt16(13)
	for i, k := range unsignedData {
		m.Put(uint(k), int16(signedData[i]))
	}
	for _, k := range unsignedData {
		m.Put(uint(k), m.Get(uint(k))+3)
	}
	for i, k := range unsignedData {
		if g := m.Get(uint(k)); g != int16(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int16(signedData[i])+3)
		}
	}
}

func TestMapUintToInt16Visit(t *testing.T) {
	m := NewMapUintToInt16(13)
	for i, k := range unsignedData {
		m.Put(uint(k), int16(signedData[i]))
	}

	n := make(map[uint]int16, len(unsignedData))

	for i, k := range unsignedData {
		n[uint(k)] = int16(signedData[i])
	}

	m.Visit(func(k uint, v int16) {
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

func TestMapUintToInt32Size(t *testing.T) {
	m := NewMapUintToInt32(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint(k), int32(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUintToInt32Contains(t *testing.T) {
	m := NewMapUintToInt32(13)
	for i, k := range unsignedData {
		m.Put(uint(k), int32(signedData[i]))
		if !m.Contains(uint(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint(k))
		if m.Contains(uint(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUintToInt32Get(t *testing.T) {
	m := NewMapUintToInt32(13)
	for i, k := range unsignedData {
		m.Put(uint(k), int32(signedData[i]))
		if g := m.Get(uint(k)); g != int32(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapUintToInt32Clear(t *testing.T) {
	m := NewMapUintToInt32(13)
	for i, k := range unsignedData {
		m.Put(uint(k), int32(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint(k), int32(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUintToInt32Inc(t *testing.T) {
	m := NewMapUintToInt32(13)
	for i, k := range unsignedData {
		m.Put(uint(k), int32(signedData[i]))
	}
	for i, k := range unsignedData {
		m.Inc(uint(k))
		m.Inc(uint(k))
		m.Inc(uint(k))
		if g := m.Get(uint(k)); g != int32(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int32(signedData[i])+3)
		}
	}
	m = NewMapUintToInt32(13)
	for _, k := range unsignedData {
		m.Inc(uint(k))
		if g := m.Get(uint(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapUintToInt32Add(t *testing.T) {
	m := NewMapUintToInt32(13)
	for i, k := range unsignedData {
		m.Put(uint(k), int32(signedData[i]))
	}
	for i, k := range unsignedData {
		m.Add(uint(k), 3)
		if g := m.Get(uint(k)); g != int32(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int32(signedData[i])+3)
		}
	}
	m = NewMapUintToInt32(13)
	for _, k := range unsignedData {
		m.Add(uint(k), 42)
		if g := m.Get(uint(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapUintToInt32Modify(t *testing.T) {
	m := NewMapUintToInt32(13)
	for i, k := range unsignedData {
		m.Put(uint(k), int32(signedData[i]))
	}
	for _, k := range unsignedData {
		m.Modify(uint(k), func(v *int32) {
			*v += 3
		})
	}
	for i, k := range unsignedData {
		if g := m.Get(uint(k)); g != int32(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int32(signedData[i])+3)
		}
	}
}

func TestMapUintToInt32Put(t *testing.T) {
	m := NewMapUintToInt32(13)
	for i, k := range unsignedData {
		m.Put(uint(k), int32(signedData[i]))
	}
	for _, k := range unsignedData {
		m.Put(uint(k), m.Get(uint(k))+3)
	}
	for i, k := range unsignedData {
		if g := m.Get(uint(k)); g != int32(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int32(signedData[i])+3)
		}
	}
}

func TestMapUintToInt32Visit(t *testing.T) {
	m := NewMapUintToInt32(13)
	for i, k := range unsignedData {
		m.Put(uint(k), int32(signedData[i]))
	}

	n := make(map[uint]int32, len(unsignedData))

	for i, k := range unsignedData {
		n[uint(k)] = int32(signedData[i])
	}

	m.Visit(func(k uint, v int32) {
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

func TestMapUintToInt64Size(t *testing.T) {
	m := NewMapUintToInt64(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint(k), int64(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUintToInt64Contains(t *testing.T) {
	m := NewMapUintToInt64(13)
	for i, k := range unsignedData {
		m.Put(uint(k), int64(signedData[i]))
		if !m.Contains(uint(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint(k))
		if m.Contains(uint(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUintToInt64Get(t *testing.T) {
	m := NewMapUintToInt64(13)
	for i, k := range unsignedData {
		m.Put(uint(k), int64(signedData[i]))
		if g := m.Get(uint(k)); g != int64(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapUintToInt64Clear(t *testing.T) {
	m := NewMapUintToInt64(13)
	for i, k := range unsignedData {
		m.Put(uint(k), int64(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint(k), int64(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUintToInt64Inc(t *testing.T) {
	m := NewMapUintToInt64(13)
	for i, k := range unsignedData {
		m.Put(uint(k), int64(signedData[i]))
	}
	for i, k := range unsignedData {
		m.Inc(uint(k))
		m.Inc(uint(k))
		m.Inc(uint(k))
		if g := m.Get(uint(k)); g != int64(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int64(signedData[i])+3)
		}
	}
	m = NewMapUintToInt64(13)
	for _, k := range unsignedData {
		m.Inc(uint(k))
		if g := m.Get(uint(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapUintToInt64Add(t *testing.T) {
	m := NewMapUintToInt64(13)
	for i, k := range unsignedData {
		m.Put(uint(k), int64(signedData[i]))
	}
	for i, k := range unsignedData {
		m.Add(uint(k), 3)
		if g := m.Get(uint(k)); g != int64(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int64(signedData[i])+3)
		}
	}
	m = NewMapUintToInt64(13)
	for _, k := range unsignedData {
		m.Add(uint(k), 42)
		if g := m.Get(uint(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapUintToInt64Modify(t *testing.T) {
	m := NewMapUintToInt64(13)
	for i, k := range unsignedData {
		m.Put(uint(k), int64(signedData[i]))
	}
	for _, k := range unsignedData {
		m.Modify(uint(k), func(v *int64) {
			*v += 3
		})
	}
	for i, k := range unsignedData {
		if g := m.Get(uint(k)); g != int64(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int64(signedData[i])+3)
		}
	}
}

func TestMapUintToInt64Put(t *testing.T) {
	m := NewMapUintToInt64(13)
	for i, k := range unsignedData {
		m.Put(uint(k), int64(signedData[i]))
	}
	for _, k := range unsignedData {
		m.Put(uint(k), m.Get(uint(k))+3)
	}
	for i, k := range unsignedData {
		if g := m.Get(uint(k)); g != int64(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int64(signedData[i])+3)
		}
	}
}

func TestMapUintToInt64Visit(t *testing.T) {
	m := NewMapUintToInt64(13)
	for i, k := range unsignedData {
		m.Put(uint(k), int64(signedData[i]))
	}

	n := make(map[uint]int64, len(unsignedData))

	for i, k := range unsignedData {
		n[uint(k)] = int64(signedData[i])
	}

	m.Visit(func(k uint, v int64) {
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

func TestMapUintToUintSize(t *testing.T) {
	m := NewMapUintToUint(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint(k), uint(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUintToUintContains(t *testing.T) {
	m := NewMapUintToUint(13)
	for i, k := range unsignedData {
		m.Put(uint(k), uint(unsignedData[i]))
		if !m.Contains(uint(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint(k))
		if m.Contains(uint(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUintToUintGet(t *testing.T) {
	m := NewMapUintToUint(13)
	for i, k := range unsignedData {
		m.Put(uint(k), uint(unsignedData[i]))
		if g := m.Get(uint(k)); g != uint(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapUintToUintClear(t *testing.T) {
	m := NewMapUintToUint(13)
	for i, k := range unsignedData {
		m.Put(uint(k), uint(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint(k), uint(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUintToUintInc(t *testing.T) {
	m := NewMapUintToUint(13)
	for i, k := range unsignedData {
		m.Put(uint(k), uint(unsignedData[i]))
	}
	for i, k := range unsignedData {
		m.Inc(uint(k))
		m.Inc(uint(k))
		m.Inc(uint(k))
		if g := m.Get(uint(k)); g != uint(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint(unsignedData[i])+3)
		}
	}
	m = NewMapUintToUint(13)
	for _, k := range unsignedData {
		m.Inc(uint(k))
		if g := m.Get(uint(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapUintToUintAdd(t *testing.T) {
	m := NewMapUintToUint(13)
	for i, k := range unsignedData {
		m.Put(uint(k), uint(unsignedData[i]))
	}
	for i, k := range unsignedData {
		m.Add(uint(k), 3)
		if g := m.Get(uint(k)); g != uint(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint(unsignedData[i])+3)
		}
	}
	m = NewMapUintToUint(13)
	for _, k := range unsignedData {
		m.Add(uint(k), 42)
		if g := m.Get(uint(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapUintToUintModify(t *testing.T) {
	m := NewMapUintToUint(13)
	for i, k := range unsignedData {
		m.Put(uint(k), uint(unsignedData[i]))
	}
	for _, k := range unsignedData {
		m.Modify(uint(k), func(v *uint) {
			*v += 3
		})
	}
	for i, k := range unsignedData {
		if g := m.Get(uint(k)); g != uint(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint(unsignedData[i])+3)
		}
	}
}

func TestMapUintToUintPut(t *testing.T) {
	m := NewMapUintToUint(13)
	for i, k := range unsignedData {
		m.Put(uint(k), uint(unsignedData[i]))
	}
	for _, k := range unsignedData {
		m.Put(uint(k), m.Get(uint(k))+3)
	}
	for i, k := range unsignedData {
		if g := m.Get(uint(k)); g != uint(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint(unsignedData[i])+3)
		}
	}
}

func TestMapUintToUintVisit(t *testing.T) {
	m := NewMapUintToUint(13)
	for i, k := range unsignedData {
		m.Put(uint(k), uint(unsignedData[i]))
	}

	n := make(map[uint]uint, len(unsignedData))

	for i, k := range unsignedData {
		n[uint(k)] = uint(unsignedData[i])
	}

	m.Visit(func(k uint, v uint) {
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

func TestMapUintToUint8Size(t *testing.T) {
	m := NewMapUintToUint8(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint(k), uint8(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUintToUint8Contains(t *testing.T) {
	m := NewMapUintToUint8(13)
	for i, k := range unsignedData {
		m.Put(uint(k), uint8(unsignedData[i]))
		if !m.Contains(uint(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint(k))
		if m.Contains(uint(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUintToUint8Get(t *testing.T) {
	m := NewMapUintToUint8(13)
	for i, k := range unsignedData {
		m.Put(uint(k), uint8(unsignedData[i]))
		if g := m.Get(uint(k)); g != uint8(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapUintToUint8Clear(t *testing.T) {
	m := NewMapUintToUint8(13)
	for i, k := range unsignedData {
		m.Put(uint(k), uint8(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint(k), uint8(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUintToUint8Inc(t *testing.T) {
	m := NewMapUintToUint8(13)
	for i, k := range unsignedData {
		m.Put(uint(k), uint8(unsignedData[i]))
	}
	for i, k := range unsignedData {
		m.Inc(uint(k))
		m.Inc(uint(k))
		m.Inc(uint(k))
		if g := m.Get(uint(k)); g != uint8(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint8(unsignedData[i])+3)
		}
	}
	m = NewMapUintToUint8(13)
	for _, k := range unsignedData {
		m.Inc(uint(k))
		if g := m.Get(uint(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapUintToUint8Add(t *testing.T) {
	m := NewMapUintToUint8(13)
	for i, k := range unsignedData {
		m.Put(uint(k), uint8(unsignedData[i]))
	}
	for i, k := range unsignedData {
		m.Add(uint(k), 3)
		if g := m.Get(uint(k)); g != uint8(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint8(unsignedData[i])+3)
		}
	}
	m = NewMapUintToUint8(13)
	for _, k := range unsignedData {
		m.Add(uint(k), 42)
		if g := m.Get(uint(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapUintToUint8Modify(t *testing.T) {
	m := NewMapUintToUint8(13)
	for i, k := range unsignedData {
		m.Put(uint(k), uint8(unsignedData[i]))
	}
	for _, k := range unsignedData {
		m.Modify(uint(k), func(v *uint8) {
			*v += 3
		})
	}
	for i, k := range unsignedData {
		if g := m.Get(uint(k)); g != uint8(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint8(unsignedData[i])+3)
		}
	}
}

func TestMapUintToUint8Put(t *testing.T) {
	m := NewMapUintToUint8(13)
	for i, k := range unsignedData {
		m.Put(uint(k), uint8(unsignedData[i]))
	}
	for _, k := range unsignedData {
		m.Put(uint(k), m.Get(uint(k))+3)
	}
	for i, k := range unsignedData {
		if g := m.Get(uint(k)); g != uint8(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint8(unsignedData[i])+3)
		}
	}
}

func TestMapUintToUint8Visit(t *testing.T) {
	m := NewMapUintToUint8(13)
	for i, k := range unsignedData {
		m.Put(uint(k), uint8(unsignedData[i]))
	}

	n := make(map[uint]uint8, len(unsignedData))

	for i, k := range unsignedData {
		n[uint(k)] = uint8(unsignedData[i])
	}

	m.Visit(func(k uint, v uint8) {
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

func TestMapUintToUint16Size(t *testing.T) {
	m := NewMapUintToUint16(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint(k), uint16(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUintToUint16Contains(t *testing.T) {
	m := NewMapUintToUint16(13)
	for i, k := range unsignedData {
		m.Put(uint(k), uint16(unsignedData[i]))
		if !m.Contains(uint(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint(k))
		if m.Contains(uint(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUintToUint16Get(t *testing.T) {
	m := NewMapUintToUint16(13)
	for i, k := range unsignedData {
		m.Put(uint(k), uint16(unsignedData[i]))
		if g := m.Get(uint(k)); g != uint16(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapUintToUint16Clear(t *testing.T) {
	m := NewMapUintToUint16(13)
	for i, k := range unsignedData {
		m.Put(uint(k), uint16(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint(k), uint16(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUintToUint16Inc(t *testing.T) {
	m := NewMapUintToUint16(13)
	for i, k := range unsignedData {
		m.Put(uint(k), uint16(unsignedData[i]))
	}
	for i, k := range unsignedData {
		m.Inc(uint(k))
		m.Inc(uint(k))
		m.Inc(uint(k))
		if g := m.Get(uint(k)); g != uint16(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint16(unsignedData[i])+3)
		}
	}
	m = NewMapUintToUint16(13)
	for _, k := range unsignedData {
		m.Inc(uint(k))
		if g := m.Get(uint(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapUintToUint16Add(t *testing.T) {
	m := NewMapUintToUint16(13)
	for i, k := range unsignedData {
		m.Put(uint(k), uint16(unsignedData[i]))
	}
	for i, k := range unsignedData {
		m.Add(uint(k), 3)
		if g := m.Get(uint(k)); g != uint16(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint16(unsignedData[i])+3)
		}
	}
	m = NewMapUintToUint16(13)
	for _, k := range unsignedData {
		m.Add(uint(k), 42)
		if g := m.Get(uint(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapUintToUint16Modify(t *testing.T) {
	m := NewMapUintToUint16(13)
	for i, k := range unsignedData {
		m.Put(uint(k), uint16(unsignedData[i]))
	}
	for _, k := range unsignedData {
		m.Modify(uint(k), func(v *uint16) {
			*v += 3
		})
	}
	for i, k := range unsignedData {
		if g := m.Get(uint(k)); g != uint16(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint16(unsignedData[i])+3)
		}
	}
}

func TestMapUintToUint16Put(t *testing.T) {
	m := NewMapUintToUint16(13)
	for i, k := range unsignedData {
		m.Put(uint(k), uint16(unsignedData[i]))
	}
	for _, k := range unsignedData {
		m.Put(uint(k), m.Get(uint(k))+3)
	}
	for i, k := range unsignedData {
		if g := m.Get(uint(k)); g != uint16(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint16(unsignedData[i])+3)
		}
	}
}

func TestMapUintToUint16Visit(t *testing.T) {
	m := NewMapUintToUint16(13)
	for i, k := range unsignedData {
		m.Put(uint(k), uint16(unsignedData[i]))
	}

	n := make(map[uint]uint16, len(unsignedData))

	for i, k := range unsignedData {
		n[uint(k)] = uint16(unsignedData[i])
	}

	m.Visit(func(k uint, v uint16) {
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

func TestMapUintToUint32Size(t *testing.T) {
	m := NewMapUintToUint32(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint(k), uint32(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUintToUint32Contains(t *testing.T) {
	m := NewMapUintToUint32(13)
	for i, k := range unsignedData {
		m.Put(uint(k), uint32(unsignedData[i]))
		if !m.Contains(uint(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint(k))
		if m.Contains(uint(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUintToUint32Get(t *testing.T) {
	m := NewMapUintToUint32(13)
	for i, k := range unsignedData {
		m.Put(uint(k), uint32(unsignedData[i]))
		if g := m.Get(uint(k)); g != uint32(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapUintToUint32Clear(t *testing.T) {
	m := NewMapUintToUint32(13)
	for i, k := range unsignedData {
		m.Put(uint(k), uint32(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint(k), uint32(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUintToUint32Inc(t *testing.T) {
	m := NewMapUintToUint32(13)
	for i, k := range unsignedData {
		m.Put(uint(k), uint32(unsignedData[i]))
	}
	for i, k := range unsignedData {
		m.Inc(uint(k))
		m.Inc(uint(k))
		m.Inc(uint(k))
		if g := m.Get(uint(k)); g != uint32(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint32(unsignedData[i])+3)
		}
	}
	m = NewMapUintToUint32(13)
	for _, k := range unsignedData {
		m.Inc(uint(k))
		if g := m.Get(uint(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapUintToUint32Add(t *testing.T) {
	m := NewMapUintToUint32(13)
	for i, k := range unsignedData {
		m.Put(uint(k), uint32(unsignedData[i]))
	}
	for i, k := range unsignedData {
		m.Add(uint(k), 3)
		if g := m.Get(uint(k)); g != uint32(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint32(unsignedData[i])+3)
		}
	}
	m = NewMapUintToUint32(13)
	for _, k := range unsignedData {
		m.Add(uint(k), 42)
		if g := m.Get(uint(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapUintToUint32Modify(t *testing.T) {
	m := NewMapUintToUint32(13)
	for i, k := range unsignedData {
		m.Put(uint(k), uint32(unsignedData[i]))
	}
	for _, k := range unsignedData {
		m.Modify(uint(k), func(v *uint32) {
			*v += 3
		})
	}
	for i, k := range unsignedData {
		if g := m.Get(uint(k)); g != uint32(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint32(unsignedData[i])+3)
		}
	}
}

func TestMapUintToUint32Put(t *testing.T) {
	m := NewMapUintToUint32(13)
	for i, k := range unsignedData {
		m.Put(uint(k), uint32(unsignedData[i]))
	}
	for _, k := range unsignedData {
		m.Put(uint(k), m.Get(uint(k))+3)
	}
	for i, k := range unsignedData {
		if g := m.Get(uint(k)); g != uint32(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint32(unsignedData[i])+3)
		}
	}
}

func TestMapUintToUint32Visit(t *testing.T) {
	m := NewMapUintToUint32(13)
	for i, k := range unsignedData {
		m.Put(uint(k), uint32(unsignedData[i]))
	}

	n := make(map[uint]uint32, len(unsignedData))

	for i, k := range unsignedData {
		n[uint(k)] = uint32(unsignedData[i])
	}

	m.Visit(func(k uint, v uint32) {
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

func TestMapUintToUint64Size(t *testing.T) {
	m := NewMapUintToUint64(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint(k), uint64(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUintToUint64Contains(t *testing.T) {
	m := NewMapUintToUint64(13)
	for i, k := range unsignedData {
		m.Put(uint(k), uint64(unsignedData[i]))
		if !m.Contains(uint(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint(k))
		if m.Contains(uint(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUintToUint64Get(t *testing.T) {
	m := NewMapUintToUint64(13)
	for i, k := range unsignedData {
		m.Put(uint(k), uint64(unsignedData[i]))
		if g := m.Get(uint(k)); g != uint64(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapUintToUint64Clear(t *testing.T) {
	m := NewMapUintToUint64(13)
	for i, k := range unsignedData {
		m.Put(uint(k), uint64(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint(k), uint64(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUintToUint64Inc(t *testing.T) {
	m := NewMapUintToUint64(13)
	for i, k := range unsignedData {
		m.Put(uint(k), uint64(unsignedData[i]))
	}
	for i, k := range unsignedData {
		m.Inc(uint(k))
		m.Inc(uint(k))
		m.Inc(uint(k))
		if g := m.Get(uint(k)); g != uint64(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint64(unsignedData[i])+3)
		}
	}
	m = NewMapUintToUint64(13)
	for _, k := range unsignedData {
		m.Inc(uint(k))
		if g := m.Get(uint(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapUintToUint64Add(t *testing.T) {
	m := NewMapUintToUint64(13)
	for i, k := range unsignedData {
		m.Put(uint(k), uint64(unsignedData[i]))
	}
	for i, k := range unsignedData {
		m.Add(uint(k), 3)
		if g := m.Get(uint(k)); g != uint64(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint64(unsignedData[i])+3)
		}
	}
	m = NewMapUintToUint64(13)
	for _, k := range unsignedData {
		m.Add(uint(k), 42)
		if g := m.Get(uint(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapUintToUint64Modify(t *testing.T) {
	m := NewMapUintToUint64(13)
	for i, k := range unsignedData {
		m.Put(uint(k), uint64(unsignedData[i]))
	}
	for _, k := range unsignedData {
		m.Modify(uint(k), func(v *uint64) {
			*v += 3
		})
	}
	for i, k := range unsignedData {
		if g := m.Get(uint(k)); g != uint64(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint64(unsignedData[i])+3)
		}
	}
}

func TestMapUintToUint64Put(t *testing.T) {
	m := NewMapUintToUint64(13)
	for i, k := range unsignedData {
		m.Put(uint(k), uint64(unsignedData[i]))
	}
	for _, k := range unsignedData {
		m.Put(uint(k), m.Get(uint(k))+3)
	}
	for i, k := range unsignedData {
		if g := m.Get(uint(k)); g != uint64(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint64(unsignedData[i])+3)
		}
	}
}

func TestMapUintToUint64Visit(t *testing.T) {
	m := NewMapUintToUint64(13)
	for i, k := range unsignedData {
		m.Put(uint(k), uint64(unsignedData[i]))
	}

	n := make(map[uint]uint64, len(unsignedData))

	for i, k := range unsignedData {
		n[uint(k)] = uint64(unsignedData[i])
	}

	m.Visit(func(k uint, v uint64) {
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
