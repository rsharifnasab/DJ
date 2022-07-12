package run

import (
	"fmt"
	"os/exec"
)

// check if specified program is installed on path or not
func isInstalled(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}

// check that which program in provided list is not installed
// if all of them are installed, return an empty string
func WhichNotInstalled(names []string) string {
	for _, name := range names {
		if !isInstalled(name) {
			return name
		}
	}
	return ""
}

// prompt user that the program is not installed
func CheckAndErrorRequirements(names []string) {
	notInstalled := WhichNotInstalled(names)
	if notInstalled != "" {
		panic(fmt.Sprintf("program [%v] is not installed on your path\n", notInstalled))

	}
}
