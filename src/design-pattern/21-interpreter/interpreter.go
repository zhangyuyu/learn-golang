package interpreter

import (
	"strings"
)

// Expression 表达式接口
type Expression interface {
	interpret(content string) bool
}

type TerminalExpression struct {
	data string
}

func (t *TerminalExpression) interpret(content string) bool {
	return strings.Contains(content, t.data)
}

type OrExpression struct {
	expr1 Expression
	expr2 Expression
}

func (o *OrExpression) interpret(content string) bool {
	return o.expr1.interpret(content) || o.expr2.interpret(content)
}

type AndExpression struct {
	expr1 Expression
	expr2 Expression
}

func (a *AndExpression) interpret(content string) bool {
	return a.expr1.interpret(content) && a.expr2.interpret(content)
}
