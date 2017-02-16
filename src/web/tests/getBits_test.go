package test

import (
	"testing"
	"math"
	"log"
)

func Test_getBits(t *testing.T) {
	log.Printf("%b",getBits(0,59,1))
}

func getBits(min, max, step uint) uint64 {
	var bits uint64

	// If step is 1, use shifts.
	if step == 1 {
		return ^(math.MaxUint64 << (max + 1)) & (math.MaxUint64 << min)
	}

	// Else, use a simple loop.
	for i := min; i <= max; i += step {
		bits |= 1 << i
	}
	return bits
}
