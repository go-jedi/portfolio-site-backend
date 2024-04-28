package tests

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/go-jedi/portfolio/internal/model/review"
	"github.com/go-jedi/portfolio/internal/service"
	servMocks "github.com/go-jedi/portfolio/internal/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	t.Parallel()
	//	Arrange
	type reviewServiceMockFunc func(mc *gomock.Controller) service.ReviewService

	mc := gomock.NewController(t)
	defer mc.Finish()

	type input struct {
		ctx   context.Context
		page  int
		limit int
	}

	type expected struct {
		reviews []review.Review
		params  review.Params
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

		params = review.Params{
			PageCount: gofakeit.IntRange(1, 1000),
			Limit:     gofakeit.IntRange(1, 20),
		}

		servErr = fmt.Errorf("service error")
	)

	tests := []struct {
		name                  string
		input                 input
		expected              expected
		reviewServiceMockFunc reviewServiceMockFunc
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
				params:  params,
				err:     nil,
			},
			reviewServiceMockFunc: func(mc *gomock.Controller) service.ReviewService {
				mock := servMocks.NewMockReviewService(mc)
				mock.EXPECT().Get(ctx, page, limit).Return(reviews, params, nil)
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
				reviews: nil,
				params:  review.Params{},
				err:     servErr,
			},
			reviewServiceMockFunc: func(mc *gomock.Controller) service.ReviewService {
				mock := servMocks.NewMockReviewService(mc)
				mock.EXPECT().Get(ctx, page, limit).Return(nil, review.Params{}, servErr)
				return mock
			},
		},
	}
	//	Act
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			reviewServiceMock := test.reviewServiceMockFunc(mc)
			result, params, err := reviewServiceMock.Get(test.input.ctx, test.input.page, test.input.limit)

			const caseOk = "OK (Get)"
			const caseError = "Service error case"

			switch test.name {
			case caseOk:
				//	Assert
				require.Equal(t, test.expected.reviews, result)
				require.Equal(t, test.expected.params, params)
				require.Equal(t, test.expected.err, err)
				require.NoError(t, err)
			case caseError:
				//	Assert
				require.Equal(t, test.expected.reviews, result)
				require.Equal(t, test.expected.params, params)
				require.Equal(t, test.expected.err, err)
				require.Error(t, err)
			}
		})
	}
}
