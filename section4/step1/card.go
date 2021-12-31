package main

type Rarity string

const (
	RarityN  Rarity = "N"
	RarityR  Rarity = "R"
	RaritySR Rarity = "SR"
	RarityXR Rarity = "XR"
)

func (r Rarity) String() string {
	return string(r)
}

type Card struct {
	Rarity Rarity // レア度
	name   string // 名前
}

func (c *Card) String() string {
	// TODO: レア度:名前のように文字列を作る
	// 例："SR:ドラゴン"
	return c.Rarity.String() + ":" + c.name
}
