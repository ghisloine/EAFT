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
	"strconv"

	"github.com/MaxHalford/eaopt"
	uuid "github.com/satori/go.uuid"
)

// A Vector of SingleBench.
type Vector []float64

type SingleBench struct {
	Id        string
	cmd       string
	problem   string
	runtime   int
	binPath   string
	execPath  string
	opt_level string
	Flagset   map[string]float64
	BinVec    Vector
}

var availableFlags []string = tools.Flags

func MatchBinaryWithFlags(X SingleBench) string {
	// First collect all available Flags
	cmd := "-O" + X.opt_level + " "
	// Replace with -f or -fno according to X

	for key, elem := range X.Flagset {
		if elem == 0 {
			cmd += "-fno-" + key + " "

		} else {
			cmd += "-f" + key + " "
		}
	}
	return cmd
}

func addPolybenchDependencies(command string, problem string, out_file string) string {
	command += path.Join(utils.Files, problem) + `.c` + ` -I` + utils.Utilities + ` --include ` + `polybench.c` + ` -o ` + path.Join(utils.ResultsPath, os.Args[1], "bin", out_file)
	return command
}

// Fitness function burasi
func (X SingleBench) Evaluate() (float64, error) {
	// Changing Binary Array to GCC command with corresponding open / close flag
	cmd := MatchBinaryWithFlags(X)

	// Adding some polybench information to run cmd
	cmd = addPolybenchDependencies(cmd, os.Args[1], X.Id)

	// Total is Execution time of Code.
	total := CompileCode(cmd, X.Id)

	WriteJsonFile(X, total)
	return total, nil
}

// Mutate a Vector by resampling each element from a normal distribution with
// probability 0.8.

func (X SingleBench) Mutate(rng *rand.Rand) {
	MutNormalFloat64(X.BinVec, 0.8, rng)
}

// TODO : Paper'da olup burada olmayan crossover metodlari neler var ona bak.
// Crossover a Vector with another Vector by applying uniform crossover.
func (X SingleBench) Crossover(Y eaopt.Genome, rng *rand.Rand) {
	eaopt.CrossGNXFloat64(X.BinVec, Y.(SingleBench).BinVec, 2, rng)
}

// Clone a Vector to produce a new one that points to a different slice.
func (X SingleBench) Clone() eaopt.Genome {
	var Y SingleBench = X
	return Y
}

// VectorFactory returns a random vector by generating 2 values uniformally
// distributed between -10 and 10.

// TODO : BinVec could be changed with Key : Value pair. Key may be flag, Value may be 0-1.
func VectorFactory(rng *rand.Rand) eaopt.Genome {
	f, v := InitBinaryFloat64(5, 0, 2, rng)
	opt := strconv.Itoa(2 + rng.Intn(int(2)))
	return SingleBench{
		Id:        uuid.NewV4().String(),
		cmd:       "",
		problem:   os.Args[1],
		runtime:   0,
		binPath:   "",
		execPath:  "",
		opt_level: opt,
		Flagset:   f,
		BinVec:    v,
	}
}

func WriteJsonFile(X SingleBench, runtime float64) {
	outJson := make(map[string]interface{})

	outJson["id"] = X.Id
	outJson["runtime"] = runtime
	outJson["problem"] = X.problem
	outJson["optlevel"] = X.opt_level
	outJson["flagset"] = X.Flagset
	outJson["binary_values"] = X.BinVec

	jsonStr, _ := json.Marshal(outJson)
	ioutil.WriteFile(filepath.Join(utils.ResultsPath, os.Args[1], "results", X.Id+".json"), jsonStr, os.ModePerm)

}
