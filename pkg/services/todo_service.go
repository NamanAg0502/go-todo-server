package services

import (
	"context"
	"fmt"
	"time"

	"github.com/namanag0502/go-todo-server/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoService struct {
	c *mongo.Collection
}

func NewTodoService(c *mongo.Collection) *TodoService {
	return &TodoService{c: c}
}

func GetUserID(ctx context.Context) (primitive.ObjectID, error) {
	userID, ok := ctx.Value(models.UserContextKey).(string)
	if !ok {
		return primitive.NilObjectID, fmt.Errorf("user ID not found")
	}
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return objID, nil
}

func (s *TodoService) Find(ctx context.Context) (*[]models.Todo, error) {
	objID, err := GetUserID(ctx)
	if err != nil {
		return nil, err
	}
	var todos []models.Todo
	cur, err := s.c.Find(ctx, bson.M{"user_id": objID})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var todo models.Todo
		if err := cur.Decode(&todo); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return &todos, nil
}

func (s *TodoService) CreateOne(ctx context.Context, req models.TodoRequest) (*models.Todo, error) {
	objID, err := GetUserID(ctx)
	if err != nil {
		return nil, err
	}
	todo := models.Todo{
		ID:          primitive.NewObjectID(),
		UserID:      objID,
		Title:       req.Title,
		IsCompleted: req.IsCompleted,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	_, err = s.c.InsertOne(ctx, todo)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (s *TodoService) UpdateOne(ctx context.Context, id string, req models.TodoRequest) (int64, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, err
	}

	filter := bson.M{"_id": objID}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "title", Value: req.Title},
			{Key: "is_completed", Value: req.IsCompleted},
			{Key: "updated_at", Value: time.Now()},
		}},
	}

	result, err := s.c.UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, err
	}
	return result.ModifiedCount, nil
}

func (s *TodoService) DeleteOne(ctx context.Context, id string) (int64, error) {
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
