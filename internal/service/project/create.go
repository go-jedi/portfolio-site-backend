package project

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/go-jedi/portfolio/internal/model/project"
	"github.com/go-jedi/portfolio/pkg/logger"
	"github.com/go-jedi/portfolio/pkg/utils/dir"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func saveFile(fileHeader *multipart.FileHeader) (string, error) {
	filename := ""

	// читаем путь до папки хранения файлов
	fileServerDir := os.Getenv("FILE_SERVER_DIR")
	if fileServerDir == "" {
		return "", ErrFileServerDirEmpty
	}

	// открываем файл переданный
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}

	// проверяем создана ли папка для хранения файлов
	isExistDir := dir.CheckExistDir(fileServerDir)
	if !isExistDir {
		err = dir.CreateDir(fileServerDir)
		if err != nil {
			return "", err
		}
	}

	// генерируем UUID для создания уникального имени файла
	id := uuid.New()

	// создаем имя файла
	filename = fmt.Sprintf("%s%s",
		id.String(),
		filepath.Ext(fileHeader.Filename),
	)

	// создаем файл по нужному пути
	osFile, err := os.Create(
		fmt.Sprintf("%s/%s",
			fileServerDir,
			filename,
		),
	)
	if err != nil {
		return "", err
	}

	// читаем переданный файл и получаем его байты
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	// записываем в созданный файл данные из переданного файла
	_, err = osFile.Write(fileBytes)
	if err != nil {
		return "", err
	}

	// закрываем переданный файл
	err = file.Close()
	if err != nil {
		return "", err
	}

	// закрываем созданный файл
	err = osFile.Close()
	if err != nil {
		return "", err
	}

	return filename, nil
}

func (s *serv) Create(ctx context.Context, dto project.Create, files []*multipart.FileHeader) error {
	logger.Info(
		"(SERVICE PROJECT) Create...",
		zap.String("title", dto.Title),
		zap.String("description", dto.Description),
		zap.String("technology", dto.Technology),
	)

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		// Загрузить фотографию(и) в директорию
		paths := make([]string, 0, len(files))

		for key := range files {
			result, err := saveFile(files[key])
			if err != nil {
				return err
			}
			paths = append(paths, result)
		}

		// сохранить данные в таблицу projects и вернуть id записи
		result, err := s.projectRepository.Create(ctx, dto)
		if err != nil {
			return err
		}

		// сохранить данные в таблицу images
		err = s.imageRepository.Create(ctx, result, paths)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
