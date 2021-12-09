package template

import "fmt"

type Cooker interface {
	fire()
	cook()
	outfire()
}

// CookMenu 类似于一个抽象类
type CookMenu struct {
}

func (c *CookMenu) fire() {
	fmt.Println("开火")
}

// 做菜，留给具体的子类实现
func (c *CookMenu) cook() {

}

func (c *CookMenu) outfire() {
	fmt.Println("关火")
}

// 封装具体步骤
func doCook(cooker Cooker) {
	cooker.fire()
	cooker.cook()
	cooker.outfire()
}

type Tomato struct {
	CookMenu
}

func (t *Tomato) cook() {
	fmt.Println("做西红柿")
}

type Egg struct {
	CookMenu
}

func (t *Egg) cook() {
	fmt.Println("做鸡蛋")
}
