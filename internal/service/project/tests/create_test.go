package tests

import (
	"context"
	"fmt"
	"mime/multipart"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/go-jedi/portfolio/internal/model/project"
	"github.com/go-jedi/portfolio/internal/service"
	servMocks "github.com/go-jedi/portfolio/internal/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	// Arrange
	type projectServiceMockFunc func(mc *gomock.Controller) service.ProjectService

	mc := gomock.NewController(t)
	defer mc.Finish()

	type input struct {
		ctx   context.Context
		dto   project.Create
		files []*multipart.FileHeader
	}

	var (
		ctx = context.TODO()

		title       = gofakeit.JobTitle()
		description = gofakeit.ProductDescription()
		technology  = gofakeit.ProductCategory()

		dto = project.Create{
			Title:       title,
			Description: description,
			Technology:  technology,
		}

		servErr = fmt.Errorf("service error")
	)

	tests := []struct {
		name                   string
		input                  input
		expected               error
		projectServiceMockFunc projectServiceMockFunc
	}{
		{
			name: "OK (Create)",
			input: input{
				ctx:   ctx,
				dto:   dto,
				files: []*multipart.FileHeader{},
			},
			expected: nil,
			projectServiceMockFunc: func(mc *gomock.Controller) service.ProjectService {
				mock := servMocks.NewMockProjectService(mc)
				mock.EXPECT().Create(ctx, dto, []*multipart.FileHeader{}).Return(nil)
				return mock
			},
		},
		{
			name: "Service error case",
			input: input{
				ctx:   ctx,
				dto:   dto,
				files: []*multipart.FileHeader{},
			},
			expected: servErr,
			projectServiceMockFunc: func(mc *gomock.Controller) service.ProjectService {
				mock := servMocks.NewMockProjectService(mc)
				mock.EXPECT().Create(ctx, dto, []*multipart.FileHeader{}).Return(servErr)
				return mock
			},
		},
	}
	//	Act
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			projectServiceMock := test.projectServiceMockFunc(mc)
			err := projectServiceMock.Create(test.input.ctx, test.input.dto, test.input.files)

			const caseOk = "OK (Create)"
			const caseError = "Service error case"

			switch test.name {
			case caseOk:
				//	Assert
				require.Equal(t, test.expected, err)
				require.NoError(t, err)
			case caseError:
				//	Assert
				require.Equal(t, test.expected, err)
				require.Error(t, err)
			}
		})
	}
}
