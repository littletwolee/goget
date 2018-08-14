package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	ccmd "github.com/littletwolee/commons/cmd"
)

type cmd struct {
	cmd *exec.Cmd
	p   string
}

func newCmd() *cmd {
	p, err := getPath()
	if err != nil {
		return nil
	}
	projectName := p[strings.LastIndex(p, "/")+1:]
	command := fmt.Sprintf("./%s", projectName)
	return &cmd{cmd: ccmd.GetCmd().Command(command, "./"), p: p}
}
func getPath() (string, error) {
	return filepath.Abs(filepath.Dir(os.Args[0])) //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
}
func (c *cmd) run() error {
	if c.cmd.Process != nil {
		if err := c.cmd.Process.Kill(); err != nil {
			return err
		}
	}
	return c.cmd.Run()
}

func (c *cmd) stop() error {
	if c.cmd.Process != nil {
		return c.cmd.Process.Kill()
	}
	return nil
}
