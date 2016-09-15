// THIS IS A MACHINE GENERATED FILE!
// BE CAREFUL WITH EDITING BY HAND.
//
// This is Free Software covered by the terms of the MIT license.
// See LICENSE file for details.
// (c) 2016 by Sascha L. Teichmann.
// See the full list of contributors in the CONTRIBUTORS file.

package genmap

import "testing"

func TestMapInt8ToIntSize(t *testing.T) {
	m := NewMapInt8ToInt(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int8(k), int(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int8(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt8ToIntContains(t *testing.T) {
	m := NewMapInt8ToInt(13)
	for i, k := range signedData {
		m.Put(int8(k), int(signedData[i]))
		if !m.Contains(int8(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int8(k))
		if m.Contains(int8(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt8ToIntGet(t *testing.T) {
	m := NewMapInt8ToInt(13)
	for i, k := range signedData {
		m.Put(int8(k), int(signedData[i]))
		if g := m.Get(int8(k)); g != int(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
}

func TestMapInt8ToInt8Size(t *testing.T) {
	m := NewMapInt8ToInt8(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int8(k), int8(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int8(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt8ToInt8Contains(t *testing.T) {
	m := NewMapInt8ToInt8(13)
	for i, k := range signedData {
		m.Put(int8(k), int8(signedData[i]))
		if !m.Contains(int8(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int8(k))
		if m.Contains(int8(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt8ToInt8Get(t *testing.T) {
	m := NewMapInt8ToInt8(13)
	for i, k := range signedData {
		m.Put(int8(k), int8(signedData[i]))
		if g := m.Get(int8(k)); g != int8(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
}

func TestMapInt8ToInt16Size(t *testing.T) {
	m := NewMapInt8ToInt16(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int8(k), int16(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int8(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt8ToInt16Contains(t *testing.T) {
	m := NewMapInt8ToInt16(13)
	for i, k := range signedData {
		m.Put(int8(k), int16(signedData[i]))
		if !m.Contains(int8(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int8(k))
		if m.Contains(int8(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt8ToInt16Get(t *testing.T) {
	m := NewMapInt8ToInt16(13)
	for i, k := range signedData {
		m.Put(int8(k), int16(signedData[i]))
		if g := m.Get(int8(k)); g != int16(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
}

func TestMapInt8ToInt32Size(t *testing.T) {
	m := NewMapInt8ToInt32(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int8(k), int32(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int8(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt8ToInt32Contains(t *testing.T) {
	m := NewMapInt8ToInt32(13)
	for i, k := range signedData {
		m.Put(int8(k), int32(signedData[i]))
		if !m.Contains(int8(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int8(k))
		if m.Contains(int8(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt8ToInt32Get(t *testing.T) {
	m := NewMapInt8ToInt32(13)
	for i, k := range signedData {
		m.Put(int8(k), int32(signedData[i]))
		if g := m.Get(int8(k)); g != int32(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
}

func TestMapInt8ToInt64Size(t *testing.T) {
	m := NewMapInt8ToInt64(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int8(k), int64(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int8(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt8ToInt64Contains(t *testing.T) {
	m := NewMapInt8ToInt64(13)
	for i, k := range signedData {
		m.Put(int8(k), int64(signedData[i]))
		if !m.Contains(int8(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int8(k))
		if m.Contains(int8(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt8ToInt64Get(t *testing.T) {
	m := NewMapInt8ToInt64(13)
	for i, k := range signedData {
		m.Put(int8(k), int64(signedData[i]))
		if g := m.Get(int8(k)); g != int64(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
}

func TestMapInt8ToUintSize(t *testing.T) {
	m := NewMapInt8ToUint(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int8(k), uint(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int8(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt8ToUintContains(t *testing.T) {
	m := NewMapInt8ToUint(13)
	for i, k := range signedData {
		m.Put(int8(k), uint(unsignedData[i]))
		if !m.Contains(int8(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int8(k))
		if m.Contains(int8(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt8ToUintGet(t *testing.T) {
	m := NewMapInt8ToUint(13)
	for i, k := range signedData {
		m.Put(int8(k), uint(unsignedData[i]))
		if g := m.Get(int8(k)); g != uint(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
}

func TestMapInt8ToUint8Size(t *testing.T) {
	m := NewMapInt8ToUint8(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int8(k), uint8(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int8(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt8ToUint8Contains(t *testing.T) {
	m := NewMapInt8ToUint8(13)
	for i, k := range signedData {
		m.Put(int8(k), uint8(unsignedData[i]))
		if !m.Contains(int8(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int8(k))
		if m.Contains(int8(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt8ToUint8Get(t *testing.T) {
	m := NewMapInt8ToUint8(13)
	for i, k := range signedData {
		m.Put(int8(k), uint8(unsignedData[i]))
		if g := m.Get(int8(k)); g != uint8(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
}

func TestMapInt8ToUint16Size(t *testing.T) {
	m := NewMapInt8ToUint16(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int8(k), uint16(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int8(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt8ToUint16Contains(t *testing.T) {
	m := NewMapInt8ToUint16(13)
	for i, k := range signedData {
		m.Put(int8(k), uint16(unsignedData[i]))
		if !m.Contains(int8(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int8(k))
		if m.Contains(int8(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt8ToUint16Get(t *testing.T) {
	m := NewMapInt8ToUint16(13)
	for i, k := range signedData {
		m.Put(int8(k), uint16(unsignedData[i]))
		if g := m.Get(int8(k)); g != uint16(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
}

func TestMapInt8ToUint32Size(t *testing.T) {
	m := NewMapInt8ToUint32(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int8(k), uint32(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int8(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt8ToUint32Contains(t *testing.T) {
	m := NewMapInt8ToUint32(13)
	for i, k := range signedData {
		m.Put(int8(k), uint32(unsignedData[i]))
		if !m.Contains(int8(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int8(k))
		if m.Contains(int8(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt8ToUint32Get(t *testing.T) {
	m := NewMapInt8ToUint32(13)
	for i, k := range signedData {
		m.Put(int8(k), uint32(unsignedData[i]))
		if g := m.Get(int8(k)); g != uint32(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
}

func TestMapInt8ToUint64Size(t *testing.T) {
	m := NewMapInt8ToUint64(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int8(k), uint64(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int8(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt8ToUint64Contains(t *testing.T) {
	m := NewMapInt8ToUint64(13)
	for i, k := range signedData {
		m.Put(int8(k), uint64(unsignedData[i]))
		if !m.Contains(int8(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int8(k))
		if m.Contains(int8(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt8ToUint64Get(t *testing.T) {
	m := NewMapInt8ToUint64(13)
	for i, k := range signedData {
		m.Put(int8(k), uint64(unsignedData[i]))
		if g := m.Get(int8(k)); g != uint64(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
}
