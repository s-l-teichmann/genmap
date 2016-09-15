// THIS IS A MACHINE GENERATED FILE!
// BE CAREFUL WITH EDITING BY HAND.
//
// This is Free Software covered by the terms of the MIT license.
// See LICENSE file for details.
// (c) 2016 by Sascha L. Teichmann.
// See the full list of contributors in the CONTRIBUTORS file.

package genmap

import "testing"

func TestMapUint64ToIntSize(t *testing.T) {
	m := NewMapUint64ToInt(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint64(k), int(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint64(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint64ToIntContains(t *testing.T) {
	m := NewMapUint64ToInt(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), int(signedData[i]))
		if !m.Contains(uint64(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint64(k))
		if m.Contains(uint64(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint64ToIntGet(t *testing.T) {
	m := NewMapUint64ToInt(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), int(signedData[i]))
		if g := m.Get(uint64(k)); g != int(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
}

func TestMapUint64ToIntClear(t *testing.T) {
	m := NewMapUint64ToInt(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), int(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint64(k), int(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUint64ToInt8Size(t *testing.T) {
	m := NewMapUint64ToInt8(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint64(k), int8(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint64(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint64ToInt8Contains(t *testing.T) {
	m := NewMapUint64ToInt8(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), int8(signedData[i]))
		if !m.Contains(uint64(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint64(k))
		if m.Contains(uint64(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint64ToInt8Get(t *testing.T) {
	m := NewMapUint64ToInt8(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), int8(signedData[i]))
		if g := m.Get(uint64(k)); g != int8(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
}

func TestMapUint64ToInt8Clear(t *testing.T) {
	m := NewMapUint64ToInt8(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), int8(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint64(k), int8(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUint64ToInt16Size(t *testing.T) {
	m := NewMapUint64ToInt16(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint64(k), int16(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint64(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint64ToInt16Contains(t *testing.T) {
	m := NewMapUint64ToInt16(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), int16(signedData[i]))
		if !m.Contains(uint64(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint64(k))
		if m.Contains(uint64(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint64ToInt16Get(t *testing.T) {
	m := NewMapUint64ToInt16(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), int16(signedData[i]))
		if g := m.Get(uint64(k)); g != int16(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
}

func TestMapUint64ToInt16Clear(t *testing.T) {
	m := NewMapUint64ToInt16(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), int16(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint64(k), int16(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUint64ToInt32Size(t *testing.T) {
	m := NewMapUint64ToInt32(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint64(k), int32(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint64(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint64ToInt32Contains(t *testing.T) {
	m := NewMapUint64ToInt32(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), int32(signedData[i]))
		if !m.Contains(uint64(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint64(k))
		if m.Contains(uint64(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint64ToInt32Get(t *testing.T) {
	m := NewMapUint64ToInt32(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), int32(signedData[i]))
		if g := m.Get(uint64(k)); g != int32(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
}

func TestMapUint64ToInt32Clear(t *testing.T) {
	m := NewMapUint64ToInt32(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), int32(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint64(k), int32(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUint64ToInt64Size(t *testing.T) {
	m := NewMapUint64ToInt64(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint64(k), int64(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint64(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint64ToInt64Contains(t *testing.T) {
	m := NewMapUint64ToInt64(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), int64(signedData[i]))
		if !m.Contains(uint64(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint64(k))
		if m.Contains(uint64(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint64ToInt64Get(t *testing.T) {
	m := NewMapUint64ToInt64(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), int64(signedData[i]))
		if g := m.Get(uint64(k)); g != int64(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
}

func TestMapUint64ToInt64Clear(t *testing.T) {
	m := NewMapUint64ToInt64(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), int64(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint64(k), int64(signedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUint64ToUintSize(t *testing.T) {
	m := NewMapUint64ToUint(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint64(k), uint(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint64(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint64ToUintContains(t *testing.T) {
	m := NewMapUint64ToUint(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), uint(unsignedData[i]))
		if !m.Contains(uint64(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint64(k))
		if m.Contains(uint64(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint64ToUintGet(t *testing.T) {
	m := NewMapUint64ToUint(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), uint(unsignedData[i]))
		if g := m.Get(uint64(k)); g != uint(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
}

func TestMapUint64ToUintClear(t *testing.T) {
	m := NewMapUint64ToUint(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), uint(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint64(k), uint(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUint64ToUint8Size(t *testing.T) {
	m := NewMapUint64ToUint8(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint64(k), uint8(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint64(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint64ToUint8Contains(t *testing.T) {
	m := NewMapUint64ToUint8(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), uint8(unsignedData[i]))
		if !m.Contains(uint64(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint64(k))
		if m.Contains(uint64(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint64ToUint8Get(t *testing.T) {
	m := NewMapUint64ToUint8(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), uint8(unsignedData[i]))
		if g := m.Get(uint64(k)); g != uint8(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
}

func TestMapUint64ToUint8Clear(t *testing.T) {
	m := NewMapUint64ToUint8(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), uint8(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint64(k), uint8(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUint64ToUint16Size(t *testing.T) {
	m := NewMapUint64ToUint16(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint64(k), uint16(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint64(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint64ToUint16Contains(t *testing.T) {
	m := NewMapUint64ToUint16(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), uint16(unsignedData[i]))
		if !m.Contains(uint64(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint64(k))
		if m.Contains(uint64(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint64ToUint16Get(t *testing.T) {
	m := NewMapUint64ToUint16(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), uint16(unsignedData[i]))
		if g := m.Get(uint64(k)); g != uint16(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
}

func TestMapUint64ToUint16Clear(t *testing.T) {
	m := NewMapUint64ToUint16(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), uint16(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint64(k), uint16(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUint64ToUint32Size(t *testing.T) {
	m := NewMapUint64ToUint32(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint64(k), uint32(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint64(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint64ToUint32Contains(t *testing.T) {
	m := NewMapUint64ToUint32(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), uint32(unsignedData[i]))
		if !m.Contains(uint64(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint64(k))
		if m.Contains(uint64(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint64ToUint32Get(t *testing.T) {
	m := NewMapUint64ToUint32(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), uint32(unsignedData[i]))
		if g := m.Get(uint64(k)); g != uint32(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
}

func TestMapUint64ToUint32Clear(t *testing.T) {
	m := NewMapUint64ToUint32(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), uint32(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint64(k), uint32(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}

func TestMapUint64ToUint64Size(t *testing.T) {
	m := NewMapUint64ToUint64(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint64(k), uint64(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint64(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint64ToUint64Contains(t *testing.T) {
	m := NewMapUint64ToUint64(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), uint64(unsignedData[i]))
		if !m.Contains(uint64(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint64(k))
		if m.Contains(uint64(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint64ToUint64Get(t *testing.T) {
	m := NewMapUint64ToUint64(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), uint64(unsignedData[i]))
		if g := m.Get(uint64(k)); g != uint64(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
}

func TestMapUint64ToUint64Clear(t *testing.T) {
	m := NewMapUint64ToUint64(13)
	for i, k := range unsignedData {
		m.Put(uint64(k), uint64(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint64(k), uint64(unsignedData[i]))
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
}
