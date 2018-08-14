package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/littletwolee/commons"
)

func build() (string, error) {
	_, strerr, err := commons.GetCmd().ExecCommand("go", "./", []string{"build"}, false)
	if err != nil {
		return "", err
	}
	return strerr, nil
}

func getPkg(pkg string) error {
	command := "git"
	pars := []string{"clone", "--progress", fmt.Sprintf(`https://%s`, pkg), pkg}
	rootdir := fmt.Sprintf(`%s/src/`, os.ExpandEnv("$GOPATH"))
	fmt.Printf("Downloading %s", pkg)
	_, _, err := commons.GetCmd().ExecCommand(command, rootdir, pars, true)
	return err
}

func getPkgName(errStr string) string {
	pkgKeyWord := "in any of"
	if errStr != "" && strings.Contains(errStr, pkgKeyWord) {
		pkg := strings.Split(errStr, `"`)[1]
		pkgSplit := strings.Split(pkg, `/`)
		if len(pkgSplit) > 3 {
			pkg = strings.Join(pkgSplit[:3], `/`)
		}
		return pkg
	}
	return ""
}
func run(outChan chan bool) {
	//pkgChan := make(chan string)
	for {
		errStr, err := build()
		if err != nil {
			commons.Console().Panic(err)
		}
		if pkg := getPkgName(errStr); pkg != "" {
			if err := getPkg(pkg); err != nil {
				commons.Console().Panic(err)
			}
			fmt.Printf("pkg \"%s\" download success!\n", pkg)
			continue
		}
		outChan <- true
	}
}
