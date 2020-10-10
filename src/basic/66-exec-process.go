package main

import (
	"os"
	"os/exec"
	"syscall"

	"learn-golang/src/utils"
)

func main() {
	binary, err := exec.LookPath("ls")
	utils.Check(err)

	args := []string{"ls", "-a", "-l", "-h"}

	env := os.Environ()

	err = syscall.Exec(binary, args, env)
	utils.Check(err)
}

// exec，完全替代当前的 Go 进程
