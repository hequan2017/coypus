package s_user

import (
	"github.com/casbin/casbin"
	"github.com/gogf/gf/crypto/gsha1"
	"github.com/gogf/gf/os/glog"
	"github.com/hequan2017/coypus/app/model"
	"github.com/hequan2017/coypus/library/e"
)

type User struct {
	ID       int
	Username string
	Password string
	Role     []int

	CreatedBy  string
	ModifiedBy string

	PageNum  int
	PageSize int

	Enforcer *casbin.Enforcer `inject:""`
}

func (a *User) Check() (bool, error) {
	return model.CheckUser(a.Username, gsha1.Encrypt(a.Password))
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

func (a *User) Add() (id int, err int) {
	menu := map[string]interface{}{
		"username": a.Username,
		"password": gsha1.Encrypt(a.Password),
		"role_id":  a.Role,
	}
	username, _ := model.CheckUserUsername(a.Username)

	if username {
		return 0, e.ERROR_USER_EXIST
	}
	for _, v := range a.Role {
		if v > 0 {
			roles, _ := model.ExistRoleByID(v)
			if !roles {
				return 0, e.ERROR_ROLE_EXIST_FAIL
			}
		}
	}

	if id, err := model.AddUser(menu); err != nil {
		return 0, e.ERROR_USER_ADD_FAIL
	} else {
		return id, e.SUCCESS
	}
}

func (a *User) Edit() (id, error int) {
	data := map[string]interface{}{
		"username": a.Username,
		"password": gsha1.Encrypt(a.Password),
		"role_id":  a.Role,
	}
	for _, v := range a.Role {
		if v > 0 {
			roles, _ := model.ExistRoleByID(v)
			if !roles {
				return 0, e.ERROR_ROLE_EXIST_FAIL
			}
		}
	}

	id, err := model.EditUser(a.ID, data)
	if err != nil {
		return 0, e.INVALID_PARAMS
	}
	return id, e.SUCCESS
}

func (a *User) Get() (*model.User, error) {

	user, err := model.GetUser(a.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (a *User) GetAll() ([]*model.User, error) {
	if a.Username != "" {
		maps := make(map[string]interface{})
		maps["deleted_on"] = 0
		maps["username"] = a.Username
		user, err := model.GetUsers(a.PageNum, a.PageSize, maps)
		if err != nil {
			return nil, err
		}
		return user, nil
	} else {
		user, err := model.GetUsers(a.PageNum, a.PageSize, a.getMaps())
		if err != nil {
			return nil, err
		}
		return user, nil
	}
}

func (a *User) Delete() error {
	err := model.DeleteUser(a.ID)
	if err != nil {
		return err
	}
	return nil
}

func (a *User) ExistByID() (bool, error) {
	return model.ExistUserByID(a.ID)
}

func (a *User) Count() (int, error) {
	return model.GetUserTotal(a.getMaps())
}

func (a *User) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_on"] = 0
	return maps
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
