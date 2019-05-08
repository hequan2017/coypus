package main

import (
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/hequan2017/coypus/library/jwt"
	"github.com/hequan2017/coypus/library/permission"
	_ "github.com/hequan2017/coypus/boot"

)

func main() {
	s := g.Server()
	s.BindHookHandlerByMap("/api/*any", map[string]ghttp.HandlerFunc{
		"BeforeServe": func(r *ghttp.Request) {
			jwt.JWT(r)     // 验证 token 是否正确
			permission.CasbinMiddleware(r)  // 权限验证
		}})
	s.BindHookHandlerByMap("/*any", map[string]ghttp.HandlerFunc{
		"BeforeServe": func(r *ghttp.Request) {
			r.Response.CORSDefault()
		}})

	_ = s.Run()
}
