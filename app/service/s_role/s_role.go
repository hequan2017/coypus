package s_role

import (
	"github.com/casbin/casbin"
	"github.com/gogf/gf/os/glog"
	"github.com/hequan2017/coypus/app/model"
	"github.com/hequan2017/coypus/library/e"
)

type Role struct {
	ID   int
	Name string
	Menu []int

	CreatedBy  string
	ModifiedBy string

	PageNum  int
	PageSize int

	Enforcer *casbin.Enforcer `inject:""`
}

func (a *Role) Add() (id int, err int) {
	role := map[string]interface{}{
		"name":    a.Name,
		"menu_id": a.Menu,
	}
	name, _ := model.CheckRoleName(a.Name)

	if name {
		return 0, e.ERROR_ROLE_EXIST
	}

	if id, err := model.AddRole(role); err != nil {
		return 0, e.ERROR_ROLE_ADD_FAIL
	} else {
		return id, e.SUCCESS
	}

}

func (a *Role) Edit() (id, error int) {
	data := map[string]interface{}{
		"name":    a.Name,
		"menu_id": a.Menu,
	}
	name, _ := model.CheckRoleNameId(a.Name, a.ID)

	if name {
		return 0, e.ERROR_ROLE_EXIST
	}
	err := model.EditRole(a.ID, data)

	if err != nil {
		glog.Error(err)
		return 0, e.ERROR_ROLE_EDIT_FAIL
	}

	return a.ID, e.SUCCESS
}

func (a *Role) Get() (*model.Role, error) {

	Role, err := model.GetRole(a.ID)
	if err != nil {
		return nil, err
	}

	return Role, nil
}

func (a *Role) GetAll() ([]*model.Role, error) {
	if a.Name != "" {
		maps := make(map[string]interface{})
		maps["deleted_on"] = 0
		maps["name"] = a.Name
		Role, err := model.GetRoles(a.PageNum, a.PageSize, maps)
		if err != nil {
			return nil, err
		}

		return Role, nil
	} else {
		Role, err := model.GetRoles(a.PageNum, a.PageSize, a.getMaps())
		if err != nil {
			return nil, err
		}
		return Role, nil
	}
}

func (a *Role) Delete() error {
	err := model.DeleteRole(a.ID)
	if err != nil {
		return err
	}
	a.Enforcer.DeletePermissionsForUser(a.Name)
	return nil
}

func (a *Role) ExistByID() (bool, error) {
	return model.ExistRoleByID(a.ID)
}

func (a *Role) Count() (int, error) {
	return model.GetRoleTotal(a.getMaps())
}

func (a *Role) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_on"] = 0
	return maps
}

// LoadAllPolicy 加载所有的角色策略
func (a *Role) LoadAllPolicy() error {
	roles, err := model.GetRolesAll()
	if err != nil {
		return err
	}

	for _, role := range roles {
		err = a.LoadPolicy(role.ID)
		if err != nil {
			return err
		}
	}
	return nil
}

// LoadPolicy 加载角色权限策略
func (a *Role) LoadPolicy(id int) error {

	role, err := model.GetRole(id)
	if err != nil {
		return err
	}
	a.Enforcer.DeleteRole(role.Name)

	for _, menu := range role.Menu {
		if menu.Path == "" || menu.Method == "" {
			continue
		}
		a.Enforcer.AddPermissionForUser(role.Name, menu.Path, menu.Method)
	}
	return nil
}
