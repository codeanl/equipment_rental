package service

import (
	"outdoor_rental/dao"
	"outdoor_rental/model"
	"outdoor_rental/model/req"
	"outdoor_rental/model/resp"
	"outdoor_rental/utils"
	"outdoor_rental/utils/r"
)

type Address struct{}

//GetCartAddress 列表
func (*Address) GetCartAddress(req req.AddressList) resp.PageResult[[]resp.AddressListVO] {
	list, count := addressDao.AddressList(req)
	return resp.PageResult[[]resp.AddressListVO]{
		PageSize: req.PageSize,
		PageNum:  req.PageNum,
		Total:    count,
		List:     list,
	}
}

func (*Address) SaveOrUpdateAddress(req req.SaveOrUpdateAddress) (code int) {
	if req.ID != 0 {
		cart := model.Address{
			Public:    model.Public{ID: req.ID},
			MemberID:  req.MemberID,
			Name:      req.Name,
			Phone:     req.Phone,
			Address:   req.Address,
			IsDefault: req.IsDefault,
		}
		dao.Update(&cart)
	} else {
		data := utils.CopyProperties[model.Address](req)
		dao.Create(&data)
	}
	return r.OK
}

//DeleteAddress
func (*Address) DeleteAddress(req req.Delete) (code int) {
	for _, i := range req.ID {
		// 检查要删除的菜单是否存在
		existMenuById := dao.GetOne(model.Address{}, "id", i)
		if existMenuById.ID == 0 {
			return r.ERROR_CATE_NOT_EXIST
		}
	}
	// 删除菜单
	dao.Delete(model.Address{}, "id in (?)", req.ID)
	return r.OK
}
