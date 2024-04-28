package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/go-jedi/portfolio/internal/model/project"
	"github.com/go-jedi/portfolio/internal/repository"
	repoMocks "github.com/go-jedi/portfolio/internal/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	// Arrange
	type projectRepositoryMockFunc func(mc *gomock.Controller) repository.ProjectRepository

	mc := gomock.NewController(t)
	defer mc.Finish()

	type input struct {
		ctx context.Context
		dto project.Create
	}

	type expected struct {
		id  int
		err error
	}

	var (
		ctx = context.Background()

		title       = gofakeit.JobTitle()
		description = gofakeit.ProductDescription()
		technology  = gofakeit.ProductCategory()

		repoRes = gofakeit.IntRange(1, 10000)
		repoErr = fmt.Errorf("repository error")

		dto = project.Create{
			Title:       title,
			Description: description,
			Technology:  technology,
		}
	)

	tests := []struct {
		name                      string
		input                     input
		expected                  expected
		projectRepositoryMockFunc projectRepositoryMockFunc
	}{
		{
			name: "OK (Create)",
			input: input{
				ctx: ctx,
				dto: dto,
			},
			expected: expected{
				id:  repoRes,
				err: nil,
			},
			projectRepositoryMockFunc: func(mc *gomock.Controller) repository.ProjectRepository {
				mock := repoMocks.NewMockProjectRepository(mc)
				mock.EXPECT().Create(ctx, dto).Return(repoRes, nil)
				return mock
			},
		},
		{
			name: "Repository error case",
			input: input{
				ctx: ctx,
				dto: dto,
			},
			expected: expected{
				id:  0,
				err: repoErr,
			},
			projectRepositoryMockFunc: func(mc *gomock.Controller) repository.ProjectRepository {
				mock := repoMocks.NewMockProjectRepository(mc)
				mock.EXPECT().Create(ctx, dto).Return(0, repoErr)
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
			result, err := projectRepositoryMock.Create(test.input.ctx, test.input.dto)

			const caseOk = "OK (Create)"
			const caseError = "Repository error case"

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
