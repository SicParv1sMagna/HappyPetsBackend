package main

import (
	"github.com/SicParv1sMagna/HappyPetsBackend/internal/dsn"
	"github.com/SicParv1sMagna/HappyPetsBackend/internal/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
    _ = godotenv.Load()
    db, err := gorm.Open(postgres.Open(dsn.FromEnv()), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // Явно мигрировать только нужные таблицы
    err = db.AutoMigrate(&model.Pet{})
    if err != nil {
        panic("cant migrate db")
    }
}