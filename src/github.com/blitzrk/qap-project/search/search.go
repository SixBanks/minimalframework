
package search

import (
	"github.com/blitzrk/qap-project/matrix"
	"log"
	"math"
	"os"
	"runtime"
)

var (
	logger *log.Logger
)

func init() {
	file, err := os.OpenFile("go.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", os.Stderr, ":", err)
	}
	logger = log.New(file, "logger: ", log.Lshortfile)
}

type Runner struct {
	NumCPU    int
	Cost      matrix.Matrix4D
	VarCutoff float64
	ProbSize  uint
	fs        *fastStore
}

func (r *Runner) Run(stop <-chan int, resultChan chan<- *Result, complete chan<- bool) {
	// maximize CPU usage
	runtime.GOMAXPROCS(r.NumCPU)
	limit := make(chan int, r.NumCPU)

	r.fs = NewFS(r.ProbSize)
	n := len(r.Cost)
	done := make(chan *Result)

loop:
	for {
		select {
		case limit <- 1:
			p := RandPerm(n)
			go r.search(p, done)
		case res := <-done:
			resultChan <- res
			<-limit
			// Check if entire solution space traversed