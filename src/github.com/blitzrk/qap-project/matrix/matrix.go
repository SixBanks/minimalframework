
package matrix

import (
	"errors"
	"fmt"
	"strings"
)

var (
	SizeMismatchError = errors.New("Matrix sizes do not match")
	NotSquareError    = errors.New("One or more matrices are not square")
)

type Element float64

type Matrix [][]Element

type Matrix4D [][][][]Element

func (m Matrix) String() string {
	var s []string

	for _, row := range m {
		s = append(s, fmt.Sprint(row))
	}

	return strings.Join(s, "\n")
}

func (m Matrix) At(r, c int) Element {
	return m[r][c]
}

func (m1 Matrix) Combine(m2 Matrix) (Matrix4D, error) {
	if len(m1) != len(m2) {
		return nil, SizeMismatchError
	}

	n := len(m1)