package b

import "github.com/bitgaming/packr/v2"

func init() {
	packr.New("b-box", "../c")
	packr.New("cb-box", "../c")
}
