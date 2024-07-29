// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameInternalFeedbackThread = "internal_feedback_threads"

// InternalFeedbackThread mapped from table <internal_feedback_threads>
type InternalFeedbackThread struct {
	ID                      string    `gorm:"column:id;primaryKey;default:gen_random_uuid()" json:"id"`
	Title                   string    `gorm:"column:title;not null" json:"title"`
	Content                 string    `gorm:"column:content;not null" json:"content"`
	UserID                  string    `gorm:"column:user_id;not null" json:"user_id"`
	CreatedAt               time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt               time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	Priority                string    `gorm:"column:priority;not null;default:low" json:"priority"`
	Type                    string    `gorm:"column:type;not null;default:general" json:"type"`
	Status                  string    `gorm:"column:status;not null;default:open" json:"status"`
	AddedToRoadmap          bool      `gorm:"column:added_to_roadmap;not null" json:"added_to_roadmap"`
	OpenForPublicDiscussion bool      `gorm:"column:open_for_public_discussion;not null" json:"open_for_public_discussion"`
	IsPubliclyVisible       bool      `gorm:"column:is_publicly_visible;not null" json:"is_publicly_visible"`
}

// TableName InternalFeedbackThread's table name
func (*InternalFeedbackThread) TableName() string {
	return TableNameInternalFeedbackThread
}