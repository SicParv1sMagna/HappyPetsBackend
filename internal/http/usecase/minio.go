package usecase

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
)

type MinIOUseCase interface {
	UploadImage(userID, petID uint64, imageFileHeader *multipart.FileHeader) (string, error)
	RemoveImage(userID, petID uint64) error
}

func (uc *UseCase) UploadImage(userID, petID uint64, imageFileHeader *multipart.FileHeader) (string, error) {
	// Проверка наличия файла изображения
	if imageFileHeader == nil {
		return "", errors.New("файл изображения не найден")
	}

	// Открытие файла изображения
	imageFile, err := imageFileHeader.Open()
	if err != nil {
		return "", fmt.Errorf("ошибка при открытии файла изображения: %v", err)
	}
	defer imageFile.Close()

	// Чтение изображения в байты
	imageBytes, err := io.ReadAll(imageFile)
	if err != nil {
		return "", fmt.Errorf("ошибка при чтении файла изображения: %v", err)
	}

	// Вызов метода репозитория для загрузки изображения в MinIO
	imageURL, err := uc.Repository.UploadServiceImage(userID, petID, imageBytes, imageFileHeader.Header.Get("Content-Type"))
	if err != nil {
		return "", err
	}

	return imageURL, nil
}

func (uc *UseCase) RemoveImage(userID uint64, petID uint64) error {
	// Вызов метода репозитория для удаления изображения из MinIO
	err := uc.Repository.RemoveServiceImage(userID, petID)
	if err != nil {
		return err
	}

	return nil
}
