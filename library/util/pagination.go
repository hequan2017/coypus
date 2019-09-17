package util

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func GetPage(r *ghttp.Request) int {
	result := 0
	page := r.GetInt("page")
	if page > 0 {
		result = (page - 1) * g.Config().GetInt("setting.PageSize")
	}

	return result
}
