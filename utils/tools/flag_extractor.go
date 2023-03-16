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

var Flags []string = ReadFlagsFromFile() //ReturnAllFlags(utils.Pc.GccShortcut)

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

func ReadFlagsFromFile() []string {
	toolsPath := utils.GetPath("utils")
	file, err := os.Open(filepath.Join(toolsPath, "tools", "flags"))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Okuma işlemini gerçekleştir
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Okunan satırlardan yeni bir liste oluştur
	var myList []string
	for _, line := range lines {
		myList = append(myList, line)
	}
	return myList
}
