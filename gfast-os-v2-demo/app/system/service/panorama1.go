package service

import (
	comModel "gfast/app/common/model"
	"gfast/app/system/dao"
	"gfast/app/system/model"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
)

type panorama1 struct{}

var Panorama1 = new(panorama1)

func (s *panorama1) List(req *dao.Panorama1SearchParams) (total, page int, list []*model.Panorama1, err error) {
	model := dao.Panorama1.M
	// if req != nil {
	// 	if req.Username != "" {
	// 		model.Where("username like ?", "%"+req.Username+"%")
	// 	}

	// 	if req.Password != "" {
	// 		model.Where("password like ?", "%"+req.Password+"%")
	// 	}
	// }

	total, err = model.Count()

	if err != nil {
		g.Log().Error(err)
		err = gerror.New("获取总行数失败")
	}

	if req.PageNum == 0 {
		req.PageNum = 1
	}

	page = req.PageNum

	if req.PageSize == 0 {
		req.PageSize = comModel.PageSize
	}
	err = model.Page(page, req.PageSize).Order("panorama1_id asc").Scan(&list)
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("获取数据失败")
	}
	return
}

// func (s *panorama1) Add(params *dao.Panorama1AddParams) error {
// 	_, err := dao.Panorama1.Insert(params)
// 	return err
// }

// func (s *panorama1) GetOneById(id int64) (post *model.Panorama1, err error) {
// 	err = dao.Panorama1.WherePri(id).Scan(&post)
// 	return
// }

// func (s *panorama1) Edit(params *dao.Panorama1EditParams) (err error) {
// 	_, err = dao.Panorama1.FieldsEx(dao.Panorama1.C.Panorama1Id).
// 		WherePri(params.Panorama1Id).Update(params)
// 	return err
// }

// func (s *panorama1) Delete(ids []int) error {
// 	_, err := dao.Panorama1.Where(dao.Panorama1.C.Panorama1Id+" in(?)", ids).Delete()
// 	return err
// }
