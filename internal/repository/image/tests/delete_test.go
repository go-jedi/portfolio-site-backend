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

	var (
		ctx = context.Background()

		id = gofakeit.IntRange(1, 10000)

		repoErr = fmt.Errorf("repository error")
	)

	tests := []struct {
		name                    string
		input                   input
		expected                error
		imageRepositoryMockFunc imageRepositoryMockFunc
	}{
		{
			name: "OK (Delete)",
			input: input{
				ctx: ctx,
				id:  id,
			},
			expected: nil,
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
			expected: repoErr,
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
				require.Equal(t, test.expected, err)
				require.Equal(t, test.input.id, result)
			case caseError:
				//	Assert
				require.Equal(t, test.expected, err)
				require.Equal(t, 0, result)
			}
		})
	}
}
