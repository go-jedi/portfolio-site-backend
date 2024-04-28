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

func TestGetByID(t *testing.T) {
	t.Parallel()
	//	Arrange
	type projectRepositoryMockFunc func(mc *gomock.Controller) repository.ProjectRepository

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
		ctx = context.Background()

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

		repoErr = fmt.Errorf("repository error")
	)

	tests := []struct {
		name                      string
		input                     input
		expected                  expected
		projectRepositoryMockFunc projectRepositoryMockFunc
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
			projectRepositoryMockFunc: func(mc *gomock.Controller) repository.ProjectRepository {
				mock := repoMocks.NewMockProjectRepository(mc)
				mock.EXPECT().GetByID(ctx, id).Return(proj, nil)
				return mock
			},
		},
		{
			name: "Repository error case",
			input: input{
				ctx: ctx,
				id:  id,
			},
			expected: expected{
				proj: project.Get{},
				err:  repoErr,
			},
			projectRepositoryMockFunc: func(mc *gomock.Controller) repository.ProjectRepository {
				mock := repoMocks.NewMockProjectRepository(mc)
				mock.EXPECT().GetByID(ctx, id).Return(project.Get{}, repoErr)
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
			result, err := projectRepositoryMock.GetByID(test.input.ctx, test.input.id)

			const caseOk = "OK (GetByID)"
			const caseError = "Repository error case"

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
