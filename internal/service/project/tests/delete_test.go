package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/go-jedi/portfolio/internal/service"
	servMocks "github.com/go-jedi/portfolio/internal/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestDelete(t *testing.T) {
	t.Parallel()
	// Arrange
	type projectServiceMockFunc func(mc *gomock.Controller) service.ProjectService

	mc := gomock.NewController(t)
	defer mc.Finish()

	type input struct {
		ctx context.Context
		id  int
	}

	type expected struct {
		id  int
		err error
	}

	var (
		ctx = context.TODO()

		id = gofakeit.IntRange(1, 10000)

		servErr = fmt.Errorf("service error")
	)

	tests := []struct {
		name                   string
		input                  input
		expected               expected
		projectServiceMockFunc projectServiceMockFunc
	}{
		{
			name: "OK (Delete)",
			input: input{
				ctx: ctx,
				id:  id,
			},
			expected: expected{
				id:  id,
				err: nil,
			},
			projectServiceMockFunc: func(mc *gomock.Controller) service.ProjectService {
				mock := servMocks.NewMockProjectService(mc)
				mock.EXPECT().Delete(ctx, id).Return(id, nil)
				return mock
			},
		},
		{
			name: "Service error case",
			input: input{
				ctx: ctx,
				id:  id,
			},
			expected: expected{
				id:  0,
				err: servErr,
			},
			projectServiceMockFunc: func(mc *gomock.Controller) service.ProjectService {
				mock := servMocks.NewMockProjectService(mc)
				mock.EXPECT().Delete(ctx, id).Return(0, servErr)
				return mock
			},
		},
	}
	//	Act
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			projectServiceMock := test.projectServiceMockFunc(mc)
			result, err := projectServiceMock.Delete(test.input.ctx, test.input.id)

			const caseOk = "OK (Delete)"
			const caseError = "Service error case"

			switch test.name {
			case caseOk:
				//	Assert
				require.Equal(t, test.expected.id, result)
				require.Equal(t, test.expected.err, err)
				require.NoError(t, err)
			case caseError:
				//	Assert
				require.Equal(t, test.expected.id, result)
				require.Equal(t, test.expected.err, err)
				require.Error(t, err)
			}
		})
	}
}
