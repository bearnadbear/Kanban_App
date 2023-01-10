package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"
	"errors"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, id int) (entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
	CreateUser(ctx context.Context, user entity.User) (entity.User, error)
	UpdateUser(ctx context.Context, user entity.User) (entity.User, error)
	DeleteUser(ctx context.Context, id int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUserByID(ctx context.Context, id int) (entity.User, error) {
	var user entity.User
	err := r.db.WithContext(ctx).Model(&entity.User{}).Where("id = ?", id).Find(&user).Error
	if err != nil {
		return entity.User{}, ctx.Err()
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return entity.User{}, nil
	}
	return user, nil // TODO: replace this
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	var user entity.User
	err := r.db.WithContext(ctx).Model(&entity.User{}).Where("email = ?", email).Find(&user).Error
	if err != nil {
		return entity.User{}, ctx.Err()
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return entity.User{}, nil
	}
	return user, nil // TODO: replace this
}

func (r *userRepository) CreateUser(ctx context.Context, user entity.User) (entity.User, error) {
	if err := r.db.WithContext(ctx).Create(&user).Error; err != nil {
		return entity.User{}, ctx.Err()
	}
	return user, nil // TODO: replace this
}

func (r *userRepository) UpdateUser(ctx context.Context, user entity.User) (entity.User, error) {
	if err := r.db.WithContext(ctx).Model(&entity.User{}).Where("id = ?", user.ID).Updates(&user).Error; err != nil {
		return entity.User{}, ctx.Err()
	}
	return user, nil // TODO: replace this
}

func (r *userRepository) DeleteUser(ctx context.Context, id int) error {
	if err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&entity.User{}).Error; err != nil {
		return ctx.Err()
	}
	return nil // TODO: replace this
}
