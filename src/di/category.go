package di

import (
	"github.com/Gabriel-Macedogmc/grpc-golang/src/infra/gorm/interfaces"
	"github.com/Gabriel-Macedogmc/grpc-golang/src/infra/gorm/repositories"
	"gorm.io/gorm"
)

func ConfigCategoryDI(conn *gorm.DB) interfaces.CategoryRepository {
	categoryRepository := repositories.New(conn)
	return categoryRepository
}
