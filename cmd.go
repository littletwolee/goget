package main

import (
	"fmt"
	"os/exec"
	"strings"
	"unsafe"

	ccmd "github.com/littletwolee/commons/cmd"
)

type cmd struct {
	exec.Cmd
}

func newCmd() *cmd {
	p, err := getPath()
	if err != nil {
		return nil
	}
	projectName := p[strings.LastIndex(p, "/")+1:]
	command := fmt.Sprintf("./%s", projectName)
	return (*cmd)(unsafe.Pointer(ccmd.GetCmd().Command(command, "./")))
}

func (c *cmd) start() error {
	if c.Process != nil {
		if err := c.Process.Kill(); err != nil {
			return err
		}
		c = newCmd()
	}
	return c.Run()
}

func (c *cmd) stop() error {
	if c.Process != nil {
		return c.Process.Kill()
	}
	return nil
}
