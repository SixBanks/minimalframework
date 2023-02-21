
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