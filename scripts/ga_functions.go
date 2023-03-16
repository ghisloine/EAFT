package scripts

import (
	"bytes"
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

	logPath := filepath.Join(utils.ResultsPath, utils.Pc.ResultFolderName, utils.Pc.ExperimentDate)

	levelTwo := calculateBaseline(logPath)

	utils.Pc.ObjectStruct.Callback = func(ga *eaopt.GA) {
		utils.TextBox = fmt.Sprintf("Best fitness at generation %d: ID:  %s, Fitness : %f, Improvement : %f\n", ga.Generations, ga.HallOfFame[0].ID, ga.HallOfFame[0].Fitness, 1-ga.HallOfFame[0].Fitness/levelTwo)
		utils.Progress = (float64(ga.Generations+1) / float64(ga.NGenerations)) * float64(100)
		utils.HallOfFame = ga.HallOfFame[0].Fitness
		utils.HofList = append(utils.HofList, ga.HallOfFame[0].Fitness)
		utils.BestOfPops = append(utils.BestOfPops, ga.HallOfFame[0].Fitness)
		utils.Stats = []float64{math.Floor(ga.HallOfFame.FitMin()*100) / 100, math.Floor(ga.HallOfFame.FitMax()*100) / 100, math.Floor(ga.HallOfFame.FitAvg()*100) / 100}

		utils.AllResults = append(utils.AllResults, *ga)
		utils.AppendToFile(filepath.Join(logPath, "hof"), fmt.Sprintf("%f", ga.HallOfFame[0].Fitness))
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

	utils.WriteAllResult(utils.AllResults, filepath.Join(logPath, "log"))

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
	outCompile, err := exec.Command(app, command...).Output()
	executeFilePath := filepath.Join(utils.ResultsPath, utils.Pc.ResultFolderName, utils.Pc.ExperimentDate, "bin", id)
	if err != nil {
		log.Print(string(outCompile))
		Total = 99999
		return Total
	}

	// EXECUTION
	TotalExecTime := 0.0
	for i := 0; i < count; i++ {
		commandExec := exec.Command(executeFilePath)
		var outExec bytes.Buffer
		// set the output to our variable
		commandExec.Stdout = &outExec
		start := time.Now()
		err = commandExec.Run()
		TotalExecTime += time.Since(start).Seconds()
		if err != nil {
			Total = math.Inf(10)
			return Total
		}
	}
	// CALC AVERAGE OF TOTAL RUN TIME
	Total = TotalExecTime / float64(count)
	utils.TotalRunTimes = append(utils.TotalRunTimes, math.Floor(Total*100)/100)
	deleteBinaryFile(executeFilePath)
	utils.Bar.Add(1)
	return
}

func deleteBinaryFile(fullPath string) {
	_ = os.RemoveAll(fullPath)
}

func calculateBaseline(logPath string) (levelTwo float64) {
	levelTwo = CollectBaseline("O2")
	levelThree := CollectBaseline("O3")

	levelTwoNotification := fmt.Sprintf("O2 : %f", levelTwo)
	levelThreeNotification := fmt.Sprintf("O3 : %f", levelThree)

	utils.Notifications = append(utils.Notifications, levelTwoNotification)
	utils.Notifications = append(utils.Notifications, levelThreeNotification)

	utils.AppendToFile(filepath.Join(logPath, "hof"), fmt.Sprintf("%f", levelTwo))
	utils.AppendToFile(filepath.Join(logPath, "hof"), fmt.Sprintf("%f", levelThree))

	return
}
