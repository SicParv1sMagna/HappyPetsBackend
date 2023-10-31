package app

import (
	"github.com/SicParv1sMagna/HappyPetsBackend/internal/dsn"
	"github.com/SicParv1sMagna/HappyPetsBackend/internal/http/delivery"
	"github.com/SicParv1sMagna/HappyPetsBackend/internal/http/repository"
	"github.com/SicParv1sMagna/HappyPetsBackend/internal/http/usecase"

	"github.com/joho/godotenv"
)

type Application struct {
	repository *repository.Repository
	usecase    *usecase.UseCase
	handler    *delivery.Handler
}

func New() (Application, error) {
	_ = godotenv.Load()
	repo, err := repository.New(dsn.FromEnv())
	uc := usecase.NewUseCase(*repo)
	h := delivery.NewHandler(*uc)
	if err != nil {
		return Application{}, err
	}

	return Application{
			repository: repo,
			usecase:    uc,
			handler:    h,
		},
		nil
}
