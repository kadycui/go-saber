package utils


import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func GetParentPth() {
	exePath, err := exec.LookPath(os.Args[0])
	if err != nil {
		panic(err)
	}

	exeDir := filepath.Dir(exePath)
	parentDir := filepath.Dir(exeDir)

	fmt.Println("上级目录:", parentDir)
}