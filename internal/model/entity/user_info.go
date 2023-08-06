// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2023-08-06 17:07:07
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserInfo is the golang structure for table user_info.
type UserInfo struct {
	Id         string      `json:"id"         ` //
	Username   string      `json:"username"   ` //
	Nickname   string      `json:"nickname"   ` //
	Password   string      `json:"password"   ` //
	Email      string      `json:"email"      ` //
	Phone      string      `json:"phone"      ` //
	RegDate    *gtime.Time `json:"regDate"    ` //
	UpdateDate *gtime.Time `json:"updateDate" ` //
	Avatar     string      `json:"avatar"     ` //
}
