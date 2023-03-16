package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/MaxHalford/eaopt"
)

var DataPath = GetPath("data")
var ResultsPath = GetPath("results")
var Utilities = filepath.Join(DataPath, "Polybench", "utilities")
var Files = filepath.Join(DataPath, "Polybench", "datamining")
var PolybenchC = filepath.Join(Utilities, "polybench.c")

type GeneticObject struct {
	ObjectType       string
	ObjectStruct     eaopt.GAConfig
	ResultFolderName string
	ExperimentDate   string
	GccShortcut      string
	ModelName        string
	SelectorName     string
	MutationRate     float64
}

type ConfigurationsObject struct {
	NumberOfPopulation uint
	PopulationSize     uint
	NumberOfGeneration uint
	ModelName          string
	SelectorName       string
	MutationRate       float64
	CrossoverRate      float64
}

var MainAlgorithms = []string{"Genetic Algorithm", "Particle Swarm Optimization"}
var Models = []string{"Steady State", "Generational", "Down to Size", "Ring", "Mutation Only"}
var Selections = []string{"Tournament", "Roulette", "Elitism"}

func GetDirFiles(dataPath string) []string {
	files, err := ioutil.ReadDir(dataPath)
	if err != nil {
		log.Fatal(err)
	}
	var Directory []string
	for _, f := range files {
		Directory = append(Directory, f.Name())
	}
	return Directory
}

func GetPath(newPath string) string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	data := filepath.Join(pwd, newPath)
	return data
}
