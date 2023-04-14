package utils

import (
	"log"
	"os"
	"os/exec"
)

func Run(cmd *exec.Cmd) error {
	log.Printf("Running %s", cmd)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func SH(c string) (string, error) {
	log.Printf("Running %s", c)
	cmd := exec.Command("/bin/sh", "-c", c)
	cmd.Env = os.Environ()
	o, err := cmd.CombinedOutput()
	return string(o), err
}

func RunInDir(c, dir string, envs ...string) (string, error) {
	log.Printf("Running %s", c)
	cmd := exec.Command("/bin/sh", "-c", c)
	cmd.Env = append(os.Environ(), envs...)
	cmd.Dir = dir
	o, err := cmd.CombinedOutput()
	return string(o), err
}
