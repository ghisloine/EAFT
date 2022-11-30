package main

import (
	"ga_tuner/utils"
)

func init() {

}

func main() {

	// TOOD : Change this side to dynamic. User should have different options for finding optimal
	// Solution for their code. For Example : GA or Particle Swarm or Differential Evolution or OpenAI.
	// Use TERMUI for making Terminal interface. Plotting results and other details will be in there.
	// Instantiate a GA with a GAConfig

	// TODO : Set initial run time of given code. Compare it with best of population.
	// Show improvement with percentage. Ex : Fitness Value : 3 sn. %34 less then normal.

	// Set the number of generations to run for
	// Args[1] => Folder name for results
	// Args[2] => GCC-11
	// Args[3] => 2mm.json -> Result file.
	// Args[4] => GA or PSO

	// TODO : Make arguments more generic and take them in CLI as an option.

	// Example full run -> go run main.go 2mm gcc-11 2mm.json GA

	_ = utils.SelectConfigurations()
	// fmt.Println(pc)
	// if pc.ObjectType == "Genetic Algorithm" {
	// 	go scripts.GARunner()
	// } else if pc.ObjectType == "Particle Swarm Optimization" {
	// 	go scripts.PSORunner()
	// }
	utils.CLI()

}
