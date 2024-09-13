package rest

import (
	"context"

	inboundports "github.com/churmd/to-do-lists/internal/domain/inbound_ports"
	"github.com/churmd/to-do-lists/internal/domain/models"
)

type StrictChiServer struct {
	TodoLister inboundports.TodoLister
}

func NewStrictChiServer(todoLister inboundports.TodoLister) StrictChiServer {
	return StrictChiServer{
		TodoLister: todoLister,
	}
}

// GetApiV1UsersUserIdTodos
// Returns a list of todos for the given user. (GET /api/v1/users/{user_id}/todos)
func (s StrictChiServer) GetApiV1UsersUserIdTodos(ctx context.Context, request GetApiV1UsersUserIdTodosRequestObject) (GetApiV1UsersUserIdTodosResponseObject, error) {
	todos, err := s.TodoLister.GetTodos(ctx, models.UserID(request.UserId))
	if err != nil {
		panic(err)
	}

	respTodos := make([]Todo, 0)
	for _, todo := range todos {
		respTodos = append(respTodos, Todo{
			Completed:   todo.Complete,
			Created:     todo.Created,
			Description: todo.Description,
			Id:          todo.ID,
		})
	}

	return GetApiV1UsersUserIdTodos200JSONResponse(respTodos), nil
}
