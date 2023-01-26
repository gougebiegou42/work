// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)

// Panorama1Dao is the manager for logic model data accessing and custom defined data operations functions management.
type Panorama1Dao struct {
	gmvc.M                // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	C      panorama1Columns // C is the short type for Columns, which contains all the column names of Table for convenient usage.
	DB     gdb.DB         // DB is the raw underlying database management object.
	Table  string         // Table is the underlying table name of the DAO.
}

// Panorama1Columns defines and stores column names for table sys_post.
type panorama1Columns struct {
	Panorama1Id    string // panorama1ID
	// Username  string // 账号名称
	// Password  string // 账号密码
	// CreatedAt string // 创建时间
	// UpdatedAt string // 修改时间
	// DeletedAt string // 删除时间
}

// NewPanorama1Dao creates and returns a new DAO object for table data access.
func NewPanorama1Dao() *Panorama1Dao {
	columns := panorama1Columns{
		Panorama1Id:    "panorama1_id",
		// Username:  "username",
		// Password:  "password",
		// CreatedAt: "created_at",
		// UpdatedAt: "updated_at",
		// DeletedAt: "deleted_at",
	}
	return &Panorama1Dao{
		C:     columns,
		M:     g.DB("default").Model("panorama1").Safe(),
		DB:    g.DB("default"),
		Table: "panorama1",
	}
}
