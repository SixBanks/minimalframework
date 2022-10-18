
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
