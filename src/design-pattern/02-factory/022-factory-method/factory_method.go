package factory

// 工厂方法在简单工厂基础上，给每种Fruit增加了一个工厂

// IFruitFactory 水果工厂接口
type IFruitFactory interface {
	// CreateFruit 生产水果
	CreateFruit() Fruit
}

//AppleFactory 苹果工厂，实现FruitFactory接口
type AppleFactory struct{}

func (a *AppleFactory) CreateFruit() Fruit {
	return &Apple{}
}

//OrangeFactory 橘子工厂，实现FruitFactory接口
type OrangeFactory struct{}

func (o *OrangeFactory) CreateFruit() Fruit {
	return &Orange{}
}

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

func GetFruitFactory(fruitType string) IFruitFactory {
	switch fruitType {
	case "apple":
		return &AppleFactory{}
	case "orange":
		return &OrangeFactory{}
	}
	return nil
}
