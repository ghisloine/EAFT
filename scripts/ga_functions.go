package scripts

import (
	"bytes"
	"ga_tuner/utils"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func MutNormalFloat64(genome []float64, rate float64, rng *rand.Rand) {
	for i := range genome {
		// Flip a coin and decide to mutate or not
		if rng.Float64() < rate {
			//genome[i] += rng.NormFloat64() * genome[i]
			if genome[i] < 0.5 {
				genome[i] = 0
			} else {
				genome[i] = 1
			}
		}
	}
}

// Returns a Binary Vector with just only 0 - 1.
func InitBinaryFloat64(n uint, lower, upper float64, rng *rand.Rand) (floats []float64) {
	floats = make([]float64, n)
	for i := range floats {
		floats[i] = lower + float64(rng.Intn(int(upper)))
	}
	return
}

func CompileCode(cmd string, id string) (Total float64) {
	// COMPILE
	command := strings.Split(cmd, " ")
	app := os.Args[2]
	out_compile, err := exec.Command(app, command...).Output()
	if err != nil {
		log.Print(string(out_compile))
		log.Fatal(err)
	}

	// EXECUTION
	exec_file := filepath.Join(utils.ResultsPath, os.Args[1], "bin", id)

	command_exec := exec.Command(exec_file)
	var out_exec bytes.Buffer
	// set the output to our variable
	command_exec.Stdout = &out_exec
	start := time.Now()
	err = command_exec.Run()
	if err != nil {
		log.Println(err)
	}

	Total = time.Since(start).Seconds()
	return
}
