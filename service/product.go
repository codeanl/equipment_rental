package service

import (
	"fmt"
	"outdoor_rental/dao"
	"outdoor_rental/model"
	"outdoor_rental/model/req"
	"outdoor_rental/model/resp"
	"outdoor_rental/utils"
	"outdoor_rental/utils/r"
)

type Product struct{}

// SaveOrUpdate 添加｜｜更新菜单
func (*Product) SaveOrUpdate(req req.SaveOrUpdateProduct) (code int) {
	//请求的sku的tags
	size := generateCombinations(req.Size, 0, "")
	var pId int
	if req.ID != 0 {
		pId = req.ID
		menu := model.Product{
			Public:     model.Public{ID: req.ID},
			Name:       req.Name,
			Pic:        req.Pic,
			Desc:       req.Desc,
			Price:      req.Price,
			CategoryID: req.CategoryId,
		}
		dao.Update(&menu)
		//判断sku进行过滤
		//过滤新添加和需要删除的
		nowSku := productDao.GetProductSkuTag(req.ID)
		add, del := filterArrays(size, nowSku)
		//fmt.Println(add, "----", del)
		//添加新sku
		for _, i := range add {
			dao.Create(&model.ProductSku{
				ProductID: int64(pId),
				Name:      req.Name,
				Pic:       req.Pic,
				Desc:      req.Desc,
				Price:     req.Price,
				Stock:     0,
				Sale:      0,
				Tag:       i,
			})
		}
		//删除新sku
		dao.Delete(&model.ProductSku{}, "tag in (?)", del)
	} else {
		existByName := dao.GetOne(model.Product{}, "name", req.Name) // 检查菜单名已经存在
		if existByName.ID != 0 && existByName.ID != req.ID {
			return r.ERROR_MENU_NAME_EXIST
		}
		data := utils.CopyProperties[model.Product](req)
		dao.Create(&data)
		pId = data.ID
		//添加skus
		for _, i := range size {
			dao.Create(&model.ProductSku{
				ProductID: int64(pId),
				Name:      req.Name,
				Pic:       req.Pic,
				Desc:      req.Desc,
				Price:     req.Price,
				Stock:     0,
				Sale:      0,
				Tag:       i,
			})
		}
	}
	//添加商品的图片
	dao.Delete(model.ProductImg{}, "product_id = ?", pId)
	for _, i := range req.Img {
		dao.Create(&model.ProductImg{
			ProductID: pId,
			Path:      i,
		})
	}
	return r.OK
}

// ProductDelete 删除菜单
func (*Product) ProductDelete(req req.Delete) (code int) {
	for _, i := range req.ID {
		// 检查要删除的菜单是否存在
		existMenuById := dao.GetOne(model.Product{}, "id", i)
		if existMenuById.ID == 0 {
			return r.ERROR_CATE_NOT_EXIST
		}
	}
	// 删除菜单
	dao.Delete(model.Product{}, "id in (?)", req.ID)
	dao.Delete(model.ProductImg{}, "product_id in (?)", req.ID)
	dao.Delete(model.ProductSku{}, "product_id in (?)", req.ID)
	return r.OK
}

// ProductList 用户列表
func (*Product) ProductList(req req.ProductList) resp.PageResult[[]resp.ProductListVO] {
	list, count := productDao.ProductList(req)
	return resp.PageResult[[]resp.ProductListVO]{
		PageSize: req.PageSize,
		PageNum:  req.PageNum,
		Total:    count,
		List:     list,
	}
}

// SkuUpdate 更新sku
func (*Product) SkuUpdate(req req.UpdateSku) (code int) {
	sku := model.ProductSku{
		Public: model.Public{ID: req.ID},
		Name:   req.Name,
		Pic:    req.Pic,
		Desc:   req.Desc,
		Price:  req.Price,
		Stock:  req.Stock,
		Sale:   req.Sale,
	}
	dao.Update(&sku)
	return r.OK
}

// SkuList  Sku列表
func (*Product) SkuList(pid int) []model.ProductSku {
	skuList := dao.List([]model.ProductSku{}, "*", "", "product_id =  ?", pid)
	return skuList
}

func filterArrays(reqSku, nowSku []string) ([]string, []string) {
	reqSkuMap := make(map[string]bool)
	nowSkuMap := make(map[string]bool)
	// 将 aa 中的元素添加到 aaMap 中
	for _, num := range reqSku {
		reqSkuMap[num] = true
	}
	// 将 bb 中的元素添加到 bbMap 中
	for _, num := range nowSku {
		nowSkuMap[num] = true
	}
	// 过滤 aa 中存在但 bb 中不存在的元素
	aaOnly := make([]string, 0)
	for _, num := range reqSku {
		if _, exists := nowSkuMap[num]; !exists {
			aaOnly = append(aaOnly, num)
		}
	}
	// 过滤 bb 中存在但 aa 中不存在的元素
	bbOnly := make([]string, 0)
	for _, num := range nowSku {
		if _, exists := reqSkuMap[num]; !exists {
			bbOnly = append(bbOnly, num)
		}
	}
	return aaOnly, bbOnly
}
func generateCombinations(items []req.Size, index int, current string) []string {
	if index == len(items) {
		return []string{current}
	}
	var combinations []string
	item := items[index]
	for _, size := range item.SizeData {
		newCurrent := current + fmt.Sprintf(`"%s: "%s"`, item.Name, size)
		if index < len(items)-1 {
			newCurrent += ", "
		}
		subCombinations := generateCombinations(items, index+1, newCurrent)
		combinations = append(combinations, subCombinations...)
	}
	return combinations
}
