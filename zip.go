package main

import (
	"fmt"
	"os"
	"os/exec"
)

func zip(source, target string) error {
	// remove file before zip
	os.Remove(target)

	cmd := exec.Command("zip", "-r", target, "./")
	cmd.Dir = source

	stdoutStderr, err := cmd.CombinedOutput()
	fmt.Printf("%s\n", string(stdoutStderr))

	if err != nil {
		return err
	}

	return nil
}
