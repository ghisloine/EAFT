package scripts

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"

	"github.com/MaxHalford/eaopt"
	uuid "github.com/satori/go.uuid"
)

func PSORunner() {
	var spso, err = eaopt.NewSPSO(10, 10, 0, 1, 0, true, nil)

	f, _ := os.Create(os.Args[3])
	defer f.Close()
	w := bufio.NewWriter(f)
	fmt.Fprint(w, "")
	w.Flush()

	f, _ = os.OpenFile(os.Args[3], os.O_APPEND|os.O_WRONLY, 0666)
	defer f.Close()

	var bytes, _ = json.Marshal(spso.GA)
	f.WriteString(string(bytes) + "\n")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Collecting Baselines of Code
	level_two := CollectBaseline("O2")
	level_three := CollectBaseline("O3")

	fmt.Printf("O2 BASELINE IS : %f\n", level_two)
	fmt.Printf("O3 BASELINE IS : %f\n", level_three)

	// Fix random number generation
	spso.GA.RNG = rand.New(rand.NewSource(42))
	spso.GA.Callback = func(ga *eaopt.GA) {
		fmt.Printf("Best fitness at generation %d: ID:  %s, Fitness : %f, Improvement : %f\n", ga.Generations, ga.HallOfFame[0].ID, ga.HallOfFame[0].Fitness, 1-ga.HallOfFame[0].Fitness/level_two)
		var bytes, err = json.Marshal(ga)
		if err != nil {
			fmt.Println(err)
		}
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
	cmd, _ := MatchBinaryWithFlags(X, "O2")

	// Adding some polybench information to run cmd
	cmd = addPolybenchDependencies(cmd, os.Args[1], output)

	// Total is Execution time of Code.
	total := CompileCode(cmd, output, 3)
	return total
}
