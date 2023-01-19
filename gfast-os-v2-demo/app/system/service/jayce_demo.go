package service

import (
	comModel "gfast/app/common/model"
	"gfast/app/system/dao"
	"gfast/app/system/model"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
)

type jayceDemo struct{}

var JayceDemo = new(jayceDemo)

func (s *jayceDemo) List(req *dao.JayceDemoSearchParams) (total, page int, list []*model.JayceDemo, err error) {
	model := dao.JayceDemo.M
	if req != nil {
		if req.PostCode != "" {
			model.Where("post_code like ?", "%"+req.PostCode+"%")
		}

		if req.PostName != "" {
			model.Where("post_name like ?", "%"+req.PostName+"%")
		}

		if req.Status != "" {
			model.Where("status", req.Status)
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
	err = model.Page(page, req.PageSize).Order("post_sort asc,post_id asc").Scan(&list)
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("获取数据失败")
	}
	return
}

func (s *jayceDemo) Add(params *dao.JayceDemoAddParams) error {
	_, err := dao.JayceDemo.Insert(params)
	return err
}

func (s *jayceDemo) GetOneById(id int64) (post *model.JayceDemo, err error) {
	err = dao.JayceDemo.WherePri(id).Scan(&post)
	return
}

func (s *jayceDemo) Edit(params *dao.JayceDemoEditParams) (err error) {
	_, err = dao.JayceDemo.FieldsEx(dao.JayceDemo.C.PostId, dao.JayceDemo.C.CreatedBy).
		WherePri(params.PostId).Update(params)
	return err
}

func (s *jayceDemo) Delete(ids []int) error {
	_, err := dao.JayceDemo.Where(dao.JayceDemo.C.PostId+" in(?)", ids).Delete()
	return err
}
