package utils

import (
	"errors"
	"log"
	"os"
	"path"
)

// TODO : Add UnitTest for this function
func Initialization(folderName string) {
	newPath := path.Join(ResultsPath, folderName)
	if _, err := os.Stat(newPath); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(newPath, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("File Generation Successful")
	} else {
		log.Panicln(" File Exist ! See You ...")
	}

}
