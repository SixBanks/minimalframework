package search

import ()

func (r *Runner) Objective(p *permutation) float64 {
	var sum float64
	n := len(p.Seq