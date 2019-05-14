package s_menu

import (
	"fmt"
	"github.com/casbin/casbin"
	"github.com/hequan2017/coypus/app/model"
	"github.com/hequan2017/coypus/app/service/s_role"
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

	Menu     *model.Menu      `inject:""`
	Enforcer *casbin.Enforcer `inject:""`
}

func (a *Menu) Add() error {
	menu := map[string]interface{}{
		"name":   a.Name,
		"path":   a.Path,
		"method": a.Method,
	}
	if err := model.AddMenu(menu); err != nil {
		return err
	}

	return nil
}

func (a *Menu) Edit() error {

	err := model.EditMenu(a.ID, map[string]interface{}{
		"name":   a.Name,
		"path":   a.Path,
		"method": a.Method,
	})

	if err != nil {
		return err
	}

	roleList := model.EditMenuGetRoles(a.ID)

	roleService := s_role.Role{}

	for _, v := range roleList {
		fmt.Println(v)
		err := roleService.LoadPolicy(v)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *Menu) Get() (*model.Menu, error) {

	menu, err := model.GetMenu(a.ID)
	if err != nil {
		return nil, err
	}

	return menu, nil
}

func (a *Menu) GetAll() ([]*model.Menu, error) {
	Menu, err := model.GetMenus(a.PageNum, a.PageSize, a.getMaps())
	if err != nil {
		return nil, err
	}

	return Menu, nil
}

func (a *Menu) Delete() error {
	err := model.DeleteMenu(a.ID)
	if err != nil {
		return err
	}
	roleList := model.EditMenuGetRoles(a.ID)
	roleService := s_role.Role{}
	for _, v := range roleList {
		err := roleService.LoadPolicy(v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *Menu) ExistByID() (bool, error) {
	return model.ExistMenuByID(a.ID)
}

func (a *Menu) Count() (int, error) {
	return model.GetMenuTotal(a.getMaps())
}

func (a *Menu) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_on"] = 0
	return maps
}
