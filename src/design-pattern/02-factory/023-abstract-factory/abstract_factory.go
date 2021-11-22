package factory

// 抽象工厂方法，在工厂工厂基础上，合并了工厂。
// 如果每种水果都需要一个工厂，工厂也会膨胀，所以可以对某些水果的工厂进行合并

// IFruitFactory 水果工厂接口
type IFruitFactory interface {
	// CreateFruit 生产水果
	CreateFruit() Fruit
}

//MyFruitFactory 自己开的水果工厂，实现FruitFactory接口
type MyFruitFactory struct{}

// Fruit 水果接口
type Fruit interface {
	name() string
}

// Apple 苹果，实现Fruit接口
type Apple struct{}

func (f *Apple) name() string {
	return "Apple"
}

// Orange 橘子，实现Fruit接口
type Orange struct{}

func (f *Orange) name() string {
	return "Orange"
}

func (m *MyFruitFactory) GetApple() Fruit {
	return &Apple{}
}

func (m *MyFruitFactory) GetOrange() Fruit {
	return &Orange{}
}
