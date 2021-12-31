// STEP01: ファイルを分けよう

package main

import (
	"fmt"
)

func main() {
	// TODO: ガチャ券10枚、コイン100枚を持ったプレイヤーを作る
	p = NewPlayer(10, 100)

	n := inputN(p)
	// TODO: gotcha.DrawN関数を呼び、変数resultsとsummaryに結果を代入する
	[results,summary] := gotcha.DrawN(p,n)

	fmt.Println(results)
	fmt.Println(summary)
}

// TODO: 引数の型をgotcha.Playerのポインタにする
func inputN(p player) int {

	max := p.DrawableNum()
	fmt.Printf("ガチャを引く回数を入力してください（最大:%d回）\n", max)

	var n int
	for {
		fmt.Print("ガチャを引く回数>")
		fmt.Scanln(&n)
		if n > 0 && n <= max {
			break
		}
		fmt.Printf("1以上%d以下の数を入力してください\n", max)
	}

	return n
}
