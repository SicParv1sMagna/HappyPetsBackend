package delivery

import (
	"net/http"
	"strconv"

	"github.com/SicParv1sMagna/HappyPetsBackend/internal/model"
	"github.com/gin-gonic/gin"
)

// @Summary Uploads an image for a specific pet associated with a user.
// @Description Uploads an image for a specific pet associated with a user.
// @Tags Image
// @Accept mpfd
// @Produce json
// @Param userID path string true "User ID"
// @Param petID path string true "Pet ID"
// @Param image formData file true "Image file"
// @Success 200 {object} model.UploadImageResponse "Image uploaded successfully"
// @Failure 400 {object} model.UploadImageResponse "Bad request, invalid input data"
// @Failure 500 {object} model.UploadImageResponse "Internal server error"
// @Router /api/image/upload/{userID}/{petID} [post]
func (h *Handler) UploadImage(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("userID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	petID, err := strconv.Atoi(ctx.Param("petID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pet ID"})
		return
	}
	image, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image data"})
		return
	}
	// Call the use case method to upload the image
	imageURL, err := h.UseCase.UploadImage(uint64(userID), uint64(petID), image)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image"})
		return
	}

	ctx.JSON(http.StatusOK, model.UploadImageResponse{ImageURL: imageURL, Message: "Image uploaded successfully"})
}

// @Summary Removes an image for a specific pet associated with a user.
// @Description Removes an image for a specific pet associated with a user.
// @Tags Image
// @Accept json
// @Produce json
// @Param userID path string true "User ID"
// @Param petID path string true "Pet ID"
// @Success 204 {object} model.RemoveImageResponse "Image removed successfully"
// @Failure 400 {object} model.RemoveImageResponse "Bad request, invalid input data"
// @Failure 404 {object} model.RemoveImageResponse "Image not found"
// @Failure 500 {object} model.RemoveImageResponse "Internal server error"
// @Router /api/image/remove/{userID}/{petID} [delete]
func (h *Handler) RemoveImage(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("userID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	petID,err := strconv.Atoi(ctx.Param("petID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pet ID"})
		return
	}

	// Call the use case method to remove the image
	err = h.UseCase.RemoveImage(uint64(userID), uint64(petID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
		return
	}

	ctx.JSON(http.StatusNoContent, model.RemoveImageResponse{})
}
