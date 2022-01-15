package scripts

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"runtime"

	"github.com/MaxHalford/eaopt"
	uuid "github.com/satori/go.uuid"
)

func PSORunner() {
	fmt.Printf("Particle Swarm Optimization running parallel with %d Processor\n", runtime.NumCPU())
	var spso, err = eaopt.NewSPSO(50, 10, 0, 1, 0, true, nil)
	f := WriteJson(spso.GA)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Fix random number generation
	spso.GA.RNG = rand.New(rand.NewSource(42))
	spso.GA.Callback = func(ga *eaopt.GA) {
		fmt.Printf("Best fitness at generation %d: ID:  %s, Fitness : %f", ga.Generations, ga.HallOfFame[0].ID, ga.HallOfFame[0].Fitness)
		var bytes, _ = json.Marshal(ga)
		f.WriteString(string(bytes) + "\n")
	}
	_, y, err := spso.Minimize(FitnessFunction, 50)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(y)
}

func FitnessFunction(X []float64) (y float64) {
	output := uuid.NewV4().String()
	flattenSlice := make([]float64, 0)
	for _, v := range X {
		if v < 0.5 {
			flattenSlice = append(flattenSlice, 0)
		} else {
			flattenSlice = append(flattenSlice, 1)
		}
	}
	cmd, _ := MatchBinaryWithFlags(flattenSlice)

	// Adding some polybench information to run cmd
	cmd = addPolybenchDependencies(cmd, os.Args[1], output)

	// Total is Execution time of Code.
	total := CompileCode(cmd, output)
	return total
}
