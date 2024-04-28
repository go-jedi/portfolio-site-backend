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

func TestGet(t *testing.T) {
	t.Parallel()
	//	Arrange
	type reviewRepositoryMockFunc func(mc *gomock.Controller) repository.ReviewRepository

	mc := gomock.NewController(t)
	defer mc.Finish()

	type input struct {
		ctx   context.Context
		page  int
		limit int
	}

	type expected struct {
		reviews []review.Review
		err     error
	}

	var (
		ctx = context.TODO()

		page  = gofakeit.IntRange(1, 100)
		limit = gofakeit.IntRange(1, 10)

		reviews = []review.Review{
			{
				ID:        gofakeit.IntRange(1, 10000),
				Username:  gofakeit.Username(),
				Message:   gofakeit.ProductDescription(),
				Rating:    gofakeit.IntRange(0, 5),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
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
			name: "OK (Get)",
			input: input{
				ctx:   ctx,
				page:  page,
				limit: limit,
			},
			expected: expected{
				reviews: reviews,
				err:     nil,
			},
			reviewRepositoryMockFunc: func(mc *gomock.Controller) repository.ReviewRepository {
				mock := repoMocks.NewMockReviewRepository(mc)
				mock.EXPECT().Get(ctx, page, limit).Return(reviews, nil)
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
				reviews: nil,
				err:     repoErr,
			},
			reviewRepositoryMockFunc: func(mc *gomock.Controller) repository.ReviewRepository {
				mock := repoMocks.NewMockReviewRepository(mc)
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
			reviewRepositoryMock := test.reviewRepositoryMockFunc(mc)
			result, err := reviewRepositoryMock.Get(test.input.ctx, test.input.page, test.input.limit)

			const caseOk = "OK (Get)"
			const caseError = "Repository error case"

			switch test.name {
			case caseOk:
				//	Assert
				require.Equal(t, test.expected.reviews, result)
				require.Equal(t, test.expected.err, err)
				require.NoError(t, err)
			case caseError:
				//	Assert
				require.Equal(t, test.expected.reviews, result)
				require.Equal(t, test.expected.err, err)
				require.Error(t, err)
			}
		})
	}
}
