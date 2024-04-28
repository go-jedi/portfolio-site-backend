package tests

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/go-jedi/portfolio/internal/model/image"
	"github.com/go-jedi/portfolio/internal/model/project"
	"github.com/go-jedi/portfolio/internal/repository"
	repoMocks "github.com/go-jedi/portfolio/internal/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	t.Parallel()
	//	Arrange
	type projectRepositoryMockFunc func(mc *gomock.Controller) repository.ProjectRepository

	mc := gomock.NewController(t)
	defer mc.Finish()

	type input struct {
		ctx   context.Context
		page  int
		limit int
	}

	type expected struct {
		projects []project.Get
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

		repoErr = fmt.Errorf("repository error")
	)

	tests := []struct {
		name                      string
		input                     input
		expected                  expected
		projectRepositoryMockFunc projectRepositoryMockFunc
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
				err:      nil,
			},
			projectRepositoryMockFunc: func(mc *gomock.Controller) repository.ProjectRepository {
				mock := repoMocks.NewMockProjectRepository(mc)
				mock.EXPECT().Get(ctx, page, limit).Return(projects, nil)
				return mock
			},
		},
		{
			name: "Repository error case",
			input: input{
				ctx:   ctx,
				page:  page,
				limit: limit,
			},
			expected: expected{
				projects: nil,
				err:      repoErr,
			},
			projectRepositoryMockFunc: func(mc *gomock.Controller) repository.ProjectRepository {
				mock := repoMocks.NewMockProjectRepository(mc)
				mock.EXPECT().Get(ctx, page, limit).Return(nil, repoErr)
				return mock
			},
		},
	}
	//	Act
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			projectRepositoryMock := test.projectRepositoryMockFunc(mc)
			result, err := projectRepositoryMock.Get(test.input.ctx, test.input.page, test.input.limit)

			const caseOk = "OK (Get)"
			const caseError = "Repository error case"

			switch test.name {
			case caseOk:
				//	Assert
				require.Equal(t, test.expected.projects, result)
				require.Equal(t, test.expected.err, err)
				require.NoError(t, err)
			case caseError:
				//	Assert
				require.Equal(t, test.expected.projects, result)
				require.Equal(t, test.expected.err, err)
				require.Error(t, err)
			}
		})
	}
}
