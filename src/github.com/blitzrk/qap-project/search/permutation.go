
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
			perm := make([]uint8, p.length)
			copy(perm, p.Seq)
			perm[j], perm[i] = p.Seq[i], p.Seq[j]
			perms = append(perms, NewPerm(perm))
		}
	}

	return perms
}

// Returns the next permutation in a 2-exchange neighborhood of p
func (p *permutation) NextNeighbor() *permutation {
	// Cycle position 1
	p.j++
	if p.j == p.length {
		p.i++
		if p.i == p.length {
			p.i = 0
		}
		p.j = 0
	}

	// Perform swaps
	s := make([]uint8, p.length)
	copy(s, p.Seq)
	s[p.j], s[p.i] = s[p.i], s[p.j]

	return NewPerm(s)
}

// After extensive research, no efficient algorithm for enumerating all permutations within
// a given Hamming distance could be found. As such, an approximation through sampling is used.
//
// IRRELEVANT: The cardinality for n=13, d=2 is 78. For d=3, it's 1,352 and for d=4, it's 15,093.
// An increase of 1 in the Hamming distance appears to approximately lead to an order of magnitude
// increase of 1 near n=13.
//
// Returns a random permutations within approximate Hamming distance d
func (p *permutation) NextHamming(d int) *permutation {
	if d < 2 {
		panic(errors.New("No permutations have a Hamming distance less than 2"))
		return nil