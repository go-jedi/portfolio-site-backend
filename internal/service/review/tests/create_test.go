package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/go-jedi/portfolio/internal/model/review"
	"github.com/go-jedi/portfolio/internal/service"
	servMocks "github.com/go-jedi/portfolio/internal/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	//	Arrange
	type reviewServiceMockFunc func(mc *gomock.Controller) service.ReviewService

	mc := gomock.NewController(t)
	defer mc.Finish()

	type input struct {
		ctx context.Context
		dto review.Create
	}

	type expected struct {
		id  int
		err error
	}

	var (
		ctx = context.TODO()

		dto = review.Create{
			Username: gofakeit.Username(),
			Message:  gofakeit.ProductDescription(),
			Rating:   gofakeit.IntRange(0, 5),
		}

		servRes = gofakeit.IntRange(1, 10000)
		servErr = fmt.Errorf("service error")
	)

	tests := []struct {
		name                  string
		input                 input
		expected              expected
		reviewServiceMockFunc reviewServiceMockFunc
	}{
		{
			name: "OK (Create)",
			input: input{
				ctx: ctx,
				dto: dto,
			},
			expected: expected{
				id:  servRes,
				err: nil,
			},
			reviewServiceMockFunc: func(mc *gomock.Controller) service.ReviewService {
				mock := servMocks.NewMockReviewService(mc)
				mock.EXPECT().Create(ctx, dto).Return(servRes, nil)
				return mock
			},
		},
		{
			name: "Service error case",
			input: input{
				ctx: ctx,
				dto: dto,
			},
			expected: expected{
				id:  0,
				err: servErr,
			},
			reviewServiceMockFunc: func(mc *gomock.Controller) service.ReviewService {
				mock := servMocks.NewMockReviewService(mc)
				mock.EXPECT().Create(ctx, dto).Return(0, servErr)
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
			result, err := reviewServiceMock.Create(test.input.ctx, test.input.dto)

			const caseOk = "OK (Create)"
			const caseError = "Service error case"

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
