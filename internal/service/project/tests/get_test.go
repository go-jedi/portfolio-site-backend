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

func TestGet(t *testing.T) {
	t.Parallel()
	// Arrange
	type projectServiceMockFunc func(mc *gomock.Controller) service.ProjectService

	mc := gomock.NewController(t)
	defer mc.Finish()

	type input struct {
		ctx   context.Context
		page  int
		limit int
	}

	type expected struct {
		projects []project.Get
		params   project.Params
		err      error
	}

	var (
		ctx = context.TODO()

		page  = gofakeit.IntRange(1, 100)
		limit = gofakeit.IntRange(1, 10)

		projects = []project.Get{
			{
				ID:          gofakeit.IntRange(1, 10000),
				Title:       gofakeit.JobTitle(),
				Description: gofakeit.ProductDescription(),
				Technology:  gofakeit.ProductCategory(),
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
				Paths: []image.Get{
					{
						ID:        gofakeit.IntRange(1, 10000),
						ProjectID: gofakeit.IntRange(1, 10000),
						CreatedAt: time.Now().String(),
						UpdatedAt: time.Now().String(),
						Filename:  gofakeit.Animal(),
					},
				},
			},
		}

		params = project.Params{
			PageCount: gofakeit.IntRange(1, 1000),
			Limit:     gofakeit.IntRange(1, 20),
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
			name: "OK (Get)",
			input: input{
				ctx:   ctx,
				page:  page,
				limit: limit,
			},
			expected: expected{
				projects: projects,
				params:   params,
				err:      nil,
			},
			projectServiceMockFunc: func(mc *gomock.Controller) service.ProjectService {
				mock := servMocks.NewMockProjectService(mc)
				mock.EXPECT().Get(ctx, page, limit).Return(projects, params, nil)
				return mock
			},
		},
		{
			name: "Service error case",
			input: input{
				ctx:   ctx,
				page:  page,
				limit: limit,
			},
			expected: expected{
				projects: nil,
				params:   project.Params{},
				err:      servErr,
			},
			projectServiceMockFunc: func(mc *gomock.Controller) service.ProjectService {
				mock := servMocks.NewMockProjectService(mc)
				mock.EXPECT().Get(ctx, page, limit).Return(nil, project.Params{}, servErr)
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
			result, params, err := projectServiceMock.Get(test.input.ctx, test.input.page, test.input.limit)

			const caseOk = "OK (Get)"
			const caseError = "Service error case"

			switch test.name {
			case caseOk:
				//	Assert
				require.Equal(t, test.expected.projects, result)
				require.Equal(t, test.expected.params, params)
				require.Equal(t, test.expected.err, err)
				require.NoError(t, err)
			case caseError:
				//	Assert
				require.Equal(t, test.expected.projects, result)
				require.Equal(t, test.expected.params, params)
				require.Equal(t, test.expected.err, err)
				require.Error(t, err)
			}
		})
	}
}
