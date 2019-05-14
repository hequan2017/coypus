package a_user

import (
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/gogf/gf/g/os/glog"
	"github.com/gogf/gf/g/util/gvalid"
	"github.com/hequan2017/coypus/app/service/s_user"
	"github.com/hequan2017/coypus/library/e"
	"github.com/hequan2017/coypus/library/inject"
	"github.com/hequan2017/coypus/library/jwt"
	"github.com/hequan2017/coypus/library/response"
	"github.com/hequan2017/coypus/library/util"
	"net/http"
)

// 用户API管理对象
type Controller struct{}

var rules = map[string]string{
	"username": "required",
	"password": "required",
	"role":     "required-with|integer",
}

var msgs = map[string]interface{}{
	"username": "账号不能为空",
	"password": "密码不能为空",
	"role":     "权限组 id 必须为 整数 列表",
}

// 用户登录接口
func Login(r *ghttp.Request) {

	data := r.GetJson()

	if err := gvalid.CheckMap(data.ToMap(), rules, msgs); err != nil {
		response.Json(r, http.StatusBadRequest, e.INVALID_PARAMS, err.String())
	}

	authService := s_user.User{Username: data.GetString("username"), Password: data.GetString("password")}
	_, err := authService.Check()
	if err != nil {
		response.Json(r, http.StatusBadRequest, e.ERROR_USER_NOT_EXIST, "")
	} else {
		token, _ := jwt.GenerateToken(data.GetString("username"))
		data := map[string]string{
			"token": token,
		}
		response.Json(r, http.StatusOK, e.SUCCESS, data)
	}
}
func UserInfo(r *ghttp.Request) {

	data := map[string]interface{}{
		"name":    "admin",
		"user_id": "1",
		"access":  []string{"admin"},
		"token":   "token",
		"avatar":  "https://file.iviewui.com/dist/a0e88e83800f138b94d2414621bd9704.png",
	}
	response.Json(r, http.StatusOK, e.SUCCESS, data)
}

func Menu(r *ghttp.Request) {

	data := []interface{}{
		map[string]interface{}{
			"path": "/assets",
			"name": "assets",
			"meta": map[string]string{
				"icon":  "md-menu",
				"title": "资产管理",
			},
			"component": "Main",
			"children": []interface{}{map[string]interface{}{
				"path": "ecs",
				"name": "ecs",
				"meta": map[string]string{
					"icon":  "md-funnel",
					"title": "ecs",
				},
				"component": "assets/ecs/ecs-list",
			},
			},
		},
	}

	response.Json(r, http.StatusOK, e.SUCCESS, data)
}

// RESTFul - GET
func (c *Controller) Get(r *ghttp.Request) {
	userService := s_user.User{
		Username: r.GetString("username"),
		PageNum:  util.GetPage(r),
		PageSize: g.Config().GetInt("setting.PageSize"),
	}
	id := r.GetInt("id")

	if id != 0 {
		userService.ID = id
		user, err := userService.Get()
		if err != nil {
			response.Json(r, http.StatusBadRequest, e.ERROR_USER_EXIST_FAIL, "")
			return
		}
		user.Password = ""
		data := make(map[string]interface{})
		data["lists"] = user

		response.Json(r, http.StatusOK, e.SUCCESS, data)

	} else {

		total, err := userService.Count()
		if err != nil {
			response.Json(r, http.StatusBadRequest, e.ERROR_USER_GET_S_FAIL, "")
			return
		}
		users, err := userService.GetAll()
		if err != nil {
			response.Json(r, http.StatusBadRequest, e.ERROR_USER_GET_S_FAIL, "")
			return
		}
		for _, v := range users {
			v.Password = ""
		}

		data := make(map[string]interface{})
		data["lists"] = users
		data["total"] = total

		response.Json(r, http.StatusOK, e.SUCCESS, data)
	}
}

// RESTFul - POST
func (c *Controller) Post(r *ghttp.Request) {

	data := r.GetJson()

	if err := gvalid.CheckMap(data.ToMap(), rules, msgs); err != nil {
		response.Json(r, http.StatusBadRequest, e.INVALID_PARAMS, err.String())
	}
	userService := s_user.User{
		Username: data.GetString("username"),
		Password: data.GetString("password"),
		Role:     data.GetInts("role"),
	}

	if id, err := userService.Add(); err != e.SUCCESS {
		response.Json(r, http.StatusBadRequest, err, "")
	} else {
		err := inject.Obj.Common.UserAPI.LoadPolicy(id)
		if err != nil {
			response.Json(r, http.StatusBadRequest, e.ERROR_LOAD_CASBIN_FAIL, "")
			r.ExitAll()
		}
		response.Json(r, http.StatusOK, e.SUCCESS, nil)
	}

}

// RESTFul - Put
func (c *Controller) Put(r *ghttp.Request) {
	data := r.GetJson()
	if id := r.GetInt("id"); id <= 0 {
		response.Json(r, http.StatusBadRequest, e.ERROR_USER_EDIT_FAIL, "")
		r.ExitAll()
	}

	if err := gvalid.CheckMap(data.ToMap(), rules, msgs); err != nil {
		response.Json(r, http.StatusBadRequest, e.INVALID_PARAMS, err.String())
	}
	userService := s_user.User{
		ID:       r.GetInt("id"),
		Username: data.GetString("username"),
		Password: data.GetString("password"),
		Role:     data.GetInts("role"),
	}

	if id, err := userService.Edit(); err != e.SUCCESS {
		response.Json(r, http.StatusBadRequest, e.ERROR_USER_EDIT_FAIL, "")

	} else {
		err := inject.Obj.Common.UserAPI.LoadPolicy(id)
		if err != nil {
			glog.Error(err)
			response.Json(r, http.StatusBadRequest, e.ERROR_LOAD_CASBIN_FAIL, "")
			r.ExitAll()
		}
		response.Json(r, http.StatusOK, e.SUCCESS, nil)
	}
}

// RESTFul - DELETE
func (c *Controller) Delete(r *ghttp.Request) {
	if id := r.GetInt("id"); id <= 0 {
		response.Json(r, http.StatusBadRequest, e.ERROR_USER_DELETE_FAIL, "")
	} else {
		userService := s_user.User{ID: id}
		_, err := userService.ExistByID()
		if err != nil {
			response.Json(r, http.StatusBadRequest, e.ERROR_USER_DELETE_FAIL, "")
			r.ExitAll()
		}
		user, err := userService.Get()
		err = userService.Delete()

		if err != nil {
			response.Json(r, http.StatusBadRequest, e.ERROR_USER_DELETE_FAIL, "")
			r.ExitAll()
		} else {
			inject.Obj.Enforcer.DeleteUser(user.Username)
			response.Json(r, http.StatusOK, e.SUCCESS, nil)
		}
	}
}
