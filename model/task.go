package model

import "github.com/jinzhu/gorm"

type Task struct {
	gorm.Model
	Id           uint64 `json:"id" form:"-"`
	JobId        uint64 `json:"job_id" form:"job_id"`
	SubjectId    uint64 `json:"subject_id" form:"subject_id"`
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

type Tasks []Task