// This is Free Software covered by the terms of the MIT license.
// See LICENSE file for details.
// (c) 2016 by Sascha L. Teichmann.
// See the full list of contributors in the CONTRIBUTORS file.

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
