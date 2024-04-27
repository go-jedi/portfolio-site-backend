package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/go-jedi/portfolio/internal/repository"

	repoMocks "github.com/go-jedi/portfolio/internal/repository/mocks"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	// Arrange
	type imageRepositoryMockFunc func(mc *gomock.Controller) repository.ImageRepository

	mc := gomock.NewController(t)
	defer mc.Finish()

	type input struct {
		ctx       context.Context
		id        int
		filenames []string
	}

	var (
		ctx = context.Background()

		id        = gofakeit.IntRange(1, 10000)
		filenames = []string{"test.jpg", "test.jpeg", "test.png", "test.svg"}

		repoErr = fmt.Errorf("repository error")
	)

	tests := []struct {
		name                    string
		input                   input
		expected                error
		imageRepositoryMockFunc imageRepositoryMockFunc
	}{
		{
			name: "OK (Create)",
			input: input{
				ctx:       ctx,
				id:        id,
				filenames: filenames,
			},
			expected: nil,
			imageRepositoryMockFunc: func(mc *gomock.Controller) repository.ImageRepository {
				mock := repoMocks.NewMockImageRepository(mc)
				mock.EXPECT().Create(ctx, id, filenames).Return(nil)
				return mock
			},
		},
		{
			name: "Repository error case",
			input: input{
				ctx:       ctx,
				id:        id,
				filenames: filenames,
			},
			expected: repoErr,
			imageRepositoryMockFunc: func(mc *gomock.Controller) repository.ImageRepository {
				mock := repoMocks.NewMockImageRepository(mc)
				mock.EXPECT().Create(ctx, id, filenames).Return(repoErr)
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
			err := imageRepositoryMock.Create(test.input.ctx, test.input.id, test.input.filenames)

			//	Assert
			require.Equal(t, test.expected, err)
		})
	}
}
