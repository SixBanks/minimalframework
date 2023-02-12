
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
