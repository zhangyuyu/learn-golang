package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"custom_page"`
	Fruits []string `json:"custom_fruits"`
}

func main() {
	marshall()
	unMarshall()

	encoder := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	err := encoder.Encode(d)
	if err != nil {
		panic(err)
	}
}

func marshall() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcB, _ := json.Marshal([]string{"apple", "peach", "pear"})
	fmt.Println(string(slcB))

	mapB, _ := json.Marshal(map[string]int{"apple": 5, "lettuce": 7})
	fmt.Println(string(mapB))

	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))
}

func unMarshall() {
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}

	err := json.Unmarshal(byt, &dat)
	if err != nil {
		panic(err)
	}
	fmt.Println(dat)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	str := `{"Page": 1, "Fruits": ["apple", "peach"]}`
	res := response1{}
	err = json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])
}
