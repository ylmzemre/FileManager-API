package models

import (
	"time"

	"github.com/google/uuid"
)

type File struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID     uint
	OriginName string
	Path       string
	MimeType   string
	Size       int64
	CreatedAt  time.Time
}
