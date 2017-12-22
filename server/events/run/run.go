// Package run handles running commands prior and following the
// regular Atlantis commands.
package run

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/hashicorp/go-version"
	"github.com/hootsuite/atlantis/server/logging"
	"github.com/pkg/errors"
)

const inlineShebang = "#!/bin/sh -e"

//go:generate pegomock generate -m --use-experimental-model-gen --package mocks -o mocks/mock_runner.go Runner

// Runner is an interface for running bash commands.
type Runner interface {
	// Execute runs each command in commands. It sets the following environment
	// variables:
	// WORKSPACE to path
	// ENVIRONMENT to environment
	// ATLANTIS_TERRAFORM_VERSION to terraformVersion
	// stage is whether this is a pre_plan, post_plan, etc. It's only used for
	// logging.
	Execute(log *logging.SimpleLogger, commands []string, path string, environment string, terraformVersion *version.Version, stage string) (string, error)
}

// DefaultRunner implements Runner.
type DefaultRunner struct{}

// Execute see Runner.Execute.
func (p *DefaultRunner) Execute(
	log *logging.SimpleLogger,
	commands []string,
	path string,
	environment string,
	terraformVersion *version.Version,
	stage string) (string, error) {
	// Create a script from the commands provided.
	if len(commands) == 0 {
		return "", errors.Errorf("%s commands cannot be empty", stage)
	}

	s, err := createScript(commands, stage)
	if err != nil {
		return "", err
	}
	defer os.Remove(s) // nolint: errcheck

	log.Info("running %s commands: %v", stage, commands)
	os.Setenv("ENVIRONMENT", environment)                              // nolint: errcheck
	os.Setenv("ATLANTIS_TERRAFORM_VERSION", terraformVersion.String()) // nolint: errcheck
	os.Setenv("WORKSPACE", path)                                       // nolint: errcheck
	return execute(s)
}

func createScript(cmds []string, stage string) (string, error) {
	// Write out the contents to a bash script that we execute. We do this
	// so we can ensure we don't have any weird execution issues when using
	// Exec() like unexpected escaping.
	tmp, err := ioutil.TempFile("/tmp", "atlantis-temp-script")
	if err != nil {
		return "", errors.Wrapf(err, "preparing %s shell script", stage)
	}

	scriptName := tmp.Name()

	writer := bufio.NewWriter(tmp)
	if _, err = writer.WriteString(fmt.Sprintf("%s\n", inlineShebang)); err != nil {
		return "", errors.Wrapf(err, "writing to %q", tmp.Name())
	}
	cmdsJoined := strings.Join(cmds, "\n")
	if _, err := writer.WriteString(cmdsJoined); err != nil {
		return "", errors.Wrapf(err, "preparing %s", stage)
	}

	if err := writer.Flush(); err != nil {
		return "", errors.Wrap(err, "flushing contents to file")
	}
	tmp.Close() // nolint: errcheck

	if err := os.Chmod(scriptName, 0700); err != nil { // nolint: gas
		return "", errors.Wrapf(err, "making %s script executable", stage)
	}

	return scriptName, nil
}

func execute(script string) (string, error) {
	localCmd := exec.Command("sh", "-c", script) // #nosec
	out, err := localCmd.CombinedOutput()
	output := string(out)
	if err != nil {
		return output, errors.Wrapf(err, "running script %s: %s", script, output)
	}

	return output, nil
}
