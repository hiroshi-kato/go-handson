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
	// カンマ区切りで複数の値を出力する。
	// 各値間にスペースが入る
	fmt.Println("A", 100, true)

	// あたりはずれ
	// 現在日時をキーにする
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
		fmt.Println("u is 0")
	case 1, 2:
		fmt.Println("u is 1 or 2")
	}

	// レア度ごとに出る確率を買えてみよう
	switch {
	case u == 1:
		fmt.Println("u is 1")
	case u == 2:
		fmt.Println("u is 2")
	default:
		fmt.Println("default")
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for i, v := range []int{1, 2, 3} {
		fmt.Println(i, v)
	}

	var price int
	fmt.Print("値段")
	fmt.Scanln(&price)
	fmt.Printf("%d円\n", price)
}
