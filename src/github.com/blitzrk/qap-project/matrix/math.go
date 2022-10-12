
package matrix

import ()

func (m Matrix) Apply(f func(Element) Element) Matrix {
	res := make(Matrix, len(m))
	for ri, r := range m {
		res[ri] = make([]Element, len(r))
		for ci, c := range r {
			res[ri][ci] = f(c)