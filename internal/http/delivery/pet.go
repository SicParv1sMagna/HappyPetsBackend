package delivery

import (
	"net/http"
	"strconv"

	"github.com/SicParv1sMagna/HappyPetsBackend/internal/model"
	"github.com/gin-gonic/gin"
)

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

	ctx.JSON(http.StatusCreated, gin.H{"pet": pet, "message": "Питомец успешно создан"})
}

// @Summary Обновление информации о питомце.
// @Description Обновляет информацию о питомце с предоставленными данными.
// @Tags Питомец
// @Accept json
// @Produce json
// @Param id path integer true "ID питомца"
// @Param pet body model.PetUpdateRequest true "Объект PetUpdateRequest в формате JSON"
// @Success 200 {object} model.PetUpdateRequest "Питомец успешно обновлен"
// @Failure 400 {object} model.PetUpdateRequest "Неверный запрос"
// @Router /api/pet/update/{id} [put]
func (h *Handler) UpdatePet(ctx *gin.Context) {
	petID, err := strconv.ParseUint(ctx.Param("petID"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "неверный ID питомца"})
		return
	}

	var petJSON model.PetUpdateRequest
	if err := ctx.ShouldBindJSON(&petJSON); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ошибка при чтении JSON файла"})
		return
	}

	petJSON.ID = petID
	pet, err := h.UseCase.UpdatePet(petJSON)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ошибка при обновлении питомца"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"pet": pet, "message": "Питомец успешно обновлен"})
}

// @Summary Удаление питомца.
// @Description Удаляет питомца по заданному идентификатору.
// @Tags Питомец
// @Accept json
// @Produce json
// @Param petID path integer true "ID питомца"
// @Success 204 "Питомец успешно удален"
// @Failure 400 "Неверный запрос"
// @Failure 404 "Питомец не найден"
// @Router /api/pet/delete/{petID} [delete]
func (h *Handler) DeletePet(ctx *gin.Context) {
	petID, err := strconv.ParseUint(ctx.Param("petID"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "неверный ID питомца"})
		return
	}

	err = h.UseCase.DeletePet(uint(petID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "питомец не найден"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Питомец успешно удален"})
}
