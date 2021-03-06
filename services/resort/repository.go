package resort

import (
	"context"
	"github.com/models"
)

type Repository interface {
	GetAll(ctx context.Context, id []string,capacity int,limit ,offset int) ([]*models.ResortJoin, error)
	GetById(ctx context.Context, id string) ([]*models.ResortJoinDetail, error)
	GetAllCount(ctx context.Context, id []string,capacity int) (int, error)
}