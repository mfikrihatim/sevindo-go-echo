package accessibility

import (
	"github.com/models"
	"golang.org/x/net/context"
)

type Usecase interface {
	Delete(ctx context.Context, id int, token string) (*models.ResponseDelete, error)
	Update(ctx context.Context, ar *models.NewCommandAccessibility,  token string) error
	List(ctx context.Context, page, limit, offset int, search string) (*models.AccessibilityWithPagination, error)
	Create(ctx context.Context, ar *models.NewCommandAccessibility, token string) (*models.NewCommandAccessibility, error)
	GetById(ctx context.Context, id int, token string) (*models.AccessibilityDto, error)
}
