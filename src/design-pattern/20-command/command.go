package command

import (
	"fmt"
)

// ICommand 命令接口
type ICommand interface {
	Execute() error
}

// StartCommand 游戏开始命令
type StartCommand struct{}

func (s *StartCommand) Execute() error {
	fmt.Println("Game Start...")
	return nil
}

// ArchiveCommand 游戏存档命令
type ArchiveCommand struct{}

func (a *ArchiveCommand) Execute() error {
	fmt.Println("Game archieve...")
	return nil
}
