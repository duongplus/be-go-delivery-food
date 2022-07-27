package common

import "time"

type SQLModel struct {
	Id        int        `json:"id,omitempty" gorm:"column:id;"`
	Status    int        `json:"status,omitempty" gorm:"column:status;"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at;"`
}
