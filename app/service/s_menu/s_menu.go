package s_menu

import (
	"github.com/casbin/casbin"
	"github.com/hequan2017/coypus/app/model"
)

type Menu struct {
	ID     int
	Name   string
	Path   string
	Method string

	CreatedBy  string
	ModifiedBy string

	PageNum  int
	PageSize int

	Menu     *model.Menu     `inject:""`
	Enforcer *casbin.Enforcer `inject:""`
}

