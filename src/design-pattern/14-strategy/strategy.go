package strategy

// IMethod 实现此接口，即为一个策略
type IMethod interface {
	do(int, int) int
}

type add struct {
}

func (*add) do(a int, b int) int {
	return a + b
}

type sub struct {
}

func (*sub) do(a int, b int) int {
	return a - b
}

type Operator struct {
	method IMethod
}

func (o *Operator) setStrategy(strategy string) {
	if strategy == "add" {
		o.method = &add{}
	} else {
		o.method = &sub{}
	}
}
func (o *Operator) calculate(a, b int) int {
	return o.method.do(a, b)
}
