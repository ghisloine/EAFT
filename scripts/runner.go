package scripts

import (
	"ga_tuner/utils"
	"ga_tuner/utils/tools"
	"math/rand"
	"os"
	"path"

	"github.com/MaxHalford/eaopt"
	uuid "github.com/satori/go.uuid"
)

// A Vector of SingleBench.
type Vector []float64

type SingleBench struct {
	uuid       uuid.UUID
	cmd        string
	problem    string
	runtime    int
	binPath    string
	execPath   string
	binary_vec Vector
}

var availableFlags []string = tools.Flags

func MatchBinaryWithFlags(X SingleBench) string {
	// First collect all available Flags
	cmd := ""
	// Replace with -f or -fno according to X
	for idx := range X.binary_vec {
		if X.binary_vec[idx] == 0 {
			cmd += "-fno-" + availableFlags[idx] + " "
		} else {
			cmd += "-f" + availableFlags[idx] + " "
		}
	}
	return cmd
}

func addPolybenchDependencies(command string, problem string, out_file string) string {
	command += path.Join(utils.Files, problem) + `.c` + ` -I` + utils.Utilities + ` --include ` + `polybench.c` + ` -o ` + path.Join(utils.ResultsPath, out_file)
	return command
}

// Fitness function burasi
func (X SingleBench) Evaluate() (float64, error) {
	// Changing Binary Array to GCC command with corresponding open / close flag
	cmd := MatchBinaryWithFlags(X)

	// Adding some polybench information to run cmd
	cmd = addPolybenchDependencies(cmd, os.Args[1], X.uuid.String())

	total := CompileCode(cmd, X.uuid)

	return total, nil
}

// Mutate a Vector by resampling each element from a normal distribution with
// probability 0.8.

func (X SingleBench) Mutate(rng *rand.Rand) {
	MutNormalFloat64(X.binary_vec, 0.8, rng)
}

// Crossover a Vector with another Vector by applying uniform crossover.
// TODO : Paper'da olup burada olmayan crossover metodlari neler var ona bak.
func (X SingleBench) Crossover(Y eaopt.Genome, rng *rand.Rand) {
	eaopt.CrossGNXFloat64(X.binary_vec, Y.(SingleBench).binary_vec, 2, rng)
}

// Clone a Vector to produce a new one that points to a different slice.
func (X SingleBench) Clone() eaopt.Genome {
	var Y SingleBench = X
	return Y
}

// VectorFactory returns a random vector by generating 2 values uniformally
// distributed between -10 and 10.
func VectorFactory(rng *rand.Rand) eaopt.Genome {
	return SingleBench{
		uuid:       uuid.NewV4(),
		cmd:        "",
		problem:    os.Args[1],
		runtime:    0,
		binPath:    "",
		execPath:   "",
		binary_vec: InitBinaryFloat64(5, 0, 2, rng),
	}
}
