package repository

import (
	"github.com/RafaelFernando12/api-golang/domain"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewDb(url string) (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	if url != "" {
		db, err = gorm.Open(postgres.Open(url), &gorm.Config{})
	} else {
		db, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	}
	return db, err
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	db.Table("user").AutoMigrate(&domain.User{})
	return &userRepository{
		db: db,
	}
}

func (urep *userRepository) Save(user *domain.User) error {
	return nil
}
