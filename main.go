package main

import (
	"github.com/gogf/gf/g"
	_ "github.com/hequan2017/coypus/boot"
)

func main() {
	s := g.Server()
	s.SetPort(8001)
	_ = s.Run()
}
