package rest_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/churmd/to-do-lists/internal/adapters/inbound/rest"
	"github.com/churmd/to-do-lists/internal/domain/models"
	inboundports "github.com/churmd/to-do-lists/internal/mocks/internal_/domain/inbound_ports"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestStrictChiServer_GetApiV1UsersUserIdTodos(t *testing.T) {
	testCases := []GetApiV1UsersUserIdTodosTestCase{
		{
			name:          "No Todos returned",
			returnedTodos: []models.Todo{},
			expectedTodos: rest.GetApiV1UsersUserIdTodos200JSONResponse{},
		},
	}
	testCases = append(testCases, generateGetApiV1UsersUserIdTodosTestCases()...)

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mockTodoLister := inboundports.NewTodoLister(t)

			mockTodoLister.EXPECT().
				GetTodos(mock.Anything, mock.Anything).
				Return(testCase.returnedTodos, testCase.returnedError)

			server := rest.NewStrictChiServer(mockTodoLister)
			resp, err := server.GetApiV1UsersUserIdTodos(context.Background(), rest.GetApiV1UsersUserIdTodosRequestObject{UserId: "1"})

			assert.ErrorIs(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedTodos, resp)
		})
	}
}

type GetApiV1UsersUserIdTodosTestCase struct {
	name          string
	returnedTodos []models.Todo
	returnedError error
	expectedTodos rest.GetApiV1UsersUserIdTodosResponseObject
	expectedError error
}

func generateGetApiV1UsersUserIdTodosTestCases() []GetApiV1UsersUserIdTodosTestCase {
	testCases := make([]GetApiV1UsersUserIdTodosTestCase, 0)
	for i := range 50 {
		respTodos := make([]rest.Todo, 0)
		for j := range i {
			respTodos = append(respTodos, rest.Todo{
				Id:          fmt.Sprintf("ID %d", j),
				Description: fmt.Sprintf("Desc %d", j),
				Created:     time.Now(),
				Completed:   j%2 == 0,
			})
		}

		domainTodos := make([]models.Todo, 0)
		for _, respTodo := range respTodos {
			domainTodos = append(domainTodos, models.Todo{
				ID:          respTodo.Id,
				UserID:      models.UserID(fmt.Sprintf("userid-%s", respTodo.Id)),
				Description: respTodo.Description,
				Created:     respTodo.Created,
				Complete:    respTodo.Completed,
			})
		}

		testCase := GetApiV1UsersUserIdTodosTestCase{
			name:          fmt.Sprintf("%d number of todos", i),
			expectedTodos: rest.GetApiV1UsersUserIdTodos200JSONResponse(respTodos),
			returnedTodos: domainTodos,
		}

		testCases = append(testCases, testCase)
	}
	return testCases
}
