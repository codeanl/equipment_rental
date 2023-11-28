package service

import (
	"outdoor_rental/dao"
	"outdoor_rental/model"
	"outdoor_rental/model/req"
	"outdoor_rental/model/resp"
	"outdoor_rental/utils"
	"outdoor_rental/utils/r"
)

type Category struct{}

// SaveOrUpdate 添加｜｜更新菜单
func (*Category) SaveOrUpdate(req req.SaveOrUpdateCategory) (code int) {
	if req.ID != 0 {
		cate := model.Category{
			Public:   model.Public{ID: req.ID},
			Name:     req.Name,
			Pic:      req.Pic,
			ParentId: req.ParentId,
			Sort:     req.Sort,
			Status:   req.Status,
		}
		dao.Update(&cate)
	} else {
		// 检查菜单名已经存在
		existByName := dao.GetOne(model.Category{}, "name", req.Name)
		if existByName.ID != 0 && existByName.ID != req.ID {
			return r.ERROR_CATE_NAME_USED
		}
		data := utils.CopyProperties[model.Category](req)
		dao.Create(&data)
	}
	return r.OK
}

// Delete 删除菜单
func (*Category) Delete(req req.Delete) (code int) {
	for _, i := range req.ID {
		// 检查要删除的菜单是否存在
		existMenuById := dao.GetOne(model.Category{}, "id", i)
		if existMenuById.ID == 0 {
			return r.ERROR_CATE_NOT_EXIST
		}
		// * 如果是一级菜单, 检查其是否有子菜单
		if existMenuById.ParentId == 0 {
			if dao.Count(model.Category{}, "parent_id", i) != 0 {
				return r.ERROR_CATE_ART_EXIST
			}
		}
	}
	// 删除菜单
	dao.Delete(model.Category{}, "id in (?)", req.ID)
	return r.OK
}

// GetTreeList 获取列表(树形)
func (s *Category) GetTreeList() []resp.CategoryListVO {
	// 从数据库中查询出菜单列表(非树形)
	menuList := dao.List([]model.Category{}, "*", "", "")
	data := utils.CopyProperties[[]resp.CategoryListVO](menuList)
	menuTree := s.buildMenuTree(data, 0)
	return menuTree
}
func (s *Category) buildMenuTree(menuItems []resp.CategoryListVO, parentID int) []resp.CategoryListVO {
	tree := make([]resp.CategoryListVO, 0)
	for _, item := range menuItems {
		if item.ParentId == parentID {
			children := s.buildMenuTree(menuItems, item.ID)
			item.Children = children
			tree = append(tree, item)
		}
	}
	return tree
}

//GetList 获取菜单列表（非树形）
func (s *Category) GetList(req req.FrontCategoryList) []resp.FrontCategoryListVO {
	// 从数据库中查询出菜单列表(非树形)
	list, _ := categoryDao.CategoryList(req)
	data := utils.CopyProperties[[]resp.FrontCategoryListVO](list)
	return data
}

//GetListNextCateAndSpu 获取某个一级菜单下面的自分类及商品
func (s *Category) GetListNextCateAndSpu(reqd req.FrontCategoryList) []resp.GetListNextCateAndSpu {
	// 从数据库中查询出菜单列表(非树形)
	list, _ := categoryDao.CategoryList(reqd)
	data := utils.CopyProperties[[]resp.GetListNextCateAndSpu](list)
	for index, i := range data {
		pro, _ := productDao.ProductList(req.ProductList{CategoryId: int64(i.ID)})
		data[index].Product = pro
	}
	return data
}
