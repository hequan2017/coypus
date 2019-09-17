package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/hequan2017/coypus/app/api/a_menu"
	"github.com/hequan2017/coypus/app/api/a_role"
	"github.com/hequan2017/coypus/app/api/a_user"
	"github.com/hequan2017/coypus/library/jwt"
	"github.com/hequan2017/coypus/library/permission"
	"net/http"
)

// 统一路由注册.
func init() {
	// 用户模块 路由注册 - 使用执行对象注册方式
	s := g.Server()
	s.BindHookHandlerByMap("/*any", map[string]ghttp.HandlerFunc{
		"BeforeServe": func(r *ghttp.Request) {

			r.Response.CORS(ghttp.CORSOptions{
				AllowOrigin:      "*",
				AllowMethods:     "POST, GET, OPTIONS, PUT, DELETE,UPDATE,",
				AllowCredentials: "false",
				MaxAge:           1728000,
				AllowHeaders:     "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma",
				ExposeHeaders:    "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar",
			})

			if r.Request.Method == "OPTIONS" {
				_ = r.Response.WriteJson(g.Map{
					"code": http.StatusOK,
					"msg":  "",
					"data": nil,
				})
				r.ExitAll()
			}

			jwt.JWT(r)                     // 验证 token 是否正确
			permission.CasbinMiddleware(r) // 权限验证
		}})

	s.BindHandler("/token", a_user.Login)
	s.BindHandler("/userInfo", a_user.UserInfo)
	s.BindHandler("/menu", a_user.Menu)

	s.BindObjectRest("/api/v1/users/*id", new(a_user.Controller))
	s.BindObjectRest("/api/v1/roles/*id", new(a_role.Controller))
	s.BindObjectRest("/api/v1/menus/*id", new(a_menu.Controller))

}
