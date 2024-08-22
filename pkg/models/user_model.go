package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id, omitempty"`
	Name      string             `json:"name" bson:"name"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"password" bson:"password"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

type UserRequest struct {
	Name     string `json:"name" bson:"name"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type UserRepository interface {
	Find(ctx context.Context) (*[]User, error)
	FindOne(ctx context.Context, id string) (*User, error)
	UpdateOne(ctx context.Context, id string, req UserRequest) (int64, error)
	DeleteOne(ctx context.Context, id string) (int64, error)
	FindMe(ctx context.Context) (*User, error)
}
