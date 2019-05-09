package a_user

import (
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/gogf/gf/g/util/gvalid"
	"github.com/hequan2017/coypus/app/service/s_user"
	"github.com/hequan2017/coypus/library/e"
	"github.com/hequan2017/coypus/library/jwt"
	"github.com/hequan2017/coypus/library/response"
	"github.com/hequan2017/coypus/library/util"
	"net/http"
)

// 用户API管理对象
type Controller struct{}

// 用户登录接口
func Login(r *ghttp.Request) {

	data := r.GetJson()

	rules := map[string]string{
		"username": "required",
		"password": "required",
	}
	msgs := map[string]interface{}{
		"username": "账号不能为空",
		"password": "密码不能为空",
	}

	if err := gvalid.CheckMap(data.ToMap(), rules, msgs); err != nil {
		response.Json(r, http.StatusOK, e.ERROR_NOT_EXIST, "")
	}

	authService := s_user.User{Username: data.GetString("username"), Password: data.GetString("password")}
	_, err := authService.Check()

	if err != nil {
		response.Json(r, http.StatusOK, e.ERROR_NOT_EXIST, "")
	} else {
		token, _ := jwt.GenerateToken(data.GetString("username"))
		data := map[string]string{
			"token": token,
		}
		response.Json(r, http.StatusOK, e.SUCCESS, data)
	}

}

// RESTFul - GET
func (c *Controller) Get(r *ghttp.Request) {

	userService := s_user.User{
		Username: r.GetString("username"),
		PageNum:  util.GetPage(r),
		PageSize: g.Config().GetInt("setting.PageSize"),
	}

	total, err := userService.Count()
	if err != nil {
		response.Json(r, http.StatusInternalServerError, e.ERROR_COUNT_FAIL, "")
		return
	}

	user, err := userService.GetAll()
	if err != nil {
		response.Json(r, http.StatusInternalServerError, e.ERROR_COUNT_FAIL, "")
		return
	}
	for _, v := range user {
		v.Password = ""
	}

	data := make(map[string]interface{})
	data["lists"] = user
	data["total"] = total

	response.Json(r, http.StatusOK, e.SUCCESS, data)
}

// RESTFul - POST
func (c *Controller) Post(r *ghttp.Request) {
	r.Response.Write("RESTFul HTTP Method POST")
}

// RESTFul - Put
func (c *Controller) Put(r *ghttp.Request) {
	r.Response.Write("RESTFul HTTP Method Put")
}

// RESTFul - DELETE
func (c *Controller) Delete(r *ghttp.Request) {
	r.Response.Write("RESTFul HTTP Method DELETE")
}
