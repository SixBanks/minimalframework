package main

import (
	"fmt"
	"github.com/blitzrk/qap-project/data"
	"github.com/blitzrk/qap-project/search"
	"log"
	"os"
	"runtime"
	"time"
)

var (
	fact   func(int) uint
	logger *log.Logger
)

func init() {
	fact = factorial()

	file, err := os.OpenFile("data.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", os.Stderr, ":", err)
	}
	logger = log.New(file, "", log.Lshortfile)
}

func main() {
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

	logger.Println("Distance matrix:")
	logger.Println(dist)
	logger.Println()
	logger.Println("Flow matrix:")
	logger.Println(flow)

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
		case res :=