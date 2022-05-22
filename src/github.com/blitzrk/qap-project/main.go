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
	l