package repository

import (
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
	mc *minio.Client
}

func New(dsn string) (*Repository, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	accessKey := os.Getenv("MINIO_ACCESS_KEY")
	secretKey := os.Getenv("MINIO_SECRET_KEY")
	endpoint  := os.Getenv("MINIO_ENDPOINT")
	minio_client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: false,
	})
	if err != nil {
		return nil, err
	}
	return &Repository{
		db: db,
		mc: minio_client,
	}, nil
}
