// ==========================================================================
// This is auto-generated by gf cli tool. Fill this file as you wish.
// ==========================================================================

package model

import (
	"gea/app/model/internal"
	"github.com/gogf/gf/os/gtime"
)

// SysPost is the golang structure for table sys_post.
type SysPost internal.SysPost

// Fill with you ideas below.

type SysPostFlag struct {
	PostId     int64       `orm:"post_id,primary" json:"post_id"`     // 岗位ID
	PostCode   string      `orm:"post_code"       json:"post_code"`   // 岗位编码
	PostName   string      `orm:"post_name"       json:"post_name"`   // 岗位名称
	PostSort   int         `orm:"post_sort"       json:"post_sort"`   // 显示顺序
	Status     string      `orm:"status"          json:"status"`      // 状态（0正常 1停用）
	CreateBy   string      `orm:"create_by"       json:"create_by"`   // 创建者
	CreateTime *gtime.Time `orm:"create_time"     json:"create_time"` // 创建时间
	UpdateBy   string      `orm:"update_by"       json:"update_by"`   // 更新者
	UpdateTime *gtime.Time `orm:"update_time"     json:"update_time"` // 更新时间
	Remark     string      `orm:"remark"          json:"remark"`      // 备注
	Flag       bool        `json:"flag"`                              // 标记
}