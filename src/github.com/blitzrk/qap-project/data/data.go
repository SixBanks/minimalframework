
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