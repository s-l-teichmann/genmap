// THIS IS A MACHINE GENERATED FILE!
// BE CAREFUL WITH EDITING BY HAND.
//
// This is Free Software covered by the terms of the MIT license.
// See LICENSE file for details.
// (c) 2016 by Sascha L. Teichmann.
// See the full list of contributors in the CONTRIBUTORS file.

package genmap

import "testing"

func TestMapInt16ToIntSize(t *testing.T) {
	m := NewMapInt16ToInt(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int16(k), int(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int16(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt16ToIntContains(t *testing.T) {
	m := NewMapInt16ToInt(13)
	for i, k := range signedData {
		m.Put(int16(k), int(signedData[i]))
		if !m.Contains(int16(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int16(k))
		if m.Contains(int16(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt16ToIntGet(t *testing.T) {
	m := NewMapInt16ToInt(13)
	for i, k := range signedData {
		m.Put(int16(k), int(signedData[i]))
		if g := m.Get(int16(k)); g != int(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapInt16ToIntClear(t *testing.T) {
	m := NewMapInt16ToInt(13)
	for i, k := range signedData {
		m.Put(int16(k), int(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int16(k), int(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapInt16ToIntInc(t *testing.T) {
	m := NewMapInt16ToInt(13)
	for i, k := range signedData {
		m.Put(int16(k), int(signedData[i]))
	}
	for i, k := range signedData {
		m.Inc(int16(k))
		m.Inc(int16(k))
		m.Inc(int16(k))
		if g := m.Get(int16(k)); g != int(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int(signedData[i])+3)
		}
	}
	m = NewMapInt16ToInt(13)
	for _, k := range signedData {
		m.Inc(int16(k))
		if g := m.Get(int16(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapInt16ToIntAdd(t *testing.T) {
	m := NewMapInt16ToInt(13)
	for i, k := range signedData {
		m.Put(int16(k), int(signedData[i]))
	}
	for i, k := range signedData {
		m.Add(int16(k), 3)
		if g := m.Get(int16(k)); g != int(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int(signedData[i])+3)
		}
	}
	m = NewMapInt16ToInt(13)
	for _, k := range signedData {
		m.Add(int16(k), 42)
		if g := m.Get(int16(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapInt16ToIntPut(t *testing.T) {
	m := NewMapInt16ToInt(13)
	for i, k := range signedData {
		m.Put(int16(k), int(signedData[i]))
	}
	for _, k := range signedData {
		m.Put(int16(k), m.Get(int16(k))+3)
	}
	for i, k := range signedData {
		if g := m.Get(int16(k)); g != int(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int(signedData[i])+3)
		}
	}
}

func TestMapInt16ToIntVisit(t *testing.T) {
	m := NewMapInt16ToInt(13)
	for i, k := range signedData {
		m.Put(int16(k), int(signedData[i]))
	}

	n := make(map[int16]int, len(signedData))

	for i, k := range signedData {
		n[int16(k)] = int(signedData[i])
	}

	m.Visit(func(k int16, v int) {
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

func TestMapInt16ToInt8Size(t *testing.T) {
	m := NewMapInt16ToInt8(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int16(k), int8(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int16(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt16ToInt8Contains(t *testing.T) {
	m := NewMapInt16ToInt8(13)
	for i, k := range signedData {
		m.Put(int16(k), int8(signedData[i]))
		if !m.Contains(int16(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int16(k))
		if m.Contains(int16(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt16ToInt8Get(t *testing.T) {
	m := NewMapInt16ToInt8(13)
	for i, k := range signedData {
		m.Put(int16(k), int8(signedData[i]))
		if g := m.Get(int16(k)); g != int8(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapInt16ToInt8Clear(t *testing.T) {
	m := NewMapInt16ToInt8(13)
	for i, k := range signedData {
		m.Put(int16(k), int8(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int16(k), int8(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapInt16ToInt8Inc(t *testing.T) {
	m := NewMapInt16ToInt8(13)
	for i, k := range signedData {
		m.Put(int16(k), int8(signedData[i]))
	}
	for i, k := range signedData {
		m.Inc(int16(k))
		m.Inc(int16(k))
		m.Inc(int16(k))
		if g := m.Get(int16(k)); g != int8(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int8(signedData[i])+3)
		}
	}
	m = NewMapInt16ToInt8(13)
	for _, k := range signedData {
		m.Inc(int16(k))
		if g := m.Get(int16(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapInt16ToInt8Add(t *testing.T) {
	m := NewMapInt16ToInt8(13)
	for i, k := range signedData {
		m.Put(int16(k), int8(signedData[i]))
	}
	for i, k := range signedData {
		m.Add(int16(k), 3)
		if g := m.Get(int16(k)); g != int8(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int8(signedData[i])+3)
		}
	}
	m = NewMapInt16ToInt8(13)
	for _, k := range signedData {
		m.Add(int16(k), 42)
		if g := m.Get(int16(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapInt16ToInt8Put(t *testing.T) {
	m := NewMapInt16ToInt8(13)
	for i, k := range signedData {
		m.Put(int16(k), int8(signedData[i]))
	}
	for _, k := range signedData {
		m.Put(int16(k), m.Get(int16(k))+3)
	}
	for i, k := range signedData {
		if g := m.Get(int16(k)); g != int8(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int8(signedData[i])+3)
		}
	}
}

func TestMapInt16ToInt8Visit(t *testing.T) {
	m := NewMapInt16ToInt8(13)
	for i, k := range signedData {
		m.Put(int16(k), int8(signedData[i]))
	}

	n := make(map[int16]int8, len(signedData))

	for i, k := range signedData {
		n[int16(k)] = int8(signedData[i])
	}

	m.Visit(func(k int16, v int8) {
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

func TestMapInt16ToInt16Size(t *testing.T) {
	m := NewMapInt16ToInt16(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int16(k), int16(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int16(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt16ToInt16Contains(t *testing.T) {
	m := NewMapInt16ToInt16(13)
	for i, k := range signedData {
		m.Put(int16(k), int16(signedData[i]))
		if !m.Contains(int16(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int16(k))
		if m.Contains(int16(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt16ToInt16Get(t *testing.T) {
	m := NewMapInt16ToInt16(13)
	for i, k := range signedData {
		m.Put(int16(k), int16(signedData[i]))
		if g := m.Get(int16(k)); g != int16(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapInt16ToInt16Clear(t *testing.T) {
	m := NewMapInt16ToInt16(13)
	for i, k := range signedData {
		m.Put(int16(k), int16(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int16(k), int16(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapInt16ToInt16Inc(t *testing.T) {
	m := NewMapInt16ToInt16(13)
	for i, k := range signedData {
		m.Put(int16(k), int16(signedData[i]))
	}
	for i, k := range signedData {
		m.Inc(int16(k))
		m.Inc(int16(k))
		m.Inc(int16(k))
		if g := m.Get(int16(k)); g != int16(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int16(signedData[i])+3)
		}
	}
	m = NewMapInt16ToInt16(13)
	for _, k := range signedData {
		m.Inc(int16(k))
		if g := m.Get(int16(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapInt16ToInt16Add(t *testing.T) {
	m := NewMapInt16ToInt16(13)
	for i, k := range signedData {
		m.Put(int16(k), int16(signedData[i]))
	}
	for i, k := range signedData {
		m.Add(int16(k), 3)
		if g := m.Get(int16(k)); g != int16(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int16(signedData[i])+3)
		}
	}
	m = NewMapInt16ToInt16(13)
	for _, k := range signedData {
		m.Add(int16(k), 42)
		if g := m.Get(int16(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapInt16ToInt16Put(t *testing.T) {
	m := NewMapInt16ToInt16(13)
	for i, k := range signedData {
		m.Put(int16(k), int16(signedData[i]))
	}
	for _, k := range signedData {
		m.Put(int16(k), m.Get(int16(k))+3)
	}
	for i, k := range signedData {
		if g := m.Get(int16(k)); g != int16(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int16(signedData[i])+3)
		}
	}
}

func TestMapInt16ToInt16Visit(t *testing.T) {
	m := NewMapInt16ToInt16(13)
	for i, k := range signedData {
		m.Put(int16(k), int16(signedData[i]))
	}

	n := make(map[int16]int16, len(signedData))

	for i, k := range signedData {
		n[int16(k)] = int16(signedData[i])
	}

	m.Visit(func(k int16, v int16) {
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

func TestMapInt16ToInt32Size(t *testing.T) {
	m := NewMapInt16ToInt32(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int16(k), int32(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int16(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt16ToInt32Contains(t *testing.T) {
	m := NewMapInt16ToInt32(13)
	for i, k := range signedData {
		m.Put(int16(k), int32(signedData[i]))
		if !m.Contains(int16(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int16(k))
		if m.Contains(int16(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt16ToInt32Get(t *testing.T) {
	m := NewMapInt16ToInt32(13)
	for i, k := range signedData {
		m.Put(int16(k), int32(signedData[i]))
		if g := m.Get(int16(k)); g != int32(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapInt16ToInt32Clear(t *testing.T) {
	m := NewMapInt16ToInt32(13)
	for i, k := range signedData {
		m.Put(int16(k), int32(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int16(k), int32(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapInt16ToInt32Inc(t *testing.T) {
	m := NewMapInt16ToInt32(13)
	for i, k := range signedData {
		m.Put(int16(k), int32(signedData[i]))
	}
	for i, k := range signedData {
		m.Inc(int16(k))
		m.Inc(int16(k))
		m.Inc(int16(k))
		if g := m.Get(int16(k)); g != int32(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int32(signedData[i])+3)
		}
	}
	m = NewMapInt16ToInt32(13)
	for _, k := range signedData {
		m.Inc(int16(k))
		if g := m.Get(int16(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapInt16ToInt32Add(t *testing.T) {
	m := NewMapInt16ToInt32(13)
	for i, k := range signedData {
		m.Put(int16(k), int32(signedData[i]))
	}
	for i, k := range signedData {
		m.Add(int16(k), 3)
		if g := m.Get(int16(k)); g != int32(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int32(signedData[i])+3)
		}
	}
	m = NewMapInt16ToInt32(13)
	for _, k := range signedData {
		m.Add(int16(k), 42)
		if g := m.Get(int16(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapInt16ToInt32Put(t *testing.T) {
	m := NewMapInt16ToInt32(13)
	for i, k := range signedData {
		m.Put(int16(k), int32(signedData[i]))
	}
	for _, k := range signedData {
		m.Put(int16(k), m.Get(int16(k))+3)
	}
	for i, k := range signedData {
		if g := m.Get(int16(k)); g != int32(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int32(signedData[i])+3)
		}
	}
}

func TestMapInt16ToInt32Visit(t *testing.T) {
	m := NewMapInt16ToInt32(13)
	for i, k := range signedData {
		m.Put(int16(k), int32(signedData[i]))
	}

	n := make(map[int16]int32, len(signedData))

	for i, k := range signedData {
		n[int16(k)] = int32(signedData[i])
	}

	m.Visit(func(k int16, v int32) {
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

func TestMapInt16ToInt64Size(t *testing.T) {
	m := NewMapInt16ToInt64(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int16(k), int64(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int16(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt16ToInt64Contains(t *testing.T) {
	m := NewMapInt16ToInt64(13)
	for i, k := range signedData {
		m.Put(int16(k), int64(signedData[i]))
		if !m.Contains(int16(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int16(k))
		if m.Contains(int16(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt16ToInt64Get(t *testing.T) {
	m := NewMapInt16ToInt64(13)
	for i, k := range signedData {
		m.Put(int16(k), int64(signedData[i]))
		if g := m.Get(int16(k)); g != int64(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapInt16ToInt64Clear(t *testing.T) {
	m := NewMapInt16ToInt64(13)
	for i, k := range signedData {
		m.Put(int16(k), int64(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int16(k), int64(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapInt16ToInt64Inc(t *testing.T) {
	m := NewMapInt16ToInt64(13)
	for i, k := range signedData {
		m.Put(int16(k), int64(signedData[i]))
	}
	for i, k := range signedData {
		m.Inc(int16(k))
		m.Inc(int16(k))
		m.Inc(int16(k))
		if g := m.Get(int16(k)); g != int64(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int64(signedData[i])+3)
		}
	}
	m = NewMapInt16ToInt64(13)
	for _, k := range signedData {
		m.Inc(int16(k))
		if g := m.Get(int16(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapInt16ToInt64Add(t *testing.T) {
	m := NewMapInt16ToInt64(13)
	for i, k := range signedData {
		m.Put(int16(k), int64(signedData[i]))
	}
	for i, k := range signedData {
		m.Add(int16(k), 3)
		if g := m.Get(int16(k)); g != int64(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int64(signedData[i])+3)
		}
	}
	m = NewMapInt16ToInt64(13)
	for _, k := range signedData {
		m.Add(int16(k), 42)
		if g := m.Get(int16(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapInt16ToInt64Put(t *testing.T) {
	m := NewMapInt16ToInt64(13)
	for i, k := range signedData {
		m.Put(int16(k), int64(signedData[i]))
	}
	for _, k := range signedData {
		m.Put(int16(k), m.Get(int16(k))+3)
	}
	for i, k := range signedData {
		if g := m.Get(int16(k)); g != int64(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int64(signedData[i])+3)
		}
	}
}

func TestMapInt16ToInt64Visit(t *testing.T) {
	m := NewMapInt16ToInt64(13)
	for i, k := range signedData {
		m.Put(int16(k), int64(signedData[i]))
	}

	n := make(map[int16]int64, len(signedData))

	for i, k := range signedData {
		n[int16(k)] = int64(signedData[i])
	}

	m.Visit(func(k int16, v int64) {
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

func TestMapInt16ToUintSize(t *testing.T) {
	m := NewMapInt16ToUint(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int16(k), uint(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int16(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt16ToUintContains(t *testing.T) {
	m := NewMapInt16ToUint(13)
	for i, k := range signedData {
		m.Put(int16(k), uint(unsignedData[i]))
		if !m.Contains(int16(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int16(k))
		if m.Contains(int16(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt16ToUintGet(t *testing.T) {
	m := NewMapInt16ToUint(13)
	for i, k := range signedData {
		m.Put(int16(k), uint(unsignedData[i]))
		if g := m.Get(int16(k)); g != uint(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapInt16ToUintClear(t *testing.T) {
	m := NewMapInt16ToUint(13)
	for i, k := range signedData {
		m.Put(int16(k), uint(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int16(k), uint(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapInt16ToUintInc(t *testing.T) {
	m := NewMapInt16ToUint(13)
	for i, k := range signedData {
		m.Put(int16(k), uint(unsignedData[i]))
	}
	for i, k := range signedData {
		m.Inc(int16(k))
		m.Inc(int16(k))
		m.Inc(int16(k))
		if g := m.Get(int16(k)); g != uint(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint(unsignedData[i])+3)
		}
	}
	m = NewMapInt16ToUint(13)
	for _, k := range signedData {
		m.Inc(int16(k))
		if g := m.Get(int16(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapInt16ToUintAdd(t *testing.T) {
	m := NewMapInt16ToUint(13)
	for i, k := range signedData {
		m.Put(int16(k), uint(unsignedData[i]))
	}
	for i, k := range signedData {
		m.Add(int16(k), 3)
		if g := m.Get(int16(k)); g != uint(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint(unsignedData[i])+3)
		}
	}
	m = NewMapInt16ToUint(13)
	for _, k := range signedData {
		m.Add(int16(k), 42)
		if g := m.Get(int16(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapInt16ToUintPut(t *testing.T) {
	m := NewMapInt16ToUint(13)
	for i, k := range signedData {
		m.Put(int16(k), uint(unsignedData[i]))
	}
	for _, k := range signedData {
		m.Put(int16(k), m.Get(int16(k))+3)
	}
	for i, k := range signedData {
		if g := m.Get(int16(k)); g != uint(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint(unsignedData[i])+3)
		}
	}
}

func TestMapInt16ToUintVisit(t *testing.T) {
	m := NewMapInt16ToUint(13)
	for i, k := range signedData {
		m.Put(int16(k), uint(unsignedData[i]))
	}

	n := make(map[int16]uint, len(signedData))

	for i, k := range signedData {
		n[int16(k)] = uint(unsignedData[i])
	}

	m.Visit(func(k int16, v uint) {
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

func TestMapInt16ToUint8Size(t *testing.T) {
	m := NewMapInt16ToUint8(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int16(k), uint8(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int16(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt16ToUint8Contains(t *testing.T) {
	m := NewMapInt16ToUint8(13)
	for i, k := range signedData {
		m.Put(int16(k), uint8(unsignedData[i]))
		if !m.Contains(int16(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int16(k))
		if m.Contains(int16(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt16ToUint8Get(t *testing.T) {
	m := NewMapInt16ToUint8(13)
	for i, k := range signedData {
		m.Put(int16(k), uint8(unsignedData[i]))
		if g := m.Get(int16(k)); g != uint8(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapInt16ToUint8Clear(t *testing.T) {
	m := NewMapInt16ToUint8(13)
	for i, k := range signedData {
		m.Put(int16(k), uint8(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int16(k), uint8(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapInt16ToUint8Inc(t *testing.T) {
	m := NewMapInt16ToUint8(13)
	for i, k := range signedData {
		m.Put(int16(k), uint8(unsignedData[i]))
	}
	for i, k := range signedData {
		m.Inc(int16(k))
		m.Inc(int16(k))
		m.Inc(int16(k))
		if g := m.Get(int16(k)); g != uint8(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint8(unsignedData[i])+3)
		}
	}
	m = NewMapInt16ToUint8(13)
	for _, k := range signedData {
		m.Inc(int16(k))
		if g := m.Get(int16(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapInt16ToUint8Add(t *testing.T) {
	m := NewMapInt16ToUint8(13)
	for i, k := range signedData {
		m.Put(int16(k), uint8(unsignedData[i]))
	}
	for i, k := range signedData {
		m.Add(int16(k), 3)
		if g := m.Get(int16(k)); g != uint8(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint8(unsignedData[i])+3)
		}
	}
	m = NewMapInt16ToUint8(13)
	for _, k := range signedData {
		m.Add(int16(k), 42)
		if g := m.Get(int16(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapInt16ToUint8Put(t *testing.T) {
	m := NewMapInt16ToUint8(13)
	for i, k := range signedData {
		m.Put(int16(k), uint8(unsignedData[i]))
	}
	for _, k := range signedData {
		m.Put(int16(k), m.Get(int16(k))+3)
	}
	for i, k := range signedData {
		if g := m.Get(int16(k)); g != uint8(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint8(unsignedData[i])+3)
		}
	}
}

func TestMapInt16ToUint8Visit(t *testing.T) {
	m := NewMapInt16ToUint8(13)
	for i, k := range signedData {
		m.Put(int16(k), uint8(unsignedData[i]))
	}

	n := make(map[int16]uint8, len(signedData))

	for i, k := range signedData {
		n[int16(k)] = uint8(unsignedData[i])
	}

	m.Visit(func(k int16, v uint8) {
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

func TestMapInt16ToUint16Size(t *testing.T) {
	m := NewMapInt16ToUint16(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int16(k), uint16(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int16(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt16ToUint16Contains(t *testing.T) {
	m := NewMapInt16ToUint16(13)
	for i, k := range signedData {
		m.Put(int16(k), uint16(unsignedData[i]))
		if !m.Contains(int16(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int16(k))
		if m.Contains(int16(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt16ToUint16Get(t *testing.T) {
	m := NewMapInt16ToUint16(13)
	for i, k := range signedData {
		m.Put(int16(k), uint16(unsignedData[i]))
		if g := m.Get(int16(k)); g != uint16(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapInt16ToUint16Clear(t *testing.T) {
	m := NewMapInt16ToUint16(13)
	for i, k := range signedData {
		m.Put(int16(k), uint16(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int16(k), uint16(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapInt16ToUint16Inc(t *testing.T) {
	m := NewMapInt16ToUint16(13)
	for i, k := range signedData {
		m.Put(int16(k), uint16(unsignedData[i]))
	}
	for i, k := range signedData {
		m.Inc(int16(k))
		m.Inc(int16(k))
		m.Inc(int16(k))
		if g := m.Get(int16(k)); g != uint16(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint16(unsignedData[i])+3)
		}
	}
	m = NewMapInt16ToUint16(13)
	for _, k := range signedData {
		m.Inc(int16(k))
		if g := m.Get(int16(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapInt16ToUint16Add(t *testing.T) {
	m := NewMapInt16ToUint16(13)
	for i, k := range signedData {
		m.Put(int16(k), uint16(unsignedData[i]))
	}
	for i, k := range signedData {
		m.Add(int16(k), 3)
		if g := m.Get(int16(k)); g != uint16(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint16(unsignedData[i])+3)
		}
	}
	m = NewMapInt16ToUint16(13)
	for _, k := range signedData {
		m.Add(int16(k), 42)
		if g := m.Get(int16(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapInt16ToUint16Put(t *testing.T) {
	m := NewMapInt16ToUint16(13)
	for i, k := range signedData {
		m.Put(int16(k), uint16(unsignedData[i]))
	}
	for _, k := range signedData {
		m.Put(int16(k), m.Get(int16(k))+3)
	}
	for i, k := range signedData {
		if g := m.Get(int16(k)); g != uint16(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint16(unsignedData[i])+3)
		}
	}
}

func TestMapInt16ToUint16Visit(t *testing.T) {
	m := NewMapInt16ToUint16(13)
	for i, k := range signedData {
		m.Put(int16(k), uint16(unsignedData[i]))
	}

	n := make(map[int16]uint16, len(signedData))

	for i, k := range signedData {
		n[int16(k)] = uint16(unsignedData[i])
	}

	m.Visit(func(k int16, v uint16) {
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

func TestMapInt16ToUint32Size(t *testing.T) {
	m := NewMapInt16ToUint32(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int16(k), uint32(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int16(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt16ToUint32Contains(t *testing.T) {
	m := NewMapInt16ToUint32(13)
	for i, k := range signedData {
		m.Put(int16(k), uint32(unsignedData[i]))
		if !m.Contains(int16(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int16(k))
		if m.Contains(int16(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt16ToUint32Get(t *testing.T) {
	m := NewMapInt16ToUint32(13)
	for i, k := range signedData {
		m.Put(int16(k), uint32(unsignedData[i]))
		if g := m.Get(int16(k)); g != uint32(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapInt16ToUint32Clear(t *testing.T) {
	m := NewMapInt16ToUint32(13)
	for i, k := range signedData {
		m.Put(int16(k), uint32(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int16(k), uint32(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapInt16ToUint32Inc(t *testing.T) {
	m := NewMapInt16ToUint32(13)
	for i, k := range signedData {
		m.Put(int16(k), uint32(unsignedData[i]))
	}
	for i, k := range signedData {
		m.Inc(int16(k))
		m.Inc(int16(k))
		m.Inc(int16(k))
		if g := m.Get(int16(k)); g != uint32(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint32(unsignedData[i])+3)
		}
	}
	m = NewMapInt16ToUint32(13)
	for _, k := range signedData {
		m.Inc(int16(k))
		if g := m.Get(int16(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapInt16ToUint32Add(t *testing.T) {
	m := NewMapInt16ToUint32(13)
	for i, k := range signedData {
		m.Put(int16(k), uint32(unsignedData[i]))
	}
	for i, k := range signedData {
		m.Add(int16(k), 3)
		if g := m.Get(int16(k)); g != uint32(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint32(unsignedData[i])+3)
		}
	}
	m = NewMapInt16ToUint32(13)
	for _, k := range signedData {
		m.Add(int16(k), 42)
		if g := m.Get(int16(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapInt16ToUint32Put(t *testing.T) {
	m := NewMapInt16ToUint32(13)
	for i, k := range signedData {
		m.Put(int16(k), uint32(unsignedData[i]))
	}
	for _, k := range signedData {
		m.Put(int16(k), m.Get(int16(k))+3)
	}
	for i, k := range signedData {
		if g := m.Get(int16(k)); g != uint32(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint32(unsignedData[i])+3)
		}
	}
}

func TestMapInt16ToUint32Visit(t *testing.T) {
	m := NewMapInt16ToUint32(13)
	for i, k := range signedData {
		m.Put(int16(k), uint32(unsignedData[i]))
	}

	n := make(map[int16]uint32, len(signedData))

	for i, k := range signedData {
		n[int16(k)] = uint32(unsignedData[i])
	}

	m.Visit(func(k int16, v uint32) {
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

func TestMapInt16ToUint64Size(t *testing.T) {
	m := NewMapInt16ToUint64(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int16(k), uint64(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int16(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt16ToUint64Contains(t *testing.T) {
	m := NewMapInt16ToUint64(13)
	for i, k := range signedData {
		m.Put(int16(k), uint64(unsignedData[i]))
		if !m.Contains(int16(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int16(k))
		if m.Contains(int16(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt16ToUint64Get(t *testing.T) {
	m := NewMapInt16ToUint64(13)
	for i, k := range signedData {
		m.Put(int16(k), uint64(unsignedData[i]))
		if g := m.Get(int16(k)); g != uint64(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapInt16ToUint64Clear(t *testing.T) {
	m := NewMapInt16ToUint64(13)
	for i, k := range signedData {
		m.Put(int16(k), uint64(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int16(k), uint64(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapInt16ToUint64Inc(t *testing.T) {
	m := NewMapInt16ToUint64(13)
	for i, k := range signedData {
		m.Put(int16(k), uint64(unsignedData[i]))
	}
	for i, k := range signedData {
		m.Inc(int16(k))
		m.Inc(int16(k))
		m.Inc(int16(k))
		if g := m.Get(int16(k)); g != uint64(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint64(unsignedData[i])+3)
		}
	}
	m = NewMapInt16ToUint64(13)
	for _, k := range signedData {
		m.Inc(int16(k))
		if g := m.Get(int16(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapInt16ToUint64Add(t *testing.T) {
	m := NewMapInt16ToUint64(13)
	for i, k := range signedData {
		m.Put(int16(k), uint64(unsignedData[i]))
	}
	for i, k := range signedData {
		m.Add(int16(k), 3)
		if g := m.Get(int16(k)); g != uint64(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint64(unsignedData[i])+3)
		}
	}
	m = NewMapInt16ToUint64(13)
	for _, k := range signedData {
		m.Add(int16(k), 42)
		if g := m.Get(int16(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapInt16ToUint64Put(t *testing.T) {
	m := NewMapInt16ToUint64(13)
	for i, k := range signedData {
		m.Put(int16(k), uint64(unsignedData[i]))
	}
	for _, k := range signedData {
		m.Put(int16(k), m.Get(int16(k))+3)
	}
	for i, k := range signedData {
		if g := m.Get(int16(k)); g != uint64(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint64(unsignedData[i])+3)
		}
	}
}

func TestMapInt16ToUint64Visit(t *testing.T) {
	m := NewMapInt16ToUint64(13)
	for i, k := range signedData {
		m.Put(int16(k), uint64(unsignedData[i]))
	}

	n := make(map[int16]uint64, len(signedData))

	for i, k := range signedData {
		n[int16(k)] = uint64(unsignedData[i])
	}

	m.Visit(func(k int16, v uint64) {
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
