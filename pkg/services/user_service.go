package services

import (
	"context"
	"time"

	"github.com/namanag0502/go-todo-server/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	c *mongo.Collection
}

func NewUserService(c *mongo.Collection) *UserService {
	return &UserService{c: c}
}

func (s *UserService) Find(ctx context.Context) (*[]models.User, error) {
	var users []models.User
	cur, err := s.c.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var user models.User
		if err := cur.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return &users, nil
}

func (s *UserService) FindOne(ctx context.Context, id string) (*models.User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.D{{Key: "_id", Value: objID}}
	var user models.User
	err = s.c.FindOne(ctx, filter).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserService) UpdateOne(ctx context.Context, id string, req models.UserRequest) (int64, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, err
	}

	updateData, err := bson.Marshal(req)
	if err != nil {
		return 0, err
	}
	var update bson.M
	if err := bson.Unmarshal(updateData, &update); err != nil {
		return 0, err
	}

	delete(update, "password")
	update["updated_at"] = time.Now()

	if len(update) == 0 {
		return 0, nil
	}

	filter := bson.D{{Key: "_id", Value: objID}}

	result, err := s.c.UpdateOne(ctx, filter, bson.M{"$set": update})
	if err != nil {
		return 0, err
	}
	return result.ModifiedCount, nil
}
func (s *UserService) DeleteOne(ctx context.Context, id string) (int64, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, err
	}

	filter := bson.M{"_id": objID}
	result, err := s.c.DeleteOne(ctx, filter)
	if err != nil {
		return 0, err
	}
	return result.DeletedCount, nil
}

func (s *UserService) FindMe(ctx context.Context) (*models.User, error) {
	userID, err := GetUserID(ctx)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": userID}
	var user models.User
	err = s.c.FindOne(ctx, filter).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &user, nil
}
