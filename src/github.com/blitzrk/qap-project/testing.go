
package main

import (
	"fmt"
	"github.com/blitzrk/qap-project/dat"
	"github.com/blitzrk/qap-project/data"
	"github.com/blitzrk/qap-project/matrix"
	"github.com/blitzrk/qap-project/search"
	"io/ioutil"
	"runtime"
	"time"
)

func AllTests() {
	testQAPLIBData()
	testGen()
	testPermutation()
	testSearch()
	testHash()
}

func testHash() {
	p1 := search.NewPerm([]uint8{1, 2, 4, 3})
	p2 := search.NewPerm([]uint8{4, 1, 2, 3})
	fmt.Println(p1)
	fmt.Println(p2)
}

func testSearch() {
	// Setup data generator
	n := 8
	gen := data.New(n, 100000)

	// Generate data
	dist, err := gen.Distance()
	if err != nil {
		panic(err)
	}
	flow, err := gen.Flow(1 / 3)
	if err != nil {
		panic(err)
	}
	cost, err := dist.Combine(flow)
	if err != nil {
		panic(err)
	}

	// Setup runner
	maxTime := time.NewTimer(15 * time.Minute)
	runner := &search.Runner{
		NumCPU:    runtime.NumCPU(),
		Cost:      cost,
		VarCutoff: 0,
		ProbSize:  fact(n),
	}

	// Run on all 4 cores
	quit := make(chan int)
	results := make(chan *search.Result)
	completed := make(chan bool)
	go runner.Run(quit, results, completed)

loop:
	for {
		select {
		case res := <-results:
			if res != nil {
				fmt.Println(res.Score, res.Perm)
			}
		case <-completed:
			// Bug: may lose last few solutions due to race condition
			fmt.Println("Completed entire search.")
			break loop
		case <-maxTime.C:
			quit <- 1
			fmt.Println("Time out.")
			break loop
		}
	}
}

func testPermutation() {
	fs := search.NewFS(2)
	p1 := search.NewPerm([]uint8{1, 2})
	p2 := search.NewPerm([]uint8{2, 1})
	fs.Store(p1)