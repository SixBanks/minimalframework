
package search

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
)

var (
	memo []uint64 = []uint64{1}
)

type permutation struct {
	Seq    []uint8
	hash   uint64
	length int
	i      int
	j      int
}

func (p *permutation) String() string {
	return fmt.Sprint(p.hash, ": ", p.Seq)
}

// Create a permutation of 1...n from an int slice
func NewPerm(p []uint8) *permutation {
	h := hash(p, 0)
	return &permutation{p, h, len(p), 0, 0}
}

// Create random permutation of 1...n
func RandPerm(n int) *permutation {
	p := rand.Perm(n)
	r := make([]uint8, len(p))

	for i, v := range p {
		r[i] = uint8(v + 1)
	}

	return NewPerm(r)
}

// DEPRECATED: