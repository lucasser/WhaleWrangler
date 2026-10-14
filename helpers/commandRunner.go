package helpers

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func RunCommand(working_dir string, cmd string, service DockerService) {
	// Change the current working directory to the service's home directory
	path := working_dir + service.ProjectDir
	if err := os.Chdir(path); err != nil {
		fmt.Printf("Error changing directory to %s: %v\n", path, err)
		return
	}

	// Execute the command
	commandArray := strings.Fields(cmd)
	// Execute the command
	cmdExec := exec.Command(commandArray[0], commandArray[1:]...)
	cmdExec.Stdout = os.Stdout
	cmdExec.Stderr = os.Stderr

	fmt.Printf("Executing command in %s: %s\n", path, cmd)

	err := cmdExec.Run()
	if err != nil {
		fmt.Printf("Command failed with error: %v\n", err)
	} else {
		fmt.Println("Command executed successfully.")
	}
}
