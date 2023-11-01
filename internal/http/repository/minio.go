package repository

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/minio/minio-go/v7"
)

// UploadServiceImage загружает изображение в MinIO bucket и возвращает URL изображения.
func (r *Repository) UploadServiceImage(userID, petID uint64, imageBytes []byte, contentType string) (string, error) {
    objectName := fmt.Sprintf("users/%d/pets/avatars-%d", userID, petID)

    reader := io.NopCloser(bytes.NewReader(imageBytes))

    _, err := r.mc.PutObject(context.TODO(), "happypets-image", objectName, reader, int64(len(imageBytes)), minio.PutObjectOptions{
        ContentType: contentType,
    })
    if err != nil {
        return "", err
    }

    // Формирование URL изображения
    imageURL := fmt.Sprintf("http://localhost:9000/happypets-image/%s", objectName)
    return imageURL, nil
}

// RemoveServiceImage удаляет изображение из bucket MinIO.
func (r *Repository) RemoveServiceImage(userID, petID uint64) error {
    objectName := fmt.Sprintf("users/%d/pets/avatars-%d", userID, petID)
    err := r.mc.RemoveObject(context.TODO(), "happypets-image", objectName, minio.RemoveObjectOptions{})
    if err != nil {
        fmt.Println("Failed to remove object from MinIO:", err)
        // Обработка ошибки удаления изображения из MinIO
        return err
    }
    fmt.Println("Object removed from MinIO successfully:", objectName)
    return nil
}
