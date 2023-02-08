package api

import (
	"gfast/app/system/dao"
	"gfast/app/system/service"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

type panoramaImg struct {
	SystemBase
}

var PanoramaImg = new(panoramaImg)

func (c *panoramaImg) List(r *ghttp.Request) {
	var req *dao.PanoramaImgSearchParams

	if err := r.Parse(&req); err != nil {
		c.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
	page, list, err := service.PanoramaImg.List(req)
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	result := g.Map{
		// 配置返回的json数据名
		"data": list,
		"page": page,
	}
	c.SusJsonExit(r, result)
}

func (c *panoramaImg) Add(r *ghttp.Request) {
	var addParams *dao.PanoramaImgAddParams
	if err := r.Parse(&addParams); err != nil {
		c.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
	addParams.CreatedBy = c.GetCurrentUser(r.GetCtx()).GetUserId()
	if err := service.PanoramaImg.Add(addParams); err != nil {
		c.FailJsonExit(r, err.Error())
	}
	c.SusJsonExit(r, "添加成功")
}

// func (c *panoramaImg) Get(r *ghttp.Request) {
// 	id := r.GetInt64("id")
// 	if id == 0 {
// 		c.FailJsonExit(r, "id必须")
// 	}
// 	if post, err := service.PanoramaImg.GetOneById(id); err != nil {
// 		c.FailJsonExit(r, err.Error())
// 	} else {
// 		c.SusJsonExit(r, post)
// 	}
// }

// func (c *panoramaImg) Edit(r *ghttp.Request) {
// 	var editParams *dao.PanoramaImgEditParams
// 	if err := r.Parse(&editParams); err != nil {
// 		c.FailJsonExit(r, err.(gvalid.Error).FirstString())
// 	}
// 	editParams.UpdatedBy = c.GetCurrentUser(r.GetCtx()).GetUserId()
// 	if err := service.PanoramaImg.Edit(editParams); err != nil {
// 		c.FailJsonExit(r, err.Error())
// 	}
// 	c.SusJsonExit(r, "修改成功")
// }

func (c *panoramaImg) Delete(r *ghttp.Request) {
	ids := r.GetInts("ids")
	if len(ids) == 0 {
		c.FailJsonExit(r, "删除失败")
	}
	err := service.PanoramaImg.Delete(ids)
	if err != nil {
		c.FailJsonExit(r, "删除失败")
	}
	c.SusJsonExit(r, "删除信息成功")
}
