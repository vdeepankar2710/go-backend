package repositories

import (
	"context"
	"time"
	"todo-backend/errors"
	"todo-backend/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TodoRepository struct {
	collection *mongo.Collection
}

func NewTodoRepository(db *mongo.Database) *TodoRepository {
	return &TodoRepository{
		collection: db.Collection("todos"),
	}
}

func (r *TodoRepository) CreateRepo(todo *models.Todo) error {
	_, err := r.collection.InsertOne(context.TODO(), todo)
	return err
}

func (r *TodoRepository) GetAllRepo(pageNumber int, entriesPerPage int, sort string) ([]models.Todo, error) {
	var todosArr []models.Todo

	offset := (pageNumber - 1) * entriesPerPage

	options := options.Find()
	options.SetSkip(int64(offset))
	options.SetLimit(int64(entriesPerPage))
	if sort=="ASC"{
		options.SetSort(bson.D{{Key: "created_at", Value: 1}})
	}else if sort=="DESC"{
		options.SetSort(bson.D{{Key: "created_at", Value: -1}})
	}else{
		return nil, errors.ErrInvalidSorting
	}
	
	cursor, err := r.collection.Find(context.TODO(), bson.D{}, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var todo models.Todo
		if err := cursor.Decode(&todo); err != nil {
			return nil, err
		}
		todosArr = append(todosArr, todo)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return todosArr, nil
}

func (r *TodoRepository) GetTodoByIDRepo(id primitive.ObjectID) (*models.Todo, error) {
	var todo models.Todo

	err := r.collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&todo)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *TodoRepository) UpdateTodoRepo(todo *models.Todo) error {
	filter := bson.M{"_id": todo.ID}
	update := bson.M{
		"$set": bson.M{
			"title":       todo.Title,
			"description": todo.Description,
			"status":      todo.Status,
			"updated_at":  time.Now(),
		},
	}

	_, err := r.collection.UpdateOne(context.TODO(), filter, update)
	return err
}

func (r *TodoRepository) DeleteTodoRepo(id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	return err
}

func (r *TodoRepository) GetTodoByUserIdRepo(userId int) ([]models.Todo, error) {
	var todosArr []models.Todo

	cursor, err := r.collection.Find(context.TODO(), bson.M{"user_id": userId})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var todo models.Todo
		if err := cursor.Decode(&todo); err != nil {
			return nil, err
		}
		todosArr = append(todosArr, todo)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return todosArr, nil
}
