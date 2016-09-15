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
			t.Errorf("got %d, want 0\n", g, int(signedData[i])+3)
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
			t.Errorf("got %d, want 0\n", g, int(signedData[i])+3)
		}
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
			t.Errorf("got %d, want 0\n", g, int8(signedData[i])+3)
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
			t.Errorf("got %d, want 0\n", g, int8(signedData[i])+3)
		}
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
			t.Errorf("got %d, want 0\n", g, int16(signedData[i])+3)
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
			t.Errorf("got %d, want 0\n", g, int16(signedData[i])+3)
		}
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
			t.Errorf("got %d, want 0\n", g, int32(signedData[i])+3)
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
			t.Errorf("got %d, want 0\n", g, int32(signedData[i])+3)
		}
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
			t.Errorf("got %d, want 0\n", g, int64(signedData[i])+3)
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
			t.Errorf("got %d, want 0\n", g, int64(signedData[i])+3)
		}
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
			t.Errorf("got %d, want 0\n", g, uint(unsignedData[i])+3)
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
			t.Errorf("got %d, want 0\n", g, uint(unsignedData[i])+3)
		}
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
			t.Errorf("got %d, want 0\n", g, uint8(unsignedData[i])+3)
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
			t.Errorf("got %d, want 0\n", g, uint8(unsignedData[i])+3)
		}
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
			t.Errorf("got %d, want 0\n", g, uint16(unsignedData[i])+3)
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
			t.Errorf("got %d, want 0\n", g, uint16(unsignedData[i])+3)
		}
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
			t.Errorf("got %d, want 0\n", g, uint32(unsignedData[i])+3)
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
			t.Errorf("got %d, want 0\n", g, uint32(unsignedData[i])+3)
		}
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
			t.Errorf("got %d, want 0\n", g, uint64(unsignedData[i])+3)
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
			t.Errorf("got %d, want 0\n", g, uint64(unsignedData[i])+3)
		}
	}
}
