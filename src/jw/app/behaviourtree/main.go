package main

import (
	"github.com/davyxu/golog"
	b3 "github.com/magicsea/behavior3go"
	. "github.com/magicsea/behavior3go/config"
	. "github.com/magicsea/behavior3go/core"
	. "github.com/magicsea/behavior3go/examples/share"
	. "github.com/magicsea/behavior3go/loader"
)

var logger = golog.New("bh-tree")

func main() {

	treeConfig, ok := LoadTreeCfg("tree.json")
	if !ok {
		logger.Errorf("load tree cfg failed")
	}

	maps := b3.NewRegisterStructMaps()
	maps.Register("Log", new(LogTest))

	tree := CreateBevTreeFromConfig(treeConfig, maps)
	tree.Print()

	board := NewBlackboard()

	for i := 0; i < 5; i++ {
		tree.Tick(i, board)
	}

}
