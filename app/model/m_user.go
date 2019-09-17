package model

import (
	"errors"
	"github.com/gogf/gf/os/glog"
	"github.com/jinzhu/gorm"
)

type User struct {
	Model
	Username string `json:"username"`
	Password string `json:"password"`
	Role     []Role `json:"role" gorm:"many2many:user_role;"`
}

func CheckUser(username, password string) (bool, error) {
	var user User
	err := db.Select("id").Where(User{Username: username, Password: password}).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, errors.New("账户未找到")
	}
	if user.ID > 0 {
		return true, nil
	}

	return false, errors.New("账户未找到")
}

func ExistUserByID(id int) (bool, error) {
	var user User
	err := db.Select("id").Where("id = ? AND deleted_on = ? ", id, 0).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if user.ID > 0 {
		return true, nil
	}

	return false, err
}

func GetUserTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&User{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func GetUsers(pageNum int, pageSize int, maps interface{}) ([]*User, error) {
	var user []*User
	err := db.Preload("Role").Where(maps).Offset(pageNum).Limit(pageSize).Find(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return user, nil
}

func GetUser(id int) (*User, error) {
	var user User
	err := db.Preload("Role").Where("id = ? AND deleted_on = ? ", id, 0).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &user, nil
}

func CheckUserUsername(username string) (bool, error) {
	var user User
	err := db.Where("username = ? AND deleted_on = ? ", username, 0).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if user.ID > 0 {
		return true, nil
	}

	return false, nil
}

func CheckUserUsernameId(username string, id int) (bool, error) {
	var user User
	err := db.Where("username = ? AND id != ? AND deleted_on = ? ", username, id, 0).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if user.ID > 0 {
		return true, nil
	}

	return false, nil
}

func EditUser(id int, data map[string]interface{}) (ids int, err error) {
	var role []Role
	var user User
	db.Where("id in (?)", data["role_id"]).Find(&role)
	if err := db.Where("id = ? AND deleted_on = ? ", id, 0).Find(&user).Error; err != nil {
		return 0, err
	}
	db.Model(&user).Association("Role").Replace(role)
	db.Model(&user).Update(data)

	return user.ID, nil
}

func AddUser(data map[string]interface{}) (id int, err error) {
	user := User{
		Username: data["username"].(string),
		Password: data["password"].(string),
	}
	var role []Role

	if err := db.Where("id in (?)", data["role_id"]).Find(&role).Error; err != nil {
		glog.Error(err)
		return 0, err
	}
	if err := db.Create(&user).Association("Role").Append(role).Error; err != nil {
		glog.Error(err)
		return 0, err
	}
	return user.ID, nil
}

func DeleteUser(id int) error {
	var user User
	db.Where("id = ?", id).Find(&user)
	db.Model(&user).Association("Role").Delete()
	if err := db.Where("id = ?", id).Delete(&user).Error; err != nil {
		return err
	}

	return nil
}

func CleanAllUser() error {
	if err := db.Unscoped().Where("deleted_on != ? ", 0).Delete(&User{}).Error; err != nil {
		return err
	}
	return nil
}

func GetUsersAll() ([]*User, error) {
	var user []*User
	err := db.Where("deleted_on = ? ", 0).Preload("Role").Find(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return user, nil
}
