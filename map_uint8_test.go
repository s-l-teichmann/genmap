// THIS IS A MACHINE GENERATED FILE!
// BE CAREFUL WITH EDITING BY HAND.
//
// This is Free Software covered by the terms of the MIT license.
// See LICENSE file for details.
// (c) 2016 by Sascha L. Teichmann.
// See the full list of contributors in the CONTRIBUTORS file.

package genmap

import "testing"

func TestMapUint8ToIntSize(t *testing.T) {
	m := NewMapUint8ToInt(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint8(k), int(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint8(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint8ToIntContains(t *testing.T) {
	m := NewMapUint8ToInt(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), int(signedData[i]))
		if !m.Contains(uint8(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint8(k))
		if m.Contains(uint8(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint8ToIntGet(t *testing.T) {
	m := NewMapUint8ToInt(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), int(signedData[i]))
		if g := m.Get(uint8(k)); g != int(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapUint8ToIntClear(t *testing.T) {
	m := NewMapUint8ToInt(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), int(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint8(k), int(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUint8ToIntInc(t *testing.T) {
	m := NewMapUint8ToInt(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), int(signedData[i]))
	}
	for i, k := range unsignedData {
		m.Inc(uint8(k))
		m.Inc(uint8(k))
		m.Inc(uint8(k))
		if g := m.Get(uint8(k)); g != int(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int(signedData[i])+3)
		}
	}
	m = NewMapUint8ToInt(13)
	for _, k := range unsignedData {
		m.Inc(uint8(k))
		if g := m.Get(uint8(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapUint8ToIntAdd(t *testing.T) {
	m := NewMapUint8ToInt(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), int(signedData[i]))
	}
	for i, k := range unsignedData {
		m.Add(uint8(k), 3)
		if g := m.Get(uint8(k)); g != int(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int(signedData[i])+3)
		}
	}
	m = NewMapUint8ToInt(13)
	for _, k := range unsignedData {
		m.Add(uint8(k), 42)
		if g := m.Get(uint8(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapUint8ToIntModify(t *testing.T) {
	m := NewMapUint8ToInt(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), int(signedData[i]))
	}
	for _, k := range unsignedData {
		m.Modify(uint8(k), func(v *int) {
			*v += 3
		})
	}
	for i, k := range unsignedData {
		if g := m.Get(uint8(k)); g != int(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int(signedData[i])+3)
		}
	}
}

func TestMapUint8ToIntPut(t *testing.T) {
	m := NewMapUint8ToInt(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), int(signedData[i]))
	}
	for _, k := range unsignedData {
		m.Put(uint8(k), m.Get(uint8(k))+3)
	}
	for i, k := range unsignedData {
		if g := m.Get(uint8(k)); g != int(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int(signedData[i])+3)
		}
	}
}

func TestMapUint8ToIntVisit(t *testing.T) {
	m := NewMapUint8ToInt(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), int(signedData[i]))
	}

	n := make(map[uint8]int, len(unsignedData))

	for i, k := range unsignedData {
		n[uint8(k)] = int(signedData[i])
	}

	m.Visit(func(k uint8, v int) {
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

func TestMapUint8ToInt8Size(t *testing.T) {
	m := NewMapUint8ToInt8(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint8(k), int8(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint8(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint8ToInt8Contains(t *testing.T) {
	m := NewMapUint8ToInt8(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), int8(signedData[i]))
		if !m.Contains(uint8(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint8(k))
		if m.Contains(uint8(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint8ToInt8Get(t *testing.T) {
	m := NewMapUint8ToInt8(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), int8(signedData[i]))
		if g := m.Get(uint8(k)); g != int8(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapUint8ToInt8Clear(t *testing.T) {
	m := NewMapUint8ToInt8(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), int8(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint8(k), int8(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUint8ToInt8Inc(t *testing.T) {
	m := NewMapUint8ToInt8(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), int8(signedData[i]))
	}
	for i, k := range unsignedData {
		m.Inc(uint8(k))
		m.Inc(uint8(k))
		m.Inc(uint8(k))
		if g := m.Get(uint8(k)); g != int8(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int8(signedData[i])+3)
		}
	}
	m = NewMapUint8ToInt8(13)
	for _, k := range unsignedData {
		m.Inc(uint8(k))
		if g := m.Get(uint8(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapUint8ToInt8Add(t *testing.T) {
	m := NewMapUint8ToInt8(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), int8(signedData[i]))
	}
	for i, k := range unsignedData {
		m.Add(uint8(k), 3)
		if g := m.Get(uint8(k)); g != int8(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int8(signedData[i])+3)
		}
	}
	m = NewMapUint8ToInt8(13)
	for _, k := range unsignedData {
		m.Add(uint8(k), 42)
		if g := m.Get(uint8(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapUint8ToInt8Modify(t *testing.T) {
	m := NewMapUint8ToInt8(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), int8(signedData[i]))
	}
	for _, k := range unsignedData {
		m.Modify(uint8(k), func(v *int8) {
			*v += 3
		})
	}
	for i, k := range unsignedData {
		if g := m.Get(uint8(k)); g != int8(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int8(signedData[i])+3)
		}
	}
}

func TestMapUint8ToInt8Put(t *testing.T) {
	m := NewMapUint8ToInt8(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), int8(signedData[i]))
	}
	for _, k := range unsignedData {
		m.Put(uint8(k), m.Get(uint8(k))+3)
	}
	for i, k := range unsignedData {
		if g := m.Get(uint8(k)); g != int8(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int8(signedData[i])+3)
		}
	}
}

func TestMapUint8ToInt8Visit(t *testing.T) {
	m := NewMapUint8ToInt8(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), int8(signedData[i]))
	}

	n := make(map[uint8]int8, len(unsignedData))

	for i, k := range unsignedData {
		n[uint8(k)] = int8(signedData[i])
	}

	m.Visit(func(k uint8, v int8) {
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

func TestMapUint8ToInt16Size(t *testing.T) {
	m := NewMapUint8ToInt16(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint8(k), int16(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint8(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint8ToInt16Contains(t *testing.T) {
	m := NewMapUint8ToInt16(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), int16(signedData[i]))
		if !m.Contains(uint8(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint8(k))
		if m.Contains(uint8(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint8ToInt16Get(t *testing.T) {
	m := NewMapUint8ToInt16(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), int16(signedData[i]))
		if g := m.Get(uint8(k)); g != int16(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapUint8ToInt16Clear(t *testing.T) {
	m := NewMapUint8ToInt16(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), int16(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint8(k), int16(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUint8ToInt16Inc(t *testing.T) {
	m := NewMapUint8ToInt16(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), int16(signedData[i]))
	}
	for i, k := range unsignedData {
		m.Inc(uint8(k))
		m.Inc(uint8(k))
		m.Inc(uint8(k))
		if g := m.Get(uint8(k)); g != int16(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int16(signedData[i])+3)
		}
	}
	m = NewMapUint8ToInt16(13)
	for _, k := range unsignedData {
		m.Inc(uint8(k))
		if g := m.Get(uint8(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapUint8ToInt16Add(t *testing.T) {
	m := NewMapUint8ToInt16(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), int16(signedData[i]))
	}
	for i, k := range unsignedData {
		m.Add(uint8(k), 3)
		if g := m.Get(uint8(k)); g != int16(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int16(signedData[i])+3)
		}
	}
	m = NewMapUint8ToInt16(13)
	for _, k := range unsignedData {
		m.Add(uint8(k), 42)
		if g := m.Get(uint8(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapUint8ToInt16Modify(t *testing.T) {
	m := NewMapUint8ToInt16(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), int16(signedData[i]))
	}
	for _, k := range unsignedData {
		m.Modify(uint8(k), func(v *int16) {
			*v += 3
		})
	}
	for i, k := range unsignedData {
		if g := m.Get(uint8(k)); g != int16(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int16(signedData[i])+3)
		}
	}
}

func TestMapUint8ToInt16Put(t *testing.T) {
	m := NewMapUint8ToInt16(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), int16(signedData[i]))
	}
	for _, k := range unsignedData {
		m.Put(uint8(k), m.Get(uint8(k))+3)
	}
	for i, k := range unsignedData {
		if g := m.Get(uint8(k)); g != int16(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int16(signedData[i])+3)
		}
	}
}

func TestMapUint8ToInt16Visit(t *testing.T) {
	m := NewMapUint8ToInt16(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), int16(signedData[i]))
	}

	n := make(map[uint8]int16, len(unsignedData))

	for i, k := range unsignedData {
		n[uint8(k)] = int16(signedData[i])
	}

	m.Visit(func(k uint8, v int16) {
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

func TestMapUint8ToInt32Size(t *testing.T) {
	m := NewMapUint8ToInt32(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint8(k), int32(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint8(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint8ToInt32Contains(t *testing.T) {
	m := NewMapUint8ToInt32(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), int32(signedData[i]))
		if !m.Contains(uint8(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint8(k))
		if m.Contains(uint8(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint8ToInt32Get(t *testing.T) {
	m := NewMapUint8ToInt32(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), int32(signedData[i]))
		if g := m.Get(uint8(k)); g != int32(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapUint8ToInt32Clear(t *testing.T) {
	m := NewMapUint8ToInt32(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), int32(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint8(k), int32(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUint8ToInt32Inc(t *testing.T) {
	m := NewMapUint8ToInt32(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), int32(signedData[i]))
	}
	for i, k := range unsignedData {
		m.Inc(uint8(k))
		m.Inc(uint8(k))
		m.Inc(uint8(k))
		if g := m.Get(uint8(k)); g != int32(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int32(signedData[i])+3)
		}
	}
	m = NewMapUint8ToInt32(13)
	for _, k := range unsignedData {
		m.Inc(uint8(k))
		if g := m.Get(uint8(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapUint8ToInt32Add(t *testing.T) {
	m := NewMapUint8ToInt32(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), int32(signedData[i]))
	}
	for i, k := range unsignedData {
		m.Add(uint8(k), 3)
		if g := m.Get(uint8(k)); g != int32(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int32(signedData[i])+3)
		}
	}
	m = NewMapUint8ToInt32(13)
	for _, k := range unsignedData {
		m.Add(uint8(k), 42)
		if g := m.Get(uint8(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapUint8ToInt32Modify(t *testing.T) {
	m := NewMapUint8ToInt32(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), int32(signedData[i]))
	}
	for _, k := range unsignedData {
		m.Modify(uint8(k), func(v *int32) {
			*v += 3
		})
	}
	for i, k := range unsignedData {
		if g := m.Get(uint8(k)); g != int32(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int32(signedData[i])+3)
		}
	}
}

func TestMapUint8ToInt32Put(t *testing.T) {
	m := NewMapUint8ToInt32(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), int32(signedData[i]))
	}
	for _, k := range unsignedData {
		m.Put(uint8(k), m.Get(uint8(k))+3)
	}
	for i, k := range unsignedData {
		if g := m.Get(uint8(k)); g != int32(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int32(signedData[i])+3)
		}
	}
}

func TestMapUint8ToInt32Visit(t *testing.T) {
	m := NewMapUint8ToInt32(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), int32(signedData[i]))
	}

	n := make(map[uint8]int32, len(unsignedData))

	for i, k := range unsignedData {
		n[uint8(k)] = int32(signedData[i])
	}

	m.Visit(func(k uint8, v int32) {
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

func TestMapUint8ToInt64Size(t *testing.T) {
	m := NewMapUint8ToInt64(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint8(k), int64(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint8(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint8ToInt64Contains(t *testing.T) {
	m := NewMapUint8ToInt64(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), int64(signedData[i]))
		if !m.Contains(uint8(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint8(k))
		if m.Contains(uint8(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint8ToInt64Get(t *testing.T) {
	m := NewMapUint8ToInt64(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), int64(signedData[i]))
		if g := m.Get(uint8(k)); g != int64(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapUint8ToInt64Clear(t *testing.T) {
	m := NewMapUint8ToInt64(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), int64(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint8(k), int64(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUint8ToInt64Inc(t *testing.T) {
	m := NewMapUint8ToInt64(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), int64(signedData[i]))
	}
	for i, k := range unsignedData {
		m.Inc(uint8(k))
		m.Inc(uint8(k))
		m.Inc(uint8(k))
		if g := m.Get(uint8(k)); g != int64(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int64(signedData[i])+3)
		}
	}
	m = NewMapUint8ToInt64(13)
	for _, k := range unsignedData {
		m.Inc(uint8(k))
		if g := m.Get(uint8(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapUint8ToInt64Add(t *testing.T) {
	m := NewMapUint8ToInt64(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), int64(signedData[i]))
	}
	for i, k := range unsignedData {
		m.Add(uint8(k), 3)
		if g := m.Get(uint8(k)); g != int64(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int64(signedData[i])+3)
		}
	}
	m = NewMapUint8ToInt64(13)
	for _, k := range unsignedData {
		m.Add(uint8(k), 42)
		if g := m.Get(uint8(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapUint8ToInt64Modify(t *testing.T) {
	m := NewMapUint8ToInt64(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), int64(signedData[i]))
	}
	for _, k := range unsignedData {
		m.Modify(uint8(k), func(v *int64) {
			*v += 3
		})
	}
	for i, k := range unsignedData {
		if g := m.Get(uint8(k)); g != int64(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int64(signedData[i])+3)
		}
	}
}

func TestMapUint8ToInt64Put(t *testing.T) {
	m := NewMapUint8ToInt64(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), int64(signedData[i]))
	}
	for _, k := range unsignedData {
		m.Put(uint8(k), m.Get(uint8(k))+3)
	}
	for i, k := range unsignedData {
		if g := m.Get(uint8(k)); g != int64(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int64(signedData[i])+3)
		}
	}
}

func TestMapUint8ToInt64Visit(t *testing.T) {
	m := NewMapUint8ToInt64(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), int64(signedData[i]))
	}

	n := make(map[uint8]int64, len(unsignedData))

	for i, k := range unsignedData {
		n[uint8(k)] = int64(signedData[i])
	}

	m.Visit(func(k uint8, v int64) {
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

func TestMapUint8ToUintSize(t *testing.T) {
	m := NewMapUint8ToUint(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint8(k), uint(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint8(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint8ToUintContains(t *testing.T) {
	m := NewMapUint8ToUint(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), uint(unsignedData[i]))
		if !m.Contains(uint8(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint8(k))
		if m.Contains(uint8(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint8ToUintGet(t *testing.T) {
	m := NewMapUint8ToUint(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), uint(unsignedData[i]))
		if g := m.Get(uint8(k)); g != uint(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapUint8ToUintClear(t *testing.T) {
	m := NewMapUint8ToUint(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), uint(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint8(k), uint(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUint8ToUintInc(t *testing.T) {
	m := NewMapUint8ToUint(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), uint(unsignedData[i]))
	}
	for i, k := range unsignedData {
		m.Inc(uint8(k))
		m.Inc(uint8(k))
		m.Inc(uint8(k))
		if g := m.Get(uint8(k)); g != uint(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint(unsignedData[i])+3)
		}
	}
	m = NewMapUint8ToUint(13)
	for _, k := range unsignedData {
		m.Inc(uint8(k))
		if g := m.Get(uint8(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapUint8ToUintAdd(t *testing.T) {
	m := NewMapUint8ToUint(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), uint(unsignedData[i]))
	}
	for i, k := range unsignedData {
		m.Add(uint8(k), 3)
		if g := m.Get(uint8(k)); g != uint(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint(unsignedData[i])+3)
		}
	}
	m = NewMapUint8ToUint(13)
	for _, k := range unsignedData {
		m.Add(uint8(k), 42)
		if g := m.Get(uint8(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapUint8ToUintModify(t *testing.T) {
	m := NewMapUint8ToUint(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), uint(unsignedData[i]))
	}
	for _, k := range unsignedData {
		m.Modify(uint8(k), func(v *uint) {
			*v += 3
		})
	}
	for i, k := range unsignedData {
		if g := m.Get(uint8(k)); g != uint(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint(unsignedData[i])+3)
		}
	}
}

func TestMapUint8ToUintPut(t *testing.T) {
	m := NewMapUint8ToUint(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), uint(unsignedData[i]))
	}
	for _, k := range unsignedData {
		m.Put(uint8(k), m.Get(uint8(k))+3)
	}
	for i, k := range unsignedData {
		if g := m.Get(uint8(k)); g != uint(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint(unsignedData[i])+3)
		}
	}
}

func TestMapUint8ToUintVisit(t *testing.T) {
	m := NewMapUint8ToUint(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), uint(unsignedData[i]))
	}

	n := make(map[uint8]uint, len(unsignedData))

	for i, k := range unsignedData {
		n[uint8(k)] = uint(unsignedData[i])
	}

	m.Visit(func(k uint8, v uint) {
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

func TestMapUint8ToUint8Size(t *testing.T) {
	m := NewMapUint8ToUint8(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint8(k), uint8(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint8(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint8ToUint8Contains(t *testing.T) {
	m := NewMapUint8ToUint8(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), uint8(unsignedData[i]))
		if !m.Contains(uint8(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint8(k))
		if m.Contains(uint8(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint8ToUint8Get(t *testing.T) {
	m := NewMapUint8ToUint8(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), uint8(unsignedData[i]))
		if g := m.Get(uint8(k)); g != uint8(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapUint8ToUint8Clear(t *testing.T) {
	m := NewMapUint8ToUint8(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), uint8(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint8(k), uint8(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUint8ToUint8Inc(t *testing.T) {
	m := NewMapUint8ToUint8(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), uint8(unsignedData[i]))
	}
	for i, k := range unsignedData {
		m.Inc(uint8(k))
		m.Inc(uint8(k))
		m.Inc(uint8(k))
		if g := m.Get(uint8(k)); g != uint8(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint8(unsignedData[i])+3)
		}
	}
	m = NewMapUint8ToUint8(13)
	for _, k := range unsignedData {
		m.Inc(uint8(k))
		if g := m.Get(uint8(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapUint8ToUint8Add(t *testing.T) {
	m := NewMapUint8ToUint8(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), uint8(unsignedData[i]))
	}
	for i, k := range unsignedData {
		m.Add(uint8(k), 3)
		if g := m.Get(uint8(k)); g != uint8(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint8(unsignedData[i])+3)
		}
	}
	m = NewMapUint8ToUint8(13)
	for _, k := range unsignedData {
		m.Add(uint8(k), 42)
		if g := m.Get(uint8(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapUint8ToUint8Modify(t *testing.T) {
	m := NewMapUint8ToUint8(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), uint8(unsignedData[i]))
	}
	for _, k := range unsignedData {
		m.Modify(uint8(k), func(v *uint8) {
			*v += 3
		})
	}
	for i, k := range unsignedData {
		if g := m.Get(uint8(k)); g != uint8(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint8(unsignedData[i])+3)
		}
	}
}

func TestMapUint8ToUint8Put(t *testing.T) {
	m := NewMapUint8ToUint8(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), uint8(unsignedData[i]))
	}
	for _, k := range unsignedData {
		m.Put(uint8(k), m.Get(uint8(k))+3)
	}
	for i, k := range unsignedData {
		if g := m.Get(uint8(k)); g != uint8(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint8(unsignedData[i])+3)
		}
	}
}

func TestMapUint8ToUint8Visit(t *testing.T) {
	m := NewMapUint8ToUint8(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), uint8(unsignedData[i]))
	}

	n := make(map[uint8]uint8, len(unsignedData))

	for i, k := range unsignedData {
		n[uint8(k)] = uint8(unsignedData[i])
	}

	m.Visit(func(k uint8, v uint8) {
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

func TestMapUint8ToUint16Size(t *testing.T) {
	m := NewMapUint8ToUint16(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint8(k), uint16(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint8(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint8ToUint16Contains(t *testing.T) {
	m := NewMapUint8ToUint16(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), uint16(unsignedData[i]))
		if !m.Contains(uint8(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint8(k))
		if m.Contains(uint8(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint8ToUint16Get(t *testing.T) {
	m := NewMapUint8ToUint16(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), uint16(unsignedData[i]))
		if g := m.Get(uint8(k)); g != uint16(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapUint8ToUint16Clear(t *testing.T) {
	m := NewMapUint8ToUint16(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), uint16(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint8(k), uint16(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUint8ToUint16Inc(t *testing.T) {
	m := NewMapUint8ToUint16(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), uint16(unsignedData[i]))
	}
	for i, k := range unsignedData {
		m.Inc(uint8(k))
		m.Inc(uint8(k))
		m.Inc(uint8(k))
		if g := m.Get(uint8(k)); g != uint16(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint16(unsignedData[i])+3)
		}
	}
	m = NewMapUint8ToUint16(13)
	for _, k := range unsignedData {
		m.Inc(uint8(k))
		if g := m.Get(uint8(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapUint8ToUint16Add(t *testing.T) {
	m := NewMapUint8ToUint16(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), uint16(unsignedData[i]))
	}
	for i, k := range unsignedData {
		m.Add(uint8(k), 3)
		if g := m.Get(uint8(k)); g != uint16(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint16(unsignedData[i])+3)
		}
	}
	m = NewMapUint8ToUint16(13)
	for _, k := range unsignedData {
		m.Add(uint8(k), 42)
		if g := m.Get(uint8(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapUint8ToUint16Modify(t *testing.T) {
	m := NewMapUint8ToUint16(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), uint16(unsignedData[i]))
	}
	for _, k := range unsignedData {
		m.Modify(uint8(k), func(v *uint16) {
			*v += 3
		})
	}
	for i, k := range unsignedData {
		if g := m.Get(uint8(k)); g != uint16(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint16(unsignedData[i])+3)
		}
	}
}

func TestMapUint8ToUint16Put(t *testing.T) {
	m := NewMapUint8ToUint16(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), uint16(unsignedData[i]))
	}
	for _, k := range unsignedData {
		m.Put(uint8(k), m.Get(uint8(k))+3)
	}
	for i, k := range unsignedData {
		if g := m.Get(uint8(k)); g != uint16(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint16(unsignedData[i])+3)
		}
	}
}

func TestMapUint8ToUint16Visit(t *testing.T) {
	m := NewMapUint8ToUint16(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), uint16(unsignedData[i]))
	}

	n := make(map[uint8]uint16, len(unsignedData))

	for i, k := range unsignedData {
		n[uint8(k)] = uint16(unsignedData[i])
	}

	m.Visit(func(k uint8, v uint16) {
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

func TestMapUint8ToUint32Size(t *testing.T) {
	m := NewMapUint8ToUint32(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint8(k), uint32(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint8(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint8ToUint32Contains(t *testing.T) {
	m := NewMapUint8ToUint32(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), uint32(unsignedData[i]))
		if !m.Contains(uint8(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint8(k))
		if m.Contains(uint8(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint8ToUint32Get(t *testing.T) {
	m := NewMapUint8ToUint32(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), uint32(unsignedData[i]))
		if g := m.Get(uint8(k)); g != uint32(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapUint8ToUint32Clear(t *testing.T) {
	m := NewMapUint8ToUint32(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), uint32(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint8(k), uint32(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUint8ToUint32Inc(t *testing.T) {
	m := NewMapUint8ToUint32(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), uint32(unsignedData[i]))
	}
	for i, k := range unsignedData {
		m.Inc(uint8(k))
		m.Inc(uint8(k))
		m.Inc(uint8(k))
		if g := m.Get(uint8(k)); g != uint32(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint32(unsignedData[i])+3)
		}
	}
	m = NewMapUint8ToUint32(13)
	for _, k := range unsignedData {
		m.Inc(uint8(k))
		if g := m.Get(uint8(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapUint8ToUint32Add(t *testing.T) {
	m := NewMapUint8ToUint32(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), uint32(unsignedData[i]))
	}
	for i, k := range unsignedData {
		m.Add(uint8(k), 3)
		if g := m.Get(uint8(k)); g != uint32(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint32(unsignedData[i])+3)
		}
	}
	m = NewMapUint8ToUint32(13)
	for _, k := range unsignedData {
		m.Add(uint8(k), 42)
		if g := m.Get(uint8(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapUint8ToUint32Modify(t *testing.T) {
	m := NewMapUint8ToUint32(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), uint32(unsignedData[i]))
	}
	for _, k := range unsignedData {
		m.Modify(uint8(k), func(v *uint32) {
			*v += 3
		})
	}
	for i, k := range unsignedData {
		if g := m.Get(uint8(k)); g != uint32(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint32(unsignedData[i])+3)
		}
	}
}

func TestMapUint8ToUint32Put(t *testing.T) {
	m := NewMapUint8ToUint32(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), uint32(unsignedData[i]))
	}
	for _, k := range unsignedData {
		m.Put(uint8(k), m.Get(uint8(k))+3)
	}
	for i, k := range unsignedData {
		if g := m.Get(uint8(k)); g != uint32(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint32(unsignedData[i])+3)
		}
	}
}

func TestMapUint8ToUint32Visit(t *testing.T) {
	m := NewMapUint8ToUint32(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), uint32(unsignedData[i]))
	}

	n := make(map[uint8]uint32, len(unsignedData))

	for i, k := range unsignedData {
		n[uint8(k)] = uint32(unsignedData[i])
	}

	m.Visit(func(k uint8, v uint32) {
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

func TestMapUint8ToUint64Size(t *testing.T) {
	m := NewMapUint8ToUint64(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint8(k), uint64(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint8(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint8ToUint64Contains(t *testing.T) {
	m := NewMapUint8ToUint64(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), uint64(unsignedData[i]))
		if !m.Contains(uint8(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint8(k))
		if m.Contains(uint8(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint8ToUint64Get(t *testing.T) {
	m := NewMapUint8ToUint64(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), uint64(unsignedData[i]))
		if g := m.Get(uint8(k)); g != uint64(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapUint8ToUint64Clear(t *testing.T) {
	m := NewMapUint8ToUint64(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), uint64(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint8(k), uint64(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUint8ToUint64Inc(t *testing.T) {
	m := NewMapUint8ToUint64(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), uint64(unsignedData[i]))
	}
	for i, k := range unsignedData {
		m.Inc(uint8(k))
		m.Inc(uint8(k))
		m.Inc(uint8(k))
		if g := m.Get(uint8(k)); g != uint64(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint64(unsignedData[i])+3)
		}
	}
	m = NewMapUint8ToUint64(13)
	for _, k := range unsignedData {
		m.Inc(uint8(k))
		if g := m.Get(uint8(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapUint8ToUint64Add(t *testing.T) {
	m := NewMapUint8ToUint64(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), uint64(unsignedData[i]))
	}
	for i, k := range unsignedData {
		m.Add(uint8(k), 3)
		if g := m.Get(uint8(k)); g != uint64(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint64(unsignedData[i])+3)
		}
	}
	m = NewMapUint8ToUint64(13)
	for _, k := range unsignedData {
		m.Add(uint8(k), 42)
		if g := m.Get(uint8(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapUint8ToUint64Modify(t *testing.T) {
	m := NewMapUint8ToUint64(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), uint64(unsignedData[i]))
	}
	for _, k := range unsignedData {
		m.Modify(uint8(k), func(v *uint64) {
			*v += 3
		})
	}
	for i, k := range unsignedData {
		if g := m.Get(uint8(k)); g != uint64(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint64(unsignedData[i])+3)
		}
	}
}

func TestMapUint8ToUint64Put(t *testing.T) {
	m := NewMapUint8ToUint64(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), uint64(unsignedData[i]))
	}
	for _, k := range unsignedData {
		m.Put(uint8(k), m.Get(uint8(k))+3)
	}
	for i, k := range unsignedData {
		if g := m.Get(uint8(k)); g != uint64(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint64(unsignedData[i])+3)
		}
	}
}

func TestMapUint8ToUint64Visit(t *testing.T) {
	m := NewMapUint8ToUint64(13)
	for i, k := range unsignedData {
		m.Put(uint8(k), uint64(unsignedData[i]))
	}

	n := make(map[uint8]uint64, len(unsignedData))

	for i, k := range unsignedData {
		n[uint8(k)] = uint64(unsignedData[i])
	}

	m.Visit(func(k uint8, v uint64) {
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
