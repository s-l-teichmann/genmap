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
		if g := m.Get(uint8(k)); g != int(signedData[i]) + 3 {
			t.Errorf("got %d, want 0\n", g, int(signedData[i]) + 3)
		}
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
		if g := m.Get(uint8(k)); g != int8(signedData[i]) + 3 {
			t.Errorf("got %d, want 0\n", g, int8(signedData[i]) + 3)
		}
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
		if g := m.Get(uint8(k)); g != int16(signedData[i]) + 3 {
			t.Errorf("got %d, want 0\n", g, int16(signedData[i]) + 3)
		}
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
		if g := m.Get(uint8(k)); g != int32(signedData[i]) + 3 {
			t.Errorf("got %d, want 0\n", g, int32(signedData[i]) + 3)
		}
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
		if g := m.Get(uint8(k)); g != int64(signedData[i]) + 3 {
			t.Errorf("got %d, want 0\n", g, int64(signedData[i]) + 3)
		}
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
		if g := m.Get(uint8(k)); g != uint(unsignedData[i]) + 3 {
			t.Errorf("got %d, want 0\n", g, uint(unsignedData[i]) + 3)
		}
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
		if g := m.Get(uint8(k)); g != uint8(unsignedData[i]) + 3 {
			t.Errorf("got %d, want 0\n", g, uint8(unsignedData[i]) + 3)
		}
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
		if g := m.Get(uint8(k)); g != uint16(unsignedData[i]) + 3 {
			t.Errorf("got %d, want 0\n", g, uint16(unsignedData[i]) + 3)
		}
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
		if g := m.Get(uint8(k)); g != uint32(unsignedData[i]) + 3 {
			t.Errorf("got %d, want 0\n", g, uint32(unsignedData[i]) + 3)
		}
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
		if g := m.Get(uint8(k)); g != uint64(unsignedData[i]) + 3 {
			t.Errorf("got %d, want 0\n", g, uint64(unsignedData[i]) + 3)
		}
	}
}
