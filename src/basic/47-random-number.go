package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 默认情况下，给定的种子是确定的，每次都会产生相同的随机数数字序列。
	fmt.Println(rand.Intn(100), rand.Intn(100))
	fmt.Println(rand.Float64(), rand.Float64())

	//  要产生不同的数字序列，需要给定一个不同的种子。
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	fmt.Println(r1.Intn(100), r1.Intn(100))
	fmt.Println(r1.Float64(), r1.Float64())

	// 如果使用相同种子生成的随机数生成器，会生成相同的随机数序列。
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Println(r2.Intn(100), r2.Intn(100))

	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), r3.Intn(100))
}
