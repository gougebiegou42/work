package api

import (
	"gfast/app/system/dao"
	"gfast/app/system/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

type sylas struct {
	SystemBase
}

var Sylas = new(sylas)

func (c *sylas) List(r *ghttp.Request) {
	var req *dao.SylasSearchParams

	if err := r.Parse(&req); err != nil {
		c.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
	total, page, list, err := service.Sylas.List(req)
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	result := g.Map{
		// 配置返回的json数据名
		"total": total,
		"nnn":  list,
		"page":  page,
	}
	c.SusJsonExit(r, result)
}

func (c *sylas) Add(r *ghttp.Request) {
	var addParams *dao.SylasAddParams
	if err := r.Parse(&addParams); err != nil {
		c.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
	addParams.CreatedBy = c.GetCurrentUser(r.GetCtx()).GetUserId()
	if err := service.Sylas.Add(addParams); err != nil {
		c.FailJsonExit(r, err.Error())
	}
	c.SusJsonExit(r, "添加成功")
}

func (c *sylas) Get(r *ghttp.Request) {
	id := r.GetInt64("id")
	if id == 0 {
		c.FailJsonExit(r, "id必须")
	}
	if post, err := service.Sylas.GetOneById(id); err != nil {
		c.FailJsonExit(r, err.Error())
	} else {
		c.SusJsonExit(r, post)
	}
}

func (c *sylas) Edit(r *ghttp.Request) {
	var editParams *dao.SylasEditParams
	if err := r.Parse(&editParams); err != nil {
		c.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
	editParams.UpdatedBy = c.GetCurrentUser(r.GetCtx()).GetUserId()
	if err := service.Sylas.Edit(editParams); err != nil {
		c.FailJsonExit(r, err.Error())
	}
	c.SusJsonExit(r, "修改成功")
}

func (c *sylas) Delete(r *ghttp.Request) {
	ids := r.GetInts("ids")
	if len(ids) == 0 {
		c.FailJsonExit(r, "删除失败")
	}
	err := service.Sylas.Delete(ids)
	if err != nil {
		c.FailJsonExit(r, "删除失败")
	}
	c.SusJsonExit(r, "删除信息成功")
}
