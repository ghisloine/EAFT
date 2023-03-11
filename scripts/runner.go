package scripts

import (
	"ga_tuner/utils"
	"ga_tuner/utils/tools"
	"math/rand"
	"path"

	"github.com/MaxHalford/eaopt"
	uuid "github.com/satori/go.uuid"
)

// A Vector of SingleBench.
type Vector []float64

var availableFlags []string = tools.Flags

// TODO : Change static optimization level with Dynamic one.
func MatchBinaryWithFlags(X Vector, OptLevel string) (string, map[string]int) {
	// First collect all available Flags
	cmd := "-" + OptLevel + " "
	// Replace with -f or -fno according to X
	flag_map := map[string]int{}
	for i, v := range X {
		if int(v) == 0 {
			cmd += "-fno-" + availableFlags[i] + " "
			flag_map[availableFlags[i]] = int(v)
		} else {
			cmd += "-f" + availableFlags[i] + " "
			flag_map[availableFlags[i]] = int(v)
		}
	}

	return cmd, flag_map
}

func addPolybenchDependencies(command string, problem string, out_file string) string {
	command += path.Join(utils.Files, problem) + `.c` + ` -I` + utils.Utilities + ` --include ` + `polybench.c` + ` -o ` + path.Join(utils.ResultsPath, utils.Pc.ResultFolderName, "bin", out_file)
	return command
}

// Fitness function burasi
func (X Vector) Evaluate() (float64, error) {
	// Changing Binary Array to GCC command with corresponding open / close flag
	output := uuid.NewV4().String()
	cmd, _ := MatchBinaryWithFlags(X, "O3")

	// Adding some polybench information to run cmd
	cmd = addPolybenchDependencies(cmd, utils.Pc.ResultFolderName, output)

	// Total is Execution time of Code.
	total := CompileCode(cmd, output, 3)

	return total, nil
}

// Mutate a Vector by resampling each element from a normal distribution with
// probability 0.8.
// TODO Change mutation with dynamic variable.

func (p Vector) Mutate(rng *rand.Rand) {
	MutNormalFloat64(p, 0.8, rng)
}

// Crossover a Vector with another Vector by applying uniform crossover.
func (X Vector) Crossover(Y eaopt.Genome, rng *rand.Rand) {
	eaopt.CrossGNXFloat64(X, Y.(Vector), 2, rng)
}

// Clone a Vector to produce a new one that points to a different slice.
func (X Vector) Clone() eaopt.Genome {
	var Y = make(Vector, len(X))
	copy(Y, X)
	return Y
}

// VectorFactory returns a random vector by generating 2 values uniformally
// distributed between -10 and 10.
func VectorFactory(rng *rand.Rand) eaopt.Genome {
	// NUMBER_OF_FLAGS := uint(len(availableFlags))
	return Vector(InitBinaryFloat64(50, 0, 2, rng))

}

func CollectBaseline(Baseline string) float64 {
	output := uuid.NewV4().String()
	cmd, _ := MatchBinaryWithFlags(make([]float64, 0), Baseline)
	// Adding some polybench information to run cmd
	cmd = addPolybenchDependencies(cmd, utils.Pc.ResultFolderName, output)

	// Total is Execution time of Code.
	total := CompileCode(cmd, output, 1)
	return total
}
