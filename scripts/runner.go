package scripts

import (
	"encoding/json"
	"ga_tuner/utils"
	"ga_tuner/utils/tools"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"path/filepath"

	"github.com/MaxHalford/eaopt"
	uuid "github.com/satori/go.uuid"
)

// A Vector of SingleBench.
type Vector []float64

var availableFlags []string = tools.Flags

func MatchBinaryWithFlags(X Vector) (string, map[string]int) {
	// First collect all available Flags
	cmd := "-O3" + " "
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
	command += path.Join(utils.Files, problem) + `.c` + ` -I` + utils.Utilities + ` --include ` + `polybench.c` + ` -o ` + path.Join(utils.ResultsPath, os.Args[1], "bin", out_file)
	return command
}

// Fitness function burasi
func (X Vector) Evaluate() (float64, error) {
	// Changing Binary Array to GCC command with corresponding open / close flag
	output := uuid.NewV4().String()

	cmd, flag_map := MatchBinaryWithFlags(X)

	// Adding some polybench information to run cmd
	cmd = addPolybenchDependencies(cmd, os.Args[1], output)

	// Total is Execution time of Code.
	total := CompileCode(cmd, output)

	WriteJsonFile(output, flag_map, total)

	return total, nil
}

// Mutate a Vector by resampling each element from a normal distribution with
// probability 0.8.

func (p Vector) Mutate(rng *rand.Rand) {
	MutNormalFloat64(p, 0.8, rng)
}

// TODO : Paper'da olup burada olmayan crossover metodlari neler var ona bak.
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

// TODO : BinVec could be changed with Key : Value pair. Key may be flag, Value may be 0-1.
func VectorFactory(rng *rand.Rand) eaopt.Genome {
	return Vector(InitBinaryFloat64(6, 0, 2, rng))

}

func WriteJsonFile(id string, flag_map map[string]int, runtime float64) {
	outJson := make(map[string]interface{})

	outJson["id"] = id
	outJson["runtime"] = runtime
	outJson["problem"] = os.Args[1]
	outJson["flagset"] = flag_map

	jsonStr, _ := json.Marshal(outJson)
	ioutil.WriteFile(filepath.Join(utils.ResultsPath, os.Args[1], "results", id+".json"), jsonStr, os.ModePerm)

}
