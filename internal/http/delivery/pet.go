package delivery

import (
	"net/http"

	"github.com/SicParv1sMagna/HappyPetsBackend/internal/model"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/pet/create
// @Summary Создание нового питомца.
// @Description Создайте нового питомца с предоставленной информацией.
// @Tags Питомец
// @Accept json
// @Produce json
// @Param pet body model.Pet true "Объект Pet в формате JSON"
// @Success 201 {object} model.Pet "Питомец успешно создан"
// @Router /api/pet/create [post]
func (h *Handler) CreatePet(ctx *gin.Context) {
	var petJSON model.Pet
	if err := ctx.ShouldBindJSON(&petJSON); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ошибка при чтении JSON файла"})
		return
	}

	pet, err := h.UseCase.CreatePet(petJSON)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "ошибка при создании нового питомца"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"pet": pet, "message":"Питомец успешно создан"})
}
