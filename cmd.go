package main

import (
	"fmt"
	"github.com/littletwolee/commons"
	"os"
	"strings"
)

func Build() (string, error) {
	_, strerr, err := commons.GetCmd().ExecCommand("go", "./", []string{"build"}, false)
	if err != nil {
		return "", err
	}
	return strerr, nil
}

func Get(output string) {
	for _, v := range strings.Split(output, "\n") {
		if strings.Contains(v, "in any of") {
			pack := strings.Split(v, `"`)[1]
			if len(strings.Split(pack, `/`)) > 3 {
				packarr := strings.Split(pack, `/`)
				pack = fmt.Sprintf(`%s/%s/%s`, packarr[0], packarr[1], packarr[2])
			}
			command := "git"
			pars := []string{"clone", "--progress", fmt.Sprintf(`https://%s`, pack), pack}
			rootdir := fmt.Sprintf(`%s/src/`, os.ExpandEnv("$GOPATH"))
			commons.GetLogger().OutMsg(fmt.Sprintf("Downloading %s", pack))
			commons.GetCmd().ExecCommand(command, rootdir, pars, true)
			commons.GetLogger().OutMsg(fmt.Sprintf("Download %s success", pack))
		}
	}

}

func Run() {
	dir, _, err := commons.GetCmd().ExecCommand("pwd", "./", nil, false)
	var (
		projectName string
	)
	if err != nil {
		commons.GetLogger().OutErr(err)
	}
	if strings.Contains(dir, "/") {
		projectName = dir[strings.LastIndex(dir, "/")+1:]
	}
	buildstr, err := Build()
	if buildstr != "" {
		Get(buildstr)
	}

}
