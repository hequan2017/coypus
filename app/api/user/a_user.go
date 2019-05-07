package a_user

import (
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/gogf/gf/g/util/gvalid"
	"github.com/hequan2017/coypus/app/service/user"
	"github.com/hequan2017/coypus/library/jwt"
	"github.com/hequan2017/coypus/library/response"
)

// 用户API管理对象
type Controller struct{}

// 用户注册接口
func (c *Controller) SignUp(r *ghttp.Request) {
	if err := s_user.SignUp(r.GetPostMap()); err != nil {
		response.Json(r, 1, err.Error())
	} else {
		response.Json(r, 0, "ok")
	}
}

// 用户登录接口
func (c *Controller) SignIn(r *ghttp.Request) {
	data := r.GetPostMap()
	rules := map[string]string{
		"username": "required",
		"password": "required",
	}
	msgs := map[string]interface{}{
		"username": "账号不能为空",
		"password": "密码不能为空",
	}

	if e := gvalid.CheckMap(data, rules, msgs); e != nil {
		response.Json(r, 1, e.String())
	}
	if err := s_user.SignIn(data["username"], data["password"]); err != nil {
		response.Json(r, 1, err.Error())
	} else {
		token, _ := jwt.GenerateToken(rules["username"], rules["password"])
		data := map[string]string{
			"token": token,
		}
		response.Json(r, 0, "", data)
	}
}

// 判断用户是否已经登录
//func (c *Controller) IsSignedIn(r *ghttp.Request) {
//	if s_user.IsSignedIn(r.Session) {
//		response.Json(r, 0, "ok")
//	} else {
//		response.Json(r, 1, "")
//	}
//}

// 用户注销/退出接口
func (c *Controller) SignOut(r *ghttp.Request) {
	s_user.SignOut(r.Session)
	response.Json(r, 0, "ok")
}

// 检测用户账号接口(唯一性校验)
func (c *Controller) CheckUsername(r *ghttp.Request) {
	passport := r.Get("username")
	if e := gvalid.Check(passport, "required", "请输入账号"); e != nil {
		response.Json(r, 1, e.String())
	}
	if s_user.CheckUsername(passport) {
		response.Json(r, 0, "ok")
	}
	response.Json(r, 1, "账号已经存在")
}
