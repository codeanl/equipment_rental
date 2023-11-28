package service

import (
	"outdoor_rental/dao"
	"outdoor_rental/model"
	"outdoor_rental/model/req"
	"outdoor_rental/model/resp"
	"outdoor_rental/utils"
	"outdoor_rental/utils/r"
)

type SlideShow struct{}

// Delete 删除菜单
func (*SlideShow) Delete(req req.Delete) (code int) {
	dao.Delete(model.SlideShow{}, "id in (?)", req.ID)
	return r.OK
}

// SlideshowList 列表
func (*SlideShow) SlideshowList(req req.SlideshowList) resp.PageResult[[]resp.SlideshowListVO] {
	list, count := slideShowDao.SlideShowList(req)
	return resp.PageResult[[]resp.SlideshowListVO]{
		PageSize: req.PageSize,
		PageNum:  req.PageNum,
		Total:    count,
		List:     list,
	}
}

// SaveOrUpdate 添加||编辑
func (*SlideShow) SaveOrUpdate(req req.SaveOrUpdateSlideshow) (code int) {
	if req.ID != 0 {
		menu := model.SlideShow{
			ID:     req.ID,
			Name:   req.Name,
			Url:    req.Url,
			Status: req.Status,
		}
		dao.Update(&menu)
	} else {
		// 检查菜单名已经存在
		existByName := dao.GetOne(model.SlideShow{}, "name", req.Name)
		if existByName.ID != 0 {
			return r.ERROR_MENU_NAME_EXIST
		}
		data := utils.CopyProperties[model.SlideShow](req)
		data.Status = "1"
		dao.Create(&data)
	}
	return r.OK
}
