package box_test

import (
	"testing"

	"github.com/qinhan-shu/study/go/game_award"
)

func TestOpenBox(t *testing.T) {
	b := box.Box{
		Items: []*box.Item{
			&box.Item{T: "GOLD", Des: 100, Weight: 10},
			&box.Item{T: "Diamond", Des: 1000, Weight: 110},
			&box.Item{T: box.BOXITEM, Des: 1000, Weight: 110},
		},
	}

	t.Log("开箱结果")
	for _, item := range b.Open() {
		t.Log(item)
	}
}
