package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomListNum(list []string) int {
	t := time.Now().UnixMicro()
	rand.Seed(t)
	return rand.Intn(len(list))
}

func getGatcha(gatchaList []string) string {
	return gatchaList[randomListNum(gatchaList)]
}

func startGatcha(gatchaList []string) {
	var answer int
	fmt.Println("ガチャを回しますか？（1:回す 2:やめる）")
	fmt.Scanln(&answer)
	if answer == 1 {
		var gatchaTimes int
		fmt.Println("何回回しますか？")
		fmt.Scanln(&gatchaTimes)
		for i := 0; i < gatchaTimes; i++ {
			fmt.Println(getGatcha((gatchaList)))
		}
	} else {
		fmt.Println("やめたよ")
	}
}

func main() {

	// ガチャの内容を定義
	gatchaList := []string{
		"悟空",
		"ベジータ",
		"ベジータ",
		"ピッコロ",
		"ピッコロ",
		"ピッコロ",
		"クリリン",
		"クリリン",
		"クリリン",
		"クリリン",
		"ヤムチャ",
		"ヤムチャ",
		"ヤムチャ",
		"ヤムチャ",
		"ヤムチャ",
		"ヤムチャ",
	}

	// ガチャを回す
	startGatcha(gatchaList)
}
