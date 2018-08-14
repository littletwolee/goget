package main

import (
	"fmt"
	"os"
	"strings"

	ccmd "github.com/littletwolee/commons/cmd"
	"github.com/littletwolee/commons/logger"
)

type b struct{}

func (b *b) build() (string, error) {
	_, strerr, err := ccmd.GetCmd().ExecCommand("go", "./", []string{"build"}, false)
	if err != nil {
		return "", err
	}
	return strerr, nil
}

func (b *b) getPkg(pkg string) error {
	command := "git"
	pars := []string{"clone", "--progress", fmt.Sprintf(`https://%s`, pkg), pkg}
	rootdir := fmt.Sprintf(`%s/src/`, os.ExpandEnv("$GOPATH"))
	fmt.Printf("Downloading %s", pkg)
	_, _, err := ccmd.GetCmd().ExecCommand(command, rootdir, pars, true)
	return err
}

func (b *b) getPkgName(errStr string) string {
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
func (b *b) run() {
	for {
		errStr, err := b.build()
		if err != nil {
			logger.Console().Panic(err)
		}
		if pkg := b.getPkgName(errStr); pkg != "" {
			if err := b.getPkg(pkg); err != nil {
				logger.Console().Panic(err)
			}
			fmt.Printf("pkg \"%s\" download success!\n", pkg)
			continue
		}
		break
	}
}
