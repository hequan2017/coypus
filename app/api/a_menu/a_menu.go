package a_menu

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
	"github.com/hequan2017/coypus/app/service/s_menu"
	"github.com/hequan2017/coypus/library/e"
	"github.com/hequan2017/coypus/library/response"
	"github.com/hequan2017/coypus/library/util"
	"net/http"
)

// 用户API管理对象
type Controller struct{}

var rules = map[string]string{
	"name":   "required",
	"path":   "required",
	"method": "required",
}

var msgs = map[string]interface{}{
	"name":   "菜单名称 不能为空",
	"path":   "路径 不能为空",
	"method": "方法 不能为空",
}

// RESTFul - GET
func (c *Controller) Get(r *ghttp.Request) {
	menuService := s_menu.Menu{
		Name:     r.GetString("rname"),
		PageNum:  util.GetPage(r),
		PageSize: g.Config().GetInt("setting.PageSize"),
	}
	id := r.GetInt("id")

	if id != 0 {
		menuService.ID = id
		user, err := menuService.Get()
		if err != nil {
			response.Json(r, http.StatusBadRequest, e.ERROR_USER_EXIST_FAIL, "")
			return
		}
		data := make(map[string]interface{})
		data["lists"] = user

		response.Json(r, http.StatusOK, e.SUCCESS, data)

	} else {

		total, err := menuService.Count()
		if err != nil {
			response.Json(r, http.StatusBadRequest, e.ERROR_USER_GET_S_FAIL, "")
			return
		}
		users, err := menuService.GetAll()
		if err != nil {
			response.Json(r, http.StatusBadRequest, e.ERROR_USER_GET_S_FAIL, "")
			return
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
	menuService := s_menu.Menu{
		Name:   data.GetString("name"),
		Path:   data.GetString("path"),
		Method: data.GetString("method"),
	}

	if err := menuService.Add(); err != nil {
		response.Json(r, http.StatusBadRequest, e.ERROR_MENU_ADD_FAIL, "")
	} else {
		response.Json(r, http.StatusOK, e.SUCCESS, nil)
	}

}

// RESTFul - Put
func (c *Controller) Put(r *ghttp.Request) {
	data := r.GetJson()
	if id := r.GetInt("id"); id <= 0 {
		response.Json(r, http.StatusBadRequest, e.ERROR_MENU_EDIT_FAIL, "")
		r.ExitAll()
	}

	if err := gvalid.CheckMap(data.ToMap(), rules, msgs); err != nil {
		response.Json(r, http.StatusBadRequest, e.INVALID_PARAMS, err.String())
	}

	menuService := s_menu.Menu{
		ID:     r.GetInt("id"),
		Name:   data.GetString("name"),
		Path:   data.GetString("path"),
		Method: data.GetString("method"),
	}

	if err := menuService.Edit(); err != nil {
		response.Json(r, http.StatusBadRequest, e.ERROR_MENU_EDIT_FAIL, "")
	} else {
		response.Json(r, http.StatusOK, e.SUCCESS, nil)
	}
}

// RESTFul - DELETE
func (c *Controller) Delete(r *ghttp.Request) {
	if id := r.GetInt("id"); id <= 0 {
		response.Json(r, http.StatusBadRequest, e.ERROR_MENU_DELETE_FAIL, "")
	} else {
		menuService := s_menu.Menu{ID: id}
		_, err := menuService.ExistByID()
		if err != nil {
			response.Json(r, http.StatusBadRequest, e.ERROR_MENU_DELETE_FAIL, "")
			r.ExitAll()
		}
		err = menuService.Delete()
		if err != nil {
			response.Json(r, http.StatusBadRequest, e.ERROR_MENU_DELETE_FAIL, "")
			r.ExitAll()
		} else {
			response.Json(r, http.StatusOK, e.SUCCESS, nil)
		}
	}
}
