package visitor

import "fmt"

type Customer interface {
	Accept(visitor Visitor)
}

// EnterpriseCustomer 企业级客户
type EnterpriseCustomer struct {
	name string
}

func (c *EnterpriseCustomer) Accept(visitor Visitor) {
	visitor.Visit(c)

}

// IndividualCustomer 个人客户
type IndividualCustomer struct {
	name string
}

func (c *IndividualCustomer) Accept(visitor Visitor) {
	visitor.Visit(c)

}

// CustomerCol 客户群
type CustomerCol struct {
	customers []Customer
}

func (c *CustomerCol) Add(customer Customer) {
	c.customers = append(c.customers, customer)
}
func (c *CustomerCol) Accept(visitor Visitor) {
	for _, customer := range c.customers {
		customer.Accept(visitor)
	}
}

// Visitor 访问器接口
type Visitor interface {
	Visit(customer Customer)
}

// Visitor1 访问者1
type Visitor1 struct{}

func (*Visitor1) Visit(customer Customer) {
	switch c := customer.(type) {
	case *EnterpriseCustomer:
		fmt.Printf("服务企业级客户：%s\n", c.name)
	case *IndividualCustomer:
		fmt.Printf("服务个人客户：%s\n", c.name)
	}
}

// Visitor2 访问者2
type Visitor2 struct{}

func (*Visitor2) Visit(customer Customer) {
	switch c := customer.(type) {
	case *EnterpriseCustomer:
		fmt.Printf("审计企业级客户：%s\n", c.name)
	case *IndividualCustomer:
		fmt.Printf("审计个人客户：%s\n", c.name)
	}
}
