// This is Free Software covered by the terms of the MIT license.
// See LICENSE file for details.
// (c) 2016 by the contributors.
// See CONTRIBUTORS for a full list.

package genmap

func nextShiftPowerOfTwo(x int) uint {
	var p uint
	for ; 1<<p < x; p++ {
	}
	return p
}

func maxFill(s int) int {
	return int(float32(s) * 0.75)
}
