package delivery

import (
	"net/http"
	"strconv"

	"github.com/SicParv1sMagna/HappyPetsBackend/internal/model"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/user/register
// @Summary Регистрация нового пользователя.
// @Description Регистрация нового пользователя с предоставленной информацией.
// @Tags Пользователь
// @Accept json
// @Produce json
// @Param user body model.User true "Пользовательский объект в формате JSON"
// @Success 201 {object} model.User "Успешно зарегистрированный пользователь"
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

func (h *Handler) Login(ctx *gin.Context) {
	var userJSON model.UserLoginRequest
	if err := ctx.ShouldBindJSON(&userJSON); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.UseCase.LoginUser(userJSON)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *Handler) GetUserById(ctx *gin.Context) {
	var userJSON model.User
	if err := ctx.ShouldBindJSON(&userJSON); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := strconv.Atoi(ctx.Param("userID"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user, err := h.UseCase.GetUserById(uint(userID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func (h *Handler) UpdateUserData(ctx *gin.Context) {
	var userJSON model.UserUpdateRequest
	if err := ctx.ShouldBindJSON(&userJSON); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

}
