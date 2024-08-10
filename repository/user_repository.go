package repository

import (
	"rest-api/model"

	"gorm.io/gorm"
)

// インターフェース
// IUserRepository型は，GetUserByEmailとCreateUserの2つの関数を持たなければいけない。
type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}

// 構造体
// userRepository型は，dbというフィールドを持たなければいけない。
type userRepository struct {
	db *gorm.DB
}

// コンストラクタ関数：dbを引数にとって，IUserRepositoryを返す関数。
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

// 以下，userRepositry構造体に紐づく関数を作成する。
func (ur *userRepository) GetUserByEmail(user *model.User, email string) error {
	if err := ur.db.Where("email=?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) CreateUser(user *model.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
