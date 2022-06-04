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
		log.Fatalln("Failed to open log file", o