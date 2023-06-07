package memento

//备忘录模式使用三个对象: Memento、Originator 和 CareTaker。
// Memento 包含了要被恢复的对象的状态。
// Originator 创建并在 Memento 对象中存储状态。
// Caretaker 对象负责从 Memento 中恢复对象的状态。

// Originator 用于保存数据
type Originator struct {
	state string
}

// SaveToMemento 创建快照
func (o *Originator) SaveToMemento() Memento {
	return Memento{content: o.state}
}

// RestoreFromMemento 从快照中恢复
func (o *Originator) RestoreFromMemento(s Memento) {
	o.state = s.GetText()
}

type Memento struct {
	content string
}

// GetText 获取数据
func (s *Memento) GetText() string {
	return s.content
}

type Caretaker struct {
	mementoList []Memento
}

func (c *Caretaker) Add(memento Memento) {
	c.mementoList = append(c.mementoList, memento)
}

func (c *Caretaker) Get(index int) Memento {
	return c.mementoList[index]
}
