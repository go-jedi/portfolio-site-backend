package tests

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/go-jedi/portfolio/internal/model/image"
	"github.com/go-jedi/portfolio/internal/model/project"
	"github.com/go-jedi/portfolio/internal/service"
	servMocks "github.com/go-jedi/portfolio/internal/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetByID(t *testing.T) {
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
		proj project.Get
		err  error
	}

	var (
		ctx = context.TODO()

		id = gofakeit.IntRange(1, 10000)

		proj = project.Get{
			ID:          id,
			Title:       gofakeit.JobTitle(),
			Description: gofakeit.ProductDescription(),
			Technology:  gofakeit.ProductCategory(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Paths: []image.Get{
				{
					ID:        gofakeit.IntRange(1, 10000),
					ProjectID: id,
					CreatedAt: time.Now().String(),
					UpdatedAt: time.Now().String(),
					Filename:  gofakeit.Animal(),
				},
			},
		}

		servErr = fmt.Errorf("service error")
	)

	tests := []struct {
		name                   string
		input                  input
		expected               expected
		projectServiceMockFunc projectServiceMockFunc
	}{
		{
			name: "OK (GetByID)",
			input: input{
				ctx: ctx,
				id:  id,
			},
			expected: expected{
				proj: proj,
				err:  nil,
			},
			projectServiceMockFunc: func(mc *gomock.Controller) service.ProjectService {
				mock := servMocks.NewMockProjectService(mc)
				mock.EXPECT().GetByID(ctx, id).Return(proj, nil)
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
				proj: project.Get{},
				err:  servErr,
			},
			projectServiceMockFunc: func(mc *gomock.Controller) service.ProjectService {
				mock := servMocks.NewMockProjectService(mc)
				mock.EXPECT().GetByID(ctx, id).Return(project.Get{}, servErr)
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
			result, err := projectServiceMock.GetByID(test.input.ctx, test.input.id)

			const caseOk = "OK (GetByID)"
			const caseError = "Service error case"

			switch test.name {
			case caseOk:
				//	Assert
				require.Equal(t, test.expected.proj, result)
				require.Equal(t, test.expected.err, err)
				require.NoError(t, err)
			case caseError:
				//	Assert
				require.Equal(t, test.expected.proj, result)
				require.Equal(t, test.expected.err, err)
				require.Error(t, err)
			}
		})
	}
}
