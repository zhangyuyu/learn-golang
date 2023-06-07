package interpreter

import (
	"fmt"
	"testing"
)

func TestExpression(t *testing.T) {
	robert := &TerminalExpression{data: "Robert"}
	john := &TerminalExpression{data: "John"}

	isMale := &OrExpression{expr1: robert, expr2: john}

	mary := &TerminalExpression{data: "Mary"}
	married := &TerminalExpression{data: "Married"}
	isMarriedWoman := &AndExpression{expr1: mary, expr2: married}

	fmt.Println("John is male? ", isMale.interpret("John"))
	fmt.Println("Mary is married? ", isMarriedWoman.interpret("Married Mary"))
}
