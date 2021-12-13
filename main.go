package main

import (
	"fmt"
	"math/rand"
	"time"
)

// main関数からすべてのコードが実行される
func main() {
	// 複数行コメントは /* */ で指定
	/*
		print("Hello, world!")
	*/

	// 標準出力に出力する
	fmt.Print("Hello, world!")
	// 改行あり
	fmt.Println("Hello, world!")

	// あたりはずれ
	t := time.Now().Unix()
	rand.Seed(t)
	u := rand.Intn(10)
	if u == 7 {
		fmt.Println(u)
		fmt.Println("あたり")
	} else {
		fmt.Println(u)
		fmt.Println("はずれ")
	}

	// レア度を作ってみる
	switch u {
	case 0:
		fmt.Println(("u is 0"))
	case 1, 2:
		fmt.Println(("u is 1 or 2"))
	}

	// レア度ごとに出る確率を買えてみよう
	switch {
	case u == 1:
		fmt.Println(("uis 1"))
	case u == 2:
		fmt.Println(("uis 2"))
	default:
		fmt.Println(("default"))
	}

}
