package scripts

import (
	"fmt"
	"ga_tuner/utils"
	"ga_tuner/utils/tools"
	"math/rand"
	"path"
    "github.com/satori/go.uuid"
	"github.com/MaxHalford/eaopt"
)
type SingleBench struct{
    uuid int
    cmd string
    problem string
    runtime int
    binPath string
    execPath string
}


// A Vector contains float64s.
type Vector []float64

var availableFlags []string = tools.Flags

func MatchBinaryWithFlags(X Vector) string {
	// First collect all available Flags
	cmd := "gcc-11"
	// Replace with -f or -fno according to X
	for idx := range X {
		if X[idx] == 0 {
			cmd += " -fno-" + availableFlags[idx]
		} else {
			cmd += " -f" + availableFlags[idx]
		}
	}
	return cmd
}

func addPolybenchDependencies(command string, problem string) string {
	command += " " + path.Join(utils.Files, problem) + " -I" + utils.Utilities + " --include " + "polybench.c" 
	return command
}

// Fitness function burasi
func (X Vector) Evaluate() (float64, error) {
	// var Total float64
	cmd := MatchBinaryWithFlags(X)
	cmd = addPolybenchDependencies(cmd, "atax")
	// Adding some polybench information to run cmd

	fmt.Println(cmd)
	return 1.0, nil
}

// Mutate a Vector by resampling each element from a normal distribution with
// probability 0.8.

// TODO : Bu kisimda mutation yapildiktan sonra deger binary olarak degismesi gerekiyor.
// Aslinda daha basit eger 0 ise 1 vice versa.
func (X Vector) Mutate(rng *rand.Rand) {
	eaopt.MutNormalFloat64(X, 0, rng)
}

// Crossover a Vector with another Vector by applying uniform crossover.
// TODO : Paper'da olup burada olmayan crossover metodlari neler var ona bak.
func (X Vector) Crossover(Y eaopt.Genome, rng *rand.Rand) {
	eaopt.CrossUniformFloat64(X, Y.(Vector), rng)
}

// Clone a Vector to produce a new one that points to a different slice.
func (X Vector) Clone() eaopt.Genome {
	var Y = make(Vector, len(X))
	copy(Y, X)
	return Y
}

// ReturnS a Binary Vector with just only 0 - 1.
func InitBinaryFloat64(n uint, lower, upper float64, rng *rand.Rand) (floats []float64) {
	floats = make([]float64, n)
	for i := range floats {
		floats[i] = lower + float64(rng.Intn(int(upper)))
		// fmt.Printf("Generated Number is %f\n", floats[i])
	}
	return
}

// VectorFactory returns a random vector by generating 2 values uniformally
// distributed between -10 and 10.
func VectorFactory(rng *rand.Rand) eaopt.Genome {
	return Vector(InitBinaryFloat64(5, 0, 2, rng))
}
