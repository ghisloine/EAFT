package utils

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// TODO : Add UnitTest for this function
func Initialization(folderName string) {
	newPath := filepath.Join(ResultsPath, folderName)
	if _, err := os.Stat(newPath); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(newPath, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("File Generation Successful")
	} else {
		fmt.Println("Folder Exist. Would you like to delete and generate again ?")
		fmt.Println("[1] For Yes, Enter 1")
		fmt.Println("[2] For No,  Enter 2")
		var answer string
		answer = "1"
		if answer == "1" {
			os.RemoveAll(newPath)
			err = os.MkdirAll(newPath, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
			log.Println("File Deletion -> Generation Successfull")
		} else {
			log.Panic("Quitting. Goodbye ...")
		}

	}

}
