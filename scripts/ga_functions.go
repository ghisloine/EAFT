package scripts

import (
	"fmt"
	"ga_tuner/utils"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
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
		// fmt.Printf("Generated Number is %f\n", floats[i])
	}
	return
}

func CompileCode(cmd string, id uuid.UUID) float64 {
	// COMPILE
	command := strings.Split(cmd, " ")
	app := os.Args[2]
	out, err := exec.Command(app, command...).Output()
	if err != nil {
		log.Print(string(out))
		log.Fatal(err)
	}

	// EXECUTION
	start := time.Now()
	app = "cmd"
	exec_file := filepath.Join(utils.ResultsPath, os.Args[1], id.String())
	out, err = exec.Command(exec_file).Output()
	if err != nil {
		log.Print(string(out))
		log.Fatal(err)
	}
	elapsed := time.Since(start).Seconds()
	fmt.Printf("Total Elapsed Time is : %f s\n", elapsed)

	return float64(elapsed)
}
