
package data

import (
	"fmt"
	"github.com/blitzrk/qap-project/matrix"
	"math"
	"math/rand"
)

type generator struct {
	n      int
	fscale float64
}

func New(n int, fscale float64) *generator {
	return &generator{n, fscale}
}

func (g *generator) Distance() (matrix.Matrix, error) {
	m := matrix.Matrix(make([][]matrix.Element, g.n))

	for i := 0; i < g.n; i++ {
		m[i] = make([]matrix.Element, g.n)
		for j := 0; j < g.n; j++ {
			m[i][j] = matrix.Element(rand.Float64())
		}
	}

	return m, nil
}

func (g *generator) Flow(spread float64) (matrix.Matrix, error) {
	if spread < 0 || spread >= 1 {
		return nil, fmt.Errorf("Error: spread must be between 0 and 1.")
	}

	// Create Zipf generator
	scale := 1000
	r := rand.New(rand.NewSource(0))
	zipf := rand.NewZipf(r, 1.01, float64(g.n), uint64(scale))

	// Populate frequencies of unigrams
	k := make([]float64, g.n)
	for i := 0; i < g.n; i++ {
		k[i] = float64(zipf.Uint64())
	}

	// Populate ideal bigram matrix
	m := matrix.Matrix(make([][]matrix.Element, g.n))
	for i := 0; i < g.n; i++ {
		m[i] = make([]matrix.Element, g.n)
		for j := 0; j < g.n; j++ {
			e := (rand.Float64() - 0.5) * spread
			m[i][j] = matrix.Element((k[i] * k[j]) * (1 + e))
		}
	}

	// Scale back to 100,000 total freq
	totalF := m.Sum()
	s := g.fscale / totalF
	for i := 0; i < g.n; i++ {
		for j := 0; j < g.n; j++ {
			m[i][j] = matrix.Element(math.Floor(float64(m[i][j]) * s))
		}
	}

	return m, nil
}