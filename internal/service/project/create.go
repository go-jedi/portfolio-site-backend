package project

import (
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/go-jedi/portfolio/internal/model/project"
	"github.com/go-jedi/portfolio/pkg/logger"
	"github.com/go-jedi/portfolio/pkg/utils/dir"
)

func saveFile(fileHeader *multipart.FileHeader) (string, error) {
	path := ""

	// читаем путь до папки хранения файлов
	fileServerDir := os.Getenv("FILE_SERVER_DIR")
	if fileServerDir == "" {
		return "", errors.New("file server dir is empty")
	}

	// открываем файл переданный
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}

	// проверяем создана ли папка для хранения файлов
	isExistDir := dir.CheckExistDir(fileServerDir)
	if !isExistDir {
		err := dir.CreateDir(fileServerDir)
		if err != nil {
			return "", err
		}
	}

	// читаем папку для хранения файлов
	files, err := dir.ReadDir(fileServerDir)
	if err != nil {
		return "", err
	}

	// создаем путь до файла
	path = fmt.Sprintf("%s/%d%s",
		fileServerDir,
		len(files)+1,
		filepath.Ext(fileHeader.Filename),
	)

	// создаем файл по нужному пути
	osFile, err := os.Create(path)
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

	return path, nil
}

func (s *serv) Create(ctx context.Context, dto project.Create, files []*multipart.FileHeader) error {
	logger.Info(
		"(SERVICE PROJECT) Create...",
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
