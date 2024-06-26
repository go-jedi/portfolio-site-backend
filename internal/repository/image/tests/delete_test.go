package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/go-jedi/portfolio/internal/repository"
	repoMocks "github.com/go-jedi/portfolio/internal/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestDelete(t *testing.T) {
	t.Parallel()
	//	Arrange
	type imageRepositoryMockFunc func(mc *gomock.Controller) repository.ImageRepository

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

		repoErr = fmt.Errorf("repository error")
	)

	tests := []struct {
		name                    string
		input                   input
		expected                expected
		imageRepositoryMockFunc imageRepositoryMockFunc
	}{
		{
			name: "OK (Delete)",
			input: input{
				ctx: ctx,
				id:  id,
			},
			expected: expected{
				id:  id,
				err: nil,
			},
			imageRepositoryMockFunc: func(mc *gomock.Controller) repository.ImageRepository {
				mock := repoMocks.NewMockImageRepository(mc)
				mock.EXPECT().Delete(ctx, id).Return(id, nil)
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
				id:  0,
				err: repoErr,
			},
			imageRepositoryMockFunc: func(mc *gomock.Controller) repository.ImageRepository {
				mock := repoMocks.NewMockImageRepository(mc)
				mock.EXPECT().Delete(ctx, id).Return(0, repoErr)
				return mock
			},
		},
	}
	//	Act
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			imageRepositoryMock := test.imageRepositoryMockFunc(mc)
			result, err := imageRepositoryMock.Delete(test.input.ctx, test.input.id)

			const caseOk = "OK (Delete)"
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
