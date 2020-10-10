package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"

	"learn-golang/src/utils"
)

func main() {
	dateCmd := exec.Command("date")

	dateOut, err := dateCmd.Output()
	utils.Check(err)
	fmt.Println("> date")
	fmt.Println(string((dateOut)))

	grepCmd := exec.Command("grep", "hello")
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep lalala\ngppdbye grep"))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	utils.Check(err)
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}

// 生成外部进程，当我们需要在运行的 Go 流程中访问的外部流程时，便可以执行此操作。
