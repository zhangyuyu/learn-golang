package command

import (
	"fmt"
)

type Command func() error

// StartCommandFunc 游戏开始命令
func StartCommandFunc() Command {
	return func() error {
		fmt.Println("Game Start...")
		return nil
	}
}

// ArchiveCommandFunc 游戏存档命令
func ArchiveCommandFunc() Command {
	return func() error {
		fmt.Println("Game archieve...")
		return nil
	}
}
