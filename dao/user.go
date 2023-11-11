package dao

import (
	"fmt"
	"outdoor_rental/model/req"
	"outdoor_rental/model/resp"
)

type User struct{}

func (User) UserList(req req.UserList) (list []resp.UserListVO, total int64, err error) {
	db := DB.Model(&User{}).Order("created_at DESC").Preload("Roles")
	if req.Username != "" {
		db = db.Where("username LIKE ?", fmt.Sprintf("%%%s%%", req.Username))
	}
	if req.Nickname != "" {
		db = db.Where("nickname LIKE ?", fmt.Sprintf("%%%s%%", req.Nickname))
	}
	if req.Email != "" {
		db = db.Where("email LIKE ?", fmt.Sprintf("%%%s%%", req.Email))
	}
	if req.Phone != "" {
		db = db.Where("phone LIKE ?", fmt.Sprintf("%%%s%%", req.Phone))
	}
	if req.Status != "" {
		db = db.Where("status = ?", req.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return list, total, err
	}
	if req.PageNum > 0 && req.PageSize > 0 {
		err = db.Offset((req.PageNum - 1) * req.PageSize).Limit(req.PageSize).Find(&list).Error
	} else {
		err = db.Find(&list).Error
	}
	return list, total, err
}
