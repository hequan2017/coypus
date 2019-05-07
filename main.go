package main

import (
	"github.com/gogf/gf/g"
	_ "github.com/hequan2017/coypus/boot"
	_ "github.com/hequan2017/coypus/router"
)

func main() {
	_ = g.Server().Run()
}
