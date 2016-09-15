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
}
