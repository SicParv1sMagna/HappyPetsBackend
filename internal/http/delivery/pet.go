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
// @Security ApiKeyAuth
// @Param pet body model.PetCreateRequest true "Объект PetCreateRequest в формате JSON"
// @Success 201 {object} model.Pet "Питомец успешно создан"
// @Router /api/pet/create [post]
func (h *Handler) CreatePet(ctx *gin.Context) {
	ctxUserID, exists := ctx.Get("userID")
	if !exists {
		// Если значение отсутствует в контексте, обработайте эту ситуацию
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Идентификатор пользователя отсутствует в контексте"})
		return
	}

	userID, ok := ctxUserID.(int)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при преобразовании идентификатора пользователя"})
		return
	}

	var petJSON model.PetCreateRequest
	if err := ctx.ShouldBindJSON(&petJSON); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ошибка при чтении JSON файла"})
		return
	}

	pet, err := h.UseCase.CreatePet(uint64(userID), petJSON)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"pet": pet, "message": "Питомец успешно создан"})
}

// @Summary Обновление информации о питомце.
// @Description Обновляет информацию о питомце с предоставленными данными.
// @Tags Питомец
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path integer true "ID питомца"
// @Param pet body model.PetUpdateRequest true "Объект PetUpdateRequest в формате JSON"
// @Success 200 {object} model.Pet "Питомец успешно обновлен"
// @Failure 400 {object} model.Pet "Неверный запрос"
// @Router /api/pet/update/{petID} [put]
func (h *Handler) UpdatePet(ctx *gin.Context) {
	ctxUserID, exists := ctx.Get("userID")
	if !exists {
		// Если значение отсутствует в контексте, обработайте эту ситуацию
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Идентификатор пользователя отсутствует в контексте"})
		return
	}

	userID, ok := ctxUserID.(int)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при преобразовании идентификатора пользователя"})
		return
	}

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

	pet, err := h.UseCase.UpdatePet(uint64(userID), uint64(petID), petJSON)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"pet": pet, "message": "Питомец успешно обновлен"})
}

// @Summary Удаление питомца.
// @Description Удаляет питомца по заданному идентификатору.
// @Tags Питомец
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param petID path integer true "ID питомца"
// @Success 204 "Питомец успешно удален"
// @Failure 400 "Неверный запрос"
// @Failure 404 "Питомец не найден"
// @Router /api/pet/delete/{petID} [delete]
func (h *Handler) DeletePet(ctx *gin.Context) {
	ctxUserID, exists := ctx.Get("userID")
	if !exists {
		// Если значение отсутствует в контексте, обработайте эту ситуацию
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Идентификатор пользователя отсутствует в контексте"})
		return
	}

	userID, ok := ctxUserID.(int)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при преобразовании идентификатора пользователя"})
		return
	}

	petID, err := strconv.ParseUint(ctx.Param("petID"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "неверный ID питомца"})
		return
	}

	err = h.UseCase.DeletePet(uint64(userID),uint64(petID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Питомец успешно удален"})
}

// @Summary Получение списка всех питомцев пользователя.
// @Description Возвращает список всех питомцев пользователя.
// @Tags Питомец
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} []model.Pet "Список питомцев пользователя"
// @Failure 500 {object} []model.Pet "Внутренняя ошибка сервера"
// @Router /api/pet/ [get]
func (h *Handler) GetAllPets(ctx *gin.Context) {
	ctxUserID, exists := ctx.Get("userID")
	if !exists {
		// Если значение отсутствует в контексте, обработайте эту ситуацию
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Идентификатор пользователя отсутствует в контексте"})
		return
	}

	userID, ok := ctxUserID.(int)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при преобразовании идентификатора пользователя"})
		return
	}

	pets, err := h.UseCase.GetAllPets(uint64(userID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"pets": pets})
}

// @Summary Получение информации о питомце по его ID.
// @Description Возвращает информацию о питомце по его ID.
// @Tags Питомец
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param petID path integer true "ID питомца"
// @Success 200 {object} model.Pet "Информация о питомце"
// @Failure 400 {object} model.Pet "Неверный запрос"
// @Failure 404 {object} model.Pet "Питомец не найден"
// @Failure 500 {object} model.Pet "Внутренняя ошибка сервера"
// @Router /api/pet/{petID} [get]
func (h *Handler) GetPetById(ctx *gin.Context) {
	ctxUserID, exists := ctx.Get("userID")
	if !exists {
		// Если значение отсутствует в контексте, обработайте эту ситуацию
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Идентификатор пользователя отсутствует в контексте"})
		return
	}

	userID, ok := ctxUserID.(int)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при преобразовании идентификатора пользователя"})
		return
	}

	petID, err := strconv.ParseUint(ctx.Param("petID"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "неверный ID питомца"})
		return
	}

	pet, err := h.UseCase.GetPetById(uint64(userID), uint64(petID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"pet": pet})
}

