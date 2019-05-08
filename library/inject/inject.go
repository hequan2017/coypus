package inject

import (
	"github.com/casbin/casbin"
	"github.com/facebookgo/inject"
	"github.com/hequan2017/coypus/app/service/s_common"
	"runtime"
)

// Object 注入对象
type Object struct {
	Common   *s_common.Common
	Enforcer *casbin.Enforcer
}

var Obj *Object

// init 初始化依赖注入
func init() {
	g := new(inject.Graph)

	// 注入casbin
	osType := runtime.GOOS
	var path string
	if osType == "windows" {
		path = "config\\rbac_model.conf"
	} else if osType == "linux" {
		path = "config/rbac_model.conf"
	}
	enforcer := casbin.NewEnforcer(path, false)
	_ = g.Provide(&inject.Object{Value: enforcer})

	Common := new(s_common.Common)
	_ = g.Provide(&inject.Object{Value: Common})

	if err := g.Populate(); err != nil {
		panic("初始化依赖注入发生错误：" + err.Error())
	}

	Obj = &Object{
		Enforcer: enforcer,
		Common:   Common,
	}
	return
}

// 加载casbin策略数据，包括角色权限数据、用户角色数据
func LoadCasbinPolicyData() error {
	c := Obj.Common

	err := c.RoleAPI.LoadAllPolicy()
	if err != nil {
		return err
	}
	err = c.UserAPI.LoadAllPolicy()
	if err != nil {
		return err
	}
	return nil
}

