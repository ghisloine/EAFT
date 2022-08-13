package tools

import (
	"fmt"
	"log"
	"testing"
)

func TestReturnAllFlags(t *testing.T) {
	flags := ReturnAllFlags("gcc-11")
	actual := len(flags) > 0
	if actual != true {
		t.Error("No flags available in this architect.")
	} else {
		fmt.Printf("Total Flags : %d\n", len(flags))
	}
}

func TestRunToFile(t *testing.T) {
	flags := ReturnAllFlags("gcc-11")
	// Writing Flags to Files
	if err := WriteFlagsToFile(flags, "flags/flags.txt"); err != nil {
		log.Fatalf("writeLines: %s", err)
	}
}

func TestIfFlagFileExist(t *testing.T) {
	flags := IfFlagFileExist("gcc-11")

	if len(flags) <= 0 {
		log.Fatalf("There is an error while reading flag file.")
	}
}
