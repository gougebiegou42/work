package service

import (
	// comModel "gfast/app/common/model"
	"gfast/app/system/dao"
	"gfast/app/system/model"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
)

type panoramaImg struct{}

var PanoramaImg = new(panoramaImg)

func (s *panoramaImg) List(req *dao.PanoramaImgSearchParams) (page int, list []*model.PanoramaImg, err error) {
	model := dao.PanoramaImg.M
	// if req != nil {
	// 	if req.Username != "" {
	// 		model.Where("username like ?", "%"+req.Username+"%")
	// 	}

	// 	if req.Password != "" {
	// 		model.Where("password like ?", "%"+req.Password+"%")
	// 	}
	// }

	// total, err = model.Count()

	// if err != nil {
	// 	g.Log().Error(err)
	// 	err = gerror.New("获取总行数失败")
	// }

	// if req.PageNum == 0 {
	// 	req.PageNum = 1
	// }

	// page = req.PageNum

	// if req.PageSize == 0 {
	// 	req.PageSize = comModel.PageSize
	// }
	err = model.Page(page, req.PageSize).Order("panoramaImg_id asc").Scan(&list)
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("获取数据失败")
	}
	return
}

func (s *panoramaImg) Add(params *dao.PanoramaImgAddParams) error {
	_, err := dao.PanoramaImg.Insert(params)
	return err
}

// func (s *panoramaImg) GetOneById(id int64) (post *model.PanoramaImg, err error) {
// 	err = dao.PanoramaImg.WherePri(id).Scan(&post)
// 	return
// }

// func (s *panoramaImg) Edit(params *dao.PanoramaImgEditParams) (err error) {
// 	_, err = dao.PanoramaImg.FieldsEx(dao.PanoramaImg.C.PanoramaImgId).
// 		WherePri(params.PanoramaImgId).Update(params)
// 	return err
// }

func (s *panoramaImg) Delete(ids []int) error {
	_, err := dao.PanoramaImg.Where(dao.PanoramaImg.C.PanoramaImgId+" in(?)", ids).Delete()
	return err
}
