package prototype

import (
	"encoding/json"
	"time"
)

// Keyword 搜索关键字
type Keyword struct {
	// 注意要为大写，不然json.Unmarshal不出来
	Word      string
	Visit     int
	UpdatedAt *time.Time
}

// Clone 这里使用序列化与反序列化的方式深拷贝
func (k *Keyword) Clone() *Keyword {
	var newKeyword Keyword
	b, _ := json.Marshal(k)
	_ = json.Unmarshal(b, &newKeyword)
	return &newKeyword
}
