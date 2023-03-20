package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/MaxHalford/eaopt"
)

func WriteAlgorithmConfigurations(obj GeneticObject, MutationRate float64, CrossoverRate float64) {

	logConf := ConfigurationsObject{obj.ObjectStruct.NPops, obj.ObjectStruct.PopSize, obj.ObjectStruct.NGenerations, obj.ModelName, obj.SelectorName, MutationRate, CrossoverRate}
	confFileName := "config.json"

	fileFullPath := filepath.Join(ResultsPath, obj.ResultFolderName, obj.ExperimentDate)

	// Main algorithm name folder
	createNewFolder(fileFullPath, false)
	// Creating binary output folder named bin.
	createNewFolder(filepath.Join(fileFullPath, "bin"), false)

	jsonMarshal, _ := json.Marshal(logConf)
	_ = ioutil.WriteFile(filepath.Join(fileFullPath, confFileName), jsonMarshal, os.ModePerm)
}

func PrintLines(filePath string, values []interface{}) error {
	var f, err = os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	for _, value := range values {
		fmt.Fprintln(f, value) // print values to f, one per line
	}
	return nil
}

func CreateLogFile(filePath string) *os.File {
	f, err := os.OpenFile(filePath,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	return f
}

func AppendToFile(filePath string, data string) {

	f, err := os.OpenFile(filePath,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(string(data) + "\n"); err != nil {
		log.Println(err)
	}
}

func WriteAllResult(list []eaopt.GA, filepath string) {
	jsonData, err := json.MarshalIndent(list, "", "    ")
	if err != nil {
		panic(err)
	}
	// Write the byte array to a file
	err = ioutil.WriteFile(filepath, jsonData, 0644)
	if err != nil {
		panic(err)
	}
}
