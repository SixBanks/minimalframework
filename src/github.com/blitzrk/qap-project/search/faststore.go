
package search

import (
	"math/big"
)

type fastStore struct {
	big.Int
	size uint
}

func NewFS(size uint) *fastStore {
	return &fastStore{big.Int{}, size}
}