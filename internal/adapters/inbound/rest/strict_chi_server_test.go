package rest_test

import (
	"context"
	"testing"

	"github.com/churmd/to-do-lists/internal/adapters/inbound/rest"
	inboundports "github.com/churmd/to-do-lists/internal/mocks/internal_/domain/inbound_ports"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestStrictChiServer_GetApiV1UsersUserIdTodos(t *testing.T) {
	mockTodoLister := inboundports.NewTodoLister(t)

	mockTodoLister.EXPECT().GetTodos(mock.Anything, mock.Anything).Return(nil, nil)

	server := rest.NewStrictChiServer(mockTodoLister)
	resp, err := server.GetApiV1UsersUserIdTodos(context.Background(), rest.GetApiV1UsersUserIdTodosRequestObject{UserId: "1"})

	assert.NoError(t, err)
	assert.Empty(t, resp)
}
