package scripts

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"ga_tuner/utils"
	"log"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/MaxHalford/eaopt"
)

func GARunner() {

	JSON_FILE := filepath.Join(utils.ResultsPath, utils.Pc.ResultFolderName, "log", utils.Pc.ResultFolderName)
	f, _ := os.Create(JSON_FILE)
	defer f.Close()
	w := bufio.NewWriter(f)
	fmt.Fprint(w, "")
	w.Flush()

	f, _ = os.OpenFile(JSON_FILE, os.O_APPEND|os.O_WRONLY, 0666)
	defer f.Close()

	var bytes, _ = json.Marshal(utils.Pc)
	f.WriteString(string(bytes) + "\n")

	level_two := CollectBaseline("O2")
	level_three := CollectBaseline("O3")

	level_two_notification := fmt.Sprintf("O2 : %f", level_two)
	level_three_notification := fmt.Sprintf("O3 : %f", level_three)

	utils.Notifications = append(utils.Notifications, level_two_notification)
	utils.Notifications = append(utils.Notifications, level_three_notification)
	// Add a custom print function to track progress
	utils.Pc.ObjectStruct.Callback = func(ga *eaopt.GA) {
		utils.TextBox = fmt.Sprintf("Best fitness at generation %d: ID:  %s, Fitness : %f, Improvement : %f\n", ga.Generations, ga.HallOfFame[0].ID, ga.HallOfFame[0].Fitness, 1-ga.HallOfFame[0].Fitness/level_two)
		utils.Progress = (float64(ga.Generations+1) / float64(ga.NGenerations)) * float64(100)
		utils.HallOfFame = ga.HallOfFame[0].Fitness
		utils.BestOfPops = append(utils.BestOfPops, ga.HallOfFame[0].Fitness)
		utils.Stats = []float64{math.Floor(ga.HallOfFame.FitMin()*100) / 100, math.Floor(ga.HallOfFame.FitMax()*100) / 100, math.Floor(ga.HallOfFame.FitAvg()*100) / 100}
		var bytes, err = json.Marshal(ga)

		if err != nil {
			fmt.Println(err)
		}
		f.WriteString(string(bytes) + "\n")
	}
	
	var geneticAlgorithm, err = utils.Pc.ObjectStruct.NewGA()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Find the minimum
	err = geneticAlgorithm.Minimize(VectorFactory)
	if err != nil {
		fmt.Println(err)
		return
	}
}

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

func CompileCode(cmd string, id string, count int) (Total float64) {
	// COMPILE
	command := strings.Split(cmd, " ")
	app := utils.Pc.GccShortcut
	out_compile, err := exec.Command(app, command...).Output()
	if err != nil {
		log.Print(string(out_compile))
		Total = math.Inf(10)
		return Total
	}

	// EXECUTION
	TotalExecTime := 0.0
	for i := 0; i < count; i++ {
		exec_file := filepath.Join(utils.ResultsPath, utils.Pc.ResultFolderName, "bin", id)
		command_exec := exec.Command(exec_file)
		var out_exec bytes.Buffer
		// set the output to our variable
		command_exec.Stdout = &out_exec
		start := time.Now()
		err = command_exec.Run()
		TotalExecTime += time.Since(start).Seconds()
		if err != nil {
			Total = math.Inf(10)
			return Total
		}
	}
	// CALC AVERAGE OF TOTAL RUN TIME
	Total = TotalExecTime / float64(count)
	utils.TotalRunTimes = append(utils.TotalRunTimes, math.Floor(Total*100)/100)
	return
}
