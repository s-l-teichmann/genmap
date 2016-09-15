// THIS IS A MACHINE GENERATED FILE!
// BE CAREFUL WITH EDITING BY HAND.
//
// This is Free Software covered by the terms of the MIT license.
// See LICENSE file for details.
// (c) 2016 by Sascha L. Teichmann.
// See the full list of contributors in the CONTRIBUTORS file.

package genmap

import "testing"

func TestMapInt8ToIntSize(t *testing.T) {
	m := NewMapInt8ToInt(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int8(k), int(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int8(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt8ToIntContains(t *testing.T) {
	m := NewMapInt8ToInt(13)
	for i, k := range signedData {
		m.Put(int8(k), int(signedData[i]))
		if !m.Contains(int8(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int8(k))
		if m.Contains(int8(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt8ToIntGet(t *testing.T) {
	m := NewMapInt8ToInt(13)
	for i, k := range signedData {
		m.Put(int8(k), int(signedData[i]))
		if g := m.Get(int8(k)); g != int(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapInt8ToIntClear(t *testing.T) {
	m := NewMapInt8ToInt(13)
	for i, k := range signedData {
		m.Put(int8(k), int(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int8(k), int(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapInt8ToIntInc(t *testing.T) {
	m := NewMapInt8ToInt(13)
	for i, k := range signedData {
		m.Put(int8(k), int(signedData[i]))
	}
	for i, k := range signedData {
		m.Inc(int8(k))
		m.Inc(int8(k))
		m.Inc(int8(k))
		if g := m.Get(int8(k)); g != int(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int(signedData[i])+3)
		}
	}
	m = NewMapInt8ToInt(13)
	for _, k := range signedData {
		m.Inc(int8(k))
		if g := m.Get(int8(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapInt8ToIntAdd(t *testing.T) {
	m := NewMapInt8ToInt(13)
	for i, k := range signedData {
		m.Put(int8(k), int(signedData[i]))
	}
	for i, k := range signedData {
		m.Add(int8(k), 3)
		if g := m.Get(int8(k)); g != int(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int(signedData[i])+3)
		}
	}
	m = NewMapInt8ToInt(13)
	for _, k := range signedData {
		m.Add(int8(k), 42)
		if g := m.Get(int8(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapInt8ToIntPut(t *testing.T) {
	m := NewMapInt8ToInt(13)
	for i, k := range signedData {
		m.Put(int8(k), int(signedData[i]))
	}
	for _, k := range signedData {
		m.Put(int8(k), m.Get(int8(k))+3)
	}
	for i, k := range signedData {
		if g := m.Get(int8(k)); g != int(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int(signedData[i])+3)
		}
	}
}

func TestMapInt8ToIntVisit(t *testing.T) {
	m := NewMapInt8ToInt(13)
	for i, k := range signedData {
		m.Put(int8(k), int(signedData[i]))
	}

	n := make(map[int8]int, len(signedData))

	for i, k := range signedData {
		n[int8(k)] = int(signedData[i])
	}

	m.Visit(func(k int8, v int) {
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

func TestMapInt8ToInt8Size(t *testing.T) {
	m := NewMapInt8ToInt8(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int8(k), int8(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int8(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt8ToInt8Contains(t *testing.T) {
	m := NewMapInt8ToInt8(13)
	for i, k := range signedData {
		m.Put(int8(k), int8(signedData[i]))
		if !m.Contains(int8(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int8(k))
		if m.Contains(int8(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt8ToInt8Get(t *testing.T) {
	m := NewMapInt8ToInt8(13)
	for i, k := range signedData {
		m.Put(int8(k), int8(signedData[i]))
		if g := m.Get(int8(k)); g != int8(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapInt8ToInt8Clear(t *testing.T) {
	m := NewMapInt8ToInt8(13)
	for i, k := range signedData {
		m.Put(int8(k), int8(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int8(k), int8(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapInt8ToInt8Inc(t *testing.T) {
	m := NewMapInt8ToInt8(13)
	for i, k := range signedData {
		m.Put(int8(k), int8(signedData[i]))
	}
	for i, k := range signedData {
		m.Inc(int8(k))
		m.Inc(int8(k))
		m.Inc(int8(k))
		if g := m.Get(int8(k)); g != int8(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int8(signedData[i])+3)
		}
	}
	m = NewMapInt8ToInt8(13)
	for _, k := range signedData {
		m.Inc(int8(k))
		if g := m.Get(int8(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapInt8ToInt8Add(t *testing.T) {
	m := NewMapInt8ToInt8(13)
	for i, k := range signedData {
		m.Put(int8(k), int8(signedData[i]))
	}
	for i, k := range signedData {
		m.Add(int8(k), 3)
		if g := m.Get(int8(k)); g != int8(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int8(signedData[i])+3)
		}
	}
	m = NewMapInt8ToInt8(13)
	for _, k := range signedData {
		m.Add(int8(k), 42)
		if g := m.Get(int8(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapInt8ToInt8Put(t *testing.T) {
	m := NewMapInt8ToInt8(13)
	for i, k := range signedData {
		m.Put(int8(k), int8(signedData[i]))
	}
	for _, k := range signedData {
		m.Put(int8(k), m.Get(int8(k))+3)
	}
	for i, k := range signedData {
		if g := m.Get(int8(k)); g != int8(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int8(signedData[i])+3)
		}
	}
}

func TestMapInt8ToInt8Visit(t *testing.T) {
	m := NewMapInt8ToInt8(13)
	for i, k := range signedData {
		m.Put(int8(k), int8(signedData[i]))
	}

	n := make(map[int8]int8, len(signedData))

	for i, k := range signedData {
		n[int8(k)] = int8(signedData[i])
	}

	m.Visit(func(k int8, v int8) {
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

func TestMapInt8ToInt16Size(t *testing.T) {
	m := NewMapInt8ToInt16(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int8(k), int16(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int8(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt8ToInt16Contains(t *testing.T) {
	m := NewMapInt8ToInt16(13)
	for i, k := range signedData {
		m.Put(int8(k), int16(signedData[i]))
		if !m.Contains(int8(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int8(k))
		if m.Contains(int8(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt8ToInt16Get(t *testing.T) {
	m := NewMapInt8ToInt16(13)
	for i, k := range signedData {
		m.Put(int8(k), int16(signedData[i]))
		if g := m.Get(int8(k)); g != int16(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapInt8ToInt16Clear(t *testing.T) {
	m := NewMapInt8ToInt16(13)
	for i, k := range signedData {
		m.Put(int8(k), int16(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int8(k), int16(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapInt8ToInt16Inc(t *testing.T) {
	m := NewMapInt8ToInt16(13)
	for i, k := range signedData {
		m.Put(int8(k), int16(signedData[i]))
	}
	for i, k := range signedData {
		m.Inc(int8(k))
		m.Inc(int8(k))
		m.Inc(int8(k))
		if g := m.Get(int8(k)); g != int16(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int16(signedData[i])+3)
		}
	}
	m = NewMapInt8ToInt16(13)
	for _, k := range signedData {
		m.Inc(int8(k))
		if g := m.Get(int8(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapInt8ToInt16Add(t *testing.T) {
	m := NewMapInt8ToInt16(13)
	for i, k := range signedData {
		m.Put(int8(k), int16(signedData[i]))
	}
	for i, k := range signedData {
		m.Add(int8(k), 3)
		if g := m.Get(int8(k)); g != int16(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int16(signedData[i])+3)
		}
	}
	m = NewMapInt8ToInt16(13)
	for _, k := range signedData {
		m.Add(int8(k), 42)
		if g := m.Get(int8(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapInt8ToInt16Put(t *testing.T) {
	m := NewMapInt8ToInt16(13)
	for i, k := range signedData {
		m.Put(int8(k), int16(signedData[i]))
	}
	for _, k := range signedData {
		m.Put(int8(k), m.Get(int8(k))+3)
	}
	for i, k := range signedData {
		if g := m.Get(int8(k)); g != int16(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int16(signedData[i])+3)
		}
	}
}

func TestMapInt8ToInt16Visit(t *testing.T) {
	m := NewMapInt8ToInt16(13)
	for i, k := range signedData {
		m.Put(int8(k), int16(signedData[i]))
	}

	n := make(map[int8]int16, len(signedData))

	for i, k := range signedData {
		n[int8(k)] = int16(signedData[i])
	}

	m.Visit(func(k int8, v int16) {
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

func TestMapInt8ToInt32Size(t *testing.T) {
	m := NewMapInt8ToInt32(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int8(k), int32(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int8(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt8ToInt32Contains(t *testing.T) {
	m := NewMapInt8ToInt32(13)
	for i, k := range signedData {
		m.Put(int8(k), int32(signedData[i]))
		if !m.Contains(int8(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int8(k))
		if m.Contains(int8(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt8ToInt32Get(t *testing.T) {
	m := NewMapInt8ToInt32(13)
	for i, k := range signedData {
		m.Put(int8(k), int32(signedData[i]))
		if g := m.Get(int8(k)); g != int32(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapInt8ToInt32Clear(t *testing.T) {
	m := NewMapInt8ToInt32(13)
	for i, k := range signedData {
		m.Put(int8(k), int32(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int8(k), int32(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapInt8ToInt32Inc(t *testing.T) {
	m := NewMapInt8ToInt32(13)
	for i, k := range signedData {
		m.Put(int8(k), int32(signedData[i]))
	}
	for i, k := range signedData {
		m.Inc(int8(k))
		m.Inc(int8(k))
		m.Inc(int8(k))
		if g := m.Get(int8(k)); g != int32(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int32(signedData[i])+3)
		}
	}
	m = NewMapInt8ToInt32(13)
	for _, k := range signedData {
		m.Inc(int8(k))
		if g := m.Get(int8(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapInt8ToInt32Add(t *testing.T) {
	m := NewMapInt8ToInt32(13)
	for i, k := range signedData {
		m.Put(int8(k), int32(signedData[i]))
	}
	for i, k := range signedData {
		m.Add(int8(k), 3)
		if g := m.Get(int8(k)); g != int32(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int32(signedData[i])+3)
		}
	}
	m = NewMapInt8ToInt32(13)
	for _, k := range signedData {
		m.Add(int8(k), 42)
		if g := m.Get(int8(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapInt8ToInt32Put(t *testing.T) {
	m := NewMapInt8ToInt32(13)
	for i, k := range signedData {
		m.Put(int8(k), int32(signedData[i]))
	}
	for _, k := range signedData {
		m.Put(int8(k), m.Get(int8(k))+3)
	}
	for i, k := range signedData {
		if g := m.Get(int8(k)); g != int32(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int32(signedData[i])+3)
		}
	}
}

func TestMapInt8ToInt32Visit(t *testing.T) {
	m := NewMapInt8ToInt32(13)
	for i, k := range signedData {
		m.Put(int8(k), int32(signedData[i]))
	}

	n := make(map[int8]int32, len(signedData))

	for i, k := range signedData {
		n[int8(k)] = int32(signedData[i])
	}

	m.Visit(func(k int8, v int32) {
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

func TestMapInt8ToInt64Size(t *testing.T) {
	m := NewMapInt8ToInt64(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int8(k), int64(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int8(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt8ToInt64Contains(t *testing.T) {
	m := NewMapInt8ToInt64(13)
	for i, k := range signedData {
		m.Put(int8(k), int64(signedData[i]))
		if !m.Contains(int8(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int8(k))
		if m.Contains(int8(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt8ToInt64Get(t *testing.T) {
	m := NewMapInt8ToInt64(13)
	for i, k := range signedData {
		m.Put(int8(k), int64(signedData[i]))
		if g := m.Get(int8(k)); g != int64(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapInt8ToInt64Clear(t *testing.T) {
	m := NewMapInt8ToInt64(13)
	for i, k := range signedData {
		m.Put(int8(k), int64(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int8(k), int64(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapInt8ToInt64Inc(t *testing.T) {
	m := NewMapInt8ToInt64(13)
	for i, k := range signedData {
		m.Put(int8(k), int64(signedData[i]))
	}
	for i, k := range signedData {
		m.Inc(int8(k))
		m.Inc(int8(k))
		m.Inc(int8(k))
		if g := m.Get(int8(k)); g != int64(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int64(signedData[i])+3)
		}
	}
	m = NewMapInt8ToInt64(13)
	for _, k := range signedData {
		m.Inc(int8(k))
		if g := m.Get(int8(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapInt8ToInt64Add(t *testing.T) {
	m := NewMapInt8ToInt64(13)
	for i, k := range signedData {
		m.Put(int8(k), int64(signedData[i]))
	}
	for i, k := range signedData {
		m.Add(int8(k), 3)
		if g := m.Get(int8(k)); g != int64(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int64(signedData[i])+3)
		}
	}
	m = NewMapInt8ToInt64(13)
	for _, k := range signedData {
		m.Add(int8(k), 42)
		if g := m.Get(int8(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapInt8ToInt64Put(t *testing.T) {
	m := NewMapInt8ToInt64(13)
	for i, k := range signedData {
		m.Put(int8(k), int64(signedData[i]))
	}
	for _, k := range signedData {
		m.Put(int8(k), m.Get(int8(k))+3)
	}
	for i, k := range signedData {
		if g := m.Get(int8(k)); g != int64(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int64(signedData[i])+3)
		}
	}
}

func TestMapInt8ToInt64Visit(t *testing.T) {
	m := NewMapInt8ToInt64(13)
	for i, k := range signedData {
		m.Put(int8(k), int64(signedData[i]))
	}

	n := make(map[int8]int64, len(signedData))

	for i, k := range signedData {
		n[int8(k)] = int64(signedData[i])
	}

	m.Visit(func(k int8, v int64) {
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

func TestMapInt8ToUintSize(t *testing.T) {
	m := NewMapInt8ToUint(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int8(k), uint(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int8(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt8ToUintContains(t *testing.T) {
	m := NewMapInt8ToUint(13)
	for i, k := range signedData {
		m.Put(int8(k), uint(unsignedData[i]))
		if !m.Contains(int8(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int8(k))
		if m.Contains(int8(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt8ToUintGet(t *testing.T) {
	m := NewMapInt8ToUint(13)
	for i, k := range signedData {
		m.Put(int8(k), uint(unsignedData[i]))
		if g := m.Get(int8(k)); g != uint(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapInt8ToUintClear(t *testing.T) {
	m := NewMapInt8ToUint(13)
	for i, k := range signedData {
		m.Put(int8(k), uint(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int8(k), uint(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapInt8ToUintInc(t *testing.T) {
	m := NewMapInt8ToUint(13)
	for i, k := range signedData {
		m.Put(int8(k), uint(unsignedData[i]))
	}
	for i, k := range signedData {
		m.Inc(int8(k))
		m.Inc(int8(k))
		m.Inc(int8(k))
		if g := m.Get(int8(k)); g != uint(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint(unsignedData[i])+3)
		}
	}
	m = NewMapInt8ToUint(13)
	for _, k := range signedData {
		m.Inc(int8(k))
		if g := m.Get(int8(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapInt8ToUintAdd(t *testing.T) {
	m := NewMapInt8ToUint(13)
	for i, k := range signedData {
		m.Put(int8(k), uint(unsignedData[i]))
	}
	for i, k := range signedData {
		m.Add(int8(k), 3)
		if g := m.Get(int8(k)); g != uint(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint(unsignedData[i])+3)
		}
	}
	m = NewMapInt8ToUint(13)
	for _, k := range signedData {
		m.Add(int8(k), 42)
		if g := m.Get(int8(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapInt8ToUintPut(t *testing.T) {
	m := NewMapInt8ToUint(13)
	for i, k := range signedData {
		m.Put(int8(k), uint(unsignedData[i]))
	}
	for _, k := range signedData {
		m.Put(int8(k), m.Get(int8(k))+3)
	}
	for i, k := range signedData {
		if g := m.Get(int8(k)); g != uint(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint(unsignedData[i])+3)
		}
	}
}

func TestMapInt8ToUintVisit(t *testing.T) {
	m := NewMapInt8ToUint(13)
	for i, k := range signedData {
		m.Put(int8(k), uint(unsignedData[i]))
	}

	n := make(map[int8]uint, len(signedData))

	for i, k := range signedData {
		n[int8(k)] = uint(unsignedData[i])
	}

	m.Visit(func(k int8, v uint) {
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

func TestMapInt8ToUint8Size(t *testing.T) {
	m := NewMapInt8ToUint8(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int8(k), uint8(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int8(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt8ToUint8Contains(t *testing.T) {
	m := NewMapInt8ToUint8(13)
	for i, k := range signedData {
		m.Put(int8(k), uint8(unsignedData[i]))
		if !m.Contains(int8(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int8(k))
		if m.Contains(int8(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt8ToUint8Get(t *testing.T) {
	m := NewMapInt8ToUint8(13)
	for i, k := range signedData {
		m.Put(int8(k), uint8(unsignedData[i]))
		if g := m.Get(int8(k)); g != uint8(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapInt8ToUint8Clear(t *testing.T) {
	m := NewMapInt8ToUint8(13)
	for i, k := range signedData {
		m.Put(int8(k), uint8(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int8(k), uint8(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapInt8ToUint8Inc(t *testing.T) {
	m := NewMapInt8ToUint8(13)
	for i, k := range signedData {
		m.Put(int8(k), uint8(unsignedData[i]))
	}
	for i, k := range signedData {
		m.Inc(int8(k))
		m.Inc(int8(k))
		m.Inc(int8(k))
		if g := m.Get(int8(k)); g != uint8(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint8(unsignedData[i])+3)
		}
	}
	m = NewMapInt8ToUint8(13)
	for _, k := range signedData {
		m.Inc(int8(k))
		if g := m.Get(int8(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapInt8ToUint8Add(t *testing.T) {
	m := NewMapInt8ToUint8(13)
	for i, k := range signedData {
		m.Put(int8(k), uint8(unsignedData[i]))
	}
	for i, k := range signedData {
		m.Add(int8(k), 3)
		if g := m.Get(int8(k)); g != uint8(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint8(unsignedData[i])+3)
		}
	}
	m = NewMapInt8ToUint8(13)
	for _, k := range signedData {
		m.Add(int8(k), 42)
		if g := m.Get(int8(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapInt8ToUint8Put(t *testing.T) {
	m := NewMapInt8ToUint8(13)
	for i, k := range signedData {
		m.Put(int8(k), uint8(unsignedData[i]))
	}
	for _, k := range signedData {
		m.Put(int8(k), m.Get(int8(k))+3)
	}
	for i, k := range signedData {
		if g := m.Get(int8(k)); g != uint8(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint8(unsignedData[i])+3)
		}
	}
}

func TestMapInt8ToUint8Visit(t *testing.T) {
	m := NewMapInt8ToUint8(13)
	for i, k := range signedData {
		m.Put(int8(k), uint8(unsignedData[i]))
	}

	n := make(map[int8]uint8, len(signedData))

	for i, k := range signedData {
		n[int8(k)] = uint8(unsignedData[i])
	}

	m.Visit(func(k int8, v uint8) {
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

func TestMapInt8ToUint16Size(t *testing.T) {
	m := NewMapInt8ToUint16(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int8(k), uint16(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int8(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt8ToUint16Contains(t *testing.T) {
	m := NewMapInt8ToUint16(13)
	for i, k := range signedData {
		m.Put(int8(k), uint16(unsignedData[i]))
		if !m.Contains(int8(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int8(k))
		if m.Contains(int8(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt8ToUint16Get(t *testing.T) {
	m := NewMapInt8ToUint16(13)
	for i, k := range signedData {
		m.Put(int8(k), uint16(unsignedData[i]))
		if g := m.Get(int8(k)); g != uint16(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapInt8ToUint16Clear(t *testing.T) {
	m := NewMapInt8ToUint16(13)
	for i, k := range signedData {
		m.Put(int8(k), uint16(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int8(k), uint16(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapInt8ToUint16Inc(t *testing.T) {
	m := NewMapInt8ToUint16(13)
	for i, k := range signedData {
		m.Put(int8(k), uint16(unsignedData[i]))
	}
	for i, k := range signedData {
		m.Inc(int8(k))
		m.Inc(int8(k))
		m.Inc(int8(k))
		if g := m.Get(int8(k)); g != uint16(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint16(unsignedData[i])+3)
		}
	}
	m = NewMapInt8ToUint16(13)
	for _, k := range signedData {
		m.Inc(int8(k))
		if g := m.Get(int8(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapInt8ToUint16Add(t *testing.T) {
	m := NewMapInt8ToUint16(13)
	for i, k := range signedData {
		m.Put(int8(k), uint16(unsignedData[i]))
	}
	for i, k := range signedData {
		m.Add(int8(k), 3)
		if g := m.Get(int8(k)); g != uint16(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint16(unsignedData[i])+3)
		}
	}
	m = NewMapInt8ToUint16(13)
	for _, k := range signedData {
		m.Add(int8(k), 42)
		if g := m.Get(int8(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapInt8ToUint16Put(t *testing.T) {
	m := NewMapInt8ToUint16(13)
	for i, k := range signedData {
		m.Put(int8(k), uint16(unsignedData[i]))
	}
	for _, k := range signedData {
		m.Put(int8(k), m.Get(int8(k))+3)
	}
	for i, k := range signedData {
		if g := m.Get(int8(k)); g != uint16(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint16(unsignedData[i])+3)
		}
	}
}

func TestMapInt8ToUint16Visit(t *testing.T) {
	m := NewMapInt8ToUint16(13)
	for i, k := range signedData {
		m.Put(int8(k), uint16(unsignedData[i]))
	}

	n := make(map[int8]uint16, len(signedData))

	for i, k := range signedData {
		n[int8(k)] = uint16(unsignedData[i])
	}

	m.Visit(func(k int8, v uint16) {
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

func TestMapInt8ToUint32Size(t *testing.T) {
	m := NewMapInt8ToUint32(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int8(k), uint32(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int8(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt8ToUint32Contains(t *testing.T) {
	m := NewMapInt8ToUint32(13)
	for i, k := range signedData {
		m.Put(int8(k), uint32(unsignedData[i]))
		if !m.Contains(int8(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int8(k))
		if m.Contains(int8(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt8ToUint32Get(t *testing.T) {
	m := NewMapInt8ToUint32(13)
	for i, k := range signedData {
		m.Put(int8(k), uint32(unsignedData[i]))
		if g := m.Get(int8(k)); g != uint32(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapInt8ToUint32Clear(t *testing.T) {
	m := NewMapInt8ToUint32(13)
	for i, k := range signedData {
		m.Put(int8(k), uint32(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int8(k), uint32(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapInt8ToUint32Inc(t *testing.T) {
	m := NewMapInt8ToUint32(13)
	for i, k := range signedData {
		m.Put(int8(k), uint32(unsignedData[i]))
	}
	for i, k := range signedData {
		m.Inc(int8(k))
		m.Inc(int8(k))
		m.Inc(int8(k))
		if g := m.Get(int8(k)); g != uint32(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint32(unsignedData[i])+3)
		}
	}
	m = NewMapInt8ToUint32(13)
	for _, k := range signedData {
		m.Inc(int8(k))
		if g := m.Get(int8(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapInt8ToUint32Add(t *testing.T) {
	m := NewMapInt8ToUint32(13)
	for i, k := range signedData {
		m.Put(int8(k), uint32(unsignedData[i]))
	}
	for i, k := range signedData {
		m.Add(int8(k), 3)
		if g := m.Get(int8(k)); g != uint32(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint32(unsignedData[i])+3)
		}
	}
	m = NewMapInt8ToUint32(13)
	for _, k := range signedData {
		m.Add(int8(k), 42)
		if g := m.Get(int8(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapInt8ToUint32Put(t *testing.T) {
	m := NewMapInt8ToUint32(13)
	for i, k := range signedData {
		m.Put(int8(k), uint32(unsignedData[i]))
	}
	for _, k := range signedData {
		m.Put(int8(k), m.Get(int8(k))+3)
	}
	for i, k := range signedData {
		if g := m.Get(int8(k)); g != uint32(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint32(unsignedData[i])+3)
		}
	}
}

func TestMapInt8ToUint32Visit(t *testing.T) {
	m := NewMapInt8ToUint32(13)
	for i, k := range signedData {
		m.Put(int8(k), uint32(unsignedData[i]))
	}

	n := make(map[int8]uint32, len(signedData))

	for i, k := range signedData {
		n[int8(k)] = uint32(unsignedData[i])
	}

	m.Visit(func(k int8, v uint32) {
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

func TestMapInt8ToUint64Size(t *testing.T) {
	m := NewMapInt8ToUint64(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int8(k), uint64(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int8(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt8ToUint64Contains(t *testing.T) {
	m := NewMapInt8ToUint64(13)
	for i, k := range signedData {
		m.Put(int8(k), uint64(unsignedData[i]))
		if !m.Contains(int8(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int8(k))
		if m.Contains(int8(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt8ToUint64Get(t *testing.T) {
	m := NewMapInt8ToUint64(13)
	for i, k := range signedData {
		m.Put(int8(k), uint64(unsignedData[i]))
		if g := m.Get(int8(k)); g != uint64(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapInt8ToUint64Clear(t *testing.T) {
	m := NewMapInt8ToUint64(13)
	for i, k := range signedData {
		m.Put(int8(k), uint64(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int8(k), uint64(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapInt8ToUint64Inc(t *testing.T) {
	m := NewMapInt8ToUint64(13)
	for i, k := range signedData {
		m.Put(int8(k), uint64(unsignedData[i]))
	}
	for i, k := range signedData {
		m.Inc(int8(k))
		m.Inc(int8(k))
		m.Inc(int8(k))
		if g := m.Get(int8(k)); g != uint64(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint64(unsignedData[i])+3)
		}
	}
	m = NewMapInt8ToUint64(13)
	for _, k := range signedData {
		m.Inc(int8(k))
		if g := m.Get(int8(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapInt8ToUint64Add(t *testing.T) {
	m := NewMapInt8ToUint64(13)
	for i, k := range signedData {
		m.Put(int8(k), uint64(unsignedData[i]))
	}
	for i, k := range signedData {
		m.Add(int8(k), 3)
		if g := m.Get(int8(k)); g != uint64(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint64(unsignedData[i])+3)
		}
	}
	m = NewMapInt8ToUint64(13)
	for _, k := range signedData {
		m.Add(int8(k), 42)
		if g := m.Get(int8(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapInt8ToUint64Put(t *testing.T) {
	m := NewMapInt8ToUint64(13)
	for i, k := range signedData {
		m.Put(int8(k), uint64(unsignedData[i]))
	}
	for _, k := range signedData {
		m.Put(int8(k), m.Get(int8(k))+3)
	}
	for i, k := range signedData {
		if g := m.Get(int8(k)); g != uint64(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint64(unsignedData[i])+3)
		}
	}
}

func TestMapInt8ToUint64Visit(t *testing.T) {
	m := NewMapInt8ToUint64(13)
	for i, k := range signedData {
		m.Put(int8(k), uint64(unsignedData[i]))
	}

	n := make(map[int8]uint64, len(signedData))

	for i, k := range signedData {
		n[int8(k)] = uint64(unsignedData[i])
	}

	m.Visit(func(k int8, v uint64) {
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
