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
