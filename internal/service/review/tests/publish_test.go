package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/go-jedi/portfolio/internal/service"
	servMocks "github.com/go-jedi/portfolio/internal/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestPublish(t *testing.T) {
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
		id  int
		err error
	}

	var (
		ctx = context.TODO()

		id = gofakeit.IntRange(1, 10000)

		servErr = fmt.Errorf("service error")
	)

	tests := []struct {
		name                  string
		input                 input
		expected              expected
		reviewServiceMockFunc reviewServiceMockFunc
	}{
		{
			name: "OK (Publish)",
			input: input{
				ctx: ctx,
				id:  id,
			},
			expected: expected{
				id:  id,
				err: nil,
			},
			reviewServiceMockFunc: func(mc *gomock.Controller) service.ReviewService {
				mock := servMocks.NewMockReviewService(mc)
				mock.EXPECT().Publish(ctx, id).Return(id, nil)
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
				id:  0,
				err: servErr,
			},
			reviewServiceMockFunc: func(mc *gomock.Controller) service.ReviewService {
				mock := servMocks.NewMockReviewService(mc)
				mock.EXPECT().Publish(ctx, id).Return(0, servErr)
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
			result, err := reviewServiceMock.Publish(test.input.ctx, test.input.id)

			const caseOk = "OK (Publish)"
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
