package main

import (
	"fmt"
	"os"

	"github.com/hiroshi-kato/gotcha-ja"
)

func main() {
	p := gotcha.NewPlayer(10, 100)

	n := inputN(p)
	results, summary := gotcha.DrawN(p, n)

	saveResults(results)
	saveSummary(summary)
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
		fmt.Printf("1以上&dイカの数を入力してください\n", max)
	}
	return n
}

func saveResults(results []*gotcha.Card) {
	f, err := os.Create("results.txt")
	if err != nil {
		fmt.Println((err))
		return
	}

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	for _, result := range results {
		fmt.Fprintln(f, result)
	}
}

func saveSummary(summary map[gotcha.Rarity]int) {
	f, err := os.Create("summary.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	for rarity, count := range summary {
		fmt.Fprintf(f, "%s:%d\n", rarity.String(), count)
	}
}
