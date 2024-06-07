package repositories

import (
	"github.com/Gabriel-Macedogmc/grpc-golang/src/infra/gorm/interfaces"
	"github.com/Gabriel-Macedogmc/grpc-golang/src/models"
	"github.com/Gabriel-Macedogmc/grpc-golang/src/proto"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) interfaces.CategoryRepository {
	return &repository{
		db: db,
	}
}

// Create implements interfaces.CategoryRepository.
func (r *repository) Create(category *proto.Category) (*models.Category, error) {
	categoryCreate := &models.Category{
		Name:        category.Name,
		Description: category.Description,
		IsActive:    category.IsActive,
	}

	r.db.Create(categoryCreate)

	return categoryCreate, nil
}

// List implements interfaces.CategoryRepository.
func (r *repository) List() ([]models.Category, error) {
	categories := []models.Category{}
	r.db.Find(&categories)

	return categories, nil
}
