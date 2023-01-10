package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"
	"errors"

	"gorm.io/gorm"
)

type TaskRepository interface {
	GetTasks(ctx context.Context, id int) ([]entity.Task, error)
	StoreTask(ctx context.Context, task *entity.Task) (taskId int, err error)
	GetTaskByID(ctx context.Context, id int) (entity.Task, error)
	GetTasksByCategoryID(ctx context.Context, catId int) ([]entity.Task, error)
	UpdateTask(ctx context.Context, task *entity.Task) error
	DeleteTask(ctx context.Context, id int) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db}
}

func (r *taskRepository) GetTasks(ctx context.Context, id int) ([]entity.Task, error) {
	var taks []entity.Task
	err := r.db.WithContext(ctx).Model(&entity.Task{}).Where("user_id = ?", id).Find(&taks).Error
	if err != nil {
		return []entity.Task{}, ctx.Err()
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []entity.Task{}, nil
	}
	return taks, nil // TODO: replace this
}

func (r *taskRepository) StoreTask(ctx context.Context, task *entity.Task) (taskId int, err error) {
	if err := r.db.WithContext(ctx).Create(&task).Error; err != nil {
		return 0, nil
	}
	return task.ID, nil // TODO: replace this
}

func (r *taskRepository) GetTaskByID(ctx context.Context, id int) (entity.Task, error) {
	var task entity.Task
	err := r.db.WithContext(ctx).Model(&entity.Task{}).Where("id = ?", id).Find(&task).Error
	if err != nil {
		return entity.Task{}, ctx.Err()
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return entity.Task{}, nil
	}
	return task, nil // TODO: replace this
}

func (r *taskRepository) GetTasksByCategoryID(ctx context.Context, catId int) ([]entity.Task, error) {
	var taks []entity.Task
	err := r.db.WithContext(ctx).Model(&entity.Task{}).Where("category_id = ?", catId).Find(&taks).Error
	if err != nil {
		return []entity.Task{}, ctx.Err()
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []entity.Task{}, nil
	}
	return taks, nil // TODO: replace this
}

func (r *taskRepository) UpdateTask(ctx context.Context, task *entity.Task) error {
	if err := r.db.WithContext(ctx).Model(&entity.Task{}).Where("id = ?", task.ID).Updates(&task).Error; err != nil {
		return ctx.Err()
	}
	return nil // TODO: replace this
}

func (r *taskRepository) DeleteTask(ctx context.Context, id int) error {
	if err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&entity.Task{}).Error; err != nil {
		return ctx.Err()
	}
	return nil // TODO: replace this
}
