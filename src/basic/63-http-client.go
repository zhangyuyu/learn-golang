package main

import (
	"bufio"
	"fmt"
	"net/http"

	"learn-golang/src/utils"
)

func main() {
	resp, err := http.Get("http://gobyexample.com")
	utils.Check(err)

	defer resp.Body.Close()

	fmt.Println("Response Status: ", resp.Status)

	scanner := bufio.NewScanner(resp.Body)

	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

}
