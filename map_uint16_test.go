// THIS IS A MACHINE GENERATED FILE!
// BE CAREFUL WITH EDITING BY HAND.
//
// This is Free Software covered by the terms of the MIT license.
// See LICENSE file for details.
// (c) 2016 by Sascha L. Teichmann.
// See the full list of contributors in the CONTRIBUTORS file.

package genmap

import "testing"

func TestMapUint16ToIntSize(t *testing.T) {
	m := NewMapUint16ToInt(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint16(k), int(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint16(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint16ToIntContains(t *testing.T) {
	m := NewMapUint16ToInt(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), int(signedData[i]))
		if !m.Contains(uint16(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint16(k))
		if m.Contains(uint16(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint16ToIntGet(t *testing.T) {
	m := NewMapUint16ToInt(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), int(signedData[i]))
		if g := m.Get(uint16(k)); g != int(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
}

func TestMapUint16ToIntClear(t *testing.T) {
	m := NewMapUint16ToInt(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), int(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint16(k), int(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUint16ToIntInc(t *testing.T) {
	m := NewMapUint16ToInt(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), int(signedData[i]))
	}
	for i, k := range unsignedData {
		m.Inc(uint16(k))
		m.Inc(uint16(k))
		m.Inc(uint16(k))
		if g := m.Get(uint16(k)); g != int(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int(signedData[i])+3)
		}
	}
	m.Clear()
	for _, k := range unsignedData {
		m.Inc(uint16(k))
		if g := m.Get(uint16(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapUint16ToIntAdd(t *testing.T) {
	m := NewMapUint16ToInt(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), int(signedData[i]))
	}
	for i, k := range unsignedData {
		m.Add(uint16(k), 3)
		if g := m.Get(uint16(k)); g != int(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int(signedData[i])+3)
		}
	}
	m.Clear()
	for _, k := range unsignedData {
		m.Add(uint16(k), 42)
		if g := m.Get(uint16(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapUint16ToIntVisit(t *testing.T) {
	m := NewMapUint16ToInt(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), int(signedData[i]))
	}

	n := make(map[uint16]int, len(unsignedData))

	for i, k := range unsignedData {
		n[uint16(k)] = int(signedData[i])
	}

	m.Visit(func(k uint16, v int) {
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

func TestMapUint16ToInt8Size(t *testing.T) {
	m := NewMapUint16ToInt8(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint16(k), int8(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint16(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint16ToInt8Contains(t *testing.T) {
	m := NewMapUint16ToInt8(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), int8(signedData[i]))
		if !m.Contains(uint16(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint16(k))
		if m.Contains(uint16(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint16ToInt8Get(t *testing.T) {
	m := NewMapUint16ToInt8(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), int8(signedData[i]))
		if g := m.Get(uint16(k)); g != int8(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
}

func TestMapUint16ToInt8Clear(t *testing.T) {
	m := NewMapUint16ToInt8(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), int8(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint16(k), int8(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUint16ToInt8Inc(t *testing.T) {
	m := NewMapUint16ToInt8(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), int8(signedData[i]))
	}
	for i, k := range unsignedData {
		m.Inc(uint16(k))
		m.Inc(uint16(k))
		m.Inc(uint16(k))
		if g := m.Get(uint16(k)); g != int8(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int8(signedData[i])+3)
		}
	}
	m.Clear()
	for _, k := range unsignedData {
		m.Inc(uint16(k))
		if g := m.Get(uint16(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapUint16ToInt8Add(t *testing.T) {
	m := NewMapUint16ToInt8(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), int8(signedData[i]))
	}
	for i, k := range unsignedData {
		m.Add(uint16(k), 3)
		if g := m.Get(uint16(k)); g != int8(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int8(signedData[i])+3)
		}
	}
	m.Clear()
	for _, k := range unsignedData {
		m.Add(uint16(k), 42)
		if g := m.Get(uint16(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapUint16ToInt8Visit(t *testing.T) {
	m := NewMapUint16ToInt8(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), int8(signedData[i]))
	}

	n := make(map[uint16]int8, len(unsignedData))

	for i, k := range unsignedData {
		n[uint16(k)] = int8(signedData[i])
	}

	m.Visit(func(k uint16, v int8) {
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

func TestMapUint16ToInt16Size(t *testing.T) {
	m := NewMapUint16ToInt16(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint16(k), int16(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint16(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint16ToInt16Contains(t *testing.T) {
	m := NewMapUint16ToInt16(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), int16(signedData[i]))
		if !m.Contains(uint16(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint16(k))
		if m.Contains(uint16(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint16ToInt16Get(t *testing.T) {
	m := NewMapUint16ToInt16(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), int16(signedData[i]))
		if g := m.Get(uint16(k)); g != int16(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
}

func TestMapUint16ToInt16Clear(t *testing.T) {
	m := NewMapUint16ToInt16(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), int16(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint16(k), int16(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUint16ToInt16Inc(t *testing.T) {
	m := NewMapUint16ToInt16(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), int16(signedData[i]))
	}
	for i, k := range unsignedData {
		m.Inc(uint16(k))
		m.Inc(uint16(k))
		m.Inc(uint16(k))
		if g := m.Get(uint16(k)); g != int16(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int16(signedData[i])+3)
		}
	}
	m.Clear()
	for _, k := range unsignedData {
		m.Inc(uint16(k))
		if g := m.Get(uint16(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapUint16ToInt16Add(t *testing.T) {
	m := NewMapUint16ToInt16(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), int16(signedData[i]))
	}
	for i, k := range unsignedData {
		m.Add(uint16(k), 3)
		if g := m.Get(uint16(k)); g != int16(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int16(signedData[i])+3)
		}
	}
	m.Clear()
	for _, k := range unsignedData {
		m.Add(uint16(k), 42)
		if g := m.Get(uint16(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapUint16ToInt16Visit(t *testing.T) {
	m := NewMapUint16ToInt16(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), int16(signedData[i]))
	}

	n := make(map[uint16]int16, len(unsignedData))

	for i, k := range unsignedData {
		n[uint16(k)] = int16(signedData[i])
	}

	m.Visit(func(k uint16, v int16) {
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

func TestMapUint16ToInt32Size(t *testing.T) {
	m := NewMapUint16ToInt32(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint16(k), int32(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint16(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint16ToInt32Contains(t *testing.T) {
	m := NewMapUint16ToInt32(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), int32(signedData[i]))
		if !m.Contains(uint16(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint16(k))
		if m.Contains(uint16(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint16ToInt32Get(t *testing.T) {
	m := NewMapUint16ToInt32(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), int32(signedData[i]))
		if g := m.Get(uint16(k)); g != int32(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
}

func TestMapUint16ToInt32Clear(t *testing.T) {
	m := NewMapUint16ToInt32(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), int32(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint16(k), int32(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUint16ToInt32Inc(t *testing.T) {
	m := NewMapUint16ToInt32(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), int32(signedData[i]))
	}
	for i, k := range unsignedData {
		m.Inc(uint16(k))
		m.Inc(uint16(k))
		m.Inc(uint16(k))
		if g := m.Get(uint16(k)); g != int32(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int32(signedData[i])+3)
		}
	}
	m.Clear()
	for _, k := range unsignedData {
		m.Inc(uint16(k))
		if g := m.Get(uint16(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapUint16ToInt32Add(t *testing.T) {
	m := NewMapUint16ToInt32(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), int32(signedData[i]))
	}
	for i, k := range unsignedData {
		m.Add(uint16(k), 3)
		if g := m.Get(uint16(k)); g != int32(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int32(signedData[i])+3)
		}
	}
	m.Clear()
	for _, k := range unsignedData {
		m.Add(uint16(k), 42)
		if g := m.Get(uint16(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapUint16ToInt32Visit(t *testing.T) {
	m := NewMapUint16ToInt32(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), int32(signedData[i]))
	}

	n := make(map[uint16]int32, len(unsignedData))

	for i, k := range unsignedData {
		n[uint16(k)] = int32(signedData[i])
	}

	m.Visit(func(k uint16, v int32) {
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

func TestMapUint16ToInt64Size(t *testing.T) {
	m := NewMapUint16ToInt64(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint16(k), int64(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint16(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint16ToInt64Contains(t *testing.T) {
	m := NewMapUint16ToInt64(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), int64(signedData[i]))
		if !m.Contains(uint16(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint16(k))
		if m.Contains(uint16(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint16ToInt64Get(t *testing.T) {
	m := NewMapUint16ToInt64(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), int64(signedData[i]))
		if g := m.Get(uint16(k)); g != int64(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
}

func TestMapUint16ToInt64Clear(t *testing.T) {
	m := NewMapUint16ToInt64(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), int64(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint16(k), int64(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUint16ToInt64Inc(t *testing.T) {
	m := NewMapUint16ToInt64(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), int64(signedData[i]))
	}
	for i, k := range unsignedData {
		m.Inc(uint16(k))
		m.Inc(uint16(k))
		m.Inc(uint16(k))
		if g := m.Get(uint16(k)); g != int64(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int64(signedData[i])+3)
		}
	}
	m.Clear()
	for _, k := range unsignedData {
		m.Inc(uint16(k))
		if g := m.Get(uint16(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapUint16ToInt64Add(t *testing.T) {
	m := NewMapUint16ToInt64(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), int64(signedData[i]))
	}
	for i, k := range unsignedData {
		m.Add(uint16(k), 3)
		if g := m.Get(uint16(k)); g != int64(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int64(signedData[i])+3)
		}
	}
	m.Clear()
	for _, k := range unsignedData {
		m.Add(uint16(k), 42)
		if g := m.Get(uint16(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapUint16ToInt64Visit(t *testing.T) {
	m := NewMapUint16ToInt64(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), int64(signedData[i]))
	}

	n := make(map[uint16]int64, len(unsignedData))

	for i, k := range unsignedData {
		n[uint16(k)] = int64(signedData[i])
	}

	m.Visit(func(k uint16, v int64) {
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

func TestMapUint16ToUintSize(t *testing.T) {
	m := NewMapUint16ToUint(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint16(k), uint(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint16(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint16ToUintContains(t *testing.T) {
	m := NewMapUint16ToUint(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), uint(unsignedData[i]))
		if !m.Contains(uint16(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint16(k))
		if m.Contains(uint16(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint16ToUintGet(t *testing.T) {
	m := NewMapUint16ToUint(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), uint(unsignedData[i]))
		if g := m.Get(uint16(k)); g != uint(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
}

func TestMapUint16ToUintClear(t *testing.T) {
	m := NewMapUint16ToUint(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), uint(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint16(k), uint(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUint16ToUintInc(t *testing.T) {
	m := NewMapUint16ToUint(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), uint(unsignedData[i]))
	}
	for i, k := range unsignedData {
		m.Inc(uint16(k))
		m.Inc(uint16(k))
		m.Inc(uint16(k))
		if g := m.Get(uint16(k)); g != uint(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint(unsignedData[i])+3)
		}
	}
	m.Clear()
	for _, k := range unsignedData {
		m.Inc(uint16(k))
		if g := m.Get(uint16(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapUint16ToUintAdd(t *testing.T) {
	m := NewMapUint16ToUint(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), uint(unsignedData[i]))
	}
	for i, k := range unsignedData {
		m.Add(uint16(k), 3)
		if g := m.Get(uint16(k)); g != uint(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint(unsignedData[i])+3)
		}
	}
	m.Clear()
	for _, k := range unsignedData {
		m.Add(uint16(k), 42)
		if g := m.Get(uint16(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapUint16ToUintVisit(t *testing.T) {
	m := NewMapUint16ToUint(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), uint(unsignedData[i]))
	}

	n := make(map[uint16]uint, len(unsignedData))

	for i, k := range unsignedData {
		n[uint16(k)] = uint(unsignedData[i])
	}

	m.Visit(func(k uint16, v uint) {
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

func TestMapUint16ToUint8Size(t *testing.T) {
	m := NewMapUint16ToUint8(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint16(k), uint8(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint16(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint16ToUint8Contains(t *testing.T) {
	m := NewMapUint16ToUint8(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), uint8(unsignedData[i]))
		if !m.Contains(uint16(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint16(k))
		if m.Contains(uint16(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint16ToUint8Get(t *testing.T) {
	m := NewMapUint16ToUint8(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), uint8(unsignedData[i]))
		if g := m.Get(uint16(k)); g != uint8(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
}

func TestMapUint16ToUint8Clear(t *testing.T) {
	m := NewMapUint16ToUint8(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), uint8(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint16(k), uint8(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUint16ToUint8Inc(t *testing.T) {
	m := NewMapUint16ToUint8(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), uint8(unsignedData[i]))
	}
	for i, k := range unsignedData {
		m.Inc(uint16(k))
		m.Inc(uint16(k))
		m.Inc(uint16(k))
		if g := m.Get(uint16(k)); g != uint8(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint8(unsignedData[i])+3)
		}
	}
	m.Clear()
	for _, k := range unsignedData {
		m.Inc(uint16(k))
		if g := m.Get(uint16(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapUint16ToUint8Add(t *testing.T) {
	m := NewMapUint16ToUint8(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), uint8(unsignedData[i]))
	}
	for i, k := range unsignedData {
		m.Add(uint16(k), 3)
		if g := m.Get(uint16(k)); g != uint8(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint8(unsignedData[i])+3)
		}
	}
	m.Clear()
	for _, k := range unsignedData {
		m.Add(uint16(k), 42)
		if g := m.Get(uint16(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapUint16ToUint8Visit(t *testing.T) {
	m := NewMapUint16ToUint8(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), uint8(unsignedData[i]))
	}

	n := make(map[uint16]uint8, len(unsignedData))

	for i, k := range unsignedData {
		n[uint16(k)] = uint8(unsignedData[i])
	}

	m.Visit(func(k uint16, v uint8) {
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

func TestMapUint16ToUint16Size(t *testing.T) {
	m := NewMapUint16ToUint16(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint16(k), uint16(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint16(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint16ToUint16Contains(t *testing.T) {
	m := NewMapUint16ToUint16(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), uint16(unsignedData[i]))
		if !m.Contains(uint16(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint16(k))
		if m.Contains(uint16(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint16ToUint16Get(t *testing.T) {
	m := NewMapUint16ToUint16(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), uint16(unsignedData[i]))
		if g := m.Get(uint16(k)); g != uint16(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
}

func TestMapUint16ToUint16Clear(t *testing.T) {
	m := NewMapUint16ToUint16(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), uint16(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint16(k), uint16(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUint16ToUint16Inc(t *testing.T) {
	m := NewMapUint16ToUint16(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), uint16(unsignedData[i]))
	}
	for i, k := range unsignedData {
		m.Inc(uint16(k))
		m.Inc(uint16(k))
		m.Inc(uint16(k))
		if g := m.Get(uint16(k)); g != uint16(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint16(unsignedData[i])+3)
		}
	}
	m.Clear()
	for _, k := range unsignedData {
		m.Inc(uint16(k))
		if g := m.Get(uint16(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapUint16ToUint16Add(t *testing.T) {
	m := NewMapUint16ToUint16(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), uint16(unsignedData[i]))
	}
	for i, k := range unsignedData {
		m.Add(uint16(k), 3)
		if g := m.Get(uint16(k)); g != uint16(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint16(unsignedData[i])+3)
		}
	}
	m.Clear()
	for _, k := range unsignedData {
		m.Add(uint16(k), 42)
		if g := m.Get(uint16(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapUint16ToUint16Visit(t *testing.T) {
	m := NewMapUint16ToUint16(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), uint16(unsignedData[i]))
	}

	n := make(map[uint16]uint16, len(unsignedData))

	for i, k := range unsignedData {
		n[uint16(k)] = uint16(unsignedData[i])
	}

	m.Visit(func(k uint16, v uint16) {
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

func TestMapUint16ToUint32Size(t *testing.T) {
	m := NewMapUint16ToUint32(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint16(k), uint32(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint16(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint16ToUint32Contains(t *testing.T) {
	m := NewMapUint16ToUint32(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), uint32(unsignedData[i]))
		if !m.Contains(uint16(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint16(k))
		if m.Contains(uint16(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint16ToUint32Get(t *testing.T) {
	m := NewMapUint16ToUint32(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), uint32(unsignedData[i]))
		if g := m.Get(uint16(k)); g != uint32(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
}

func TestMapUint16ToUint32Clear(t *testing.T) {
	m := NewMapUint16ToUint32(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), uint32(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint16(k), uint32(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUint16ToUint32Inc(t *testing.T) {
	m := NewMapUint16ToUint32(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), uint32(unsignedData[i]))
	}
	for i, k := range unsignedData {
		m.Inc(uint16(k))
		m.Inc(uint16(k))
		m.Inc(uint16(k))
		if g := m.Get(uint16(k)); g != uint32(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint32(unsignedData[i])+3)
		}
	}
	m.Clear()
	for _, k := range unsignedData {
		m.Inc(uint16(k))
		if g := m.Get(uint16(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapUint16ToUint32Add(t *testing.T) {
	m := NewMapUint16ToUint32(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), uint32(unsignedData[i]))
	}
	for i, k := range unsignedData {
		m.Add(uint16(k), 3)
		if g := m.Get(uint16(k)); g != uint32(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint32(unsignedData[i])+3)
		}
	}
	m.Clear()
	for _, k := range unsignedData {
		m.Add(uint16(k), 42)
		if g := m.Get(uint16(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapUint16ToUint32Visit(t *testing.T) {
	m := NewMapUint16ToUint32(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), uint32(unsignedData[i]))
	}

	n := make(map[uint16]uint32, len(unsignedData))

	for i, k := range unsignedData {
		n[uint16(k)] = uint32(unsignedData[i])
	}

	m.Visit(func(k uint16, v uint32) {
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

func TestMapUint16ToUint64Size(t *testing.T) {
	m := NewMapUint16ToUint64(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint16(k), uint64(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint16(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint16ToUint64Contains(t *testing.T) {
	m := NewMapUint16ToUint64(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), uint64(unsignedData[i]))
		if !m.Contains(uint16(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint16(k))
		if m.Contains(uint16(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint16ToUint64Get(t *testing.T) {
	m := NewMapUint16ToUint64(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), uint64(unsignedData[i]))
		if g := m.Get(uint16(k)); g != uint64(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
}

func TestMapUint16ToUint64Clear(t *testing.T) {
	m := NewMapUint16ToUint64(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), uint64(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint16(k), uint64(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUint16ToUint64Inc(t *testing.T) {
	m := NewMapUint16ToUint64(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), uint64(unsignedData[i]))
	}
	for i, k := range unsignedData {
		m.Inc(uint16(k))
		m.Inc(uint16(k))
		m.Inc(uint16(k))
		if g := m.Get(uint16(k)); g != uint64(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint64(unsignedData[i])+3)
		}
	}
	m.Clear()
	for _, k := range unsignedData {
		m.Inc(uint16(k))
		if g := m.Get(uint16(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapUint16ToUint64Add(t *testing.T) {
	m := NewMapUint16ToUint64(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), uint64(unsignedData[i]))
	}
	for i, k := range unsignedData {
		m.Add(uint16(k), 3)
		if g := m.Get(uint16(k)); g != uint64(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint64(unsignedData[i])+3)
		}
	}
	m.Clear()
	for _, k := range unsignedData {
		m.Add(uint16(k), 42)
		if g := m.Get(uint16(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapUint16ToUint64Visit(t *testing.T) {
	m := NewMapUint16ToUint64(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), uint64(unsignedData[i]))
	}

	n := make(map[uint16]uint64, len(unsignedData))

	for i, k := range unsignedData {
		n[uint16(k)] = uint64(unsignedData[i])
	}

	m.Visit(func(k uint16, v uint64) {
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
