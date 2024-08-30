package inboundports

import (
	"context"

	"github.com/churmd/to-do-lists/internal/domain/models"
)

type TodoLister interface {
	GetTodos(ctx context.Context, userID models.UserID) ([]models.Todo, error)
}
