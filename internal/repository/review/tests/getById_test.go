package tests

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/go-jedi/portfolio/internal/model/review"
	"github.com/go-jedi/portfolio/internal/repository"
	repoMocks "github.com/go-jedi/portfolio/internal/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetByID(t *testing.T) {
	t.Parallel()
	//	Arrange
	type reviewRepositoryMockFunc func(mc *gomock.Controller) repository.ReviewRepository

	mc := gomock.NewController(t)
	defer mc.Finish()

	type input struct {
		ctx context.Context
		id  int
	}

	type expected struct {
		rvw review.Review
		err error
	}

	var (
		ctx = context.TODO()

		id = gofakeit.IntRange(1, 10000)

		rvw = review.Review{
			ID:        id,
			Username:  gofakeit.Username(),
			Message:   gofakeit.ProductDescription(),
			Rating:    gofakeit.IntRange(0, 5),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
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
			name: "OK (GetByID)",
			input: input{
				ctx: ctx,
				id:  id,
			},
			expected: expected{
				rvw: rvw,
				err: nil,
			},
			reviewRepositoryMockFunc: func(mc *gomock.Controller) repository.ReviewRepository {
				mock := repoMocks.NewMockReviewRepository(mc)
				mock.EXPECT().GetByID(ctx, id).Return(rvw, nil)
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
				rvw: review.Review{},
				err: repoErr,
			},
			reviewRepositoryMockFunc: func(mc *gomock.Controller) repository.ReviewRepository {
				mock := repoMocks.NewMockReviewRepository(mc)
				mock.EXPECT().GetByID(ctx, id).Return(review.Review{}, repoErr)
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
			result, err := reviewRepositoryMock.GetByID(test.input.ctx, test.input.id)

			const caseOk = "OK (GetByID)"
			const caseError = "Repository error case"

			switch test.name {
			case caseOk:
				//	Assert
				require.Equal(t, test.expected.rvw, result)
				require.Equal(t, test.expected.err, err)
				require.NoError(t, err)
			case caseError:
				//	Assert
				require.Equal(t, test.expected.rvw, result)
				require.Equal(t, test.expected.err, err)
				require.Error(t, err)
			}
		})
	}
}
