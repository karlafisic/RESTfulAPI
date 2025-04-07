package models

import "time"

type Document struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Size       int64     `json:"size"`
	UploadedAt time.Time `json:"uploaded_at"`
}
