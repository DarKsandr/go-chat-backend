package pkg

import "time"

type Message struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Nickname  string    `json:"nickname" binding:"required"`
	Message   string    `json:"message" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
