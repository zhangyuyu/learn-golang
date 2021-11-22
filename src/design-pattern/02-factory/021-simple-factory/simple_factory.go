package factory

// IFruitFactory 水果工厂接口
type IFruitFactory interface {
	// CreateFruit 生产水果
	CreateFruit() Fruit
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

func GetFruit(fruitType string) Fruit {
	switch fruitType {
	case "apple":
		return &Apple{}
	case "orange":
		return &Orange{}
	}
	return nil
}
