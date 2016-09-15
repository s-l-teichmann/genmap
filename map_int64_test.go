// THIS IS A MACHINE GENERATED FILE!
// BE CAREFUL WITH EDITING BY HAND.
//
// This is Free Software covered by the terms of the MIT license.
// See LICENSE file for details.
// (c) 2016 by Sascha L. Teichmann.
// See the full list of contributors in the CONTRIBUTORS file.

package genmap

import "testing"

func TestMapInt64ToIntSize(t *testing.T) {
	m := NewMapInt64ToInt(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int64(k), int(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int64(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt64ToIntContains(t *testing.T) {
	m := NewMapInt64ToInt(13)
	for i, k := range signedData {
		m.Put(int64(k), int(signedData[i]))
		if !m.Contains(int64(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int64(k))
		if m.Contains(int64(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt64ToIntGet(t *testing.T) {
	m := NewMapInt64ToInt(13)
	for i, k := range signedData {
		m.Put(int64(k), int(signedData[i]))
		if g := m.Get(int64(k)); g != int(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapInt64ToIntRemove(t *testing.T) {
	m := NewMapInt64ToInt(13)
	// Used to trigger code path when all keys are
	// hashed to same slot.
	mask := ^(^0 << nextShiftPowerOfTwo(13))
	for i, k := range signedData {
		if int(k)&mask == 0 {
			m.Put(int64(k), int(signedData[i]))
		}
	}
	s := m.Size()
	for _, k := range signedData {
		if int(k)&mask == 0 {
			m.Remove(int64(k))
			s--
			if l := m.Size(); l != s {
				t.Errorf("key: %d size %d, want %d\n", k, l, s)
			}
		}
	}

}

func TestMapInt64ToIntClear(t *testing.T) {
	m := NewMapInt64ToInt(13)
	for i, k := range signedData {
		m.Put(int64(k), int(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int64(k), int(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapInt64ToIntInc(t *testing.T) {
	m := NewMapInt64ToInt(13)
	for i, k := range signedData {
		m.Put(int64(k), int(signedData[i]))
	}
	for i, k := range signedData {
		m.Inc(int64(k))
		m.Inc(int64(k))
		m.Inc(int64(k))
		if g := m.Get(int64(k)); g != int(signedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, int(signedData[i])+3)
		}
	}
	m = NewMapInt64ToInt(13)
	for _, k := range signedData {
		m.Inc(int64(k))
		if g := m.Get(int64(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapInt64ToIntAdd(t *testing.T) {
	m := NewMapInt64ToInt(13)
	for i, k := range signedData {
		m.Put(int64(k), int(signedData[i]))
	}
	for i, k := range signedData {
		m.Add(int64(k), 3)
		if g := m.Get(int64(k)); g != int(signedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, int(signedData[i])+3)
		}
	}
	m = NewMapInt64ToInt(13)
	for _, k := range signedData {
		m.Add(int64(k), 42)
		if g := m.Get(int64(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapInt64ToIntModify(t *testing.T) {
	m := NewMapInt64ToInt(13)
	for i, k := range signedData {
		m.Put(int64(k), int(signedData[i]))
	}
	for _, k := range signedData {
		m.Modify(int64(k), func(v *int) {
			*v += 3
		})
	}
	for i, k := range signedData {
		if g := m.Get(int64(k)); g != int(signedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, int(signedData[i])+3)
		}
	}
}

func TestMapInt64ToIntFind(t *testing.T) {
	m := NewMapInt64ToInt(13)

	low := signedData[:len(signedData)/2]
	hi := signedData[len(signedData)/2:]

	for i, k := range low {
		m.Put(int64(k), int(signedData[i]))
	}
	for i, k := range low {
		v, ok := m.Find(int64(k))
		if !ok {
			t.Errorf("missing value for key %d\n", k)
			continue
		}
		if want := int(signedData[i]); v != want {
			t.Errorf("key %d: got %d, want %d\n", k, v, want)
		}
	}
	for _, k := range hi {
		if v, ok := m.Find(int64(k)); ok {
			t.Errorf("key %d: got %d, want nothing\n", k, v)
		}
	}
}

func TestMapInt64ToIntPut(t *testing.T) {
	m := NewMapInt64ToInt(13)
	for i, k := range signedData {
		m.Put(int64(k), int(signedData[i]))
	}
	for _, k := range signedData {
		m.Put(int64(k), m.Get(int64(k))+3)
	}
	for i, k := range signedData {
		if g := m.Get(int64(k)); g != int(signedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, int(signedData[i])+3)
		}
	}
}

func TestMapInt64ToIntVisit(t *testing.T) {
	m := NewMapInt64ToInt(13)
	for i, k := range signedData {
		m.Put(int64(k), int(signedData[i]))
	}

	n := make(map[int64]int, len(signedData))

	for i, k := range signedData {
		n[int64(k)] = int(signedData[i])
	}

	m.Visit(func(k int64, v int) {
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

func TestMapInt64ToInt8Size(t *testing.T) {
	m := NewMapInt64ToInt8(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int64(k), int8(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int64(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt64ToInt8Contains(t *testing.T) {
	m := NewMapInt64ToInt8(13)
	for i, k := range signedData {
		m.Put(int64(k), int8(signedData[i]))
		if !m.Contains(int64(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int64(k))
		if m.Contains(int64(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt64ToInt8Get(t *testing.T) {
	m := NewMapInt64ToInt8(13)
	for i, k := range signedData {
		m.Put(int64(k), int8(signedData[i]))
		if g := m.Get(int64(k)); g != int8(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapInt64ToInt8Remove(t *testing.T) {
	m := NewMapInt64ToInt8(13)
	// Used to trigger code path when all keys are
	// hashed to same slot.
	mask := ^(^0 << nextShiftPowerOfTwo(13))
	for i, k := range signedData {
		if int(k)&mask == 0 {
			m.Put(int64(k), int8(signedData[i]))
		}
	}
	s := m.Size()
	for _, k := range signedData {
		if int(k)&mask == 0 {
			m.Remove(int64(k))
			s--
			if l := m.Size(); l != s {
				t.Errorf("key: %d size %d, want %d\n", k, l, s)
			}
		}
	}

}

func TestMapInt64ToInt8Clear(t *testing.T) {
	m := NewMapInt64ToInt8(13)
	for i, k := range signedData {
		m.Put(int64(k), int8(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int64(k), int8(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapInt64ToInt8Inc(t *testing.T) {
	m := NewMapInt64ToInt8(13)
	for i, k := range signedData {
		m.Put(int64(k), int8(signedData[i]))
	}
	for i, k := range signedData {
		m.Inc(int64(k))
		m.Inc(int64(k))
		m.Inc(int64(k))
		if g := m.Get(int64(k)); g != int8(signedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, int8(signedData[i])+3)
		}
	}
	m = NewMapInt64ToInt8(13)
	for _, k := range signedData {
		m.Inc(int64(k))
		if g := m.Get(int64(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapInt64ToInt8Add(t *testing.T) {
	m := NewMapInt64ToInt8(13)
	for i, k := range signedData {
		m.Put(int64(k), int8(signedData[i]))
	}
	for i, k := range signedData {
		m.Add(int64(k), 3)
		if g := m.Get(int64(k)); g != int8(signedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, int8(signedData[i])+3)
		}
	}
	m = NewMapInt64ToInt8(13)
	for _, k := range signedData {
		m.Add(int64(k), 42)
		if g := m.Get(int64(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapInt64ToInt8Modify(t *testing.T) {
	m := NewMapInt64ToInt8(13)
	for i, k := range signedData {
		m.Put(int64(k), int8(signedData[i]))
	}
	for _, k := range signedData {
		m.Modify(int64(k), func(v *int8) {
			*v += 3
		})
	}
	for i, k := range signedData {
		if g := m.Get(int64(k)); g != int8(signedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, int8(signedData[i])+3)
		}
	}
}

func TestMapInt64ToInt8Find(t *testing.T) {
	m := NewMapInt64ToInt8(13)

	low := signedData[:len(signedData)/2]
	hi := signedData[len(signedData)/2:]

	for i, k := range low {
		m.Put(int64(k), int8(signedData[i]))
	}
	for i, k := range low {
		v, ok := m.Find(int64(k))
		if !ok {
			t.Errorf("missing value for key %d\n", k)
			continue
		}
		if want := int8(signedData[i]); v != want {
			t.Errorf("key %d: got %d, want %d\n", k, v, want)
		}
	}
	for _, k := range hi {
		if v, ok := m.Find(int64(k)); ok {
			t.Errorf("key %d: got %d, want nothing\n", k, v)
		}
	}
}

func TestMapInt64ToInt8Put(t *testing.T) {
	m := NewMapInt64ToInt8(13)
	for i, k := range signedData {
		m.Put(int64(k), int8(signedData[i]))
	}
	for _, k := range signedData {
		m.Put(int64(k), m.Get(int64(k))+3)
	}
	for i, k := range signedData {
		if g := m.Get(int64(k)); g != int8(signedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, int8(signedData[i])+3)
		}
	}
}

func TestMapInt64ToInt8Visit(t *testing.T) {
	m := NewMapInt64ToInt8(13)
	for i, k := range signedData {
		m.Put(int64(k), int8(signedData[i]))
	}

	n := make(map[int64]int8, len(signedData))

	for i, k := range signedData {
		n[int64(k)] = int8(signedData[i])
	}

	m.Visit(func(k int64, v int8) {
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

func TestMapInt64ToInt16Size(t *testing.T) {
	m := NewMapInt64ToInt16(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int64(k), int16(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int64(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt64ToInt16Contains(t *testing.T) {
	m := NewMapInt64ToInt16(13)
	for i, k := range signedData {
		m.Put(int64(k), int16(signedData[i]))
		if !m.Contains(int64(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int64(k))
		if m.Contains(int64(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt64ToInt16Get(t *testing.T) {
	m := NewMapInt64ToInt16(13)
	for i, k := range signedData {
		m.Put(int64(k), int16(signedData[i]))
		if g := m.Get(int64(k)); g != int16(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapInt64ToInt16Remove(t *testing.T) {
	m := NewMapInt64ToInt16(13)
	// Used to trigger code path when all keys are
	// hashed to same slot.
	mask := ^(^0 << nextShiftPowerOfTwo(13))
	for i, k := range signedData {
		if int(k)&mask == 0 {
			m.Put(int64(k), int16(signedData[i]))
		}
	}
	s := m.Size()
	for _, k := range signedData {
		if int(k)&mask == 0 {
			m.Remove(int64(k))
			s--
			if l := m.Size(); l != s {
				t.Errorf("key: %d size %d, want %d\n", k, l, s)
			}
		}
	}

}

func TestMapInt64ToInt16Clear(t *testing.T) {
	m := NewMapInt64ToInt16(13)
	for i, k := range signedData {
		m.Put(int64(k), int16(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int64(k), int16(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapInt64ToInt16Inc(t *testing.T) {
	m := NewMapInt64ToInt16(13)
	for i, k := range signedData {
		m.Put(int64(k), int16(signedData[i]))
	}
	for i, k := range signedData {
		m.Inc(int64(k))
		m.Inc(int64(k))
		m.Inc(int64(k))
		if g := m.Get(int64(k)); g != int16(signedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, int16(signedData[i])+3)
		}
	}
	m = NewMapInt64ToInt16(13)
	for _, k := range signedData {
		m.Inc(int64(k))
		if g := m.Get(int64(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapInt64ToInt16Add(t *testing.T) {
	m := NewMapInt64ToInt16(13)
	for i, k := range signedData {
		m.Put(int64(k), int16(signedData[i]))
	}
	for i, k := range signedData {
		m.Add(int64(k), 3)
		if g := m.Get(int64(k)); g != int16(signedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, int16(signedData[i])+3)
		}
	}
	m = NewMapInt64ToInt16(13)
	for _, k := range signedData {
		m.Add(int64(k), 42)
		if g := m.Get(int64(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapInt64ToInt16Modify(t *testing.T) {
	m := NewMapInt64ToInt16(13)
	for i, k := range signedData {
		m.Put(int64(k), int16(signedData[i]))
	}
	for _, k := range signedData {
		m.Modify(int64(k), func(v *int16) {
			*v += 3
		})
	}
	for i, k := range signedData {
		if g := m.Get(int64(k)); g != int16(signedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, int16(signedData[i])+3)
		}
	}
}

func TestMapInt64ToInt16Find(t *testing.T) {
	m := NewMapInt64ToInt16(13)

	low := signedData[:len(signedData)/2]
	hi := signedData[len(signedData)/2:]

	for i, k := range low {
		m.Put(int64(k), int16(signedData[i]))
	}
	for i, k := range low {
		v, ok := m.Find(int64(k))
		if !ok {
			t.Errorf("missing value for key %d\n", k)
			continue
		}
		if want := int16(signedData[i]); v != want {
			t.Errorf("key %d: got %d, want %d\n", k, v, want)
		}
	}
	for _, k := range hi {
		if v, ok := m.Find(int64(k)); ok {
			t.Errorf("key %d: got %d, want nothing\n", k, v)
		}
	}
}

func TestMapInt64ToInt16Put(t *testing.T) {
	m := NewMapInt64ToInt16(13)
	for i, k := range signedData {
		m.Put(int64(k), int16(signedData[i]))
	}
	for _, k := range signedData {
		m.Put(int64(k), m.Get(int64(k))+3)
	}
	for i, k := range signedData {
		if g := m.Get(int64(k)); g != int16(signedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, int16(signedData[i])+3)
		}
	}
}

func TestMapInt64ToInt16Visit(t *testing.T) {
	m := NewMapInt64ToInt16(13)
	for i, k := range signedData {
		m.Put(int64(k), int16(signedData[i]))
	}

	n := make(map[int64]int16, len(signedData))

	for i, k := range signedData {
		n[int64(k)] = int16(signedData[i])
	}

	m.Visit(func(k int64, v int16) {
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

func TestMapInt64ToInt32Size(t *testing.T) {
	m := NewMapInt64ToInt32(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int64(k), int32(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int64(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt64ToInt32Contains(t *testing.T) {
	m := NewMapInt64ToInt32(13)
	for i, k := range signedData {
		m.Put(int64(k), int32(signedData[i]))
		if !m.Contains(int64(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int64(k))
		if m.Contains(int64(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt64ToInt32Get(t *testing.T) {
	m := NewMapInt64ToInt32(13)
	for i, k := range signedData {
		m.Put(int64(k), int32(signedData[i]))
		if g := m.Get(int64(k)); g != int32(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapInt64ToInt32Remove(t *testing.T) {
	m := NewMapInt64ToInt32(13)
	// Used to trigger code path when all keys are
	// hashed to same slot.
	mask := ^(^0 << nextShiftPowerOfTwo(13))
	for i, k := range signedData {
		if int(k)&mask == 0 {
			m.Put(int64(k), int32(signedData[i]))
		}
	}
	s := m.Size()
	for _, k := range signedData {
		if int(k)&mask == 0 {
			m.Remove(int64(k))
			s--
			if l := m.Size(); l != s {
				t.Errorf("key: %d size %d, want %d\n", k, l, s)
			}
		}
	}

}

func TestMapInt64ToInt32Clear(t *testing.T) {
	m := NewMapInt64ToInt32(13)
	for i, k := range signedData {
		m.Put(int64(k), int32(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int64(k), int32(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapInt64ToInt32Inc(t *testing.T) {
	m := NewMapInt64ToInt32(13)
	for i, k := range signedData {
		m.Put(int64(k), int32(signedData[i]))
	}
	for i, k := range signedData {
		m.Inc(int64(k))
		m.Inc(int64(k))
		m.Inc(int64(k))
		if g := m.Get(int64(k)); g != int32(signedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, int32(signedData[i])+3)
		}
	}
	m = NewMapInt64ToInt32(13)
	for _, k := range signedData {
		m.Inc(int64(k))
		if g := m.Get(int64(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapInt64ToInt32Add(t *testing.T) {
	m := NewMapInt64ToInt32(13)
	for i, k := range signedData {
		m.Put(int64(k), int32(signedData[i]))
	}
	for i, k := range signedData {
		m.Add(int64(k), 3)
		if g := m.Get(int64(k)); g != int32(signedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, int32(signedData[i])+3)
		}
	}
	m = NewMapInt64ToInt32(13)
	for _, k := range signedData {
		m.Add(int64(k), 42)
		if g := m.Get(int64(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapInt64ToInt32Modify(t *testing.T) {
	m := NewMapInt64ToInt32(13)
	for i, k := range signedData {
		m.Put(int64(k), int32(signedData[i]))
	}
	for _, k := range signedData {
		m.Modify(int64(k), func(v *int32) {
			*v += 3
		})
	}
	for i, k := range signedData {
		if g := m.Get(int64(k)); g != int32(signedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, int32(signedData[i])+3)
		}
	}
}

func TestMapInt64ToInt32Find(t *testing.T) {
	m := NewMapInt64ToInt32(13)

	low := signedData[:len(signedData)/2]
	hi := signedData[len(signedData)/2:]

	for i, k := range low {
		m.Put(int64(k), int32(signedData[i]))
	}
	for i, k := range low {
		v, ok := m.Find(int64(k))
		if !ok {
			t.Errorf("missing value for key %d\n", k)
			continue
		}
		if want := int32(signedData[i]); v != want {
			t.Errorf("key %d: got %d, want %d\n", k, v, want)
		}
	}
	for _, k := range hi {
		if v, ok := m.Find(int64(k)); ok {
			t.Errorf("key %d: got %d, want nothing\n", k, v)
		}
	}
}

func TestMapInt64ToInt32Put(t *testing.T) {
	m := NewMapInt64ToInt32(13)
	for i, k := range signedData {
		m.Put(int64(k), int32(signedData[i]))
	}
	for _, k := range signedData {
		m.Put(int64(k), m.Get(int64(k))+3)
	}
	for i, k := range signedData {
		if g := m.Get(int64(k)); g != int32(signedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, int32(signedData[i])+3)
		}
	}
}

func TestMapInt64ToInt32Visit(t *testing.T) {
	m := NewMapInt64ToInt32(13)
	for i, k := range signedData {
		m.Put(int64(k), int32(signedData[i]))
	}

	n := make(map[int64]int32, len(signedData))

	for i, k := range signedData {
		n[int64(k)] = int32(signedData[i])
	}

	m.Visit(func(k int64, v int32) {
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

func TestMapInt64ToInt64Size(t *testing.T) {
	m := NewMapInt64ToInt64(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int64(k), int64(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int64(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt64ToInt64Contains(t *testing.T) {
	m := NewMapInt64ToInt64(13)
	for i, k := range signedData {
		m.Put(int64(k), int64(signedData[i]))
		if !m.Contains(int64(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int64(k))
		if m.Contains(int64(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt64ToInt64Get(t *testing.T) {
	m := NewMapInt64ToInt64(13)
	for i, k := range signedData {
		m.Put(int64(k), int64(signedData[i]))
		if g := m.Get(int64(k)); g != int64(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapInt64ToInt64Remove(t *testing.T) {
	m := NewMapInt64ToInt64(13)
	// Used to trigger code path when all keys are
	// hashed to same slot.
	mask := ^(^0 << nextShiftPowerOfTwo(13))
	for i, k := range signedData {
		if int(k)&mask == 0 {
			m.Put(int64(k), int64(signedData[i]))
		}
	}
	s := m.Size()
	for _, k := range signedData {
		if int(k)&mask == 0 {
			m.Remove(int64(k))
			s--
			if l := m.Size(); l != s {
				t.Errorf("key: %d size %d, want %d\n", k, l, s)
			}
		}
	}

}

func TestMapInt64ToInt64Clear(t *testing.T) {
	m := NewMapInt64ToInt64(13)
	for i, k := range signedData {
		m.Put(int64(k), int64(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int64(k), int64(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapInt64ToInt64Inc(t *testing.T) {
	m := NewMapInt64ToInt64(13)
	for i, k := range signedData {
		m.Put(int64(k), int64(signedData[i]))
	}
	for i, k := range signedData {
		m.Inc(int64(k))
		m.Inc(int64(k))
		m.Inc(int64(k))
		if g := m.Get(int64(k)); g != int64(signedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, int64(signedData[i])+3)
		}
	}
	m = NewMapInt64ToInt64(13)
	for _, k := range signedData {
		m.Inc(int64(k))
		if g := m.Get(int64(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapInt64ToInt64Add(t *testing.T) {
	m := NewMapInt64ToInt64(13)
	for i, k := range signedData {
		m.Put(int64(k), int64(signedData[i]))
	}
	for i, k := range signedData {
		m.Add(int64(k), 3)
		if g := m.Get(int64(k)); g != int64(signedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, int64(signedData[i])+3)
		}
	}
	m = NewMapInt64ToInt64(13)
	for _, k := range signedData {
		m.Add(int64(k), 42)
		if g := m.Get(int64(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapInt64ToInt64Modify(t *testing.T) {
	m := NewMapInt64ToInt64(13)
	for i, k := range signedData {
		m.Put(int64(k), int64(signedData[i]))
	}
	for _, k := range signedData {
		m.Modify(int64(k), func(v *int64) {
			*v += 3
		})
	}
	for i, k := range signedData {
		if g := m.Get(int64(k)); g != int64(signedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, int64(signedData[i])+3)
		}
	}
}

func TestMapInt64ToInt64Find(t *testing.T) {
	m := NewMapInt64ToInt64(13)

	low := signedData[:len(signedData)/2]
	hi := signedData[len(signedData)/2:]

	for i, k := range low {
		m.Put(int64(k), int64(signedData[i]))
	}
	for i, k := range low {
		v, ok := m.Find(int64(k))
		if !ok {
			t.Errorf("missing value for key %d\n", k)
			continue
		}
		if want := int64(signedData[i]); v != want {
			t.Errorf("key %d: got %d, want %d\n", k, v, want)
		}
	}
	for _, k := range hi {
		if v, ok := m.Find(int64(k)); ok {
			t.Errorf("key %d: got %d, want nothing\n", k, v)
		}
	}
}

func TestMapInt64ToInt64Put(t *testing.T) {
	m := NewMapInt64ToInt64(13)
	for i, k := range signedData {
		m.Put(int64(k), int64(signedData[i]))
	}
	for _, k := range signedData {
		m.Put(int64(k), m.Get(int64(k))+3)
	}
	for i, k := range signedData {
		if g := m.Get(int64(k)); g != int64(signedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, int64(signedData[i])+3)
		}
	}
}

func TestMapInt64ToInt64Visit(t *testing.T) {
	m := NewMapInt64ToInt64(13)
	for i, k := range signedData {
		m.Put(int64(k), int64(signedData[i]))
	}

	n := make(map[int64]int64, len(signedData))

	for i, k := range signedData {
		n[int64(k)] = int64(signedData[i])
	}

	m.Visit(func(k int64, v int64) {
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

func TestMapInt64ToUintSize(t *testing.T) {
	m := NewMapInt64ToUint(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int64(k), uint(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int64(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt64ToUintContains(t *testing.T) {
	m := NewMapInt64ToUint(13)
	for i, k := range signedData {
		m.Put(int64(k), uint(unsignedData[i]))
		if !m.Contains(int64(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int64(k))
		if m.Contains(int64(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt64ToUintGet(t *testing.T) {
	m := NewMapInt64ToUint(13)
	for i, k := range signedData {
		m.Put(int64(k), uint(unsignedData[i]))
		if g := m.Get(int64(k)); g != uint(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapInt64ToUintRemove(t *testing.T) {
	m := NewMapInt64ToUint(13)
	// Used to trigger code path when all keys are
	// hashed to same slot.
	mask := ^(^0 << nextShiftPowerOfTwo(13))
	for i, k := range signedData {
		if int(k)&mask == 0 {
			m.Put(int64(k), uint(unsignedData[i]))
		}
	}
	s := m.Size()
	for _, k := range signedData {
		if int(k)&mask == 0 {
			m.Remove(int64(k))
			s--
			if l := m.Size(); l != s {
				t.Errorf("key: %d size %d, want %d\n", k, l, s)
			}
		}
	}

}

func TestMapInt64ToUintClear(t *testing.T) {
	m := NewMapInt64ToUint(13)
	for i, k := range signedData {
		m.Put(int64(k), uint(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int64(k), uint(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapInt64ToUintInc(t *testing.T) {
	m := NewMapInt64ToUint(13)
	for i, k := range signedData {
		m.Put(int64(k), uint(unsignedData[i]))
	}
	for i, k := range signedData {
		m.Inc(int64(k))
		m.Inc(int64(k))
		m.Inc(int64(k))
		if g := m.Get(int64(k)); g != uint(unsignedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, uint(unsignedData[i])+3)
		}
	}
	m = NewMapInt64ToUint(13)
	for _, k := range signedData {
		m.Inc(int64(k))
		if g := m.Get(int64(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapInt64ToUintAdd(t *testing.T) {
	m := NewMapInt64ToUint(13)
	for i, k := range signedData {
		m.Put(int64(k), uint(unsignedData[i]))
	}
	for i, k := range signedData {
		m.Add(int64(k), 3)
		if g := m.Get(int64(k)); g != uint(unsignedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, uint(unsignedData[i])+3)
		}
	}
	m = NewMapInt64ToUint(13)
	for _, k := range signedData {
		m.Add(int64(k), 42)
		if g := m.Get(int64(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapInt64ToUintModify(t *testing.T) {
	m := NewMapInt64ToUint(13)
	for i, k := range signedData {
		m.Put(int64(k), uint(unsignedData[i]))
	}
	for _, k := range signedData {
		m.Modify(int64(k), func(v *uint) {
			*v += 3
		})
	}
	for i, k := range signedData {
		if g := m.Get(int64(k)); g != uint(unsignedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, uint(unsignedData[i])+3)
		}
	}
}

func TestMapInt64ToUintFind(t *testing.T) {
	m := NewMapInt64ToUint(13)

	low := signedData[:len(signedData)/2]
	hi := signedData[len(signedData)/2:]

	for i, k := range low {
		m.Put(int64(k), uint(unsignedData[i]))
	}
	for i, k := range low {
		v, ok := m.Find(int64(k))
		if !ok {
			t.Errorf("missing value for key %d\n", k)
			continue
		}
		if want := uint(unsignedData[i]); v != want {
			t.Errorf("key %d: got %d, want %d\n", k, v, want)
		}
	}
	for _, k := range hi {
		if v, ok := m.Find(int64(k)); ok {
			t.Errorf("key %d: got %d, want nothing\n", k, v)
		}
	}
}

func TestMapInt64ToUintPut(t *testing.T) {
	m := NewMapInt64ToUint(13)
	for i, k := range signedData {
		m.Put(int64(k), uint(unsignedData[i]))
	}
	for _, k := range signedData {
		m.Put(int64(k), m.Get(int64(k))+3)
	}
	for i, k := range signedData {
		if g := m.Get(int64(k)); g != uint(unsignedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, uint(unsignedData[i])+3)
		}
	}
}

func TestMapInt64ToUintVisit(t *testing.T) {
	m := NewMapInt64ToUint(13)
	for i, k := range signedData {
		m.Put(int64(k), uint(unsignedData[i]))
	}

	n := make(map[int64]uint, len(signedData))

	for i, k := range signedData {
		n[int64(k)] = uint(unsignedData[i])
	}

	m.Visit(func(k int64, v uint) {
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

func TestMapInt64ToUint8Size(t *testing.T) {
	m := NewMapInt64ToUint8(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int64(k), uint8(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int64(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt64ToUint8Contains(t *testing.T) {
	m := NewMapInt64ToUint8(13)
	for i, k := range signedData {
		m.Put(int64(k), uint8(unsignedData[i]))
		if !m.Contains(int64(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int64(k))
		if m.Contains(int64(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt64ToUint8Get(t *testing.T) {
	m := NewMapInt64ToUint8(13)
	for i, k := range signedData {
		m.Put(int64(k), uint8(unsignedData[i]))
		if g := m.Get(int64(k)); g != uint8(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapInt64ToUint8Remove(t *testing.T) {
	m := NewMapInt64ToUint8(13)
	// Used to trigger code path when all keys are
	// hashed to same slot.
	mask := ^(^0 << nextShiftPowerOfTwo(13))
	for i, k := range signedData {
		if int(k)&mask == 0 {
			m.Put(int64(k), uint8(unsignedData[i]))
		}
	}
	s := m.Size()
	for _, k := range signedData {
		if int(k)&mask == 0 {
			m.Remove(int64(k))
			s--
			if l := m.Size(); l != s {
				t.Errorf("key: %d size %d, want %d\n", k, l, s)
			}
		}
	}

}

func TestMapInt64ToUint8Clear(t *testing.T) {
	m := NewMapInt64ToUint8(13)
	for i, k := range signedData {
		m.Put(int64(k), uint8(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int64(k), uint8(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapInt64ToUint8Inc(t *testing.T) {
	m := NewMapInt64ToUint8(13)
	for i, k := range signedData {
		m.Put(int64(k), uint8(unsignedData[i]))
	}
	for i, k := range signedData {
		m.Inc(int64(k))
		m.Inc(int64(k))
		m.Inc(int64(k))
		if g := m.Get(int64(k)); g != uint8(unsignedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, uint8(unsignedData[i])+3)
		}
	}
	m = NewMapInt64ToUint8(13)
	for _, k := range signedData {
		m.Inc(int64(k))
		if g := m.Get(int64(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapInt64ToUint8Add(t *testing.T) {
	m := NewMapInt64ToUint8(13)
	for i, k := range signedData {
		m.Put(int64(k), uint8(unsignedData[i]))
	}
	for i, k := range signedData {
		m.Add(int64(k), 3)
		if g := m.Get(int64(k)); g != uint8(unsignedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, uint8(unsignedData[i])+3)
		}
	}
	m = NewMapInt64ToUint8(13)
	for _, k := range signedData {
		m.Add(int64(k), 42)
		if g := m.Get(int64(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapInt64ToUint8Modify(t *testing.T) {
	m := NewMapInt64ToUint8(13)
	for i, k := range signedData {
		m.Put(int64(k), uint8(unsignedData[i]))
	}
	for _, k := range signedData {
		m.Modify(int64(k), func(v *uint8) {
			*v += 3
		})
	}
	for i, k := range signedData {
		if g := m.Get(int64(k)); g != uint8(unsignedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, uint8(unsignedData[i])+3)
		}
	}
}

func TestMapInt64ToUint8Find(t *testing.T) {
	m := NewMapInt64ToUint8(13)

	low := signedData[:len(signedData)/2]
	hi := signedData[len(signedData)/2:]

	for i, k := range low {
		m.Put(int64(k), uint8(unsignedData[i]))
	}
	for i, k := range low {
		v, ok := m.Find(int64(k))
		if !ok {
			t.Errorf("missing value for key %d\n", k)
			continue
		}
		if want := uint8(unsignedData[i]); v != want {
			t.Errorf("key %d: got %d, want %d\n", k, v, want)
		}
	}
	for _, k := range hi {
		if v, ok := m.Find(int64(k)); ok {
			t.Errorf("key %d: got %d, want nothing\n", k, v)
		}
	}
}

func TestMapInt64ToUint8Put(t *testing.T) {
	m := NewMapInt64ToUint8(13)
	for i, k := range signedData {
		m.Put(int64(k), uint8(unsignedData[i]))
	}
	for _, k := range signedData {
		m.Put(int64(k), m.Get(int64(k))+3)
	}
	for i, k := range signedData {
		if g := m.Get(int64(k)); g != uint8(unsignedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, uint8(unsignedData[i])+3)
		}
	}
}

func TestMapInt64ToUint8Visit(t *testing.T) {
	m := NewMapInt64ToUint8(13)
	for i, k := range signedData {
		m.Put(int64(k), uint8(unsignedData[i]))
	}

	n := make(map[int64]uint8, len(signedData))

	for i, k := range signedData {
		n[int64(k)] = uint8(unsignedData[i])
	}

	m.Visit(func(k int64, v uint8) {
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

func TestMapInt64ToUint16Size(t *testing.T) {
	m := NewMapInt64ToUint16(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int64(k), uint16(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int64(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt64ToUint16Contains(t *testing.T) {
	m := NewMapInt64ToUint16(13)
	for i, k := range signedData {
		m.Put(int64(k), uint16(unsignedData[i]))
		if !m.Contains(int64(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int64(k))
		if m.Contains(int64(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt64ToUint16Get(t *testing.T) {
	m := NewMapInt64ToUint16(13)
	for i, k := range signedData {
		m.Put(int64(k), uint16(unsignedData[i]))
		if g := m.Get(int64(k)); g != uint16(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapInt64ToUint16Remove(t *testing.T) {
	m := NewMapInt64ToUint16(13)
	// Used to trigger code path when all keys are
	// hashed to same slot.
	mask := ^(^0 << nextShiftPowerOfTwo(13))
	for i, k := range signedData {
		if int(k)&mask == 0 {
			m.Put(int64(k), uint16(unsignedData[i]))
		}
	}
	s := m.Size()
	for _, k := range signedData {
		if int(k)&mask == 0 {
			m.Remove(int64(k))
			s--
			if l := m.Size(); l != s {
				t.Errorf("key: %d size %d, want %d\n", k, l, s)
			}
		}
	}

}

func TestMapInt64ToUint16Clear(t *testing.T) {
	m := NewMapInt64ToUint16(13)
	for i, k := range signedData {
		m.Put(int64(k), uint16(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int64(k), uint16(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapInt64ToUint16Inc(t *testing.T) {
	m := NewMapInt64ToUint16(13)
	for i, k := range signedData {
		m.Put(int64(k), uint16(unsignedData[i]))
	}
	for i, k := range signedData {
		m.Inc(int64(k))
		m.Inc(int64(k))
		m.Inc(int64(k))
		if g := m.Get(int64(k)); g != uint16(unsignedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, uint16(unsignedData[i])+3)
		}
	}
	m = NewMapInt64ToUint16(13)
	for _, k := range signedData {
		m.Inc(int64(k))
		if g := m.Get(int64(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapInt64ToUint16Add(t *testing.T) {
	m := NewMapInt64ToUint16(13)
	for i, k := range signedData {
		m.Put(int64(k), uint16(unsignedData[i]))
	}
	for i, k := range signedData {
		m.Add(int64(k), 3)
		if g := m.Get(int64(k)); g != uint16(unsignedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, uint16(unsignedData[i])+3)
		}
	}
	m = NewMapInt64ToUint16(13)
	for _, k := range signedData {
		m.Add(int64(k), 42)
		if g := m.Get(int64(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapInt64ToUint16Modify(t *testing.T) {
	m := NewMapInt64ToUint16(13)
	for i, k := range signedData {
		m.Put(int64(k), uint16(unsignedData[i]))
	}
	for _, k := range signedData {
		m.Modify(int64(k), func(v *uint16) {
			*v += 3
		})
	}
	for i, k := range signedData {
		if g := m.Get(int64(k)); g != uint16(unsignedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, uint16(unsignedData[i])+3)
		}
	}
}

func TestMapInt64ToUint16Find(t *testing.T) {
	m := NewMapInt64ToUint16(13)

	low := signedData[:len(signedData)/2]
	hi := signedData[len(signedData)/2:]

	for i, k := range low {
		m.Put(int64(k), uint16(unsignedData[i]))
	}
	for i, k := range low {
		v, ok := m.Find(int64(k))
		if !ok {
			t.Errorf("missing value for key %d\n", k)
			continue
		}
		if want := uint16(unsignedData[i]); v != want {
			t.Errorf("key %d: got %d, want %d\n", k, v, want)
		}
	}
	for _, k := range hi {
		if v, ok := m.Find(int64(k)); ok {
			t.Errorf("key %d: got %d, want nothing\n", k, v)
		}
	}
}

func TestMapInt64ToUint16Put(t *testing.T) {
	m := NewMapInt64ToUint16(13)
	for i, k := range signedData {
		m.Put(int64(k), uint16(unsignedData[i]))
	}
	for _, k := range signedData {
		m.Put(int64(k), m.Get(int64(k))+3)
	}
	for i, k := range signedData {
		if g := m.Get(int64(k)); g != uint16(unsignedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, uint16(unsignedData[i])+3)
		}
	}
}

func TestMapInt64ToUint16Visit(t *testing.T) {
	m := NewMapInt64ToUint16(13)
	for i, k := range signedData {
		m.Put(int64(k), uint16(unsignedData[i]))
	}

	n := make(map[int64]uint16, len(signedData))

	for i, k := range signedData {
		n[int64(k)] = uint16(unsignedData[i])
	}

	m.Visit(func(k int64, v uint16) {
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

func TestMapInt64ToUint32Size(t *testing.T) {
	m := NewMapInt64ToUint32(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int64(k), uint32(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int64(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt64ToUint32Contains(t *testing.T) {
	m := NewMapInt64ToUint32(13)
	for i, k := range signedData {
		m.Put(int64(k), uint32(unsignedData[i]))
		if !m.Contains(int64(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int64(k))
		if m.Contains(int64(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt64ToUint32Get(t *testing.T) {
	m := NewMapInt64ToUint32(13)
	for i, k := range signedData {
		m.Put(int64(k), uint32(unsignedData[i]))
		if g := m.Get(int64(k)); g != uint32(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapInt64ToUint32Remove(t *testing.T) {
	m := NewMapInt64ToUint32(13)
	// Used to trigger code path when all keys are
	// hashed to same slot.
	mask := ^(^0 << nextShiftPowerOfTwo(13))
	for i, k := range signedData {
		if int(k)&mask == 0 {
			m.Put(int64(k), uint32(unsignedData[i]))
		}
	}
	s := m.Size()
	for _, k := range signedData {
		if int(k)&mask == 0 {
			m.Remove(int64(k))
			s--
			if l := m.Size(); l != s {
				t.Errorf("key: %d size %d, want %d\n", k, l, s)
			}
		}
	}

}

func TestMapInt64ToUint32Clear(t *testing.T) {
	m := NewMapInt64ToUint32(13)
	for i, k := range signedData {
		m.Put(int64(k), uint32(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int64(k), uint32(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapInt64ToUint32Inc(t *testing.T) {
	m := NewMapInt64ToUint32(13)
	for i, k := range signedData {
		m.Put(int64(k), uint32(unsignedData[i]))
	}
	for i, k := range signedData {
		m.Inc(int64(k))
		m.Inc(int64(k))
		m.Inc(int64(k))
		if g := m.Get(int64(k)); g != uint32(unsignedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, uint32(unsignedData[i])+3)
		}
	}
	m = NewMapInt64ToUint32(13)
	for _, k := range signedData {
		m.Inc(int64(k))
		if g := m.Get(int64(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapInt64ToUint32Add(t *testing.T) {
	m := NewMapInt64ToUint32(13)
	for i, k := range signedData {
		m.Put(int64(k), uint32(unsignedData[i]))
	}
	for i, k := range signedData {
		m.Add(int64(k), 3)
		if g := m.Get(int64(k)); g != uint32(unsignedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, uint32(unsignedData[i])+3)
		}
	}
	m = NewMapInt64ToUint32(13)
	for _, k := range signedData {
		m.Add(int64(k), 42)
		if g := m.Get(int64(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapInt64ToUint32Modify(t *testing.T) {
	m := NewMapInt64ToUint32(13)
	for i, k := range signedData {
		m.Put(int64(k), uint32(unsignedData[i]))
	}
	for _, k := range signedData {
		m.Modify(int64(k), func(v *uint32) {
			*v += 3
		})
	}
	for i, k := range signedData {
		if g := m.Get(int64(k)); g != uint32(unsignedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, uint32(unsignedData[i])+3)
		}
	}
}

func TestMapInt64ToUint32Find(t *testing.T) {
	m := NewMapInt64ToUint32(13)

	low := signedData[:len(signedData)/2]
	hi := signedData[len(signedData)/2:]

	for i, k := range low {
		m.Put(int64(k), uint32(unsignedData[i]))
	}
	for i, k := range low {
		v, ok := m.Find(int64(k))
		if !ok {
			t.Errorf("missing value for key %d\n", k)
			continue
		}
		if want := uint32(unsignedData[i]); v != want {
			t.Errorf("key %d: got %d, want %d\n", k, v, want)
		}
	}
	for _, k := range hi {
		if v, ok := m.Find(int64(k)); ok {
			t.Errorf("key %d: got %d, want nothing\n", k, v)
		}
	}
}

func TestMapInt64ToUint32Put(t *testing.T) {
	m := NewMapInt64ToUint32(13)
	for i, k := range signedData {
		m.Put(int64(k), uint32(unsignedData[i]))
	}
	for _, k := range signedData {
		m.Put(int64(k), m.Get(int64(k))+3)
	}
	for i, k := range signedData {
		if g := m.Get(int64(k)); g != uint32(unsignedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, uint32(unsignedData[i])+3)
		}
	}
}

func TestMapInt64ToUint32Visit(t *testing.T) {
	m := NewMapInt64ToUint32(13)
	for i, k := range signedData {
		m.Put(int64(k), uint32(unsignedData[i]))
	}

	n := make(map[int64]uint32, len(signedData))

	for i, k := range signedData {
		n[int64(k)] = uint32(unsignedData[i])
	}

	m.Visit(func(k int64, v uint32) {
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

func TestMapInt64ToUint64Size(t *testing.T) {
	m := NewMapInt64ToUint64(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int64(k), uint64(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int64(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt64ToUint64Contains(t *testing.T) {
	m := NewMapInt64ToUint64(13)
	for i, k := range signedData {
		m.Put(int64(k), uint64(unsignedData[i]))
		if !m.Contains(int64(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int64(k))
		if m.Contains(int64(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt64ToUint64Get(t *testing.T) {
	m := NewMapInt64ToUint64(13)
	for i, k := range signedData {
		m.Put(int64(k), uint64(unsignedData[i]))
		if g := m.Get(int64(k)); g != uint64(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
	m.Clear()
	if g := m.Get(0); g != 0 {
		t.Errorf("Got %d, want 0\n", g)
	}
}

func TestMapInt64ToUint64Remove(t *testing.T) {
	m := NewMapInt64ToUint64(13)
	// Used to trigger code path when all keys are
	// hashed to same slot.
	mask := ^(^0 << nextShiftPowerOfTwo(13))
	for i, k := range signedData {
		if int(k)&mask == 0 {
			m.Put(int64(k), uint64(unsignedData[i]))
		}
	}
	s := m.Size()
	for _, k := range signedData {
		if int(k)&mask == 0 {
			m.Remove(int64(k))
			s--
			if l := m.Size(); l != s {
				t.Errorf("key: %d size %d, want %d\n", k, l, s)
			}
		}
	}

}

func TestMapInt64ToUint64Clear(t *testing.T) {
	m := NewMapInt64ToUint64(13)
	for i, k := range signedData {
		m.Put(int64(k), uint64(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int64(k), uint64(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapInt64ToUint64Inc(t *testing.T) {
	m := NewMapInt64ToUint64(13)
	for i, k := range signedData {
		m.Put(int64(k), uint64(unsignedData[i]))
	}
	for i, k := range signedData {
		m.Inc(int64(k))
		m.Inc(int64(k))
		m.Inc(int64(k))
		if g := m.Get(int64(k)); g != uint64(unsignedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, uint64(unsignedData[i])+3)
		}
	}
	m = NewMapInt64ToUint64(13)
	for _, k := range signedData {
		m.Inc(int64(k))
		if g := m.Get(int64(k)); g != 1 {
			t.Errorf("got %d, want 1\n", g)
		}
	}
}

func TestMapInt64ToUint64Add(t *testing.T) {
	m := NewMapInt64ToUint64(13)
	for i, k := range signedData {
		m.Put(int64(k), uint64(unsignedData[i]))
	}
	for i, k := range signedData {
		m.Add(int64(k), 3)
		if g := m.Get(int64(k)); g != uint64(unsignedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, uint64(unsignedData[i])+3)
		}
	}
	m = NewMapInt64ToUint64(13)
	for _, k := range signedData {
		m.Add(int64(k), 42)
		if g := m.Get(int64(k)); g != 42 {
			t.Errorf("got %d, want 42\n", g)
		}
	}
}

func TestMapInt64ToUint64Modify(t *testing.T) {
	m := NewMapInt64ToUint64(13)
	for i, k := range signedData {
		m.Put(int64(k), uint64(unsignedData[i]))
	}
	for _, k := range signedData {
		m.Modify(int64(k), func(v *uint64) {
			*v += 3
		})
	}
	for i, k := range signedData {
		if g := m.Get(int64(k)); g != uint64(unsignedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, uint64(unsignedData[i])+3)
		}
	}
}

func TestMapInt64ToUint64Find(t *testing.T) {
	m := NewMapInt64ToUint64(13)

	low := signedData[:len(signedData)/2]
	hi := signedData[len(signedData)/2:]

	for i, k := range low {
		m.Put(int64(k), uint64(unsignedData[i]))
	}
	for i, k := range low {
		v, ok := m.Find(int64(k))
		if !ok {
			t.Errorf("missing value for key %d\n", k)
			continue
		}
		if want := uint64(unsignedData[i]); v != want {
			t.Errorf("key %d: got %d, want %d\n", k, v, want)
		}
	}
	for _, k := range hi {
		if v, ok := m.Find(int64(k)); ok {
			t.Errorf("key %d: got %d, want nothing\n", k, v)
		}
	}
}

func TestMapInt64ToUint64Put(t *testing.T) {
	m := NewMapInt64ToUint64(13)
	for i, k := range signedData {
		m.Put(int64(k), uint64(unsignedData[i]))
	}
	for _, k := range signedData {
		m.Put(int64(k), m.Get(int64(k))+3)
	}
	for i, k := range signedData {
		if g := m.Get(int64(k)); g != uint64(unsignedData[i])+3 {
			t.Errorf("got %d, want %d\n", g, uint64(unsignedData[i])+3)
		}
	}
}

func TestMapInt64ToUint64Visit(t *testing.T) {
	m := NewMapInt64ToUint64(13)
	for i, k := range signedData {
		m.Put(int64(k), uint64(unsignedData[i]))
	}

	n := make(map[int64]uint64, len(signedData))

	for i, k := range signedData {
		n[int64(k)] = uint64(unsignedData[i])
	}

	m.Visit(func(k int64, v uint64) {
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
