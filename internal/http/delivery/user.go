package delivery

import (
	"net/http"
	"strconv"

	"github.com/SicParv1sMagna/HappyPetsBackend/internal/model"
	"github.com/gin-gonic/gin"
)

// @BasePath /user/register
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

// @BasePath /user/login
// @Summary Вход пользователя
// @Description Авторизация пользователя и генерация JWT-токена
// @Tags Пользователь
// @Accept json
// @Produce json
// @Param body body model.UserLoginRequest true "Данные для входа"
// @Success 200 {object} map[string]string "Успешный ответ"
// @Failure 400 {object} map[string]string "Неверный запрос"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router /login [post]
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

// @BasePath /api/user
// @Summary Получить пользователя по идентификатору
// @Description Получение данных пользователя по его идентификатору
// @Produce json
// @Param userID path int true "Идентификатор пользователя"
// @Success 200 {object} map[string]interface{} "Успешный ответ"
// @Failure 400 {object} map[string]string "Неверный запрос"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router /users/{userID} [get]
func (h *Handler) GetUserById(ctx *gin.Context) {
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

// @BasePath /api/user
// @Summary Обновить данные пользователя
// @Description Обновление данных пользователя по его идентификатору
// @Accept json
// @Produce json
// @Param userID path int true "Идентификатор пользователя"
// @Param body body model.UserUpdateRequest true "Данные для обновления"
// @Success 200 {object} map[string]interface{} "Успешный ответ"
// @Failure 400 {object} map[string]string "Неверный запрос"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router /users/{userID} [put]
func (h *Handler) UpdateUserData(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("userID"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var userJSON model.UserUpdateRequest
	if err = ctx.ShouldBindJSON(&userJSON); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.UseCase.UpdateUserData(userJSON, uint(userID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}
