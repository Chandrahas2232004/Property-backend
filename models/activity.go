package models

import "time"

// ActivityItem represents a recent activity entry for dashboard
// It is a lightweight DTO, not a database table.
type ActivityItem struct {
	Type         string    `json:"type"`
	TimeCreated  time.Time `json:"time_created"`
	RespectiveID uint      `json:"respective_id"`
}
