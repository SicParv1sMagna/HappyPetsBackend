package repository

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/SicParv1sMagna/HappyPetsBackend/internal/model"
	"github.com/minio/minio-go/v7"
)

type MinioRepository interface{
    UploadServiceImage(userID, petID uint64, imageBytes []byte, contentType string) (string, error)
    RemoveServiceImage(userID, petID uint64) error
}


// UploadServiceImage загружает изображение в MinIO bucket и возвращает URL изображения.
func (r *Repository) UploadServiceImage(userID, petID uint64, imageBytes []byte, contentType string) (string, error) {
    objectName := fmt.Sprintf("users/%d/pets/avatars-%d", userID, petID)

    reader := io.NopCloser(bytes.NewReader(imageBytes))

    _, err := r.mc.PutObject(context.TODO(), "happypets-image", objectName, reader, int64(len(imageBytes)), minio.PutObjectOptions{
        ContentType: contentType,
    })
    if err != nil {
        return "", errors.New("ошибка при добавлении изображения в минио бакет")
    }

    // Формирование URL изображения
    imageURL := fmt.Sprintf("http://localhost:9000/happypets-image/%s", objectName)

    // Обновление поля photo в таблице Pet с новым URL изображения
    if err := r.db.Table("pet").Where("id = ? AND status = ?", petID, model.PET_STATUS_ACTIVE).Update("photo", imageURL).Error; err != nil {
        return "", errors.New("ошибка при обновлении URL изображения в базе данных")
    }

    return imageURL, nil
}

// RemoveServiceImage удаляет изображение из bucket MinIO.
func (r *Repository) RemoveServiceImage(userID, petID uint64) error {
    objectName := fmt.Sprintf("users/%d/pets/avatars-%d", userID, petID)
    err := r.mc.RemoveObject(context.TODO(), "happypets-image", objectName, minio.RemoveObjectOptions{})
    if err != nil {
        return errors.New("ошибка при удалении изображения из минио бакета")
    }

    // Обновление поля photo в таблице Pet на NULL
    if err := r.db.Table("pet").Where("id = ?", petID).Update("photo", nil).Error; err != nil {
        return errors.New("ошибка при обновлении URL изображения в базе данных")
    }

    return nil
}
