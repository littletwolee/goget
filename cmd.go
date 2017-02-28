package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func ExecCommand(cmd *exec.Cmd, isStdout bool) (string, error) {
	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	if isStdout {
		cmd.Stdout = os.Stdout // 重定向标准输出
		cmd.Stderr = os.Stderr // 重定向标准输出
	}
	err := cmd.Run()
	out := buf.Bytes()
	if err != nil {
		return fmt.Sprintf("%s", out), err
	}
	return fmt.Sprintf("%s", out), nil
}
