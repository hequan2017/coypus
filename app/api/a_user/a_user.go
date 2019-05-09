package a_user

import (
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/gogf/gf/g/util/gvalid"
	"github.com/hequan2017/coypus/app/service/s_user"
	"github.com/hequan2017/coypus/library/jwt"
	"github.com/hequan2017/coypus/library/response"
)

// 用户API管理对象
type Controller struct{}

// 用户登录接口
func (c *Controller) Login(r *ghttp.Request) {

	data := r.GetJson()

	rules := map[string]string{
		"username": "required",
		"password": "required",
	}
	msgs := map[string]interface{}{
		"username": "账号不能为空",
		"password": "密码不能为空",
	}

	if e := gvalid.CheckMap(data.ToMap(), rules, msgs); e != nil {
		response.Json(r, 1, e.String())
	}

	authService := s_user.User{Username: data.GetString("username"), Password: data.GetString("password")}
	_, err := authService.Check()

	if err != nil {
		response.Json(r, 1, err.Error())
	} else {
		token, _ := jwt.GenerateToken(data.GetString("username"))
		data := map[string]string{
			"token": token,
		}
		response.Json(r, 0, "", data)
	}

}
