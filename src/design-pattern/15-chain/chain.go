package chain

import "fmt"

// SensitiveWordFilter 敏感词过滤器接口
type SensitiveWordFilter interface {
	Filter(content string) bool
}

// WordFilter 职责链
type WordFilter struct {
	filters []SensitiveWordFilter
}

func (w *WordFilter) AddFilter(filter SensitiveWordFilter) {
	w.filters = append(w.filters, filter)
}

// Filter 执行过滤
func (w *WordFilter) Filter(content string) bool {
	for _, filter := range w.filters {
		// 如果发现敏感直接返回结果
		if filter.Filter(content) {
			return true
		}
	}
	return false
}

// AdSensitiveWordFilter 广告
type AdSensitiveWordFilter struct{}

func (a *AdSensitiveWordFilter) Filter(content string) bool {
	fmt.Println("Filter Ad content...")
	return false
}

// PoliticalWordFilter 政治敏感
type PoliticalWordFilter struct{}

func (p *PoliticalWordFilter) Filter(content string) bool {
	fmt.Println("Filter Political content...")
	return true
}
