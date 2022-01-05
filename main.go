package main

import (
	"fmt"
	"ga_tuner/scripts"
	"ga_tuner/utils"
	"log"
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
		fmt.Printf("Best fitness at generation %d: ID:  %s, Fitness : %f\n", ga.Generations, ga.HallOfFame[0].ID, ga.HallOfFame[0].Fitness)
	}
	ga.Logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// Find the minimum
	err = ga.Minimize(scripts.VectorFactory)
	if err != nil {
		fmt.Println(err)
		return
	}
}
