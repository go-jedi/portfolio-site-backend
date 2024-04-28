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

func TestCreate(t *testing.T) {
	t.Parallel()
	//	Arrange
	type reviewRepositoryMockFunc func(mc *gomock.Controller) repository.ReviewRepository

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

		repoRes = gofakeit.IntRange(1, 10000)
		repoErr = fmt.Errorf("repository error")
	)

	tests := []struct {
		name                     string
		input                    input
		expected                 expected
		reviewRepositoryMockFunc reviewRepositoryMockFunc
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
			reviewRepositoryMockFunc: func(mc *gomock.Controller) repository.ReviewRepository {
				mock := repoMocks.NewMockReviewRepository(mc)
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
			reviewRepositoryMockFunc: func(mc *gomock.Controller) repository.ReviewRepository {
				mock := repoMocks.NewMockReviewRepository(mc)
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
			reviewRepositoryMock := test.reviewRepositoryMockFunc(mc)
			result, err := reviewRepositoryMock.Create(test.input.ctx, test.input.dto)

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
