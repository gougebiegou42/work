package service

import (
	comModel "gfast/app/common/model"
	"gfast/app/system/dao"
	"gfast/app/system/model"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
)

type sylas struct{}

var Sylas = new(sylas)

func (s *sylas) List(req *dao.SylasSearchParams) (total, page int, list []*model.Sylas, err error) {
	model := dao.Sylas.M
	if req != nil {
		if req.Username != "" {
			model.Where("username like ?", "%"+req.Username+"%")
		}

		if req.Password != "" {
			model.Where("password like ?", "%"+req.Password+"%")
		}
	}

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
	err = model.Page(page, req.PageSize).Order("sylas_id asc").Scan(&list)
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("获取数据失败")
	}
	return
}

func (s *sylas) Add(params *dao.SylasAddParams) error {
	_, err := dao.Sylas.Insert(params)
	return err
}

func (s *sylas) GetOneById(id int64) (post *model.Sylas, err error) {
	err = dao.Sylas.WherePri(id).Scan(&post)
	return
}

func (s *sylas) Edit(params *dao.SylasEditParams) (err error) {
	_, err = dao.Sylas.FieldsEx(dao.Sylas.C.SylasId).
		WherePri(params.SylasId).Update(params)
	return err
}

func (s *sylas) Delete(ids []int) error {
	_, err := dao.Sylas.Where(dao.Sylas.C.SylasId+" in(?)", ids).Delete()
	return err
}
