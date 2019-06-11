package main

import (
	"github.com/gogf/gf/g"
	_ "github.com/hequan2017/coypus/boot"
)

func main() {
	s := g.Server()
	_ = s.Run()
}
