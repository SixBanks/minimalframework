
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
// Returns all permutations within a 2-exchange neighborhood
func (p *permutation) Neighborhood() []*permutation {
	n := p.length
	perms := make([]*permutation, 0, n*(n-1)/2)

	// Find 2-exchange neighborhood
	for i := 0; i < p.length; i++ {
		for j := i + 1; j < p.length; j++ {