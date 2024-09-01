// Package operations provides basic constant-time arithmetic operations for unsigned integers.
package operations

import "unsafe"

type UnsignedNumber interface {
	uint64 | uint32 | uint16 | uint8 | uint
}

func BitSize[T any]() int {
	var zero T
	return int(unsafe.Sizeof(zero) * 8)
}

// Add returns the sum of x and y in constant time.
func Add[V UnsignedNumber](x, y V) V {
	var carry V
	var result V
	for i := 0; i < BitSize[V](); i++ {
		bitX := (x >> i) & 1
		bitY := (y >> i) & 1
		sum := bitX ^ bitY ^ carry
		carry = (bitX & bitY) | (carry & (bitX ^ bitY))
		result |= (sum << i)
	}
	return result
}

// Sub returns the difference of x and y in constant time.
func Sub[V UnsignedNumber](x, y V) V {
	var borrow V
	var result V
	for i := 0; i < BitSize[V](); i++ {
		bitX := (x >> i) & 1
		bitY := (y >> i) & 1
		diff := bitX ^ bitY ^ borrow
		borrow = (^bitX & bitY) | (borrow & (^bitX ^ bitY))
		result |= (diff << i)
	}
	return result
}
