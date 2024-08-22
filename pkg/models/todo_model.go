package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoRequest struct {
	Title       string `json:"title" bson:"title"`
	IsCompleted bool   `json:"is_completed" bson:"is_completed"`
}

type Todo struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	UserID      primitive.ObjectID `json:"user_id" bson:"user_id"`
	IsCompleted bool               `json:"is_completed" bson:"is_completed"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}

type TodoRepository interface {
	Find(ctx context.Context) (*[]Todo, error)
	CreateOne(ctx context.Context, req TodoRequest) (*Todo, error)
	UpdateOne(ctx context.Context, id string, req TodoRequest) (int64, error)
	DeleteOne(ctx context.Context, id string) (int64, error)
}
