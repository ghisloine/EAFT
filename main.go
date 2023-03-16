package main

import (
	"fmt"
	"ga_tuner/scripts"
	"ga_tuner/utils"
)

func main() {

	for i := 1.0; i <= 10.0; i++ {
		fmt.Sprintf("ITERATION : %f", i)
		utils.Pc = utils.ManuelConfiguration("2mm", 0.8, uint(50*i))
		scripts.GARunner()
		utils.Bar.Reset()
	}

	for i := 1.0; i <= 10.0; i++ {
		fmt.Sprintf("ITERATION : %f", i)
		utils.Pc = utils.ManuelConfiguration("floyd-warshall", 0.8, uint(50*i))
		scripts.GARunner()
		utils.Bar.Reset()
	}

	//if utils.Pc.ObjectType == "Genetic Algorithm" {
	//	go scripts.GARunner()
	//} else if utils.Pc.ObjectType == "Particle Swarm Optimization" {
	//	go scripts.PSORunner()
	//}
	//utils.CLI()

}
