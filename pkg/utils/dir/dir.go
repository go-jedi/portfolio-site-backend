package dir

import "os"

// CheckExistDir проверка существует ли папка или нет.
func CheckExistDir(path string) bool {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return true
	}

	return false
}

func ReadDir(path string) ([]os.DirEntry, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	return files, nil
}

// CreateDir создать директорию.
func CreateDir(path string) error {
	err := os.Mkdir(path, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// CreateDirAll создать иерархию директорий (a/b/c/d).
func CreateDirAll(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
