package task

import (
	"context"
	"timesheet-manager-backend/api/presenter"
	"timesheet-manager-backend/pkg/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Repository interface allows us to access the CRUD Operations in mongo here.
type Repository interface {
	CreateTask(task *entities.Task) (*entities.Task, error)
	ReadTask() (*[]presenter.Task, error)
	UpdateTask(task *entities.Task) (*entities.Task, error)
	DeleteTask(ID string) error
}
type repository struct {
	Collection *mongo.Collection
}

// NewRepo is the single instance repo that is being created.
func NewRepo(collection *mongo.Collection) Repository {
	return &repository{
		Collection: collection,
	}
}

// CreateTask is a mongo repository that helps to create tasks
func (r *repository) CreateTask(task *entities.Task) (*entities.Task, error) {
	task.ID = primitive.NewObjectID()
	// task.CreatedAt = time.Now()
	// task.UpdatedAt = time.Now()
	_, err := r.Collection.InsertOne(context.Background(), task)
	if err != nil {
		return nil, err
	}
	return task, nil
}

// ReadTask is a mongo repository that helps to fetch tasks
func (r *repository) ReadTask() (*[]presenter.Task, error) {
	var tasks []presenter.Task
	cursor, err := r.Collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var task presenter.Task
		_ = cursor.Decode(&task)
		tasks = append(tasks, task)
	}
	return &tasks, nil
}

// UpdateTask is a mongo repository that helps to update tasks
func (r *repository) UpdateTask(task *entities.Task) (*entities.Task, error) {
	// task.UpdatedAt = time.Now()
	_, err := r.Collection.UpdateOne(context.Background(), bson.M{"_id": task.ID}, bson.M{"$set": task})
	if err != nil {
		return nil, err
	}
	return task, nil
}

// DeleteTask is a mongo repository that helps to delete tasks
func (r *repository) DeleteTask(ID string) error {
	taskID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = r.Collection.DeleteOne(context.Background(), bson.M{"_id": taskID})
	if err != nil {
		return err
	}
	return nil
}
