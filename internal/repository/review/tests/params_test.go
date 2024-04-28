package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/go-jedi/portfolio/internal/model/review"
	"github.com/go-jedi/portfolio/internal/repository"
	repoMocks "github.com/go-jedi/portfolio/internal/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestParams(t *testing.T) {
	t.Parallel()
	//	Arrange
	type reviewRepositoryMockFunc func(mc *gomock.Controller) repository.ReviewRepository

	mc := gomock.NewController(t)
	defer mc.Finish()

	type input struct {
		ctx context.Context
	}

	type expected struct {
		params review.Params
		err    error
	}

	var (
		ctx = context.TODO()

		params = review.Params{
			PageCount: gofakeit.IntRange(1, 1000),
			Limit:     gofakeit.IntRange(1, 20),
		}

		repoErr = fmt.Errorf("repository error")
	)

	tests := []struct {
		name                     string
		input                    input
		expected                 expected
		reviewRepositoryMockFunc reviewRepositoryMockFunc
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
			reviewRepositoryMockFunc: func(mc *gomock.Controller) repository.ReviewRepository {
				mock := repoMocks.NewMockReviewRepository(mc)
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
				params: review.Params{},
				err:    repoErr,
			},
			reviewRepositoryMockFunc: func(mc *gomock.Controller) repository.ReviewRepository {
				mock := repoMocks.NewMockReviewRepository(mc)
				mock.EXPECT().Params(ctx).Return(review.Params{}, repoErr)
				return mock
			},
		},
	}
	//	Act
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			reviewRepositoryMock := test.reviewRepositoryMockFunc(mc)
			result, err := reviewRepositoryMock.Params(test.input.ctx)

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
