package main

import (
	"testing"
)

func IntMin(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		t.Errorf("imtMin(2, -2) = %d; want -2", ans)
	}
}

// 文件名格式必须是name_test.go，测试函数名称必须以Test开头，并且传入参数必须是*testing.T
/*
$ go test -v src/feature/58-testing_test.go
=== RUN   TestIntMinBasic
--- PASS: TestIntMinBasic (0.00s)
*/
