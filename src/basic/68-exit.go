package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("!")

	os.Exit(3)
}

// 注意，程序中的 ! 永远不会被打印出来。
/*
自定义退出状态码，可以在脚本中定义自己的退出状态代码，然后使用echo $?检查。

0 命令成功完成
1 通常的未知错误
2 误用Shell命令
*/
