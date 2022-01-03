package main

import (
	"fmt"
	"ga_tuner/scripts"
	"ga_tuner/utils"
	"os"

	"github.com/MaxHalford/eaopt"
)

func init() {
	utils.Initialization(os.Args[1])
}

func main() {

	// Instantiate a GA with a GAConfig
	var ga, err = eaopt.NewDefaultGAConfig().NewGA()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Set the number of generations to run for
	ga.NGenerations = 10
	ga.ParallelEval = true
	// Add a custom print function to track progress
	ga.Callback = func(ga *eaopt.GA) {
		hof := ga.HallOfFame[0].Genome
		fmt.Printf("Best Combination of Flags : %f\n", hof.(scripts.SingleBench).BinVec)
		fmt.Printf("Best fitness at generation %d: %f\n", ga.Generations, ga.HallOfFame[0].Fitness)
	}

	// Find the minimum
	err = ga.Minimize(scripts.VectorFactory)
	if err != nil {
		fmt.Println(err)
		return
	}
}
