
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