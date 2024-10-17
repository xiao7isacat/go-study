package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

type shellExec struct {
	Command string
	stdout  *bytes.Buffer
	stderr  *bytes.Buffer
}

func (this *shellExec) Run() error {
	cmd := exec.Command("/bin/bash", "-c", this.Command)
	cmd.Stdout = this.stdout
	cmd.Stderr = this.stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func main() {
	var test shellExec
	test.Command = "ls /opt"
	test.stdout = new(bytes.Buffer)
	test.stderr = new(bytes.Buffer)
	if err := test.Run(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("err", test.stderr.Bytes())
	fmt.Println("out", string(test.stdout.Bytes()))

}
