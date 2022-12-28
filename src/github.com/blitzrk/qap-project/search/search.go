
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