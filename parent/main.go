package main

import (
	"bytes"
	"envtest/params"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	var (
		stderr bytes.Buffer
		stdout bytes.Buffer
	)

	os.Setenv(params.ParentSecretEnvKey, params.ParentSecretEnvValue)

	cmd := exec.Command("./child/child", "arg1", "arg2")
	cmd.Env = append(cmd.Env, params.ParentChildSecretEnvKey, params.ParentChildSecretEnvValue)

	cmd.Stderr = &stderr
	cmd.Stdout = &stdout

	err := cmd.Run()

	if err != nil {
		fmt.Printf("Child returned an error: %v", err)
		return
	}

	stderrBytes := stderr.Bytes()

	if len(stderrBytes) > 0 {
		fmt.Printf("Child returned error strings: %v", string(stderrBytes))
		return
	}

	stdoutBytes := stdout.Bytes()

	if len(stdoutBytes) > 0 {
		fmt.Printf("Child gave some output: %v", string(stdoutBytes))
		return
	}
}
