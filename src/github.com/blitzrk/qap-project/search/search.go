
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
			if r.fs.Full() {
				// Let remaining processes finish
				for i := 0; i < r.NumCPU-1; i++ {
					resultChan <- <-done
					<-limit
				}
				complete <- true
			}
		case <-stop:
			break loop
		}
	}
}

type Result struct {
	Score float64
	Perm  []uint8
}

type runResult struct {
	Perm   *permutation
	Score  float64
	Opt    bool
	Var    float64
	Center *permutation
	FinalR int
}

func (r *Runner) search(perm *permutation, done chan<- *Result) {
	// Check if already been to the proposed next step
	if r.fs.Test(perm) {
		// No need to continue further
		done <- nil
		return
	}
	r.fs.Store(perm)

	collect := make(chan *runResult)
	go r.findBestNeighbor(perm, collect)

	// Change what gets sent here
	result := <-collect
	go r.interpret(result, done)
}

func (r *Runner) searchHamming(perm *permutation, dist int, done chan<- *Result) {
	collect := make(chan *runResult)
	go r.sampleHammingRegion(perm, dist, collect)

	// Change what gets sent here
	result := <-collect
	go r.interpret(result, done)
}

// Find best permutation
func (r *Runner) findBestNeighbor(center *permutation, done chan<- *runResult) {
	n := len(center.Seq)
	size := (n * (n - 1) / 2)

	var bestPerm *permutation
	bestScore := math.Inf(1)
	scores := make([]float64, size)

	for i := 0; i < size; i++ {
		neighbor := center.NextNeighbor()
		r.fs.Store(neighbor)
		score := r.Objective(neighbor)
		scores[i] = score

		if score < bestScore {
			bestScore = score
			bestPerm = neighbor
		}
	}