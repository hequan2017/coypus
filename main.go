package main

import (
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/net/ghttp"
	_ "github.com/hequan2017/coypus/boot"
	"github.com/hequan2017/coypus/library/jwt"
	"github.com/hequan2017/coypus/library/permission"
)

func main() {
	s := g.Server()
	s.BindHookHandlerByMap("/*any", map[string]ghttp.HandlerFunc{
		"BeforeServe": func(r *ghttp.Request) {
			r.Response.CORSDefault()       //关闭 CORS
			jwt.JWT(r)                     // 验证 token 是否正确
			permission.CasbinMiddleware(r) // 权限验证
		}})

	_ = s.Run()
}
