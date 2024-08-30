package rest

import "context"

type StrictChiServer struct{}

// Returns a list of todos for the given user.
// (GET /api/v1/users/{user_id}/todos)
func (s StrictChiServer) GetApiV1UsersUserIdTodos(ctx context.Context, request GetApiV1UsersUserIdTodosRequestObject) (GetApiV1UsersUserIdTodosResponseObject, error) {
	panic("not implemented") // TODO: Implement
}
