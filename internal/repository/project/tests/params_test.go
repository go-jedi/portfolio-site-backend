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

func TestParams(t *testing.T) {
	t.Parallel()
	//	Arrange
	type projectRepositoryMockFunc func(mc *gomock.Controller) repository.ProjectRepository

	mc := gomock.NewController(t)
	defer mc.Finish()

	type input struct {
		ctx context.Context
	}

	type expected struct {
		params project.Params
		err    error
	}

	var (
		ctx = context.Background()

		params = project.Params{
			PageCount: gofakeit.IntRange(1, 1000),
			Limit:     gofakeit.IntRange(1, 20),
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
			name: "OK (Params)",
			input: input{
				ctx: ctx,
			},
			expected: expected{
				params: params,
				err:    nil,
			},
			projectRepositoryMockFunc: func(mc *gomock.Controller) repository.ProjectRepository {
				mock := repoMocks.NewMockProjectRepository(mc)
				mock.EXPECT().Params(ctx).Return(params, nil)
				return mock
			},
		},
		{
			name: "Repository error case",
			input: input{
				ctx: ctx,
			},
			expected: expected{
				params: project.Params{},
				err:    repoErr,
			},
			projectRepositoryMockFunc: func(mc *gomock.Controller) repository.ProjectRepository {
				mock := repoMocks.NewMockProjectRepository(mc)
				mock.EXPECT().Params(ctx).Return(project.Params{}, repoErr)
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
			result, err := projectRepositoryMock.Params(test.input.ctx)

			const caseOk = "OK (Params)"
			const caseError = "Repository error case"

			switch test.name {
			case caseOk:
				//	Assert
				require.Equal(t, test.expected.params, result)
				require.Equal(t, test.expected.err, err)
				require.NoError(t, err)
			case caseError:
				//	Assert
				require.Equal(t, test.expected.params, result)
				require.Equal(t, test.expected.err, err)
				require.Error(t, err)
			}
		})
	}
}
