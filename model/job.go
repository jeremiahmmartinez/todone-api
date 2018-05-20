package model

import "github.com/jinzhu/gorm"

type Job struct {
	gorm.Model
	Id           uint64 `json:"id" form:"-"`
	OwnerUserId  uint64 `json:"owner_user_id" form:"owner_user_id"`
	OwnerGroupId uint64 `json:"owner_group_id" form:"owner_group_id"`
	OwnerType    string `json:"owner_type" form:"owner_type"`
	Label	     string `json:"label" form:"label"`
	Description  string `json:"description" form:"description"`
	Notify       bool   `json:"notify" form:"notify"`
	NotifyType   string `json:"notify_type" form:"notify_type"`
	NotifyBefore uint   `json:"notify_before" form:"notify_before"`
	StartTime    string `json:"start_time" form:"start_time"`
	EndTime      string `json:"end_time" form:"end_time"`
	Created      string `json:"created" form:"created"`
	Updated      string `json:"updated" form:"updated"`
}

type Jobs []Job