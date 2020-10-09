package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t.Format(time.RFC3339))

	t1, _ := time.Parse(time.RFC3339, "2020-09-30T11:54:33+08:00")
	fmt.Println(t1)

	fmt.Println(t.Format("3:04PM"))
	fmt.Println(t.Format("2006-01-02"))

	form := "3 04 PM"
	t2, _ := time.Parse(form, "8 41 PM")
	fmt.Println(t2)

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
}
