package main

import "math/rand"

func drawN(p *player, n int) ([]*Card, map[Rarity]int) {
	p.Draw(n)

	results := make([]*Card, n)
	summary := make(map[Rarity]int)
	for i := 0; i < n; i++ {
		results[i] = draw()
		summary[results[i].Rarity]++
	}

	// 変数resultsとsummaryの値を戻り値として返す
	return results, summary
}

func draw() *Card {
	num := rand.Intn(100)

	switch {
	case num < 80:
		return &Card{Rarity: RarityN, name: "スライム"}
	case num < 95:
		return &Card{Rarity: RarityR, name: "オーク"}
	case num < 99:
		// TODO: RarityフィールドがRaritySRで
		// nameフィールドが"ドラゴン"であるcard構造体のポインタを返す
		return &Card{Rarity: RaritySR, name: "ドラゴン"}
	default:
		return &Card{Rarity: RarityXR, name: "イフリート"}
	}
}
