// THIS IS A MACHINE GENERATED FILE!
// BE CAREFUL WITH EDITING BY HAND.
//
// This is Free Software covered by the terms of the MIT license.
// See LICENSE file for details.
// (c) 2016 by Sascha L. Teichmann.
// See the full list of contributors in the CONTRIBUTORS file.

package genmap

import "testing"

func TestMapInt32ToIntSize(t *testing.T) {
	m := NewMapInt32ToInt(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int32(k), int(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int32(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt32ToIntContains(t *testing.T) {
	m := NewMapInt32ToInt(13)
	for i, k := range signedData {
		m.Put(int32(k), int(signedData[i]))
		if !m.Contains(int32(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int32(k))
		if m.Contains(int32(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt32ToIntGet(t *testing.T) {
	m := NewMapInt32ToInt(13)
	for i, k := range signedData {
		m.Put(int32(k), int(signedData[i]))
		if g := m.Get(int32(k)); g != int(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
}

func TestMapInt32ToIntClear(t *testing.T) {
	m := NewMapInt32ToInt(13)
	for i, k := range signedData {
		m.Put(int32(k), int(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int32(k), int(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapInt32ToIntInc(t *testing.T) {
	m := NewMapInt32ToInt(13)
	for i, k := range signedData {
		m.Put(int32(k), int(signedData[i]))
	}
	for i, k := range signedData {
		m.Inc(int32(k))
		m.Inc(int32(k))
		m.Inc(int32(k))
		if g := m.Get(int32(k)); g != int(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int(signedData[i])+3)
		}
	}
	m.Clear()
	for _, k := range signedData {
		m.Inc(int32(k))
		if g := m.Get(int32(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapInt32ToIntAdd(t *testing.T) {
	m := NewMapInt32ToInt(13)
	for i, k := range signedData {
		m.Put(int32(k), int(signedData[i]))
	}
	for i, k := range signedData {
		m.Add(int32(k), 3)
		if g := m.Get(int32(k)); g != int(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int(signedData[i])+3)
		}
	}
	m.Clear()
	for _, k := range signedData {
		m.Add(int32(k), 42)
		if g := m.Get(int32(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapInt32ToIntVisit(t *testing.T) {
	m := NewMapInt32ToInt(13)
	for i, k := range signedData {
		m.Put(int32(k), int(signedData[i]))
	}

	n := make(map[int32]int, len(signedData))

	for i, k := range signedData {
		n[int32(k)] = int(signedData[i])
	}

	m.Visit(func(k int32, v int) {
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

func TestMapInt32ToInt8Size(t *testing.T) {
	m := NewMapInt32ToInt8(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int32(k), int8(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int32(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt32ToInt8Contains(t *testing.T) {
	m := NewMapInt32ToInt8(13)
	for i, k := range signedData {
		m.Put(int32(k), int8(signedData[i]))
		if !m.Contains(int32(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int32(k))
		if m.Contains(int32(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt32ToInt8Get(t *testing.T) {
	m := NewMapInt32ToInt8(13)
	for i, k := range signedData {
		m.Put(int32(k), int8(signedData[i]))
		if g := m.Get(int32(k)); g != int8(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
}

func TestMapInt32ToInt8Clear(t *testing.T) {
	m := NewMapInt32ToInt8(13)
	for i, k := range signedData {
		m.Put(int32(k), int8(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int32(k), int8(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapInt32ToInt8Inc(t *testing.T) {
	m := NewMapInt32ToInt8(13)
	for i, k := range signedData {
		m.Put(int32(k), int8(signedData[i]))
	}
	for i, k := range signedData {
		m.Inc(int32(k))
		m.Inc(int32(k))
		m.Inc(int32(k))
		if g := m.Get(int32(k)); g != int8(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int8(signedData[i])+3)
		}
	}
	m.Clear()
	for _, k := range signedData {
		m.Inc(int32(k))
		if g := m.Get(int32(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapInt32ToInt8Add(t *testing.T) {
	m := NewMapInt32ToInt8(13)
	for i, k := range signedData {
		m.Put(int32(k), int8(signedData[i]))
	}
	for i, k := range signedData {
		m.Add(int32(k), 3)
		if g := m.Get(int32(k)); g != int8(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int8(signedData[i])+3)
		}
	}
	m.Clear()
	for _, k := range signedData {
		m.Add(int32(k), 42)
		if g := m.Get(int32(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapInt32ToInt8Visit(t *testing.T) {
	m := NewMapInt32ToInt8(13)
	for i, k := range signedData {
		m.Put(int32(k), int8(signedData[i]))
	}

	n := make(map[int32]int8, len(signedData))

	for i, k := range signedData {
		n[int32(k)] = int8(signedData[i])
	}

	m.Visit(func(k int32, v int8) {
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

func TestMapInt32ToInt16Size(t *testing.T) {
	m := NewMapInt32ToInt16(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int32(k), int16(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int32(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt32ToInt16Contains(t *testing.T) {
	m := NewMapInt32ToInt16(13)
	for i, k := range signedData {
		m.Put(int32(k), int16(signedData[i]))
		if !m.Contains(int32(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int32(k))
		if m.Contains(int32(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt32ToInt16Get(t *testing.T) {
	m := NewMapInt32ToInt16(13)
	for i, k := range signedData {
		m.Put(int32(k), int16(signedData[i]))
		if g := m.Get(int32(k)); g != int16(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
}

func TestMapInt32ToInt16Clear(t *testing.T) {
	m := NewMapInt32ToInt16(13)
	for i, k := range signedData {
		m.Put(int32(k), int16(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int32(k), int16(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapInt32ToInt16Inc(t *testing.T) {
	m := NewMapInt32ToInt16(13)
	for i, k := range signedData {
		m.Put(int32(k), int16(signedData[i]))
	}
	for i, k := range signedData {
		m.Inc(int32(k))
		m.Inc(int32(k))
		m.Inc(int32(k))
		if g := m.Get(int32(k)); g != int16(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int16(signedData[i])+3)
		}
	}
	m.Clear()
	for _, k := range signedData {
		m.Inc(int32(k))
		if g := m.Get(int32(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapInt32ToInt16Add(t *testing.T) {
	m := NewMapInt32ToInt16(13)
	for i, k := range signedData {
		m.Put(int32(k), int16(signedData[i]))
	}
	for i, k := range signedData {
		m.Add(int32(k), 3)
		if g := m.Get(int32(k)); g != int16(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int16(signedData[i])+3)
		}
	}
	m.Clear()
	for _, k := range signedData {
		m.Add(int32(k), 42)
		if g := m.Get(int32(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapInt32ToInt16Visit(t *testing.T) {
	m := NewMapInt32ToInt16(13)
	for i, k := range signedData {
		m.Put(int32(k), int16(signedData[i]))
	}

	n := make(map[int32]int16, len(signedData))

	for i, k := range signedData {
		n[int32(k)] = int16(signedData[i])
	}

	m.Visit(func(k int32, v int16) {
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

func TestMapInt32ToInt32Size(t *testing.T) {
	m := NewMapInt32ToInt32(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int32(k), int32(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int32(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt32ToInt32Contains(t *testing.T) {
	m := NewMapInt32ToInt32(13)
	for i, k := range signedData {
		m.Put(int32(k), int32(signedData[i]))
		if !m.Contains(int32(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int32(k))
		if m.Contains(int32(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt32ToInt32Get(t *testing.T) {
	m := NewMapInt32ToInt32(13)
	for i, k := range signedData {
		m.Put(int32(k), int32(signedData[i]))
		if g := m.Get(int32(k)); g != int32(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
}

func TestMapInt32ToInt32Clear(t *testing.T) {
	m := NewMapInt32ToInt32(13)
	for i, k := range signedData {
		m.Put(int32(k), int32(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int32(k), int32(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapInt32ToInt32Inc(t *testing.T) {
	m := NewMapInt32ToInt32(13)
	for i, k := range signedData {
		m.Put(int32(k), int32(signedData[i]))
	}
	for i, k := range signedData {
		m.Inc(int32(k))
		m.Inc(int32(k))
		m.Inc(int32(k))
		if g := m.Get(int32(k)); g != int32(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int32(signedData[i])+3)
		}
	}
	m.Clear()
	for _, k := range signedData {
		m.Inc(int32(k))
		if g := m.Get(int32(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapInt32ToInt32Add(t *testing.T) {
	m := NewMapInt32ToInt32(13)
	for i, k := range signedData {
		m.Put(int32(k), int32(signedData[i]))
	}
	for i, k := range signedData {
		m.Add(int32(k), 3)
		if g := m.Get(int32(k)); g != int32(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int32(signedData[i])+3)
		}
	}
	m.Clear()
	for _, k := range signedData {
		m.Add(int32(k), 42)
		if g := m.Get(int32(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapInt32ToInt32Visit(t *testing.T) {
	m := NewMapInt32ToInt32(13)
	for i, k := range signedData {
		m.Put(int32(k), int32(signedData[i]))
	}

	n := make(map[int32]int32, len(signedData))

	for i, k := range signedData {
		n[int32(k)] = int32(signedData[i])
	}

	m.Visit(func(k int32, v int32) {
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

func TestMapInt32ToInt64Size(t *testing.T) {
	m := NewMapInt32ToInt64(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int32(k), int64(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int32(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt32ToInt64Contains(t *testing.T) {
	m := NewMapInt32ToInt64(13)
	for i, k := range signedData {
		m.Put(int32(k), int64(signedData[i]))
		if !m.Contains(int32(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int32(k))
		if m.Contains(int32(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt32ToInt64Get(t *testing.T) {
	m := NewMapInt32ToInt64(13)
	for i, k := range signedData {
		m.Put(int32(k), int64(signedData[i]))
		if g := m.Get(int32(k)); g != int64(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
}

func TestMapInt32ToInt64Clear(t *testing.T) {
	m := NewMapInt32ToInt64(13)
	for i, k := range signedData {
		m.Put(int32(k), int64(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int32(k), int64(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapInt32ToInt64Inc(t *testing.T) {
	m := NewMapInt32ToInt64(13)
	for i, k := range signedData {
		m.Put(int32(k), int64(signedData[i]))
	}
	for i, k := range signedData {
		m.Inc(int32(k))
		m.Inc(int32(k))
		m.Inc(int32(k))
		if g := m.Get(int32(k)); g != int64(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int64(signedData[i])+3)
		}
	}
	m.Clear()
	for _, k := range signedData {
		m.Inc(int32(k))
		if g := m.Get(int32(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapInt32ToInt64Add(t *testing.T) {
	m := NewMapInt32ToInt64(13)
	for i, k := range signedData {
		m.Put(int32(k), int64(signedData[i]))
	}
	for i, k := range signedData {
		m.Add(int32(k), 3)
		if g := m.Get(int32(k)); g != int64(signedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, int64(signedData[i])+3)
		}
	}
	m.Clear()
	for _, k := range signedData {
		m.Add(int32(k), 42)
		if g := m.Get(int32(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapInt32ToInt64Visit(t *testing.T) {
	m := NewMapInt32ToInt64(13)
	for i, k := range signedData {
		m.Put(int32(k), int64(signedData[i]))
	}

	n := make(map[int32]int64, len(signedData))

	for i, k := range signedData {
		n[int32(k)] = int64(signedData[i])
	}

	m.Visit(func(k int32, v int64) {
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

func TestMapInt32ToUintSize(t *testing.T) {
	m := NewMapInt32ToUint(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int32(k), uint(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int32(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt32ToUintContains(t *testing.T) {
	m := NewMapInt32ToUint(13)
	for i, k := range signedData {
		m.Put(int32(k), uint(unsignedData[i]))
		if !m.Contains(int32(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int32(k))
		if m.Contains(int32(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt32ToUintGet(t *testing.T) {
	m := NewMapInt32ToUint(13)
	for i, k := range signedData {
		m.Put(int32(k), uint(unsignedData[i]))
		if g := m.Get(int32(k)); g != uint(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
}

func TestMapInt32ToUintClear(t *testing.T) {
	m := NewMapInt32ToUint(13)
	for i, k := range signedData {
		m.Put(int32(k), uint(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int32(k), uint(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapInt32ToUintInc(t *testing.T) {
	m := NewMapInt32ToUint(13)
	for i, k := range signedData {
		m.Put(int32(k), uint(unsignedData[i]))
	}
	for i, k := range signedData {
		m.Inc(int32(k))
		m.Inc(int32(k))
		m.Inc(int32(k))
		if g := m.Get(int32(k)); g != uint(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint(unsignedData[i])+3)
		}
	}
	m.Clear()
	for _, k := range signedData {
		m.Inc(int32(k))
		if g := m.Get(int32(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapInt32ToUintAdd(t *testing.T) {
	m := NewMapInt32ToUint(13)
	for i, k := range signedData {
		m.Put(int32(k), uint(unsignedData[i]))
	}
	for i, k := range signedData {
		m.Add(int32(k), 3)
		if g := m.Get(int32(k)); g != uint(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint(unsignedData[i])+3)
		}
	}
	m.Clear()
	for _, k := range signedData {
		m.Add(int32(k), 42)
		if g := m.Get(int32(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapInt32ToUintVisit(t *testing.T) {
	m := NewMapInt32ToUint(13)
	for i, k := range signedData {
		m.Put(int32(k), uint(unsignedData[i]))
	}

	n := make(map[int32]uint, len(signedData))

	for i, k := range signedData {
		n[int32(k)] = uint(unsignedData[i])
	}

	m.Visit(func(k int32, v uint) {
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

func TestMapInt32ToUint8Size(t *testing.T) {
	m := NewMapInt32ToUint8(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int32(k), uint8(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int32(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt32ToUint8Contains(t *testing.T) {
	m := NewMapInt32ToUint8(13)
	for i, k := range signedData {
		m.Put(int32(k), uint8(unsignedData[i]))
		if !m.Contains(int32(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int32(k))
		if m.Contains(int32(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt32ToUint8Get(t *testing.T) {
	m := NewMapInt32ToUint8(13)
	for i, k := range signedData {
		m.Put(int32(k), uint8(unsignedData[i]))
		if g := m.Get(int32(k)); g != uint8(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
}

func TestMapInt32ToUint8Clear(t *testing.T) {
	m := NewMapInt32ToUint8(13)
	for i, k := range signedData {
		m.Put(int32(k), uint8(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int32(k), uint8(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapInt32ToUint8Inc(t *testing.T) {
	m := NewMapInt32ToUint8(13)
	for i, k := range signedData {
		m.Put(int32(k), uint8(unsignedData[i]))
	}
	for i, k := range signedData {
		m.Inc(int32(k))
		m.Inc(int32(k))
		m.Inc(int32(k))
		if g := m.Get(int32(k)); g != uint8(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint8(unsignedData[i])+3)
		}
	}
	m.Clear()
	for _, k := range signedData {
		m.Inc(int32(k))
		if g := m.Get(int32(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapInt32ToUint8Add(t *testing.T) {
	m := NewMapInt32ToUint8(13)
	for i, k := range signedData {
		m.Put(int32(k), uint8(unsignedData[i]))
	}
	for i, k := range signedData {
		m.Add(int32(k), 3)
		if g := m.Get(int32(k)); g != uint8(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint8(unsignedData[i])+3)
		}
	}
	m.Clear()
	for _, k := range signedData {
		m.Add(int32(k), 42)
		if g := m.Get(int32(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapInt32ToUint8Visit(t *testing.T) {
	m := NewMapInt32ToUint8(13)
	for i, k := range signedData {
		m.Put(int32(k), uint8(unsignedData[i]))
	}

	n := make(map[int32]uint8, len(signedData))

	for i, k := range signedData {
		n[int32(k)] = uint8(unsignedData[i])
	}

	m.Visit(func(k int32, v uint8) {
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

func TestMapInt32ToUint16Size(t *testing.T) {
	m := NewMapInt32ToUint16(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int32(k), uint16(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int32(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt32ToUint16Contains(t *testing.T) {
	m := NewMapInt32ToUint16(13)
	for i, k := range signedData {
		m.Put(int32(k), uint16(unsignedData[i]))
		if !m.Contains(int32(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int32(k))
		if m.Contains(int32(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt32ToUint16Get(t *testing.T) {
	m := NewMapInt32ToUint16(13)
	for i, k := range signedData {
		m.Put(int32(k), uint16(unsignedData[i]))
		if g := m.Get(int32(k)); g != uint16(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
}

func TestMapInt32ToUint16Clear(t *testing.T) {
	m := NewMapInt32ToUint16(13)
	for i, k := range signedData {
		m.Put(int32(k), uint16(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int32(k), uint16(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapInt32ToUint16Inc(t *testing.T) {
	m := NewMapInt32ToUint16(13)
	for i, k := range signedData {
		m.Put(int32(k), uint16(unsignedData[i]))
	}
	for i, k := range signedData {
		m.Inc(int32(k))
		m.Inc(int32(k))
		m.Inc(int32(k))
		if g := m.Get(int32(k)); g != uint16(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint16(unsignedData[i])+3)
		}
	}
	m.Clear()
	for _, k := range signedData {
		m.Inc(int32(k))
		if g := m.Get(int32(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapInt32ToUint16Add(t *testing.T) {
	m := NewMapInt32ToUint16(13)
	for i, k := range signedData {
		m.Put(int32(k), uint16(unsignedData[i]))
	}
	for i, k := range signedData {
		m.Add(int32(k), 3)
		if g := m.Get(int32(k)); g != uint16(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint16(unsignedData[i])+3)
		}
	}
	m.Clear()
	for _, k := range signedData {
		m.Add(int32(k), 42)
		if g := m.Get(int32(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapInt32ToUint16Visit(t *testing.T) {
	m := NewMapInt32ToUint16(13)
	for i, k := range signedData {
		m.Put(int32(k), uint16(unsignedData[i]))
	}

	n := make(map[int32]uint16, len(signedData))

	for i, k := range signedData {
		n[int32(k)] = uint16(unsignedData[i])
	}

	m.Visit(func(k int32, v uint16) {
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

func TestMapInt32ToUint32Size(t *testing.T) {
	m := NewMapInt32ToUint32(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int32(k), uint32(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int32(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt32ToUint32Contains(t *testing.T) {
	m := NewMapInt32ToUint32(13)
	for i, k := range signedData {
		m.Put(int32(k), uint32(unsignedData[i]))
		if !m.Contains(int32(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int32(k))
		if m.Contains(int32(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt32ToUint32Get(t *testing.T) {
	m := NewMapInt32ToUint32(13)
	for i, k := range signedData {
		m.Put(int32(k), uint32(unsignedData[i]))
		if g := m.Get(int32(k)); g != uint32(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
}

func TestMapInt32ToUint32Clear(t *testing.T) {
	m := NewMapInt32ToUint32(13)
	for i, k := range signedData {
		m.Put(int32(k), uint32(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int32(k), uint32(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapInt32ToUint32Inc(t *testing.T) {
	m := NewMapInt32ToUint32(13)
	for i, k := range signedData {
		m.Put(int32(k), uint32(unsignedData[i]))
	}
	for i, k := range signedData {
		m.Inc(int32(k))
		m.Inc(int32(k))
		m.Inc(int32(k))
		if g := m.Get(int32(k)); g != uint32(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint32(unsignedData[i])+3)
		}
	}
	m.Clear()
	for _, k := range signedData {
		m.Inc(int32(k))
		if g := m.Get(int32(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapInt32ToUint32Add(t *testing.T) {
	m := NewMapInt32ToUint32(13)
	for i, k := range signedData {
		m.Put(int32(k), uint32(unsignedData[i]))
	}
	for i, k := range signedData {
		m.Add(int32(k), 3)
		if g := m.Get(int32(k)); g != uint32(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint32(unsignedData[i])+3)
		}
	}
	m.Clear()
	for _, k := range signedData {
		m.Add(int32(k), 42)
		if g := m.Get(int32(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapInt32ToUint32Visit(t *testing.T) {
	m := NewMapInt32ToUint32(13)
	for i, k := range signedData {
		m.Put(int32(k), uint32(unsignedData[i]))
	}

	n := make(map[int32]uint32, len(signedData))

	for i, k := range signedData {
		n[int32(k)] = uint32(unsignedData[i])
	}

	m.Visit(func(k int32, v uint32) {
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

func TestMapInt32ToUint64Size(t *testing.T) {
	m := NewMapInt32ToUint64(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int32(k), uint64(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int32(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt32ToUint64Contains(t *testing.T) {
	m := NewMapInt32ToUint64(13)
	for i, k := range signedData {
		m.Put(int32(k), uint64(unsignedData[i]))
		if !m.Contains(int32(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int32(k))
		if m.Contains(int32(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt32ToUint64Get(t *testing.T) {
	m := NewMapInt32ToUint64(13)
	for i, k := range signedData {
		m.Put(int32(k), uint64(unsignedData[i]))
		if g := m.Get(int32(k)); g != uint64(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
}

func TestMapInt32ToUint64Clear(t *testing.T) {
	m := NewMapInt32ToUint64(13)
	for i, k := range signedData {
		m.Put(int32(k), uint64(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int32(k), uint64(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapInt32ToUint64Inc(t *testing.T) {
	m := NewMapInt32ToUint64(13)
	for i, k := range signedData {
		m.Put(int32(k), uint64(unsignedData[i]))
	}
	for i, k := range signedData {
		m.Inc(int32(k))
		m.Inc(int32(k))
		m.Inc(int32(k))
		if g := m.Get(int32(k)); g != uint64(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint64(unsignedData[i])+3)
		}
	}
	m.Clear()
	for _, k := range signedData {
		m.Inc(int32(k))
		if g := m.Get(int32(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapInt32ToUint64Add(t *testing.T) {
	m := NewMapInt32ToUint64(13)
	for i, k := range signedData {
		m.Put(int32(k), uint64(unsignedData[i]))
	}
	for i, k := range signedData {
		m.Add(int32(k), 3)
		if g := m.Get(int32(k)); g != uint64(unsignedData[i])+3 {
			t.Errorf("got %d, want 0\n", g, uint64(unsignedData[i])+3)
		}
	}
	m.Clear()
	for _, k := range signedData {
		m.Add(int32(k), 42)
		if g := m.Get(int32(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapInt32ToUint64Visit(t *testing.T) {
	m := NewMapInt32ToUint64(13)
	for i, k := range signedData {
		m.Put(int32(k), uint64(unsignedData[i]))
	}

	n := make(map[int32]uint64, len(signedData))

	for i, k := range signedData {
		n[int32(k)] = uint64(unsignedData[i])
	}

	m.Visit(func(k int32, v uint64) {
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
