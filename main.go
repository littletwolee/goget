package main

import (
	"bytes"
	"fmt"
	constant "goget/constant"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	cmdstrs := os.Args
	if len(cmdstrs) < 1 {
		fmt.Println(constant.CMD_INPUT_ERROR)
	}
	path, err := filepath.Abs(cmdstrs[1])
	if err != nil {
		fmt.Println(constant.PATH_UNEXISTS_ERROR)
	}
	if !checkFileIsExist(path) {
		fmt.Println(constant.PATH_UNEXISTS_ERROR)
	}
	cmd := exec.Command("go", "build", path)
	cmd.Dir = path[:strings.LastIndex(path, `/`)]
	result, err := execCommand(cmd, false)
	if err == nil {

	}
	for _, v := range strings.Split(result, "\n") {
		if strings.Contains(v, "in any of") {
			pack := strings.Split(v, `"`)[1]
			if len(strings.Split(pack, `/`)) > 3 {
				packarr := strings.Split(pack, `/`)
				pack = fmt.Sprintf(`%s/%s/%s`, packarr[0], packarr[1], packarr[2])
			}
			cmd = exec.Command("git", "clone", "--progress", fmt.Sprintf(`https://%s`, pack), pack)
			cmd.Dir = fmt.Sprintf(`%s/src/`, os.ExpandEnv("$GOPATH"))
			fmt.Printf("Downloading %s", pack)
			result, err := execCommand(cmd, true)
			if err != nil {
				os.Stderr.Write([]byte(result))
			}
			fmt.Printf("Download %s success", pack)
		}
	}
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func execCommand(cmd *exec.Cmd, isStdout bool) (string, error) {
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
