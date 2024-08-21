package services

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/namanag0502/go-todo-server/pkg/models"
	"github.com/namanag0502/go-todo-server/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthService struct {
	C *mongo.Collection
}

func NewAuthService(c *mongo.Collection) *AuthService {
	return &AuthService{C: c}
}

func (s *AuthService) Login(ctx context.Context, req *models.LoginRequest) (*models.User, error) {
	var user models.User
	if err := s.C.FindOne(ctx, bson.M{"email": req.Email}).Decode(&user); err != nil {
		return nil, err
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return nil, errors.New("invalid password")
	}

	return &user, nil
}

func (s *AuthService) Register(ctx context.Context, req *models.RegisterRequest) (*models.User, error) {
	_, err := s.C.FindOne(ctx, bson.M{"email": req.Email}).Raw()
	log.Println(err)
	if err != mongo.ErrNoDocuments {
		return nil, err
	}

	pass, _ := utils.HashPassword(req.Password)
	user := &models.User{
		ID:        primitive.NewObjectID(),
		Name:      req.Name,
		Email:     req.Email,
		Password:  pass,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_, err = s.C.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
