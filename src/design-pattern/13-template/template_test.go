package template

import "testing"

func TestTemplate_doCook(t *testing.T) {
	doCook(&Tomato{})
	doCook(&Egg{})
}
