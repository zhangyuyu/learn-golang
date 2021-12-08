package compose

import "fmt"

// IOrganization 组织接口，都实现统计人数的功能
type IOrganization interface {
	Count() int
}

type Employee struct {
	Name string
}

func (e *Employee) Count() int {
	return 1
}

// Department 部门
type Department struct {
	Name            string
	SubOrganization []IOrganization
}

func (d *Department) Count() int {
	c := 0
	for _, org := range d.SubOrganization {
		c += org.Count()
	}
	return c
}

func (d *Department) AddSub(org IOrganization) {
	d.SubOrganization = append(d.SubOrganization, org)
}

func NewOrganization() IOrganization {
	root := &Department{Name: "root"}
	for i := 0; i < 10; i++ {
		root.AddSub(&Employee{})
		root.AddSub(&Department{
			Name:            fmt.Sprintf("sub-%d", i),
			SubOrganization: []IOrganization{&Employee{}}})
	}
	return root
}
