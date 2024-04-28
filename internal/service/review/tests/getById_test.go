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

func TestGetByID(t *testing.T) {
	t.Parallel()
	//	Arrange
	type reviewServiceMockFunc func(mc *gomock.Controller) service.ReviewService

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

		servErr = fmt.Errorf("service error")
	)

	tests := []struct {
		name                  string
		input                 input
		expected              expected
		reviewServiceMockFunc reviewServiceMockFunc
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
			reviewServiceMockFunc: func(mc *gomock.Controller) service.ReviewService {
				mock := servMocks.NewMockReviewService(mc)
				mock.EXPECT().GetByID(ctx, id).Return(rvw, nil)
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
				rvw: review.Review{},
				err: servErr,
			},
			reviewServiceMockFunc: func(mc *gomock.Controller) service.ReviewService {
				mock := servMocks.NewMockReviewService(mc)
				mock.EXPECT().GetByID(ctx, id).Return(review.Review{}, servErr)
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
			result, err := reviewServiceMock.GetByID(test.input.ctx, test.input.id)

			const caseOk = "OK (GetByID)"
			const caseError = "Service error case"

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
