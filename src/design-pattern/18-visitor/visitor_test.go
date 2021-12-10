package visitor

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	c := &CustomerCol{}
	c.Add(&EnterpriseCustomer{"字节"})
	c.Add(&EnterpriseCustomer{"腾讯"})
	c.Add(&IndividualCustomer{"yukki"})
	c.Accept(&Visitor1{})
	fmt.Println()
	c.Accept(&Visitor2{})
}
