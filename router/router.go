package router

import (
	"github.com/gogf/gf/g"
	"github.com/hequan2017/coypus/app/api/a_user"
)

// 统一路由注册.
func init() {
	// 用户模块 路由注册 - 使用执行对象注册方式
	s := g.Server()
	s.BindHandler("/token", a_user.Login)
	s.BindObjectRest("/api/v1/users/*id", new(a_user.Controller))
}
