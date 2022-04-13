package wordz

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

var Hello = "This is package wordz"
var Prefix = "Random word: "
var Words = []string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten"}

func Random() string {
	max := len(Words)
	r, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
	return get(r.Int64())
}

func get(index int64) string {
	return Prefix + Words[index]
}

func init() { //"эта функция вызывается один раз при импорте пакета wordz в файле main.go"
	fmt.Println("Function init in package wordz")
}
