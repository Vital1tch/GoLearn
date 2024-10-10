package Errors

import (
	"errors"
)

func readFile(filename string) (string, error) {
	// Проверяем существование файла (симулируем)
	if filename != "heeel.txt" {
		return "", errors.New("file not found")
	}
	// Возвращаем содержимое файла (симуляция)
	return "File content: Hello World", nil
}
