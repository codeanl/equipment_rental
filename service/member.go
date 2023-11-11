package service

import (
	"outdoor_rental/dao"
	"outdoor_rental/model"
	"outdoor_rental/model/req"
	"outdoor_rental/model/resp"
	"outdoor_rental/utils"
	"outdoor_rental/utils/r"
)

type Member struct{}

// MemberInfo 用户详情
func (s *Member) MemberInfo(id int) resp.MemberInfoVO {
	var memberInfo model.Member
	dao.GetOne(&memberInfo, "id", id)
	data := utils.CopyProperties[resp.MemberInfoVO](memberInfo) //复制
	return data
}

// MemberList 用户列表
func (*Member) MemberList(req req.MemberList) resp.PageResult[[]resp.MemberListVO] {
	list, count, _ := memberDao.MemberList(req)
	return resp.PageResult[[]resp.MemberListVO]{
		PageSize: req.PageSize,
		PageNum:  req.PageNum,
		Total:    count,
		List:     list,
	}
}

// MemberUpdate 更新用户
func (*Member) MemberUpdate(req req.MemberUser) int {
	userInfo := model.Member{
		Public:   model.Public{ID: req.ID},
		Nickname: req.Nickname,
		Phone:    req.Phone,
		Email:    req.Email,
		Avatar:   req.Avatar,
		Intro:    req.Intro,
		Status:   req.Status,
	}
	dao.Update(&userInfo)
	return r.OK
}

// MemberDelete 删除用户
func (*Member) MemberDelete(req req.Delete) int {
	dao.Delete(&model.Member{}, "id in (?)", req.ID)
	return r.OK
}

//
//会员登录

//会员注册

//更新个人信息
