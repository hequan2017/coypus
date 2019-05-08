package s_user

import (
	"github.com/casbin/casbin"
	"github.com/gogf/gf/g/os/glog"
	"github.com/hequan2017/coypus/app/model"
)

type User struct {
	ID       int
	Username string
	Password string
	Role     int

	CreatedBy  string
	ModifiedBy string

	Enforcer *casbin.Enforcer `inject:""`
}


func (a *User) Check() (bool, error) {
	return model.CheckUser(a.Username, a.Password)
}

// LoadAllPolicy 加载所有的用户策略
func (a *User) LoadAllPolicy() error {
	users, err := model.GetUsersAll()
	if err != nil {
		return err
	}
	for _, user := range users {
		if len(user.Role) != 0 {
			err = a.LoadPolicy(user.ID)
			if err != nil {
				return err
			}
		}
	}
	glog.Info("角色权限关系", a.Enforcer.GetGroupingPolicy())
	return nil
}

// LoadPolicy 加载用户权限策略
func (a *User) LoadPolicy(id int) error {

	user, err := model.GetUser(id)
	if err != nil {
		return err
	}

	a.Enforcer.DeleteRolesForUser(user.Username)

	for _, ro := range user.Role {
		a.Enforcer.AddRoleForUser(user.Username, ro.Name)
	}
	glog.Info("更新角色权限关系", a.Enforcer.GetGroupingPolicy())
	return nil
}
