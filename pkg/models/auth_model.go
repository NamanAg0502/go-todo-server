package models

import (
	"context"
)

type contextKey string

const UserContextKey contextKey = "user_id"

type LoginRequest struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type RegisterRequest struct {
	Name     string `json:"name" bson:"name"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type AuthRepository interface {
	Login(ctx context.Context, req *LoginRequest) (*User, error)
	Register(ctx context.Context, req *RegisterRequest) (*User, error)
}
