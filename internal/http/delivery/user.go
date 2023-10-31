package delivery

import (
	"net/http"

	"github.com/SicParv1sMagna/HappyPetsBackend/internal/model"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/user/register
// @Summary Register a new user
// @Description Register a new user with the provided information
// @Accept json
// @Produce json
// @Param user body model.User true "User object in JSON format"
// @Success 201 {object} model.User "Successfully registered user"
// @Router /api/user/register [post]
func (h *Handler) Register(ctx *gin.Context) {
	var userJSON model.User
	if err := ctx.ShouldBindJSON(&userJSON); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.UseCase.RegisterUser(userJSON)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"user": user})
}
