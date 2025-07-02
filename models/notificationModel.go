package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Notification struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	UserID    primitive.ObjectID `bson:"user_id" json:"user_id"` // Who receives the notification
	Title     string             `bson:"title" json:"title"`     // Short title
	Message   string             `bson:"message" json:"message"` // Notification content
	Type      string             `bson:"type" json:"type"`       // e.g. "late_slip", "reminder"
	IsRead    bool               `bson:"is_read" json:"is_read"` // Read/unread status
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}
