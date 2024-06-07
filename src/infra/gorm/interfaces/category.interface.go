package interfaces

import (
	"github.com/Gabriel-Macedogmc/grpc-golang/src/models"
	pb "github.com/Gabriel-Macedogmc/grpc-golang/src/proto"
)

type CategoryRepository interface {
	Create(category *pb.Category) (*models.Category, error)
	List() ([]models.Category, error)
}
