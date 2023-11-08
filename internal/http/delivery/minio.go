package delivery

import (
	"net/http"
	"strconv"

	"github.com/SicParv1sMagna/HappyPetsBackend/internal/model"
	"github.com/gin-gonic/gin"
)

// @Summary Загружает изображение в Minio happypets-image(bucket) определенного домашнего животного, связанного с пользователем.
// @Description Загружает изображение в Minio happypets-image(bucket) определенного домашнего животного, связанного с пользователем.
// @Tags Питомец
// @Accept mpfd
// @Produce json
// @Security ApiKeyAuth
// @Param petID path string true "Pet ID"
// @Param image formData file true "Image file"
// @Success 200 {object} model.UploadImageResponse "Изображение успешно загружено"
// @Failure 400 {object} model.UploadImageResponse "Неверный запрос, неверные входные данные"
// @Failure 500 {object} model.UploadImageResponse "Внутренняя ошибка сервера"
// @Router /api/pet/image/upload/{petID} [post]
func (h *Handler) UploadImage(ctx *gin.Context) {
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
	petID, err := strconv.Atoi(ctx.Param("petID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "неверный ИД питомца"})
		return
	}
	image, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "недопустимые данные изображения"})
		return
	}
	// Call the use case method to upload the image
	imageURL, err := h.UseCase.UploadImage(uint64(userID), uint64(petID), image)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, model.UploadImageResponse{ImageURL: imageURL, Message: "Изображение успешно загружено"})
}

// @Summary Удаляет изображение из Minio happypets-image(bucket) определенного домашнего животного, связанного с пользователем.
// @Description Удаляет изображение из Minio happypets-image(bucket) определенного домашнего животного, связанного с пользователем.
// @Tags Питомец
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param petID path string true "Pet ID"
// @Success 204 {object} model.RemoveImageResponse "Изображение успешно удалено"
// @Failure 400 {object} model.RemoveImageResponse "Неверный запрос, неверные входные данные"
// @Failure 404 {object} model.RemoveImageResponse "Изображение не найдено"
// @Failure 500 {object} model.RemoveImageResponse "Внутренняя ошибка сервера"
// @Router /api/pet/image/remove/{petID} [delete]
func (h *Handler) RemoveImage(ctx *gin.Context) {
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

	petID,err := strconv.Atoi(ctx.Param("petID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "неверный ИД питомца"})
		return
	}

	// Call the use case method to remove the image
	err = h.UseCase.RemoveImage(uint64(userID), uint64(petID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, model.RemoveImageResponse{Message: "Изображение успешно удалено"})
}
