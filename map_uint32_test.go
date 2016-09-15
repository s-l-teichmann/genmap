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
