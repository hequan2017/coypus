package main

import (
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/net/ghttp"
	_ "github.com/hequan2017/coypus/boot"
	"github.com/hequan2017/coypus/library/jwt"
	_ "github.com/hequan2017/coypus/router"
)

func main() {
	p := "/*any"
	s := g.Server()
	s.BindHookHandlerByMap(p, map[string]ghttp.HandlerFunc{
		"BeforeServe": func(r *ghttp.Request) {
			r.Response.CORSDefault()
			jwt.JWT(r)
		}})

	_ = s.Run()
}
