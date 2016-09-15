// THIS IS A MACHINE GENERATED FILE!
// BE CAREFUL WITH EDITING BY HAND.
//
// This is Free Software covered by the terms of the MIT license.
// See LICENSE file for details.
// (c) 2016 by Sascha L. Teichmann.
// See the full list of contributors in the CONTRIBUTORS file.

package genmap

import "testing"

func TestMapIntToIntSize(t *testing.T) {
	m := NewMapIntToInt(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int(k), int(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapIntToIntContains(t *testing.T) {
	m := NewMapIntToInt(13)
	for i, k := range signedData {
		m.Put(int(k), int(signedData[i]))
		if !m.Contains(int(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int(k))
		if m.Contains(int(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapIntToIntGet(t *testing.T) {
	m := NewMapIntToInt(13)
	for i, k := range signedData {
		m.Put(int(k), int(signedData[i]))
		if g := m.Get(int(k)); g != int(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapIntToIntClear(t *testing.T) {
	m := NewMapIntToInt(13)
	for i, k := range signedData {
		m.Put(int(k), int(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int(k), int(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapIntToIntInc(t *testing.T) {
	m := NewMapIntToInt(13)
	for i, k := range signedData {
		m.Put(int(k), int(signedData[i]))
	}
	for i, k := range signedData {
		m.Inc(int(k))
		m.Inc(int(k))
		m.Inc(int(k))
		if g := m.Get(int(k)); g != int(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int(signedData[i])+3)
		}
	}
	m = NewMapIntToInt(13)
	for _, k := range signedData {
		m.Inc(int(k))
		if g := m.Get(int(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapIntToIntAdd(t *testing.T) {
	m := NewMapIntToInt(13)
	for i, k := range signedData {
		m.Put(int(k), int(signedData[i]))
	}
	for i, k := range signedData {
		m.Add(int(k), 3)
		if g := m.Get(int(k)); g != int(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int(signedData[i])+3)
		}
	}
	m = NewMapIntToInt(13)
	for _, k := range signedData {
		m.Add(int(k), 42)
		if g := m.Get(int(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapIntToIntModify(t *testing.T) {
	m := NewMapIntToInt(13)
	for i, k := range signedData {
		m.Put(int(k), int(signedData[i]))
	}
	for _, k := range signedData {
		m.Modify(int(k), func(v *int) {
			*v += 3
		})
	}
	for i, k := range signedData {
		if g := m.Get(int(k)); g != int(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int(signedData[i])+3)
		}
	}
}

func TestMapIntToIntPut(t *testing.T) {
	m := NewMapIntToInt(13)
	for i, k := range signedData {
		m.Put(int(k), int(signedData[i]))
	}
	for _, k := range signedData {
		m.Put(int(k), m.Get(int(k))+3)
	}
	for i, k := range signedData {
		if g := m.Get(int(k)); g != int(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int(signedData[i])+3)
		}
	}
}

func TestMapIntToIntVisit(t *testing.T) {
	m := NewMapIntToInt(13)
	for i, k := range signedData {
		m.Put(int(k), int(signedData[i]))
	}

	n := make(map[int]int, len(signedData))

	for i, k := range signedData {
		n[int(k)] = int(signedData[i])
	}

	m.Visit(func(k int, v int) {
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

func TestMapIntToInt8Size(t *testing.T) {
	m := NewMapIntToInt8(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int(k), int8(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapIntToInt8Contains(t *testing.T) {
	m := NewMapIntToInt8(13)
	for i, k := range signedData {
		m.Put(int(k), int8(signedData[i]))
		if !m.Contains(int(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int(k))
		if m.Contains(int(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapIntToInt8Get(t *testing.T) {
	m := NewMapIntToInt8(13)
	for i, k := range signedData {
		m.Put(int(k), int8(signedData[i]))
		if g := m.Get(int(k)); g != int8(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapIntToInt8Clear(t *testing.T) {
	m := NewMapIntToInt8(13)
	for i, k := range signedData {
		m.Put(int(k), int8(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int(k), int8(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapIntToInt8Inc(t *testing.T) {
	m := NewMapIntToInt8(13)
	for i, k := range signedData {
		m.Put(int(k), int8(signedData[i]))
	}
	for i, k := range signedData {
		m.Inc(int(k))
		m.Inc(int(k))
		m.Inc(int(k))
		if g := m.Get(int(k)); g != int8(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int8(signedData[i])+3)
		}
	}
	m = NewMapIntToInt8(13)
	for _, k := range signedData {
		m.Inc(int(k))
		if g := m.Get(int(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapIntToInt8Add(t *testing.T) {
	m := NewMapIntToInt8(13)
	for i, k := range signedData {
		m.Put(int(k), int8(signedData[i]))
	}
	for i, k := range signedData {
		m.Add(int(k), 3)
		if g := m.Get(int(k)); g != int8(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int8(signedData[i])+3)
		}
	}
	m = NewMapIntToInt8(13)
	for _, k := range signedData {
		m.Add(int(k), 42)
		if g := m.Get(int(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapIntToInt8Modify(t *testing.T) {
	m := NewMapIntToInt8(13)
	for i, k := range signedData {
		m.Put(int(k), int8(signedData[i]))
	}
	for _, k := range signedData {
		m.Modify(int(k), func(v *int8) {
			*v += 3
		})
	}
	for i, k := range signedData {
		if g := m.Get(int(k)); g != int8(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int8(signedData[i])+3)
		}
	}
}

func TestMapIntToInt8Put(t *testing.T) {
	m := NewMapIntToInt8(13)
	for i, k := range signedData {
		m.Put(int(k), int8(signedData[i]))
	}
	for _, k := range signedData {
		m.Put(int(k), m.Get(int(k))+3)
	}
	for i, k := range signedData {
		if g := m.Get(int(k)); g != int8(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int8(signedData[i])+3)
		}
	}
}

func TestMapIntToInt8Visit(t *testing.T) {
	m := NewMapIntToInt8(13)
	for i, k := range signedData {
		m.Put(int(k), int8(signedData[i]))
	}

	n := make(map[int]int8, len(signedData))

	for i, k := range signedData {
		n[int(k)] = int8(signedData[i])
	}

	m.Visit(func(k int, v int8) {
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

func TestMapIntToInt16Size(t *testing.T) {
	m := NewMapIntToInt16(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int(k), int16(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapIntToInt16Contains(t *testing.T) {
	m := NewMapIntToInt16(13)
	for i, k := range signedData {
		m.Put(int(k), int16(signedData[i]))
		if !m.Contains(int(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int(k))
		if m.Contains(int(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapIntToInt16Get(t *testing.T) {
	m := NewMapIntToInt16(13)
	for i, k := range signedData {
		m.Put(int(k), int16(signedData[i]))
		if g := m.Get(int(k)); g != int16(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapIntToInt16Clear(t *testing.T) {
	m := NewMapIntToInt16(13)
	for i, k := range signedData {
		m.Put(int(k), int16(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int(k), int16(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapIntToInt16Inc(t *testing.T) {
	m := NewMapIntToInt16(13)
	for i, k := range signedData {
		m.Put(int(k), int16(signedData[i]))
	}
	for i, k := range signedData {
		m.Inc(int(k))
		m.Inc(int(k))
		m.Inc(int(k))
		if g := m.Get(int(k)); g != int16(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int16(signedData[i])+3)
		}
	}
	m = NewMapIntToInt16(13)
	for _, k := range signedData {
		m.Inc(int(k))
		if g := m.Get(int(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapIntToInt16Add(t *testing.T) {
	m := NewMapIntToInt16(13)
	for i, k := range signedData {
		m.Put(int(k), int16(signedData[i]))
	}
	for i, k := range signedData {
		m.Add(int(k), 3)
		if g := m.Get(int(k)); g != int16(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int16(signedData[i])+3)
		}
	}
	m = NewMapIntToInt16(13)
	for _, k := range signedData {
		m.Add(int(k), 42)
		if g := m.Get(int(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapIntToInt16Modify(t *testing.T) {
	m := NewMapIntToInt16(13)
	for i, k := range signedData {
		m.Put(int(k), int16(signedData[i]))
	}
	for _, k := range signedData {
		m.Modify(int(k), func(v *int16) {
			*v += 3
		})
	}
	for i, k := range signedData {
		if g := m.Get(int(k)); g != int16(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int16(signedData[i])+3)
		}
	}
}

func TestMapIntToInt16Put(t *testing.T) {
	m := NewMapIntToInt16(13)
	for i, k := range signedData {
		m.Put(int(k), int16(signedData[i]))
	}
	for _, k := range signedData {
		m.Put(int(k), m.Get(int(k))+3)
	}
	for i, k := range signedData {
		if g := m.Get(int(k)); g != int16(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int16(signedData[i])+3)
		}
	}
}

func TestMapIntToInt16Visit(t *testing.T) {
	m := NewMapIntToInt16(13)
	for i, k := range signedData {
		m.Put(int(k), int16(signedData[i]))
	}

	n := make(map[int]int16, len(signedData))

	for i, k := range signedData {
		n[int(k)] = int16(signedData[i])
	}

	m.Visit(func(k int, v int16) {
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

func TestMapIntToInt32Size(t *testing.T) {
	m := NewMapIntToInt32(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int(k), int32(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapIntToInt32Contains(t *testing.T) {
	m := NewMapIntToInt32(13)
	for i, k := range signedData {
		m.Put(int(k), int32(signedData[i]))
		if !m.Contains(int(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int(k))
		if m.Contains(int(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapIntToInt32Get(t *testing.T) {
	m := NewMapIntToInt32(13)
	for i, k := range signedData {
		m.Put(int(k), int32(signedData[i]))
		if g := m.Get(int(k)); g != int32(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapIntToInt32Clear(t *testing.T) {
	m := NewMapIntToInt32(13)
	for i, k := range signedData {
		m.Put(int(k), int32(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int(k), int32(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapIntToInt32Inc(t *testing.T) {
	m := NewMapIntToInt32(13)
	for i, k := range signedData {
		m.Put(int(k), int32(signedData[i]))
	}
	for i, k := range signedData {
		m.Inc(int(k))
		m.Inc(int(k))
		m.Inc(int(k))
		if g := m.Get(int(k)); g != int32(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int32(signedData[i])+3)
		}
	}
	m = NewMapIntToInt32(13)
	for _, k := range signedData {
		m.Inc(int(k))
		if g := m.Get(int(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapIntToInt32Add(t *testing.T) {
	m := NewMapIntToInt32(13)
	for i, k := range signedData {
		m.Put(int(k), int32(signedData[i]))
	}
	for i, k := range signedData {
		m.Add(int(k), 3)
		if g := m.Get(int(k)); g != int32(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int32(signedData[i])+3)
		}
	}
	m = NewMapIntToInt32(13)
	for _, k := range signedData {
		m.Add(int(k), 42)
		if g := m.Get(int(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapIntToInt32Modify(t *testing.T) {
	m := NewMapIntToInt32(13)
	for i, k := range signedData {
		m.Put(int(k), int32(signedData[i]))
	}
	for _, k := range signedData {
		m.Modify(int(k), func(v *int32) {
			*v += 3
		})
	}
	for i, k := range signedData {
		if g := m.Get(int(k)); g != int32(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int32(signedData[i])+3)
		}
	}
}

func TestMapIntToInt32Put(t *testing.T) {
	m := NewMapIntToInt32(13)
	for i, k := range signedData {
		m.Put(int(k), int32(signedData[i]))
	}
	for _, k := range signedData {
		m.Put(int(k), m.Get(int(k))+3)
	}
	for i, k := range signedData {
		if g := m.Get(int(k)); g != int32(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int32(signedData[i])+3)
		}
	}
}

func TestMapIntToInt32Visit(t *testing.T) {
	m := NewMapIntToInt32(13)
	for i, k := range signedData {
		m.Put(int(k), int32(signedData[i]))
	}

	n := make(map[int]int32, len(signedData))

	for i, k := range signedData {
		n[int(k)] = int32(signedData[i])
	}

	m.Visit(func(k int, v int32) {
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

func TestMapIntToInt64Size(t *testing.T) {
	m := NewMapIntToInt64(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int(k), int64(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapIntToInt64Contains(t *testing.T) {
	m := NewMapIntToInt64(13)
	for i, k := range signedData {
		m.Put(int(k), int64(signedData[i]))
		if !m.Contains(int(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int(k))
		if m.Contains(int(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapIntToInt64Get(t *testing.T) {
	m := NewMapIntToInt64(13)
	for i, k := range signedData {
		m.Put(int(k), int64(signedData[i]))
		if g := m.Get(int(k)); g != int64(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapIntToInt64Clear(t *testing.T) {
	m := NewMapIntToInt64(13)
	for i, k := range signedData {
		m.Put(int(k), int64(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int(k), int64(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapIntToInt64Inc(t *testing.T) {
	m := NewMapIntToInt64(13)
	for i, k := range signedData {
		m.Put(int(k), int64(signedData[i]))
	}
	for i, k := range signedData {
		m.Inc(int(k))
		m.Inc(int(k))
		m.Inc(int(k))
		if g := m.Get(int(k)); g != int64(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int64(signedData[i])+3)
		}
	}
	m = NewMapIntToInt64(13)
	for _, k := range signedData {
		m.Inc(int(k))
		if g := m.Get(int(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapIntToInt64Add(t *testing.T) {
	m := NewMapIntToInt64(13)
	for i, k := range signedData {
		m.Put(int(k), int64(signedData[i]))
	}
	for i, k := range signedData {
		m.Add(int(k), 3)
		if g := m.Get(int(k)); g != int64(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int64(signedData[i])+3)
		}
	}
	m = NewMapIntToInt64(13)
	for _, k := range signedData {
		m.Add(int(k), 42)
		if g := m.Get(int(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapIntToInt64Modify(t *testing.T) {
	m := NewMapIntToInt64(13)
	for i, k := range signedData {
		m.Put(int(k), int64(signedData[i]))
	}
	for _, k := range signedData {
		m.Modify(int(k), func(v *int64) {
			*v += 3
		})
	}
	for i, k := range signedData {
		if g := m.Get(int(k)); g != int64(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int64(signedData[i])+3)
		}
	}
}

func TestMapIntToInt64Put(t *testing.T) {
	m := NewMapIntToInt64(13)
	for i, k := range signedData {
		m.Put(int(k), int64(signedData[i]))
	}
	for _, k := range signedData {
		m.Put(int(k), m.Get(int(k))+3)
	}
	for i, k := range signedData {
		if g := m.Get(int(k)); g != int64(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int64(signedData[i])+3)
		}
	}
}

func TestMapIntToInt64Visit(t *testing.T) {
	m := NewMapIntToInt64(13)
	for i, k := range signedData {
		m.Put(int(k), int64(signedData[i]))
	}

	n := make(map[int]int64, len(signedData))

	for i, k := range signedData {
		n[int(k)] = int64(signedData[i])
	}

	m.Visit(func(k int, v int64) {
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

func TestMapIntToUintSize(t *testing.T) {
	m := NewMapIntToUint(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int(k), uint(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapIntToUintContains(t *testing.T) {
	m := NewMapIntToUint(13)
	for i, k := range signedData {
		m.Put(int(k), uint(unsignedData[i]))
		if !m.Contains(int(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int(k))
		if m.Contains(int(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapIntToUintGet(t *testing.T) {
	m := NewMapIntToUint(13)
	for i, k := range signedData {
		m.Put(int(k), uint(unsignedData[i]))
		if g := m.Get(int(k)); g != uint(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapIntToUintClear(t *testing.T) {
	m := NewMapIntToUint(13)
	for i, k := range signedData {
		m.Put(int(k), uint(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int(k), uint(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapIntToUintInc(t *testing.T) {
	m := NewMapIntToUint(13)
	for i, k := range signedData {
		m.Put(int(k), uint(unsignedData[i]))
	}
	for i, k := range signedData {
		m.Inc(int(k))
		m.Inc(int(k))
		m.Inc(int(k))
		if g := m.Get(int(k)); g != uint(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint(unsignedData[i])+3)
		}
	}
	m = NewMapIntToUint(13)
	for _, k := range signedData {
		m.Inc(int(k))
		if g := m.Get(int(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapIntToUintAdd(t *testing.T) {
	m := NewMapIntToUint(13)
	for i, k := range signedData {
		m.Put(int(k), uint(unsignedData[i]))
	}
	for i, k := range signedData {
		m.Add(int(k), 3)
		if g := m.Get(int(k)); g != uint(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint(unsignedData[i])+3)
		}
	}
	m = NewMapIntToUint(13)
	for _, k := range signedData {
		m.Add(int(k), 42)
		if g := m.Get(int(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapIntToUintModify(t *testing.T) {
	m := NewMapIntToUint(13)
	for i, k := range signedData {
		m.Put(int(k), uint(unsignedData[i]))
	}
	for _, k := range signedData {
		m.Modify(int(k), func(v *uint) {
			*v += 3
		})
	}
	for i, k := range signedData {
		if g := m.Get(int(k)); g != uint(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint(unsignedData[i])+3)
		}
	}
}

func TestMapIntToUintPut(t *testing.T) {
	m := NewMapIntToUint(13)
	for i, k := range signedData {
		m.Put(int(k), uint(unsignedData[i]))
	}
	for _, k := range signedData {
		m.Put(int(k), m.Get(int(k))+3)
	}
	for i, k := range signedData {
		if g := m.Get(int(k)); g != uint(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint(unsignedData[i])+3)
		}
	}
}

func TestMapIntToUintVisit(t *testing.T) {
	m := NewMapIntToUint(13)
	for i, k := range signedData {
		m.Put(int(k), uint(unsignedData[i]))
	}

	n := make(map[int]uint, len(signedData))

	for i, k := range signedData {
		n[int(k)] = uint(unsignedData[i])
	}

	m.Visit(func(k int, v uint) {
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

func TestMapIntToUint8Size(t *testing.T) {
	m := NewMapIntToUint8(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int(k), uint8(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapIntToUint8Contains(t *testing.T) {
	m := NewMapIntToUint8(13)
	for i, k := range signedData {
		m.Put(int(k), uint8(unsignedData[i]))
		if !m.Contains(int(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int(k))
		if m.Contains(int(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapIntToUint8Get(t *testing.T) {
	m := NewMapIntToUint8(13)
	for i, k := range signedData {
		m.Put(int(k), uint8(unsignedData[i]))
		if g := m.Get(int(k)); g != uint8(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapIntToUint8Clear(t *testing.T) {
	m := NewMapIntToUint8(13)
	for i, k := range signedData {
		m.Put(int(k), uint8(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int(k), uint8(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapIntToUint8Inc(t *testing.T) {
	m := NewMapIntToUint8(13)
	for i, k := range signedData {
		m.Put(int(k), uint8(unsignedData[i]))
	}
	for i, k := range signedData {
		m.Inc(int(k))
		m.Inc(int(k))
		m.Inc(int(k))
		if g := m.Get(int(k)); g != uint8(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint8(unsignedData[i])+3)
		}
	}
	m = NewMapIntToUint8(13)
	for _, k := range signedData {
		m.Inc(int(k))
		if g := m.Get(int(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapIntToUint8Add(t *testing.T) {
	m := NewMapIntToUint8(13)
	for i, k := range signedData {
		m.Put(int(k), uint8(unsignedData[i]))
	}
	for i, k := range signedData {
		m.Add(int(k), 3)
		if g := m.Get(int(k)); g != uint8(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint8(unsignedData[i])+3)
		}
	}
	m = NewMapIntToUint8(13)
	for _, k := range signedData {
		m.Add(int(k), 42)
		if g := m.Get(int(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapIntToUint8Modify(t *testing.T) {
	m := NewMapIntToUint8(13)
	for i, k := range signedData {
		m.Put(int(k), uint8(unsignedData[i]))
	}
	for _, k := range signedData {
		m.Modify(int(k), func(v *uint8) {
			*v += 3
		})
	}
	for i, k := range signedData {
		if g := m.Get(int(k)); g != uint8(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint8(unsignedData[i])+3)
		}
	}
}

func TestMapIntToUint8Put(t *testing.T) {
	m := NewMapIntToUint8(13)
	for i, k := range signedData {
		m.Put(int(k), uint8(unsignedData[i]))
	}
	for _, k := range signedData {
		m.Put(int(k), m.Get(int(k))+3)
	}
	for i, k := range signedData {
		if g := m.Get(int(k)); g != uint8(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint8(unsignedData[i])+3)
		}
	}
}

func TestMapIntToUint8Visit(t *testing.T) {
	m := NewMapIntToUint8(13)
	for i, k := range signedData {
		m.Put(int(k), uint8(unsignedData[i]))
	}

	n := make(map[int]uint8, len(signedData))

	for i, k := range signedData {
		n[int(k)] = uint8(unsignedData[i])
	}

	m.Visit(func(k int, v uint8) {
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

func TestMapIntToUint16Size(t *testing.T) {
	m := NewMapIntToUint16(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int(k), uint16(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapIntToUint16Contains(t *testing.T) {
	m := NewMapIntToUint16(13)
	for i, k := range signedData {
		m.Put(int(k), uint16(unsignedData[i]))
		if !m.Contains(int(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int(k))
		if m.Contains(int(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapIntToUint16Get(t *testing.T) {
	m := NewMapIntToUint16(13)
	for i, k := range signedData {
		m.Put(int(k), uint16(unsignedData[i]))
		if g := m.Get(int(k)); g != uint16(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapIntToUint16Clear(t *testing.T) {
	m := NewMapIntToUint16(13)
	for i, k := range signedData {
		m.Put(int(k), uint16(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int(k), uint16(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapIntToUint16Inc(t *testing.T) {
	m := NewMapIntToUint16(13)
	for i, k := range signedData {
		m.Put(int(k), uint16(unsignedData[i]))
	}
	for i, k := range signedData {
		m.Inc(int(k))
		m.Inc(int(k))
		m.Inc(int(k))
		if g := m.Get(int(k)); g != uint16(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint16(unsignedData[i])+3)
		}
	}
	m = NewMapIntToUint16(13)
	for _, k := range signedData {
		m.Inc(int(k))
		if g := m.Get(int(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapIntToUint16Add(t *testing.T) {
	m := NewMapIntToUint16(13)
	for i, k := range signedData {
		m.Put(int(k), uint16(unsignedData[i]))
	}
	for i, k := range signedData {
		m.Add(int(k), 3)
		if g := m.Get(int(k)); g != uint16(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint16(unsignedData[i])+3)
		}
	}
	m = NewMapIntToUint16(13)
	for _, k := range signedData {
		m.Add(int(k), 42)
		if g := m.Get(int(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapIntToUint16Modify(t *testing.T) {
	m := NewMapIntToUint16(13)
	for i, k := range signedData {
		m.Put(int(k), uint16(unsignedData[i]))
	}
	for _, k := range signedData {
		m.Modify(int(k), func(v *uint16) {
			*v += 3
		})
	}
	for i, k := range signedData {
		if g := m.Get(int(k)); g != uint16(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint16(unsignedData[i])+3)
		}
	}
}

func TestMapIntToUint16Put(t *testing.T) {
	m := NewMapIntToUint16(13)
	for i, k := range signedData {
		m.Put(int(k), uint16(unsignedData[i]))
	}
	for _, k := range signedData {
		m.Put(int(k), m.Get(int(k))+3)
	}
	for i, k := range signedData {
		if g := m.Get(int(k)); g != uint16(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint16(unsignedData[i])+3)
		}
	}
}

func TestMapIntToUint16Visit(t *testing.T) {
	m := NewMapIntToUint16(13)
	for i, k := range signedData {
		m.Put(int(k), uint16(unsignedData[i]))
	}

	n := make(map[int]uint16, len(signedData))

	for i, k := range signedData {
		n[int(k)] = uint16(unsignedData[i])
	}

	m.Visit(func(k int, v uint16) {
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

func TestMapIntToUint32Size(t *testing.T) {
	m := NewMapIntToUint32(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int(k), uint32(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapIntToUint32Contains(t *testing.T) {
	m := NewMapIntToUint32(13)
	for i, k := range signedData {
		m.Put(int(k), uint32(unsignedData[i]))
		if !m.Contains(int(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int(k))
		if m.Contains(int(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapIntToUint32Get(t *testing.T) {
	m := NewMapIntToUint32(13)
	for i, k := range signedData {
		m.Put(int(k), uint32(unsignedData[i]))
		if g := m.Get(int(k)); g != uint32(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapIntToUint32Clear(t *testing.T) {
	m := NewMapIntToUint32(13)
	for i, k := range signedData {
		m.Put(int(k), uint32(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int(k), uint32(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapIntToUint32Inc(t *testing.T) {
	m := NewMapIntToUint32(13)
	for i, k := range signedData {
		m.Put(int(k), uint32(unsignedData[i]))
	}
	for i, k := range signedData {
		m.Inc(int(k))
		m.Inc(int(k))
		m.Inc(int(k))
		if g := m.Get(int(k)); g != uint32(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint32(unsignedData[i])+3)
		}
	}
	m = NewMapIntToUint32(13)
	for _, k := range signedData {
		m.Inc(int(k))
		if g := m.Get(int(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapIntToUint32Add(t *testing.T) {
	m := NewMapIntToUint32(13)
	for i, k := range signedData {
		m.Put(int(k), uint32(unsignedData[i]))
	}
	for i, k := range signedData {
		m.Add(int(k), 3)
		if g := m.Get(int(k)); g != uint32(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint32(unsignedData[i])+3)
		}
	}
	m = NewMapIntToUint32(13)
	for _, k := range signedData {
		m.Add(int(k), 42)
		if g := m.Get(int(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapIntToUint32Modify(t *testing.T) {
	m := NewMapIntToUint32(13)
	for i, k := range signedData {
		m.Put(int(k), uint32(unsignedData[i]))
	}
	for _, k := range signedData {
		m.Modify(int(k), func(v *uint32) {
			*v += 3
		})
	}
	for i, k := range signedData {
		if g := m.Get(int(k)); g != uint32(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint32(unsignedData[i])+3)
		}
	}
}

func TestMapIntToUint32Put(t *testing.T) {
	m := NewMapIntToUint32(13)
	for i, k := range signedData {
		m.Put(int(k), uint32(unsignedData[i]))
	}
	for _, k := range signedData {
		m.Put(int(k), m.Get(int(k))+3)
	}
	for i, k := range signedData {
		if g := m.Get(int(k)); g != uint32(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint32(unsignedData[i])+3)
		}
	}
}

func TestMapIntToUint32Visit(t *testing.T) {
	m := NewMapIntToUint32(13)
	for i, k := range signedData {
		m.Put(int(k), uint32(unsignedData[i]))
	}

	n := make(map[int]uint32, len(signedData))

	for i, k := range signedData {
		n[int(k)] = uint32(unsignedData[i])
	}

	m.Visit(func(k int, v uint32) {
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

func TestMapIntToUint64Size(t *testing.T) {
	m := NewMapIntToUint64(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int(k), uint64(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapIntToUint64Contains(t *testing.T) {
	m := NewMapIntToUint64(13)
	for i, k := range signedData {
		m.Put(int(k), uint64(unsignedData[i]))
		if !m.Contains(int(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int(k))
		if m.Contains(int(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapIntToUint64Get(t *testing.T) {
	m := NewMapIntToUint64(13)
	for i, k := range signedData {
		m.Put(int(k), uint64(unsignedData[i]))
		if g := m.Get(int(k)); g != uint64(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapIntToUint64Clear(t *testing.T) {
	m := NewMapIntToUint64(13)
	for i, k := range signedData {
		m.Put(int(k), uint64(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int(k), uint64(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapIntToUint64Inc(t *testing.T) {
	m := NewMapIntToUint64(13)
	for i, k := range signedData {
		m.Put(int(k), uint64(unsignedData[i]))
	}
	for i, k := range signedData {
		m.Inc(int(k))
		m.Inc(int(k))
		m.Inc(int(k))
		if g := m.Get(int(k)); g != uint64(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint64(unsignedData[i])+3)
		}
	}
	m = NewMapIntToUint64(13)
	for _, k := range signedData {
		m.Inc(int(k))
		if g := m.Get(int(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapIntToUint64Add(t *testing.T) {
	m := NewMapIntToUint64(13)
	for i, k := range signedData {
		m.Put(int(k), uint64(unsignedData[i]))
	}
	for i, k := range signedData {
		m.Add(int(k), 3)
		if g := m.Get(int(k)); g != uint64(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint64(unsignedData[i])+3)
		}
	}
	m = NewMapIntToUint64(13)
	for _, k := range signedData {
		m.Add(int(k), 42)
		if g := m.Get(int(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapIntToUint64Modify(t *testing.T) {
	m := NewMapIntToUint64(13)
	for i, k := range signedData {
		m.Put(int(k), uint64(unsignedData[i]))
	}
	for _, k := range signedData {
		m.Modify(int(k), func(v *uint64) {
			*v += 3
		})
	}
	for i, k := range signedData {
		if g := m.Get(int(k)); g != uint64(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint64(unsignedData[i])+3)
		}
	}
}

func TestMapIntToUint64Put(t *testing.T) {
	m := NewMapIntToUint64(13)
	for i, k := range signedData {
		m.Put(int(k), uint64(unsignedData[i]))
	}
	for _, k := range signedData {
		m.Put(int(k), m.Get(int(k))+3)
	}
	for i, k := range signedData {
		if g := m.Get(int(k)); g != uint64(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint64(unsignedData[i])+3)
		}
	}
}

func TestMapIntToUint64Visit(t *testing.T) {
	m := NewMapIntToUint64(13)
	for i, k := range signedData {
		m.Put(int(k), uint64(unsignedData[i]))
	}

	n := make(map[int]uint64, len(signedData))

	for i, k := range signedData {
		n[int(k)] = uint64(unsignedData[i])
	}

	m.Visit(func(k int, v uint64) {
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
