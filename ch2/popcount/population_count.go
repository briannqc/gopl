package popcount

import "strconv"

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountV2(x uint64) int {
	binary := strconv.FormatUint(x, 2)

	var count int
	for _, bit := range binary {
		if bit == '1' {
			count += 1
		}
	}
	return count
}

func PopCountV3(x uint64) int {
	count := int(x & 1)
	for i := 1; i < 64; i++ {
		if x <= 0 {
			break
		}

		x = x >> 1
		count += int(x & 1)
	}
	return count
}

func PopCountV4(x uint64) int {
	var count int
	for ; x > 0; x = x & (x - 1) {
		count += 1
	}
	return count
}
