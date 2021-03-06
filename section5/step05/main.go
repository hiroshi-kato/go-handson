// STEP05: 初期ガチャチケットの枚数をフラグで渡そう

package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/hiroshi-kato/gotcha-ja"
)

var flagCoin = flag.Int("coin", 0, "コインの初期枚数")

func init() {
	// TODO: flagCoinに"coin"という名前のフラグを設定する
	// デフォルト値は0で説明は"コインの初期枚数"
	// flag.IntVar(&flagCoin, "coin", 0, "コインの初期枚数")
}

func main() {
	// TODO: フラグをパースする
	flag.Parse()

	tickets := initialTickets()
	p := gotcha.NewPlayer(tickets, *flagCoin)

	n := inputN(p)
	results, summary := gotcha.DrawN(p, n)

	saveResults(results)
	saveSummary(summary)
}

func initialTickets() int {
	if flag.NArg() == 0 {
		fmt.Fprintln(os.Stderr, "ガチャチケットの枚数を入力してください")
		os.Exit(1)
	}

	// TODO: フラグを除いて1つめのプログラム引数を取得してint型に変換する
	num, err := strconv.Atoi(flag.Arg(0))

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	return num
}

func inputN(p *gotcha.Player) int {

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

func saveResults(results []*gotcha.Card) {
	f, err := os.Create("results.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}()

	for _, result := range results {
		fmt.Fprintln(f, result)
	}
}

func saveSummary(summary map[gotcha.Rarity]int) {
	f, err := os.Create("summary.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}()

	for rarity, count := range summary {
		fmt.Fprintf(f, "%s %d\n", rarity.String(), count)
	}
}
