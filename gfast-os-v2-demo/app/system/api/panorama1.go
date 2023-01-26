package api

import (
	"gfast/app/system/dao"
	"gfast/app/system/service"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

type panorama1 struct {
	SystemBase
}

var Panorama1 = new(panorama1)

func (c *panorama1) List(r *ghttp.Request) {
	var req *dao.Panorama1SearchParams

	if err := r.Parse(&req); err != nil {
		c.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
	total, page, list, err := service.Panorama1.List(req)
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	result := g.Map{
		// 配置返回的json数据名
		"total": total,
		"data":  list,
		"page":  page,
	}
	c.SusJsonExit(r, result)
}

// func (c *panorama1) Add(r *ghttp.Request) {
// 	var addParams *dao.Panorama1AddParams
// 	if err := r.Parse(&addParams); err != nil {
// 		c.FailJsonExit(r, err.(gvalid.Error).FirstString())
// 	}
// 	addParams.CreatedBy = c.GetCurrentUser(r.GetCtx()).GetUserId()
// 	if err := service.Panorama1.Add(addParams); err != nil {
// 		c.FailJsonExit(r, err.Error())
// 	}
// 	c.SusJsonExit(r, "添加成功")
// }

// func (c *panorama1) Get(r *ghttp.Request) {
// 	id := r.GetInt64("id")
// 	if id == 0 {
// 		c.FailJsonExit(r, "id必须")
// 	}
// 	if post, err := service.Panorama1.GetOneById(id); err != nil {
// 		c.FailJsonExit(r, err.Error())
// 	} else {
// 		c.SusJsonExit(r, post)
// 	}
// }

// func (c *panorama1) Edit(r *ghttp.Request) {
// 	var editParams *dao.Panorama1EditParams
// 	if err := r.Parse(&editParams); err != nil {
// 		c.FailJsonExit(r, err.(gvalid.Error).FirstString())
// 	}
// 	editParams.UpdatedBy = c.GetCurrentUser(r.GetCtx()).GetUserId()
// 	if err := service.Panorama1.Edit(editParams); err != nil {
// 		c.FailJsonExit(r, err.Error())
// 	}
// 	c.SusJsonExit(r, "修改成功")
// }

// func (c *panorama1) Delete(r *ghttp.Request) {
// 	ids := r.GetInts("ids")
// 	if len(ids) == 0 {
// 		c.FailJsonExit(r, "删除失败")
// 	}
// 	err := service.Panorama1.Delete(ids)
// 	if err != nil {
// 		c.FailJsonExit(r, "删除失败")
// 	}
// 	c.SusJsonExit(r, "删除信息成功")
// }
