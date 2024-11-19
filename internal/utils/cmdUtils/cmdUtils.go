package utils_cmd

import (
	"os"
	"os/exec"

	"github.com/charmbracelet/log"
)

// Executes a command and pipes the output to stdout. If cmdDir is an empty string, the cwd is used.
func PipeCommand(cmdDir string, commandString string, args ...string) {
	cmd := exec.Command(commandString, args...)
	if cmdDir != "" {
		cmd.Dir = cmdDir
	}
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	// stdout, _ := cmd.StdoutPipe()
	// if err := cmd.Start(); err != nil {
	// 	log.Fatal(err)
	// }
	// stdout.Read()
	// fmt.Print(stdout)
}
