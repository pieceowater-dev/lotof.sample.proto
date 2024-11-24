package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const (
	yellowColor = "\033[33m"
	resetColor  = "\033[0m" 
)

func getModuleName() string {
	data, err := os.ReadFile("go.mod")
	if err != nil {
		return "unknown"
	}

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		if strings.HasPrefix(line, "module") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module"))
		}
	}
	return "unknown"
}

func getAppVersion() string {
	cmd := exec.Command("git", "describe", "--tags", "--abbrev=0")
	cmd.Env = os.Environ()
	output, err := cmd.Output()
	if err != nil {
		return "v0.0.0"
	}
	return strings.TrimSpace(string(output))
}

func logModuleInfo() {
	moduleName := getModuleName()
	version := getAppVersion()
	fmt.Printf("%sProto from %s, release: %s%s\n", yellowColor, moduleName, version, resetColor)
}

func Info() {
	logModuleInfo()
}

func main() {
	Info()
}