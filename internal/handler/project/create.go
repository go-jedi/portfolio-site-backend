package project

import (
	"path/filepath"

	"github.com/go-jedi/portfolio/internal/model/project"
	"github.com/go-jedi/portfolio/pkg/logger"
	"github.com/go-jedi/portfolio/pkg/utils/contains"
	"github.com/gofiber/fiber/v3"
)

func (h *Handler) Create(c fiber.Ctx) error {
	logger.Info(
		"(HANDLER PROJECT) Create...",
	)

	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	files := form.File["files"]

	title := c.FormValue("title")
	description := c.FormValue("description")
	technology := c.FormValue("technology")

	err = h.validator.Var(title, "required,min=1")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err = h.validator.Var(description, "required,min=1")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err = h.validator.Var(technology, "required,min=1")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	for _, file := range files {
		//	проверка размера файла
		if file.Size > 1024*1024 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "некорректный размер переданного файла",
			})
		}

		// проверка файла на необходимое расширение
		result := contains.Contains(
			[]string{".jpg", ".jpeg", ".png", ".svg"},
			filepath.Ext(file.Filename),
		)
		if !result {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "некорректное расширение переданного файла",
			})
		}
	}

	dto := project.Create{
		Title:       title,
		Description: description,
		Technology:  technology,
	}

	err = h.projectService.Create(c.UserContext(), dto, files)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "успешное создание проекта",
	})
}
