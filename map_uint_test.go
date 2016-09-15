// THIS IS A MACHINE GENERATED FILE!
// BE CAREFUL WITH EDITING BY HAND.
//
// This is Free Software covered by the terms of the MIT license.
// See LICENSE file for details.
// (c) 2016 by Sascha L. Teichmann.
// See the full list of contributors in the CONTRIBUTORS file.

package genmap

import "testing"

func TestMapUintToIntSize(t *testing.T) {
	m := NewMapUintToInt(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint(k), int(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUintToIntContains(t *testing.T) {
	m := NewMapUintToInt(13)
	for i, k := range unsignedData {
		m.Put(uint(k), int(signedData[i]))
		if !m.Contains(uint(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint(k))
		if m.Contains(uint(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUintToIntGet(t *testing.T) {
	m := NewMapUintToInt(13)
	for i, k := range unsignedData {
		m.Put(uint(k), int(signedData[i]))
		if g := m.Get(uint(k)); g != int(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
}

func TestMapUintToInt8Size(t *testing.T) {
	m := NewMapUintToInt8(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint(k), int8(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUintToInt8Contains(t *testing.T) {
	m := NewMapUintToInt8(13)
	for i, k := range unsignedData {
		m.Put(uint(k), int8(signedData[i]))
		if !m.Contains(uint(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint(k))
		if m.Contains(uint(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUintToInt8Get(t *testing.T) {
	m := NewMapUintToInt8(13)
	for i, k := range unsignedData {
		m.Put(uint(k), int8(signedData[i]))
		if g := m.Get(uint(k)); g != int8(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
}

func TestMapUintToInt16Size(t *testing.T) {
	m := NewMapUintToInt16(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint(k), int16(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUintToInt16Contains(t *testing.T) {
	m := NewMapUintToInt16(13)
	for i, k := range unsignedData {
		m.Put(uint(k), int16(signedData[i]))
		if !m.Contains(uint(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint(k))
		if m.Contains(uint(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUintToInt16Get(t *testing.T) {
	m := NewMapUintToInt16(13)
	for i, k := range unsignedData {
		m.Put(uint(k), int16(signedData[i]))
		if g := m.Get(uint(k)); g != int16(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
}

func TestMapUintToInt32Size(t *testing.T) {
	m := NewMapUintToInt32(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint(k), int32(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUintToInt32Contains(t *testing.T) {
	m := NewMapUintToInt32(13)
	for i, k := range unsignedData {
		m.Put(uint(k), int32(signedData[i]))
		if !m.Contains(uint(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint(k))
		if m.Contains(uint(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUintToInt32Get(t *testing.T) {
	m := NewMapUintToInt32(13)
	for i, k := range unsignedData {
		m.Put(uint(k), int32(signedData[i]))
		if g := m.Get(uint(k)); g != int32(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
}

func TestMapUintToInt64Size(t *testing.T) {
	m := NewMapUintToInt64(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint(k), int64(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUintToInt64Contains(t *testing.T) {
	m := NewMapUintToInt64(13)
	for i, k := range unsignedData {
		m.Put(uint(k), int64(signedData[i]))
		if !m.Contains(uint(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint(k))
		if m.Contains(uint(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUintToInt64Get(t *testing.T) {
	m := NewMapUintToInt64(13)
	for i, k := range unsignedData {
		m.Put(uint(k), int64(signedData[i]))
		if g := m.Get(uint(k)); g != int64(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
}

func TestMapUintToUintSize(t *testing.T) {
	m := NewMapUintToUint(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint(k), uint(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUintToUintContains(t *testing.T) {
	m := NewMapUintToUint(13)
	for i, k := range unsignedData {
		m.Put(uint(k), uint(unsignedData[i]))
		if !m.Contains(uint(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint(k))
		if m.Contains(uint(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUintToUintGet(t *testing.T) {
	m := NewMapUintToUint(13)
	for i, k := range unsignedData {
		m.Put(uint(k), uint(unsignedData[i]))
		if g := m.Get(uint(k)); g != uint(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
}

func TestMapUintToUint8Size(t *testing.T) {
	m := NewMapUintToUint8(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint(k), uint8(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUintToUint8Contains(t *testing.T) {
	m := NewMapUintToUint8(13)
	for i, k := range unsignedData {
		m.Put(uint(k), uint8(unsignedData[i]))
		if !m.Contains(uint(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint(k))
		if m.Contains(uint(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUintToUint8Get(t *testing.T) {
	m := NewMapUintToUint8(13)
	for i, k := range unsignedData {
		m.Put(uint(k), uint8(unsignedData[i]))
		if g := m.Get(uint(k)); g != uint8(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
}

func TestMapUintToUint16Size(t *testing.T) {
	m := NewMapUintToUint16(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint(k), uint16(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUintToUint16Contains(t *testing.T) {
	m := NewMapUintToUint16(13)
	for i, k := range unsignedData {
		m.Put(uint(k), uint16(unsignedData[i]))
		if !m.Contains(uint(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint(k))
		if m.Contains(uint(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUintToUint16Get(t *testing.T) {
	m := NewMapUintToUint16(13)
	for i, k := range unsignedData {
		m.Put(uint(k), uint16(unsignedData[i]))
		if g := m.Get(uint(k)); g != uint16(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
}

func TestMapUintToUint32Size(t *testing.T) {
	m := NewMapUintToUint32(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint(k), uint32(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUintToUint32Contains(t *testing.T) {
	m := NewMapUintToUint32(13)
	for i, k := range unsignedData {
		m.Put(uint(k), uint32(unsignedData[i]))
		if !m.Contains(uint(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint(k))
		if m.Contains(uint(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUintToUint32Get(t *testing.T) {
	m := NewMapUintToUint32(13)
	for i, k := range unsignedData {
		m.Put(uint(k), uint32(unsignedData[i]))
		if g := m.Get(uint(k)); g != uint32(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
}

func TestMapUintToUint64Size(t *testing.T) {
	m := NewMapUintToUint64(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint(k), uint64(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUintToUint64Contains(t *testing.T) {
	m := NewMapUintToUint64(13)
	for i, k := range unsignedData {
		m.Put(uint(k), uint64(unsignedData[i]))
		if !m.Contains(uint(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint(k))
		if m.Contains(uint(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUintToUint64Get(t *testing.T) {
	m := NewMapUintToUint64(13)
	for i, k := range unsignedData {
		m.Put(uint(k), uint64(unsignedData[i]))
		if g := m.Get(uint(k)); g != uint64(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
}
