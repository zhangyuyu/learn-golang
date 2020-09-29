package main

import (
	"errors"
	"fmt"
)

func main() {
	for _, i := range []int{7, 42} {
		if r, e := handleError1(i); e != nil {
			fmt.Println("handleError1 failed: ", e)
		} else {
			fmt.Println("handleError1 worked: ", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := handleError2(i); e != nil {
			fmt.Println("handleError2 failed: ", e)
		} else {
			fmt.Println("handleError2 worked: ", r)
		}
	}

	_, e := handleError2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}

}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func handleError2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func handleError1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}

	return arg + 3, nil
}
