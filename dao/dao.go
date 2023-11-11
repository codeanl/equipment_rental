package dao

import (
	"errors"
	"gorm.io/gorm"
)

var DB *gorm.DB

// 创建数据(可以创建[单条]数据, 也可[批量]创建)
func Create[T any](data *T) {
	err := DB.Create(&data).Error
	if err != nil {
		panic(err)
	}
}

func GetOne[T any](data T, query string, args ...any) T {
	err := DB.Where(query, args...).First(&data).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) { // 记录找不到 err 不 panic
		panic(err)
	}
	return data
}

//在函数中，首先检查是否提供了选择字段列表，即 slt 是否有传递的参数。如果有传递参数，表示只更新指定的字段，通过 DB.Model(data).Select(slt).Updates(data) 来进行更新。
//如果没有传递参数，表示更新所有字段，通过 DB.Model(data).Updates(data) 进行更新。
func Update[T any](data *T, slt ...string) {
	if len(slt) > 0 {
		DB.Model(&data).Select(slt).Updates(&data)
		return
	}
	err := DB.Model(&data).Updates(&data).Error
	if err != nil {
		panic(err)
	}
}

func Delete[T any](data T, query string, args ...any) {
	err := DB.Where(query, args...).Delete(&data).Error
	if err != nil {
		panic(err)
	}
}

// 统计数量
func Count[T any](data T, query string, args ...any) int64 {
	var total int64
	db := DB.Model(data)
	if query != "" {
		db = db.Where(query, args...)
	}
	if err := db.Count(&total).Error; err != nil {
		panic(err)
	}
	return total
}

// 数据列表
func List[T any](data T, slt, order, query string, args ...any) T {
	db := DB.Model(&data).Select(slt).Order(order)
	if query != "" {
		db = db.Where(query, args...)
	}
	if err := db.Find(&data).Error; err != nil {
		panic(err)
	}
	return data
}
