package test

import (
	"testing"
	"log"
)

func Test_shift(t *testing.T) {
	var starBit uint64
	starBit = 1 << 63
	//log.Print(starBit)
	log.Printf("%b",starBit)
	log.Println(len("1000000000000000000000000000000000000000000000000000000000000000"))
}
