package tools

import (
	"bufio"
	"fmt"
	"ga_tuner/utils"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

// This will pass flags to other files.
// CHANGE THERE
var Flags []string = IfFlagFileExist("gcc-11")

// Checking Flags from terminal output. This is the best way to check flags for a system.
// Next Step is Working flags. Not all working flags are compatible with operating system.
func ReturnAllFlags(app string) []string {
	arg1 := "--help=optimizers"
	cmd := exec.Command(app, arg1)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	output := string(stdout)
	var re = regexp.MustCompile("(?m)^  (-f[a-z0-9-]+) ")

	var flags []string
	flags = append(flags, re.FindAllString(output, -1)...)
	// Removing each -f flag from begining.
	for idx, v := range flags {
		flags[idx] = strings.Replace(v, "-f", "", 1)
		flags[idx] = strings.Replace(flags[idx], " ", "", 3)

	}
	return flags
}

// If there is a file in flags folder, don't need to read all flags from terminal
// Just read the file
func IfFlagFileExist(app string) []string {
	file, err := os.Open(filepath.Join(utils.UtilsFolder, "tools", "flags", "flags.txt"))

	if err != nil {
		fmt.Println("Error While Reading Flag File")
		fmt.Println("Collecting Flags From Terminal Output, not Txt")
		return ReturnAllFlags(app)
	} else {
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		var text []string
		for scanner.Scan() {
			text = append(text, scanner.Text())
		}
		file.Close()
		return text

	}
}

// For further use this function will write all flags line by line in a txt file.
func WriteFlagsToFile(flags []string, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range flags {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}
