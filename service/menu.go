package service

import (
	"outdoor_rental/dao"
	"outdoor_rental/model"
	"outdoor_rental/model/req"
	"outdoor_rental/model/resp"
	"outdoor_rental/utils"
	"outdoor_rental/utils/r"
)

type Menu struct{}

// SaveOrUpdate 添加｜｜更新菜单
func (*Menu) SaveOrUpdate(req req.SaveOrUpdateMenu) (code int) {
	if req.ID != 0 {
		menu := model.Menu{
			Public:    model.Public{ID: req.ID},
			Name:      req.Name,
			Icon:      req.Icon,
			ParentId:  req.ParentId,
			Path:      req.Path,
			Component: req.Component,
			Sort:      req.Sort,
			Status:    req.Status,
		}
		dao.Update(&menu)
	} else {
		// 检查菜单名已经存在
		existByName := dao.GetOne(model.Menu{}, "name", req.Name)
		if existByName.ID != 0 && existByName.ID != req.ID {
			return r.ERROR_MENU_NAME_EXIST
		}
		data := utils.CopyProperties[model.Menu](req)
		dao.Create(&data)
	}
	return r.OK
}

// Delete 删除菜单
func (*Menu) Delete(req req.Delete) (code int) {
	for _, i := range req.ID {
		// 检查要删除的菜单是否存在
		existMenuById := dao.GetOne(model.Menu{}, "id", i)
		if existMenuById.ID == 0 {
			return r.ERROR_MENU_NAME_EXIST
		}
		// * 检查 role_menu 下是否有数据
		existRoleMenu := dao.GetOne(model.RoleMenu{}, "menu_id", i)
		if existRoleMenu.MenuId != 0 {
			return r.ERROR_MENU_USED_BY_ROLE
		}
		// * 如果是一级菜单, 检查其是否有子菜单
		if existMenuById.ParentId == 0 {
			if dao.Count(model.Menu{}, "parent_id", i) != 0 {
				return r.ERROR_MENU_HAS_CHILDREN
			}
		}
	}
	// 删除菜单
	dao.Delete(model.Menu{}, "id in (?)", req.ID)
	return r.OK
}

// 获取菜单列表(树形)
func (s *Menu) GetTreeList() []resp.MenuListVO {
	// 从数据库中查询出菜单列表(非树形)
	menuList := dao.List([]model.Menu{}, "*", "sort ASC", "")
	data := utils.CopyProperties[[]resp.MenuListVO](menuList)
	menuTree := s.buildMenuTree(data, 0)
	return menuTree
}

func (s *Menu) buildMenuTree(menuItems []resp.MenuListVO, parentID int) []resp.MenuListVO {
	tree := make([]resp.MenuListVO, 0)
	for _, item := range menuItems {
		if item.ParentId == parentID {
			children := s.buildMenuTree(menuItems, item.ID)
			item.Children = children
			tree = append(tree, item)
		}
	}
	return tree
}

// 获取某个用户的菜单列表(树形)
func (s *Menu) GetUserMenuTree(userInfoId int) []resp.MenuListVO {
	// 从数据库查出用户菜单列表(非树形)
	menuList := menuDao.GetMenusByUserInfoId(userInfoId)
	data := utils.CopyProperties[[]resp.MenuListVO](menuList)
	menuTree := s.buildMenuTree(data, 0)
	return menuTree
}
