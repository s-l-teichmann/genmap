// THIS IS A MACHINE GENERATED FILE!
// BE CAREFUL WITH EDITING BY HAND.
//
// This is Free Software covered by the terms of the MIT license.
// See LICENSE file for details.
// (c) 2016 by Sascha L. Teichmann.
// See the full list of contributors in the CONTRIBUTORS file.

package genmap

import "testing"

func TestMapInt16ToIntSize(t *testing.T) {
	m := NewMapInt16ToInt(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int16(k), int(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int16(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt16ToIntContains(t *testing.T) {
	m := NewMapInt16ToInt(13)
	for i, k := range signedData {
		m.Put(int16(k), int(signedData[i]))
		if !m.Contains(int16(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int16(k))
		if m.Contains(int16(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt16ToIntGet(t *testing.T) {
	m := NewMapInt16ToInt(13)
	for i, k := range signedData {
		m.Put(int16(k), int(signedData[i]))
		if g := m.Get(int16(k)); g != int(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
}

func TestMapInt16ToInt8Size(t *testing.T) {
	m := NewMapInt16ToInt8(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int16(k), int8(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int16(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt16ToInt8Contains(t *testing.T) {
	m := NewMapInt16ToInt8(13)
	for i, k := range signedData {
		m.Put(int16(k), int8(signedData[i]))
		if !m.Contains(int16(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int16(k))
		if m.Contains(int16(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt16ToInt8Get(t *testing.T) {
	m := NewMapInt16ToInt8(13)
	for i, k := range signedData {
		m.Put(int16(k), int8(signedData[i]))
		if g := m.Get(int16(k)); g != int8(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
}

func TestMapInt16ToInt16Size(t *testing.T) {
	m := NewMapInt16ToInt16(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int16(k), int16(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int16(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt16ToInt16Contains(t *testing.T) {
	m := NewMapInt16ToInt16(13)
	for i, k := range signedData {
		m.Put(int16(k), int16(signedData[i]))
		if !m.Contains(int16(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int16(k))
		if m.Contains(int16(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt16ToInt16Get(t *testing.T) {
	m := NewMapInt16ToInt16(13)
	for i, k := range signedData {
		m.Put(int16(k), int16(signedData[i]))
		if g := m.Get(int16(k)); g != int16(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
}

func TestMapInt16ToInt32Size(t *testing.T) {
	m := NewMapInt16ToInt32(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int16(k), int32(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int16(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt16ToInt32Contains(t *testing.T) {
	m := NewMapInt16ToInt32(13)
	for i, k := range signedData {
		m.Put(int16(k), int32(signedData[i]))
		if !m.Contains(int16(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int16(k))
		if m.Contains(int16(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt16ToInt32Get(t *testing.T) {
	m := NewMapInt16ToInt32(13)
	for i, k := range signedData {
		m.Put(int16(k), int32(signedData[i]))
		if g := m.Get(int16(k)); g != int32(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
}

func TestMapInt16ToInt64Size(t *testing.T) {
	m := NewMapInt16ToInt64(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int16(k), int64(signedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int16(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt16ToInt64Contains(t *testing.T) {
	m := NewMapInt16ToInt64(13)
	for i, k := range signedData {
		m.Put(int16(k), int64(signedData[i]))
		if !m.Contains(int16(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int16(k))
		if m.Contains(int16(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt16ToInt64Get(t *testing.T) {
	m := NewMapInt16ToInt64(13)
	for i, k := range signedData {
		m.Put(int16(k), int64(signedData[i]))
		if g := m.Get(int16(k)); g != int64(signedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, signedData[i])
		}
	}
}

func TestMapInt16ToUintSize(t *testing.T) {
	m := NewMapInt16ToUint(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int16(k), uint(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int16(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt16ToUintContains(t *testing.T) {
	m := NewMapInt16ToUint(13)
	for i, k := range signedData {
		m.Put(int16(k), uint(unsignedData[i]))
		if !m.Contains(int16(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int16(k))
		if m.Contains(int16(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt16ToUintGet(t *testing.T) {
	m := NewMapInt16ToUint(13)
	for i, k := range signedData {
		m.Put(int16(k), uint(unsignedData[i]))
		if g := m.Get(int16(k)); g != uint(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
}

func TestMapInt16ToUint8Size(t *testing.T) {
	m := NewMapInt16ToUint8(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int16(k), uint8(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int16(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt16ToUint8Contains(t *testing.T) {
	m := NewMapInt16ToUint8(13)
	for i, k := range signedData {
		m.Put(int16(k), uint8(unsignedData[i]))
		if !m.Contains(int16(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int16(k))
		if m.Contains(int16(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt16ToUint8Get(t *testing.T) {
	m := NewMapInt16ToUint8(13)
	for i, k := range signedData {
		m.Put(int16(k), uint8(unsignedData[i]))
		if g := m.Get(int16(k)); g != uint8(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
}

func TestMapInt16ToUint16Size(t *testing.T) {
	m := NewMapInt16ToUint16(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int16(k), uint16(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int16(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt16ToUint16Contains(t *testing.T) {
	m := NewMapInt16ToUint16(13)
	for i, k := range signedData {
		m.Put(int16(k), uint16(unsignedData[i]))
		if !m.Contains(int16(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int16(k))
		if m.Contains(int16(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt16ToUint16Get(t *testing.T) {
	m := NewMapInt16ToUint16(13)
	for i, k := range signedData {
		m.Put(int16(k), uint16(unsignedData[i]))
		if g := m.Get(int16(k)); g != uint16(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
}

func TestMapInt16ToUint32Size(t *testing.T) {
	m := NewMapInt16ToUint32(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int16(k), uint32(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int16(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt16ToUint32Contains(t *testing.T) {
	m := NewMapInt16ToUint32(13)
	for i, k := range signedData {
		m.Put(int16(k), uint32(unsignedData[i]))
		if !m.Contains(int16(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int16(k))
		if m.Contains(int16(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt16ToUint32Get(t *testing.T) {
	m := NewMapInt16ToUint32(13)
	for i, k := range signedData {
		m.Put(int16(k), uint32(unsignedData[i]))
		if g := m.Get(int16(k)); g != uint32(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
}

func TestMapInt16ToUint64Size(t *testing.T) {
	m := NewMapInt16ToUint64(13)
	if m.Size() != 0 {
		t.Errorf("map size is %d, want 0\n", m.Size())
	}
	for i, k := range signedData {
		m.Put(int16(k), uint64(unsignedData[i]))
		if m.Size() != i+1 {
			t.Errorf("map size is %d, want %d\n", m.Size(), i+1)
		}
	}
	for i, k := range signedData {
		m.Remove(int16(k))
		if want := len(signedData) - (i + 1); m.Size() != want {
			t.Errorf("map size is %d, want %d\n", m.Size(), want)
		}
	}
}

func TestMapInt16ToUint64Contains(t *testing.T) {
	m := NewMapInt16ToUint64(13)
	for i, k := range signedData {
		m.Put(int16(k), uint64(unsignedData[i]))
		if !m.Contains(int16(k)) {
			t.Errorf("map does not contain %d\n", k)
		}
	}
	for _, k := range signedData {
		m.Remove(int16(k))
		if m.Contains(int16(k)) {
			t.Errorf("map contains %d\n", k)
		}
	}
}

func TestMapInt16ToUint64Get(t *testing.T) {
	m := NewMapInt16ToUint64(13)
	for i, k := range signedData {
		m.Put(int16(k), uint64(unsignedData[i]))
		if g := m.Get(int16(k)); g != uint64(unsignedData[i]) {
			t.Errorf("Get(%d) = %d, want %d\n", k, g, unsignedData[i])
		}
	}
}
