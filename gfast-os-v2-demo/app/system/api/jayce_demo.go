 package api

 import (
	 "gfast/app/system/dao"
	 "gfast/app/system/service"
	 "github.com/gogf/gf/frame/g"
	 "github.com/gogf/gf/net/ghttp"
	 "github.com/gogf/gf/util/gvalid"
 )
 
 type jayceDemo struct {
	 SystemBase
 }
 
 var JayceDemo = new(jayceDemo)
 
 func (c *jayceDemo) List(r *ghttp.Request) {
	 var req *dao.JayceDemoSearchParams
 
	 if err := r.Parse(&req); err != nil {
		 c.FailJsonExit(r, err.(gvalid.Error).FirstString())
	 }
	 total, page, list, err := service.JayceDemo.List(req)
	 if err != nil {
		 c.FailJsonExit(r, err.Error())
	 }
	 result := g.Map{
		 "total": total,
		 "list":  list,
		 "page":  page,
	 }
	 c.SusJsonExit(r, result)
 }
 
 func (c *jayceDemo) Add(r *ghttp.Request) {
	 var addParams *dao.JayceDemoAddParams
	 if err := r.Parse(&addParams); err != nil {
		 c.FailJsonExit(r, err.(gvalid.Error).FirstString())
	 }
	 addParams.CreatedBy = c.GetCurrentUser(r.GetCtx()).GetUserId()
	 if err := service.JayceDemo.Add(addParams); err != nil {
		 c.FailJsonExit(r, err.Error())
	 }
	 c.SusJsonExit(r, "添加成功")
 }
 
 func (c *jayceDemo) Get(r *ghttp.Request) {
	 id := r.GetInt64("id")
	 if id == 0 {
		 c.FailJsonExit(r, "id必须")
	 }
	 if post, err := service.JayceDemo.GetOneById(id); err != nil {
		 c.FailJsonExit(r, err.Error())
	 } else {
		 c.SusJsonExit(r, post)
	 }
 }
 
 func (c *jayceDemo) Edit(r *ghttp.Request) {
	 var editParams *dao.JayceDemoEditParams
	 if err := r.Parse(&editParams); err != nil {
		 c.FailJsonExit(r, err.(gvalid.Error).FirstString())
	 }
	 editParams.UpdatedBy = c.GetCurrentUser(r.GetCtx()).GetUserId()
	 if err := service.JayceDemo.Edit(editParams); err != nil {
		 c.FailJsonExit(r, err.Error())
	 }
	 c.SusJsonExit(r, "修改成功")
 }
 
 func (c *jayceDemo) Delete(r *ghttp.Request) {
	 ids := r.GetInts("ids")
	 if len(ids) == 0 {
		 c.FailJsonExit(r, "删除失败")
	 }
	 err := service.JayceDemo.Delete(ids)
	 if err != nil {
		 c.FailJsonExit(r, "删除失败")
	 }
	 c.SusJsonExit(r, "删除信息成功")
 }
 