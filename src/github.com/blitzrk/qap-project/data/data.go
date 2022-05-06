
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