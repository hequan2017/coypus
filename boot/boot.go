package boot

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	_ "github.com/hequan2017/coypus/app/model"
	"github.com/hequan2017/coypus/library/inject"
	_ "github.com/hequan2017/coypus/library/inject"
	_ "github.com/hequan2017/coypus/router"
)

// 用于应用初始化。
func init() {

	_ = g.View()
	c := g.Config()
	s := g.Server()

	// 模板引擎配置
	//_ = v.AddPath("template")
	//v.SetDelimiters("${", "}")

	// glog配置
	logpath := c.GetString("setting.logpath")
	glog.SetPath(logpath)

	// Web Server配置
	//s.SetServerRoot("public")
	s.SetLogPath(logpath)
	s.SetNameToUriType(ghttp.NAME_TO_URI_TYPE_ALLLOWER)
	s.SetErrorLogEnabled(true)
	s.SetAccessLogEnabled(true)
	s.SetPort(8000)

	AppSetting.PageSize = c.GetInt("setting.PageSize")

	_ = inject.LoadCasbinPolicyData()
}
