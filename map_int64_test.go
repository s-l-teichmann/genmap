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
			t.Errorf("got %d, want 0\n", g, int(signedData[i])+3)
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
			t.Errorf("got %d, want 0\n", g, int(signedData[i])+3)
		}
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
			t.Errorf("got %d, want 0\n", g, int8(signedData[i])+3)
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
			t.Errorf("got %d, want 0\n", g, int8(signedData[i])+3)
		}
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
			t.Errorf("got %d, want 0\n", g, int16(signedData[i])+3)
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
			t.Errorf("got %d, want 0\n", g, int16(signedData[i])+3)
		}
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
			t.Errorf("got %d, want 0\n", g, int32(signedData[i])+3)
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
			t.Errorf("got %d, want 0\n", g, int32(signedData[i])+3)
		}
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
			t.Errorf("got %d, want 0\n", g, int64(signedData[i])+3)
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
			t.Errorf("got %d, want 0\n", g, int64(signedData[i])+3)
		}
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
			t.Errorf("got %d, want 0\n", g, uint(unsignedData[i])+3)
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
			t.Errorf("got %d, want 0\n", g, uint(unsignedData[i])+3)
		}
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
			t.Errorf("got %d, want 0\n", g, uint8(unsignedData[i])+3)
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
			t.Errorf("got %d, want 0\n", g, uint8(unsignedData[i])+3)
		}
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
			t.Errorf("got %d, want 0\n", g, uint16(unsignedData[i])+3)
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
			t.Errorf("got %d, want 0\n", g, uint16(unsignedData[i])+3)
		}
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
			t.Errorf("got %d, want 0\n", g, uint32(unsignedData[i])+3)
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
			t.Errorf("got %d, want 0\n", g, uint32(unsignedData[i])+3)
		}
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
			t.Errorf("got %d, want 0\n", g, uint64(unsignedData[i])+3)
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
			t.Errorf("got %d, want 0\n", g, uint64(unsignedData[i])+3)
		}
	}
}
