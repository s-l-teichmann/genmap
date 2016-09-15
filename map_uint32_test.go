// THIS IS A MACHINE GENERATED FILE!
// BE CAREFUL WITH EDITING BY HAND.
//
// This is Free Software covered by the terms of the MIT license.
// See LICENSE file for details.
// (c) 2016 by Sascha L. Teichmann.
// See the full list of contributors in the CONTRIBUTORS file.

package genmap

import "testing"

func TestMapUint32ToIntSize(t *testing.T) {
	m := NewMapUint32ToInt(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint32(k), int(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint32(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint32ToIntContains(t *testing.T) {
	m := NewMapUint32ToInt(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), int(signedData[i]))
		if !m.Contains(uint32(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint32(k))
		if m.Contains(uint32(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint32ToIntGet(t *testing.T) {
	m := NewMapUint32ToInt(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), int(signedData[i]))
		if g := m.Get(uint32(k)); g != int(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapUint32ToIntClear(t *testing.T) {
	m := NewMapUint32ToInt(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), int(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint32(k), int(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUint32ToIntInc(t *testing.T) {
	m := NewMapUint32ToInt(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), int(signedData[i]))
	}
	for i, k := range unsignedData {
		m.Inc(uint32(k))
		m.Inc(uint32(k))
		m.Inc(uint32(k))
		if g := m.Get(uint32(k)); g != int(signedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, int(signedData[i])+3)
		}
	}
	m = NewMapUint32ToInt(13)
	for _, k := range unsignedData {
		m.Inc(uint32(k))
		if g := m.Get(uint32(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapUint32ToIntAdd(t *testing.T) {
	m := NewMapUint32ToInt(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), int(signedData[i]))
	}
	for i, k := range unsignedData {
		m.Add(uint32(k), 3)
		if g := m.Get(uint32(k)); g != int(signedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, int(signedData[i])+3)
		}
	}
	m = NewMapUint32ToInt(13)
	for _, k := range unsignedData {
		m.Add(uint32(k), 42)
		if g := m.Get(uint32(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapUint32ToIntModify(t *testing.T) {
	m := NewMapUint32ToInt(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), int(signedData[i]))
	}
	for _, k := range unsignedData {
		m.Modify(uint32(k), func(v *int) {
			*v += 3
		})
	}
	for i, k := range unsignedData {
		if g := m.Get(uint32(k)); g != int(signedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, int(signedData[i])+3)
		}
	}
}

func TestMapUint32ToIntFind(t *testing.T) {
	m := NewMapUint32ToInt(13)

	low := unsignedData[:len(unsignedData)/2]
	hi := unsignedData[len(unsignedData)/2:]

	for i, k := range low {
		m.Put(uint32(k), int(signedData[i]))
	}
	for i, k := range low {
		v, ok := m.Find(uint32(k))
		if !ok {
			t.Errorf("missing value for key %d\n", k)
			continue
		}
		if want := int(signedData[i]); v != want {
			t.Errorf("key %d: got %d, want %d\n", k, v, want)
		}
	}
	for _, k := range hi {
		if v, ok := m.Find(uint32(k)); ok {
			t.Errorf("key %d: got %d, want nothing\n", k, v)
		}
	}
}

func TestMapUint32ToIntPut(t *testing.T) {
	m := NewMapUint32ToInt(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), int(signedData[i]))
	}
	for _, k := range unsignedData {
		m.Put(uint32(k), m.Get(uint32(k))+3)
	}
	for i, k := range unsignedData {
		if g := m.Get(uint32(k)); g != int(signedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, int(signedData[i])+3)
		}
	}
}

func TestMapUint32ToIntVisit(t *testing.T) {
	m := NewMapUint32ToInt(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), int(signedData[i]))
	}

	n := make(map[uint32]int, len(unsignedData))

	for i, k := range unsignedData {
		n[uint32(k)] = int(signedData[i])
	}

	m.Visit(func(k uint32, v int) {
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

func TestMapUint32ToInt8Size(t *testing.T) {
	m := NewMapUint32ToInt8(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint32(k), int8(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint32(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint32ToInt8Contains(t *testing.T) {
	m := NewMapUint32ToInt8(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), int8(signedData[i]))
		if !m.Contains(uint32(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint32(k))
		if m.Contains(uint32(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint32ToInt8Get(t *testing.T) {
	m := NewMapUint32ToInt8(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), int8(signedData[i]))
		if g := m.Get(uint32(k)); g != int8(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapUint32ToInt8Clear(t *testing.T) {
	m := NewMapUint32ToInt8(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), int8(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint32(k), int8(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUint32ToInt8Inc(t *testing.T) {
	m := NewMapUint32ToInt8(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), int8(signedData[i]))
	}
	for i, k := range unsignedData {
		m.Inc(uint32(k))
		m.Inc(uint32(k))
		m.Inc(uint32(k))
		if g := m.Get(uint32(k)); g != int8(signedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, int8(signedData[i])+3)
		}
	}
	m = NewMapUint32ToInt8(13)
	for _, k := range unsignedData {
		m.Inc(uint32(k))
		if g := m.Get(uint32(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapUint32ToInt8Add(t *testing.T) {
	m := NewMapUint32ToInt8(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), int8(signedData[i]))
	}
	for i, k := range unsignedData {
		m.Add(uint32(k), 3)
		if g := m.Get(uint32(k)); g != int8(signedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, int8(signedData[i])+3)
		}
	}
	m = NewMapUint32ToInt8(13)
	for _, k := range unsignedData {
		m.Add(uint32(k), 42)
		if g := m.Get(uint32(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapUint32ToInt8Modify(t *testing.T) {
	m := NewMapUint32ToInt8(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), int8(signedData[i]))
	}
	for _, k := range unsignedData {
		m.Modify(uint32(k), func(v *int8) {
			*v += 3
		})
	}
	for i, k := range unsignedData {
		if g := m.Get(uint32(k)); g != int8(signedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, int8(signedData[i])+3)
		}
	}
}

func TestMapUint32ToInt8Find(t *testing.T) {
	m := NewMapUint32ToInt8(13)

	low := unsignedData[:len(unsignedData)/2]
	hi := unsignedData[len(unsignedData)/2:]

	for i, k := range low {
		m.Put(uint32(k), int8(signedData[i]))
	}
	for i, k := range low {
		v, ok := m.Find(uint32(k))
		if !ok {
			t.Errorf("missing value for key %d\n", k)
			continue
		}
		if want := int8(signedData[i]); v != want {
			t.Errorf("key %d: got %d, want %d\n", k, v, want)
		}
	}
	for _, k := range hi {
		if v, ok := m.Find(uint32(k)); ok {
			t.Errorf("key %d: got %d, want nothing\n", k, v)
		}
	}
}

func TestMapUint32ToInt8Put(t *testing.T) {
	m := NewMapUint32ToInt8(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), int8(signedData[i]))
	}
	for _, k := range unsignedData {
		m.Put(uint32(k), m.Get(uint32(k))+3)
	}
	for i, k := range unsignedData {
		if g := m.Get(uint32(k)); g != int8(signedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, int8(signedData[i])+3)
		}
	}
}

func TestMapUint32ToInt8Visit(t *testing.T) {
	m := NewMapUint32ToInt8(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), int8(signedData[i]))
	}

	n := make(map[uint32]int8, len(unsignedData))

	for i, k := range unsignedData {
		n[uint32(k)] = int8(signedData[i])
	}

	m.Visit(func(k uint32, v int8) {
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

func TestMapUint32ToInt16Size(t *testing.T) {
	m := NewMapUint32ToInt16(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint32(k), int16(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint32(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint32ToInt16Contains(t *testing.T) {
	m := NewMapUint32ToInt16(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), int16(signedData[i]))
		if !m.Contains(uint32(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint32(k))
		if m.Contains(uint32(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint32ToInt16Get(t *testing.T) {
	m := NewMapUint32ToInt16(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), int16(signedData[i]))
		if g := m.Get(uint32(k)); g != int16(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapUint32ToInt16Clear(t *testing.T) {
	m := NewMapUint32ToInt16(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), int16(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint32(k), int16(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUint32ToInt16Inc(t *testing.T) {
	m := NewMapUint32ToInt16(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), int16(signedData[i]))
	}
	for i, k := range unsignedData {
		m.Inc(uint32(k))
		m.Inc(uint32(k))
		m.Inc(uint32(k))
		if g := m.Get(uint32(k)); g != int16(signedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, int16(signedData[i])+3)
		}
	}
	m = NewMapUint32ToInt16(13)
	for _, k := range unsignedData {
		m.Inc(uint32(k))
		if g := m.Get(uint32(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapUint32ToInt16Add(t *testing.T) {
	m := NewMapUint32ToInt16(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), int16(signedData[i]))
	}
	for i, k := range unsignedData {
		m.Add(uint32(k), 3)
		if g := m.Get(uint32(k)); g != int16(signedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, int16(signedData[i])+3)
		}
	}
	m = NewMapUint32ToInt16(13)
	for _, k := range unsignedData {
		m.Add(uint32(k), 42)
		if g := m.Get(uint32(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapUint32ToInt16Modify(t *testing.T) {
	m := NewMapUint32ToInt16(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), int16(signedData[i]))
	}
	for _, k := range unsignedData {
		m.Modify(uint32(k), func(v *int16) {
			*v += 3
		})
	}
	for i, k := range unsignedData {
		if g := m.Get(uint32(k)); g != int16(signedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, int16(signedData[i])+3)
		}
	}
}

func TestMapUint32ToInt16Find(t *testing.T) {
	m := NewMapUint32ToInt16(13)

	low := unsignedData[:len(unsignedData)/2]
	hi := unsignedData[len(unsignedData)/2:]

	for i, k := range low {
		m.Put(uint32(k), int16(signedData[i]))
	}
	for i, k := range low {
		v, ok := m.Find(uint32(k))
		if !ok {
			t.Errorf("missing value for key %d\n", k)
			continue
		}
		if want := int16(signedData[i]); v != want {
			t.Errorf("key %d: got %d, want %d\n", k, v, want)
		}
	}
	for _, k := range hi {
		if v, ok := m.Find(uint32(k)); ok {
			t.Errorf("key %d: got %d, want nothing\n", k, v)
		}
	}
}

func TestMapUint32ToInt16Put(t *testing.T) {
	m := NewMapUint32ToInt16(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), int16(signedData[i]))
	}
	for _, k := range unsignedData {
		m.Put(uint32(k), m.Get(uint32(k))+3)
	}
	for i, k := range unsignedData {
		if g := m.Get(uint32(k)); g != int16(signedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, int16(signedData[i])+3)
		}
	}
}

func TestMapUint32ToInt16Visit(t *testing.T) {
	m := NewMapUint32ToInt16(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), int16(signedData[i]))
	}

	n := make(map[uint32]int16, len(unsignedData))

	for i, k := range unsignedData {
		n[uint32(k)] = int16(signedData[i])
	}

	m.Visit(func(k uint32, v int16) {
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

func TestMapUint32ToInt32Size(t *testing.T) {
	m := NewMapUint32ToInt32(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint32(k), int32(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint32(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint32ToInt32Contains(t *testing.T) {
	m := NewMapUint32ToInt32(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), int32(signedData[i]))
		if !m.Contains(uint32(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint32(k))
		if m.Contains(uint32(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint32ToInt32Get(t *testing.T) {
	m := NewMapUint32ToInt32(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), int32(signedData[i]))
		if g := m.Get(uint32(k)); g != int32(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapUint32ToInt32Clear(t *testing.T) {
	m := NewMapUint32ToInt32(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), int32(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint32(k), int32(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUint32ToInt32Inc(t *testing.T) {
	m := NewMapUint32ToInt32(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), int32(signedData[i]))
	}
	for i, k := range unsignedData {
		m.Inc(uint32(k))
		m.Inc(uint32(k))
		m.Inc(uint32(k))
		if g := m.Get(uint32(k)); g != int32(signedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, int32(signedData[i])+3)
		}
	}
	m = NewMapUint32ToInt32(13)
	for _, k := range unsignedData {
		m.Inc(uint32(k))
		if g := m.Get(uint32(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapUint32ToInt32Add(t *testing.T) {
	m := NewMapUint32ToInt32(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), int32(signedData[i]))
	}
	for i, k := range unsignedData {
		m.Add(uint32(k), 3)
		if g := m.Get(uint32(k)); g != int32(signedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, int32(signedData[i])+3)
		}
	}
	m = NewMapUint32ToInt32(13)
	for _, k := range unsignedData {
		m.Add(uint32(k), 42)
		if g := m.Get(uint32(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapUint32ToInt32Modify(t *testing.T) {
	m := NewMapUint32ToInt32(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), int32(signedData[i]))
	}
	for _, k := range unsignedData {
		m.Modify(uint32(k), func(v *int32) {
			*v += 3
		})
	}
	for i, k := range unsignedData {
		if g := m.Get(uint32(k)); g != int32(signedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, int32(signedData[i])+3)
		}
	}
}

func TestMapUint32ToInt32Find(t *testing.T) {
	m := NewMapUint32ToInt32(13)

	low := unsignedData[:len(unsignedData)/2]
	hi := unsignedData[len(unsignedData)/2:]

	for i, k := range low {
		m.Put(uint32(k), int32(signedData[i]))
	}
	for i, k := range low {
		v, ok := m.Find(uint32(k))
		if !ok {
			t.Errorf("missing value for key %d\n", k)
			continue
		}
		if want := int32(signedData[i]); v != want {
			t.Errorf("key %d: got %d, want %d\n", k, v, want)
		}
	}
	for _, k := range hi {
		if v, ok := m.Find(uint32(k)); ok {
			t.Errorf("key %d: got %d, want nothing\n", k, v)
		}
	}
}

func TestMapUint32ToInt32Put(t *testing.T) {
	m := NewMapUint32ToInt32(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), int32(signedData[i]))
	}
	for _, k := range unsignedData {
		m.Put(uint32(k), m.Get(uint32(k))+3)
	}
	for i, k := range unsignedData {
		if g := m.Get(uint32(k)); g != int32(signedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, int32(signedData[i])+3)
		}
	}
}

func TestMapUint32ToInt32Visit(t *testing.T) {
	m := NewMapUint32ToInt32(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), int32(signedData[i]))
	}

	n := make(map[uint32]int32, len(unsignedData))

	for i, k := range unsignedData {
		n[uint32(k)] = int32(signedData[i])
	}

	m.Visit(func(k uint32, v int32) {
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

func TestMapUint32ToInt64Size(t *testing.T) {
	m := NewMapUint32ToInt64(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint32(k), int64(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint32(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint32ToInt64Contains(t *testing.T) {
	m := NewMapUint32ToInt64(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), int64(signedData[i]))
		if !m.Contains(uint32(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint32(k))
		if m.Contains(uint32(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint32ToInt64Get(t *testing.T) {
	m := NewMapUint32ToInt64(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), int64(signedData[i]))
		if g := m.Get(uint32(k)); g != int64(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapUint32ToInt64Clear(t *testing.T) {
	m := NewMapUint32ToInt64(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), int64(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint32(k), int64(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUint32ToInt64Inc(t *testing.T) {
	m := NewMapUint32ToInt64(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), int64(signedData[i]))
	}
	for i, k := range unsignedData {
		m.Inc(uint32(k))
		m.Inc(uint32(k))
		m.Inc(uint32(k))
		if g := m.Get(uint32(k)); g != int64(signedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, int64(signedData[i])+3)
		}
	}
	m = NewMapUint32ToInt64(13)
	for _, k := range unsignedData {
		m.Inc(uint32(k))
		if g := m.Get(uint32(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapUint32ToInt64Add(t *testing.T) {
	m := NewMapUint32ToInt64(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), int64(signedData[i]))
	}
	for i, k := range unsignedData {
		m.Add(uint32(k), 3)
		if g := m.Get(uint32(k)); g != int64(signedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, int64(signedData[i])+3)
		}
	}
	m = NewMapUint32ToInt64(13)
	for _, k := range unsignedData {
		m.Add(uint32(k), 42)
		if g := m.Get(uint32(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapUint32ToInt64Modify(t *testing.T) {
	m := NewMapUint32ToInt64(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), int64(signedData[i]))
	}
	for _, k := range unsignedData {
		m.Modify(uint32(k), func(v *int64) {
			*v += 3
		})
	}
	for i, k := range unsignedData {
		if g := m.Get(uint32(k)); g != int64(signedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, int64(signedData[i])+3)
		}
	}
}

func TestMapUint32ToInt64Find(t *testing.T) {
	m := NewMapUint32ToInt64(13)

	low := unsignedData[:len(unsignedData)/2]
	hi := unsignedData[len(unsignedData)/2:]

	for i, k := range low {
		m.Put(uint32(k), int64(signedData[i]))
	}
	for i, k := range low {
		v, ok := m.Find(uint32(k))
		if !ok {
			t.Errorf("missing value for key %d\n", k)
			continue
		}
		if want := int64(signedData[i]); v != want {
			t.Errorf("key %d: got %d, want %d\n", k, v, want)
		}
	}
	for _, k := range hi {
		if v, ok := m.Find(uint32(k)); ok {
			t.Errorf("key %d: got %d, want nothing\n", k, v)
		}
	}
}

func TestMapUint32ToInt64Put(t *testing.T) {
	m := NewMapUint32ToInt64(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), int64(signedData[i]))
	}
	for _, k := range unsignedData {
		m.Put(uint32(k), m.Get(uint32(k))+3)
	}
	for i, k := range unsignedData {
		if g := m.Get(uint32(k)); g != int64(signedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, int64(signedData[i])+3)
		}
	}
}

func TestMapUint32ToInt64Visit(t *testing.T) {
	m := NewMapUint32ToInt64(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), int64(signedData[i]))
	}

	n := make(map[uint32]int64, len(unsignedData))

	for i, k := range unsignedData {
		n[uint32(k)] = int64(signedData[i])
	}

	m.Visit(func(k uint32, v int64) {
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

func TestMapUint32ToUintSize(t *testing.T) {
	m := NewMapUint32ToUint(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint32(k), uint(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint32(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint32ToUintContains(t *testing.T) {
	m := NewMapUint32ToUint(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), uint(unsignedData[i]))
		if !m.Contains(uint32(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint32(k))
		if m.Contains(uint32(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint32ToUintGet(t *testing.T) {
	m := NewMapUint32ToUint(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), uint(unsignedData[i]))
		if g := m.Get(uint32(k)); g != uint(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapUint32ToUintClear(t *testing.T) {
	m := NewMapUint32ToUint(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), uint(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint32(k), uint(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUint32ToUintInc(t *testing.T) {
	m := NewMapUint32ToUint(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), uint(unsignedData[i]))
	}
	for i, k := range unsignedData {
		m.Inc(uint32(k))
		m.Inc(uint32(k))
		m.Inc(uint32(k))
		if g := m.Get(uint32(k)); g != uint(unsignedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, uint(unsignedData[i])+3)
		}
	}
	m = NewMapUint32ToUint(13)
	for _, k := range unsignedData {
		m.Inc(uint32(k))
		if g := m.Get(uint32(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapUint32ToUintAdd(t *testing.T) {
	m := NewMapUint32ToUint(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), uint(unsignedData[i]))
	}
	for i, k := range unsignedData {
		m.Add(uint32(k), 3)
		if g := m.Get(uint32(k)); g != uint(unsignedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, uint(unsignedData[i])+3)
		}
	}
	m = NewMapUint32ToUint(13)
	for _, k := range unsignedData {
		m.Add(uint32(k), 42)
		if g := m.Get(uint32(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapUint32ToUintModify(t *testing.T) {
	m := NewMapUint32ToUint(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), uint(unsignedData[i]))
	}
	for _, k := range unsignedData {
		m.Modify(uint32(k), func(v *uint) {
			*v += 3
		})
	}
	for i, k := range unsignedData {
		if g := m.Get(uint32(k)); g != uint(unsignedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, uint(unsignedData[i])+3)
		}
	}
}

func TestMapUint32ToUintFind(t *testing.T) {
	m := NewMapUint32ToUint(13)

	low := unsignedData[:len(unsignedData)/2]
	hi := unsignedData[len(unsignedData)/2:]

	for i, k := range low {
		m.Put(uint32(k), uint(unsignedData[i]))
	}
	for i, k := range low {
		v, ok := m.Find(uint32(k))
		if !ok {
			t.Errorf("missing value for key %d\n", k)
			continue
		}
		if want := uint(unsignedData[i]); v != want {
			t.Errorf("key %d: got %d, want %d\n", k, v, want)
		}
	}
	for _, k := range hi {
		if v, ok := m.Find(uint32(k)); ok {
			t.Errorf("key %d: got %d, want nothing\n", k, v)
		}
	}
}

func TestMapUint32ToUintPut(t *testing.T) {
	m := NewMapUint32ToUint(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), uint(unsignedData[i]))
	}
	for _, k := range unsignedData {
		m.Put(uint32(k), m.Get(uint32(k))+3)
	}
	for i, k := range unsignedData {
		if g := m.Get(uint32(k)); g != uint(unsignedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, uint(unsignedData[i])+3)
		}
	}
}

func TestMapUint32ToUintVisit(t *testing.T) {
	m := NewMapUint32ToUint(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), uint(unsignedData[i]))
	}

	n := make(map[uint32]uint, len(unsignedData))

	for i, k := range unsignedData {
		n[uint32(k)] = uint(unsignedData[i])
	}

	m.Visit(func(k uint32, v uint) {
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

func TestMapUint32ToUint8Size(t *testing.T) {
	m := NewMapUint32ToUint8(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint32(k), uint8(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint32(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint32ToUint8Contains(t *testing.T) {
	m := NewMapUint32ToUint8(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), uint8(unsignedData[i]))
		if !m.Contains(uint32(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint32(k))
		if m.Contains(uint32(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint32ToUint8Get(t *testing.T) {
	m := NewMapUint32ToUint8(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), uint8(unsignedData[i]))
		if g := m.Get(uint32(k)); g != uint8(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapUint32ToUint8Clear(t *testing.T) {
	m := NewMapUint32ToUint8(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), uint8(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint32(k), uint8(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUint32ToUint8Inc(t *testing.T) {
	m := NewMapUint32ToUint8(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), uint8(unsignedData[i]))
	}
	for i, k := range unsignedData {
		m.Inc(uint32(k))
		m.Inc(uint32(k))
		m.Inc(uint32(k))
		if g := m.Get(uint32(k)); g != uint8(unsignedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, uint8(unsignedData[i])+3)
		}
	}
	m = NewMapUint32ToUint8(13)
	for _, k := range unsignedData {
		m.Inc(uint32(k))
		if g := m.Get(uint32(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapUint32ToUint8Add(t *testing.T) {
	m := NewMapUint32ToUint8(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), uint8(unsignedData[i]))
	}
	for i, k := range unsignedData {
		m.Add(uint32(k), 3)
		if g := m.Get(uint32(k)); g != uint8(unsignedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, uint8(unsignedData[i])+3)
		}
	}
	m = NewMapUint32ToUint8(13)
	for _, k := range unsignedData {
		m.Add(uint32(k), 42)
		if g := m.Get(uint32(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapUint32ToUint8Modify(t *testing.T) {
	m := NewMapUint32ToUint8(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), uint8(unsignedData[i]))
	}
	for _, k := range unsignedData {
		m.Modify(uint32(k), func(v *uint8) {
			*v += 3
		})
	}
	for i, k := range unsignedData {
		if g := m.Get(uint32(k)); g != uint8(unsignedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, uint8(unsignedData[i])+3)
		}
	}
}

func TestMapUint32ToUint8Find(t *testing.T) {
	m := NewMapUint32ToUint8(13)

	low := unsignedData[:len(unsignedData)/2]
	hi := unsignedData[len(unsignedData)/2:]

	for i, k := range low {
		m.Put(uint32(k), uint8(unsignedData[i]))
	}
	for i, k := range low {
		v, ok := m.Find(uint32(k))
		if !ok {
			t.Errorf("missing value for key %d\n", k)
			continue
		}
		if want := uint8(unsignedData[i]); v != want {
			t.Errorf("key %d: got %d, want %d\n", k, v, want)
		}
	}
	for _, k := range hi {
		if v, ok := m.Find(uint32(k)); ok {
			t.Errorf("key %d: got %d, want nothing\n", k, v)
		}
	}
}

func TestMapUint32ToUint8Put(t *testing.T) {
	m := NewMapUint32ToUint8(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), uint8(unsignedData[i]))
	}
	for _, k := range unsignedData {
		m.Put(uint32(k), m.Get(uint32(k))+3)
	}
	for i, k := range unsignedData {
		if g := m.Get(uint32(k)); g != uint8(unsignedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, uint8(unsignedData[i])+3)
		}
	}
}

func TestMapUint32ToUint8Visit(t *testing.T) {
	m := NewMapUint32ToUint8(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), uint8(unsignedData[i]))
	}

	n := make(map[uint32]uint8, len(unsignedData))

	for i, k := range unsignedData {
		n[uint32(k)] = uint8(unsignedData[i])
	}

	m.Visit(func(k uint32, v uint8) {
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

func TestMapUint32ToUint16Size(t *testing.T) {
	m := NewMapUint32ToUint16(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint32(k), uint16(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint32(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint32ToUint16Contains(t *testing.T) {
	m := NewMapUint32ToUint16(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), uint16(unsignedData[i]))
		if !m.Contains(uint32(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint32(k))
		if m.Contains(uint32(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint32ToUint16Get(t *testing.T) {
	m := NewMapUint32ToUint16(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), uint16(unsignedData[i]))
		if g := m.Get(uint32(k)); g != uint16(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapUint32ToUint16Clear(t *testing.T) {
	m := NewMapUint32ToUint16(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), uint16(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint32(k), uint16(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUint32ToUint16Inc(t *testing.T) {
	m := NewMapUint32ToUint16(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), uint16(unsignedData[i]))
	}
	for i, k := range unsignedData {
		m.Inc(uint32(k))
		m.Inc(uint32(k))
		m.Inc(uint32(k))
		if g := m.Get(uint32(k)); g != uint16(unsignedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, uint16(unsignedData[i])+3)
		}
	}
	m = NewMapUint32ToUint16(13)
	for _, k := range unsignedData {
		m.Inc(uint32(k))
		if g := m.Get(uint32(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapUint32ToUint16Add(t *testing.T) {
	m := NewMapUint32ToUint16(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), uint16(unsignedData[i]))
	}
	for i, k := range unsignedData {
		m.Add(uint32(k), 3)
		if g := m.Get(uint32(k)); g != uint16(unsignedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, uint16(unsignedData[i])+3)
		}
	}
	m = NewMapUint32ToUint16(13)
	for _, k := range unsignedData {
		m.Add(uint32(k), 42)
		if g := m.Get(uint32(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapUint32ToUint16Modify(t *testing.T) {
	m := NewMapUint32ToUint16(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), uint16(unsignedData[i]))
	}
	for _, k := range unsignedData {
		m.Modify(uint32(k), func(v *uint16) {
			*v += 3
		})
	}
	for i, k := range unsignedData {
		if g := m.Get(uint32(k)); g != uint16(unsignedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, uint16(unsignedData[i])+3)
		}
	}
}

func TestMapUint32ToUint16Find(t *testing.T) {
	m := NewMapUint32ToUint16(13)

	low := unsignedData[:len(unsignedData)/2]
	hi := unsignedData[len(unsignedData)/2:]

	for i, k := range low {
		m.Put(uint32(k), uint16(unsignedData[i]))
	}
	for i, k := range low {
		v, ok := m.Find(uint32(k))
		if !ok {
			t.Errorf("missing value for key %d\n", k)
			continue
		}
		if want := uint16(unsignedData[i]); v != want {
			t.Errorf("key %d: got %d, want %d\n", k, v, want)
		}
	}
	for _, k := range hi {
		if v, ok := m.Find(uint32(k)); ok {
			t.Errorf("key %d: got %d, want nothing\n", k, v)
		}
	}
}

func TestMapUint32ToUint16Put(t *testing.T) {
	m := NewMapUint32ToUint16(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), uint16(unsignedData[i]))
	}
	for _, k := range unsignedData {
		m.Put(uint32(k), m.Get(uint32(k))+3)
	}
	for i, k := range unsignedData {
		if g := m.Get(uint32(k)); g != uint16(unsignedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, uint16(unsignedData[i])+3)
		}
	}
}

func TestMapUint32ToUint16Visit(t *testing.T) {
	m := NewMapUint32ToUint16(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), uint16(unsignedData[i]))
	}

	n := make(map[uint32]uint16, len(unsignedData))

	for i, k := range unsignedData {
		n[uint32(k)] = uint16(unsignedData[i])
	}

	m.Visit(func(k uint32, v uint16) {
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

func TestMapUint32ToUint32Size(t *testing.T) {
	m := NewMapUint32ToUint32(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint32(k), uint32(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint32(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint32ToUint32Contains(t *testing.T) {
	m := NewMapUint32ToUint32(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), uint32(unsignedData[i]))
		if !m.Contains(uint32(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint32(k))
		if m.Contains(uint32(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint32ToUint32Get(t *testing.T) {
	m := NewMapUint32ToUint32(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), uint32(unsignedData[i]))
		if g := m.Get(uint32(k)); g != uint32(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapUint32ToUint32Clear(t *testing.T) {
	m := NewMapUint32ToUint32(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), uint32(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint32(k), uint32(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUint32ToUint32Inc(t *testing.T) {
	m := NewMapUint32ToUint32(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), uint32(unsignedData[i]))
	}
	for i, k := range unsignedData {
		m.Inc(uint32(k))
		m.Inc(uint32(k))
		m.Inc(uint32(k))
		if g := m.Get(uint32(k)); g != uint32(unsignedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, uint32(unsignedData[i])+3)
		}
	}
	m = NewMapUint32ToUint32(13)
	for _, k := range unsignedData {
		m.Inc(uint32(k))
		if g := m.Get(uint32(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapUint32ToUint32Add(t *testing.T) {
	m := NewMapUint32ToUint32(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), uint32(unsignedData[i]))
	}
	for i, k := range unsignedData {
		m.Add(uint32(k), 3)
		if g := m.Get(uint32(k)); g != uint32(unsignedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, uint32(unsignedData[i])+3)
		}
	}
	m = NewMapUint32ToUint32(13)
	for _, k := range unsignedData {
		m.Add(uint32(k), 42)
		if g := m.Get(uint32(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapUint32ToUint32Modify(t *testing.T) {
	m := NewMapUint32ToUint32(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), uint32(unsignedData[i]))
	}
	for _, k := range unsignedData {
		m.Modify(uint32(k), func(v *uint32) {
			*v += 3
		})
	}
	for i, k := range unsignedData {
		if g := m.Get(uint32(k)); g != uint32(unsignedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, uint32(unsignedData[i])+3)
		}
	}
}

func TestMapUint32ToUint32Find(t *testing.T) {
	m := NewMapUint32ToUint32(13)

	low := unsignedData[:len(unsignedData)/2]
	hi := unsignedData[len(unsignedData)/2:]

	for i, k := range low {
		m.Put(uint32(k), uint32(unsignedData[i]))
	}
	for i, k := range low {
		v, ok := m.Find(uint32(k))
		if !ok {
			t.Errorf("missing value for key %d\n", k)
			continue
		}
		if want := uint32(unsignedData[i]); v != want {
			t.Errorf("key %d: got %d, want %d\n", k, v, want)
		}
	}
	for _, k := range hi {
		if v, ok := m.Find(uint32(k)); ok {
			t.Errorf("key %d: got %d, want nothing\n", k, v)
		}
	}
}

func TestMapUint32ToUint32Put(t *testing.T) {
	m := NewMapUint32ToUint32(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), uint32(unsignedData[i]))
	}
	for _, k := range unsignedData {
		m.Put(uint32(k), m.Get(uint32(k))+3)
	}
	for i, k := range unsignedData {
		if g := m.Get(uint32(k)); g != uint32(unsignedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, uint32(unsignedData[i])+3)
		}
	}
}

func TestMapUint32ToUint32Visit(t *testing.T) {
	m := NewMapUint32ToUint32(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), uint32(unsignedData[i]))
	}

	n := make(map[uint32]uint32, len(unsignedData))

	for i, k := range unsignedData {
		n[uint32(k)] = uint32(unsignedData[i])
	}

	m.Visit(func(k uint32, v uint32) {
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

func TestMapUint32ToUint64Size(t *testing.T) {
	m := NewMapUint32ToUint64(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint32(k), uint64(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint32(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint32ToUint64Contains(t *testing.T) {
	m := NewMapUint32ToUint64(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), uint64(unsignedData[i]))
		if !m.Contains(uint32(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint32(k))
		if m.Contains(uint32(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint32ToUint64Get(t *testing.T) {
	m := NewMapUint32ToUint64(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), uint64(unsignedData[i]))
		if g := m.Get(uint32(k)); g != uint64(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapUint32ToUint64Clear(t *testing.T) {
	m := NewMapUint32ToUint64(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), uint64(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint32(k), uint64(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUint32ToUint64Inc(t *testing.T) {
	m := NewMapUint32ToUint64(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), uint64(unsignedData[i]))
	}
	for i, k := range unsignedData {
		m.Inc(uint32(k))
		m.Inc(uint32(k))
		m.Inc(uint32(k))
		if g := m.Get(uint32(k)); g != uint64(unsignedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, uint64(unsignedData[i])+3)
		}
	}
	m = NewMapUint32ToUint64(13)
	for _, k := range unsignedData {
		m.Inc(uint32(k))
		if g := m.Get(uint32(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapUint32ToUint64Add(t *testing.T) {
	m := NewMapUint32ToUint64(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), uint64(unsignedData[i]))
	}
	for i, k := range unsignedData {
		m.Add(uint32(k), 3)
		if g := m.Get(uint32(k)); g != uint64(unsignedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, uint64(unsignedData[i])+3)
		}
	}
	m = NewMapUint32ToUint64(13)
	for _, k := range unsignedData {
		m.Add(uint32(k), 42)
		if g := m.Get(uint32(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapUint32ToUint64Modify(t *testing.T) {
	m := NewMapUint32ToUint64(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), uint64(unsignedData[i]))
	}
	for _, k := range unsignedData {
		m.Modify(uint32(k), func(v *uint64) {
			*v += 3
		})
	}
	for i, k := range unsignedData {
		if g := m.Get(uint32(k)); g != uint64(unsignedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, uint64(unsignedData[i])+3)
		}
	}
}

func TestMapUint32ToUint64Find(t *testing.T) {
	m := NewMapUint32ToUint64(13)

	low := unsignedData[:len(unsignedData)/2]
	hi := unsignedData[len(unsignedData)/2:]

	for i, k := range low {
		m.Put(uint32(k), uint64(unsignedData[i]))
	}
	for i, k := range low {
		v, ok := m.Find(uint32(k))
		if !ok {
			t.Errorf("missing value for key %d\n", k)
			continue
		}
		if want := uint64(unsignedData[i]); v != want {
			t.Errorf("key %d: got %d, want %d\n", k, v, want)
		}
	}
	for _, k := range hi {
		if v, ok := m.Find(uint32(k)); ok {
			t.Errorf("key %d: got %d, want nothing\n", k, v)
		}
	}
}

func TestMapUint32ToUint64Put(t *testing.T) {
	m := NewMapUint32ToUint64(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), uint64(unsignedData[i]))
	}
	for _, k := range unsignedData {
		m.Put(uint32(k), m.Get(uint32(k))+3)
	}
	for i, k := range unsignedData {
		if g := m.Get(uint32(k)); g != uint64(unsignedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, uint64(unsignedData[i])+3)
		}
	}
}

func TestMapUint32ToUint64Visit(t *testing.T) {
	m := NewMapUint32ToUint64(13)
	for i, k := range unsignedData {
		m.Put(uint32(k), uint64(unsignedData[i]))
	}

	n := make(map[uint32]uint64, len(unsignedData))

	for i, k := range unsignedData {
		n[uint32(k)] = uint64(unsignedData[i])
	}

	m.Visit(func(k uint32, v uint64) {
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
