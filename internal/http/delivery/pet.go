package delivery

import (
	"net/http"

	"github.com/SicParv1sMagna/HappyPetsBackend/internal/model"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/pet/create
// @Summary Create a new pet
// @Description Create a new pet with the provided information
// @Tags Pet
// @Accept json
// @Produce json
// @Param pet body model.Pet true "Pet object in JSON format"
// @Success 201 {object} model.Pet "Successfully created pet"
// @Router /api/pet/create [post]
func (h *Handler) CreatePet(ctx *gin.Context) {
	var petJSON model.Pet
	if err := ctx.ShouldBindJSON(&petJSON); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pet, err := h.UseCase.CreatePet(petJSON)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"pet": pet})
}
