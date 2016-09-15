// THIS IS A MACHINE GENERATED FILE!
// BE CAREFUL WITH EDITING BY HAND.
//
// This is Free Software covered by the terms of the MIT license.
// See LICENSE file for details.
// (c) 2016 by Sascha L. Teichmann.
// See the full list of contributors in the CONTRIBUTORS file.

package genmap

import "testing"

func TestMapUint16ToIntSize(t *testing.T) {
	m := NewMapUint16ToInt(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint16(k), int(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint16(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint16ToIntContains(t *testing.T) {
	m := NewMapUint16ToInt(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), int(signedData[i]))
		if !m.Contains(uint16(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint16(k))
		if m.Contains(uint16(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint16ToInt8Size(t *testing.T) {
	m := NewMapUint16ToInt8(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint16(k), int8(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint16(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint16ToInt8Contains(t *testing.T) {
	m := NewMapUint16ToInt8(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), int8(signedData[i]))
		if !m.Contains(uint16(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint16(k))
		if m.Contains(uint16(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint16ToInt16Size(t *testing.T) {
	m := NewMapUint16ToInt16(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint16(k), int16(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint16(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint16ToInt16Contains(t *testing.T) {
	m := NewMapUint16ToInt16(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), int16(signedData[i]))
		if !m.Contains(uint16(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint16(k))
		if m.Contains(uint16(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint16ToInt32Size(t *testing.T) {
	m := NewMapUint16ToInt32(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint16(k), int32(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint16(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint16ToInt32Contains(t *testing.T) {
	m := NewMapUint16ToInt32(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), int32(signedData[i]))
		if !m.Contains(uint16(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint16(k))
		if m.Contains(uint16(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint16ToInt64Size(t *testing.T) {
	m := NewMapUint16ToInt64(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint16(k), int64(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint16(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint16ToInt64Contains(t *testing.T) {
	m := NewMapUint16ToInt64(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), int64(signedData[i]))
		if !m.Contains(uint16(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint16(k))
		if m.Contains(uint16(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint16ToUintSize(t *testing.T) {
	m := NewMapUint16ToUint(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint16(k), uint(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint16(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint16ToUintContains(t *testing.T) {
	m := NewMapUint16ToUint(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), uint(unsignedData[i]))
		if !m.Contains(uint16(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint16(k))
		if m.Contains(uint16(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint16ToUint8Size(t *testing.T) {
	m := NewMapUint16ToUint8(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint16(k), uint8(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint16(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint16ToUint8Contains(t *testing.T) {
	m := NewMapUint16ToUint8(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), uint8(unsignedData[i]))
		if !m.Contains(uint16(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint16(k))
		if m.Contains(uint16(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint16ToUint16Size(t *testing.T) {
	m := NewMapUint16ToUint16(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint16(k), uint16(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint16(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint16ToUint16Contains(t *testing.T) {
	m := NewMapUint16ToUint16(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), uint16(unsignedData[i]))
		if !m.Contains(uint16(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint16(k))
		if m.Contains(uint16(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint16ToUint32Size(t *testing.T) {
	m := NewMapUint16ToUint32(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint16(k), uint32(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint16(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint16ToUint32Contains(t *testing.T) {
	m := NewMapUint16ToUint32(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), uint32(unsignedData[i]))
		if !m.Contains(uint16(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint16(k))
		if m.Contains(uint16(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapUint16ToUint64Size(t *testing.T) {
	m := NewMapUint16ToUint64(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range unsignedData {
		m.Put(uint16(k), uint64(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range unsignedData {
		m.Remove(uint16(k))
		if want := len(unsignedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapUint16ToUint64Contains(t *testing.T) {
	m := NewMapUint16ToUint64(13)
	for i, k := range unsignedData {
		m.Put(uint16(k), uint64(unsignedData[i]))
		if !m.Contains(uint16(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range unsignedData {
		m.Remove(uint16(k))
		if m.Contains(uint16(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}
