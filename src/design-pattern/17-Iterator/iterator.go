package iterator

// Iterator 迭代器接口
type Iterator interface {
	HasNext() bool
	Next()
	CurrentItem() interface{}
}

// ContainerIterator 迭代器实现
type ContainerIterator struct {
	container Container
	index     int
}

func (a *ContainerIterator) HasNext() bool {
	return a.index < len(a.container)
}

func (a *ContainerIterator) Next() {
	a.index++
}

func (a *ContainerIterator) CurrentItem() interface{} {
	return a.container[a.index]
}

// Container 当前容器
type Container []int

// GetIterator 获取迭代器
func (a *Container) GetIterator() Iterator {
	return &ContainerIterator{}
}
