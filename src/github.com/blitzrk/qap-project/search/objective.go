package search

import ()

func (r *Runner) Objective(p *permutation) float64 {
	var sum float64
	n := len(p.Seq)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sum += float64(r.Cost.A