package service

import (
	"outdoor_rental/dao"
	"outdoor_rental/model"
	"outdoor_rental/model/req"
	"outdoor_rental/model/resp"
	"outdoor_rental/utils"
	"outdoor_rental/utils/r"
)

type Article struct{}

//GetCartAddress 列表
func (*Article) GetArticleList(req req.ArticleList) resp.PageResult[[]resp.ArticleListVO] {
	list, count := articleDao.ArticleList(req)
	for index, i := range list {
		user := dao.GetOne(model.Member{}, "id", i.MemberID)
		list[index].Member = user
	}
	return resp.PageResult[[]resp.ArticleListVO]{
		PageSize: req.PageSize,
		PageNum:  req.PageNum,
		Total:    count,
		List:     list,
	}
}

func (*Article) SaveOrUpdateArticle(req req.SaveOrUpdateArticle) (code int) {
	if req.ID != 0 {
		cart := model.Article{
			Public:   model.Public{ID: req.ID},
			MemberID: req.MemberID,
			Title:    req.Title,
			Content:  req.Content,
			Pic:      req.Pic,
		}
		dao.Update(&cart)
	} else {
		data := utils.CopyProperties[model.Article](req)
		dao.Create(&data)
	}
	return r.OK
}

//DeleteAddress
func (*Article) DeleteArticle(req req.Delete) (code int) {
	for _, i := range req.ID {
		// 检查要删除的菜单是否存在
		existMenuById := dao.GetOne(model.Article{}, "id", i)
		if existMenuById.ID == 0 {
			return r.ERROR_CATE_NOT_EXIST
		}
	}
	// 删除菜单
	dao.Delete(model.Article{}, "id in (?)", req.ID)
	return r.OK
}
