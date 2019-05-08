package s_common

import (
	"github.com/hequan2017/coypus/app/service/s_menu"
	"github.com/hequan2017/coypus/app/service/s_role"
	"github.com/hequan2017/coypus/app/service/s_user"
)

type Common struct {
	UserAPI *s_user.User `inject:""`
	RoleAPI *s_role.Role `inject:""`
	MenuAPI *s_menu.Menu `inject:""`
}
