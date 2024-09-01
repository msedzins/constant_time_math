package operations

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBitSize(t *testing.T) {
	require.Equal(t, 32, BitSize[uint32]())
	require.Equal(t, 64, BitSize[uint64]())
	require.Equal(t, 16, BitSize[uint16]())
	require.Equal(t, 8, BitSize[uint8]())
}

func FuzzAddUint32(f *testing.F) {

	f.Add(uint32(0), uint32(0))
	f.Add(uint32(math.MaxUint32), uint32(0))
	f.Add(uint32(math.MaxUint32), uint32(1))

	f.Fuzz(func(t *testing.T, x uint32, y uint32) {
		z := Add(x, y)
		require.Equal(t, x+y, z)
	})
}

func FuzzAddUint8(f *testing.F) {

	f.Add(uint8(0), uint8(0))
	f.Add(uint8(math.MaxUint8), uint8(0))
	f.Add(uint8(math.MaxUint8), uint8(1))

	f.Fuzz(func(t *testing.T, x uint8, y uint8) {
		z := Add(x, y)
		require.Equal(t, x+y, z)
	})
}

func FuzzSubUint32(f *testing.F) {

	f.Add(uint32(0), uint32(0))
	f.Add(uint32(math.MaxUint32), uint32(0))
	f.Add(uint32(1), uint32(math.MaxUint32))

	f.Fuzz(func(t *testing.T, x uint32, y uint32) {
		z := Sub(x, y)
		require.Equal(t, x-y, z)
	})
}

func FuzzSubUint8(f *testing.F) {

	f.Add(uint8(0), uint8(0))
	f.Add(uint8(math.MaxUint8), uint8(0))
	f.Add(uint8(1), uint8(math.MaxUint8))

	f.Fuzz(func(t *testing.T, x uint8, y uint8) {
		z := Sub(x, y)
		require.Equal(t, x-y, z)
	})
}

func BenchmarkAddUint32(b *testing.B) {
	benchmarks := []struct {
		name string
		x, y uint32
	}{
		{"Add1", 0, 0},
		{"Add2", math.MaxUint32, 0},
		{"Add3", 1, 1},
		{"Add4", 292929293, 293939292},
		{"Add5", 1, 293939292},
		{"Add6", 9293, 92},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Add(bm.x, bm.y)
			}
		})
	}
}

func BenchmarkAddUint8(b *testing.B) {
	benchmarks := []struct {
		name string
		x, y uint8
	}{
		{"Add1", 0, 0},
		{"AddMaxUint32And0", math.MaxUint8, 0},
		{"Add3", 1, 1},
		{"Add4", 100, 200},
		{"Add5", 1, 250},
		{"Add6", 250, 0},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Add(bm.x, bm.y)
			}
		})
	}
}

func BenchmarkSubUint32(b *testing.B) {
	benchmarks := []struct {
		name string
		x, y uint32
	}{
		{"Sub1", 0, 0},
		{"Sub2", math.MaxUint32, 0},
		{"Sub3", 1, 1},
		{"Sub4", 292929293, 293939292},
		{"Sub5", 1, 293939292},
		{"Sub6", 9293, 92},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Sub(bm.x, bm.y)
			}
		})
	}
}

func BenchmarkSubUint8(b *testing.B) {
	benchmarks := []struct {
		name string
		x, y uint8
	}{
		{"Sub", 0, 0},
		{"Sub2", math.MaxUint8, 0},
		{"Sub3", 1, 1},
		{"Sub4", 100, 200},
		{"Sub5", 1, 250},
		{"Sub6", 250, 0},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Add(bm.x, bm.y)
			}
		})
	}
}
