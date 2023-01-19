// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/os/gtime"
)

// SysUser is the golang structure for table sys_user.
type SysUser struct {
	Id            uint64      `orm:"id,primary"       json:"id"`            //
	UserName      string      `orm:"user_name,unique" json:"userName"`      // 用户名
	Mobile        string      `orm:"mobile,unique"    json:"mobile"`        // 中国手机不带国家代码，国际手机号格式为：国家代码-手机号
	UserNickname  string      `orm:"user_nickname"    json:"userNickname"`  // 用户昵称
	Birthday      int         `orm:"birthday"         json:"birthday"`      // 生日
	UserPassword  string      `orm:"user_password"    json:"userPassword"`  // 登录密码;cmf_password加密
	UserSalt      string      `orm:"user_salt"        json:"userSalt"`      // 加密盐
	UserStatus    uint        `orm:"user_status"      json:"userStatus"`    // 用户状态;0:禁用,1:正常,2:未验证
	UserEmail     string      `orm:"user_email"       json:"userEmail"`     // 用户登录邮箱
	Sex           int         `orm:"sex"              json:"sex"`           // 性别;0:保密,1:男,2:女
	Avatar        string      `orm:"avatar"           json:"avatar"`        // 用户头像
	DeptId        uint64      `orm:"dept_id"          json:"deptId"`        // 部门id
	Remark        string      `orm:"remark"           json:"remark"`        // 备注
	IsAdmin       int         `orm:"is_admin"         json:"isAdmin"`       // 是否后台管理员 1 是  0   否
	Address       string      `orm:"address"          json:"address"`       // 联系地址
	Describe      string      `orm:"describe"         json:"describe"`      // 描述信息
	PhoneNum      string      `orm:"phone_num"        json:"phoneNum"`      // 联系电话
	LastLoginIp   string      `orm:"last_login_ip"    json:"lastLoginIp"`   // 最后登录ip
	LastLoginTime *gtime.Time `orm:"last_login_time"  json:"lastLoginTime"` // 最后登录时间
	CreatedAt     *gtime.Time `orm:"created_at"       json:"createdAt"`     // 创建时间
	UpdatedAt     *gtime.Time `orm:"updated_at"       json:"updatedAt"`     // 更新时间
	DeletedAt     *gtime.Time `orm:"deleted_at"       json:"deletedAt"`     // 删除时间
}