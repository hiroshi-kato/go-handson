package main

import (
	"math/rand"
	"time"
)

func main() {
	list := []string{"かとう", "こいずみ", "岩元"}

	rand.Seed(time.Now().Unix())
	num := rand.Intn(len(list))

	println(list[num])
}
