package dat

import (
	"bytes"
	"github.com/blitzrk/qap-project/matrix"
	"strconv"
	"strings"
)

// Reads a .dat file where the entries are matricies of numbers
// separated by empty new lines and individually aligned with
// whitespace between row entries and newlines between rows
func Read(file []byte) []matrix.Matrix {
	data := make([]matrix.Matrix, 0