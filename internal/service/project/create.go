package project

import (
	"context"

	"github.com/go-jedi/portfolio/pkg/logger"
)

func (s *serv) Create(ctx context.Context) error {
	logger.Info(
		"(SERVICE PROJECT) Create...",
	)

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		// Загрузить фото в директорию
		// сохранить данные в таблицу projects и вернуть id записи
		// сохранить данные в таблицу images

		err := s.projectRepository.Create(ctx)
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
