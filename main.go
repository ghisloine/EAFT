package main

import (
	"fmt"
	"ga_tuner/scripts"
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

	// Runner := os.Args[4]
	// if Runner == "GA" {
	// 	go scripts.GARunner()
	// } else if Runner == "PSO" {
	// 	go scripts.PSORunner()
	// }
	// utils.CLI()
	benchmark_list := []string{"nussinov", "syr2k", "durbin", "seidel-2d", "ludcmp", "adi", "jacobi-1d", "trisolv", "jacobi-2d", "3mm", "symm", "correlation", "floyd-warshall", "gramschmidt", "syrk", "heat-3d", "2mm", "bicg", "fdtd-2d", "deriche", "mvt", "gemm", "doitgen", "covariance", "cholesky", "lu", "gesummv", "trmm", "atax", "gemver"}
	for _, v := range benchmark_list {
		fmt.Printf("Starting Problem : %s\n", v)
		utils.Initialization(v)
		scripts.PSORunner(v)
	}

	// utils.Initialization(os.Args[1])
	// scripts.GARunner(v)
}
