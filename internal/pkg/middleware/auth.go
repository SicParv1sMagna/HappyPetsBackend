package middleware

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/golang-jwt/jwt"
)

func Authenticated(redisClient *redis.Client, jwtSecretKey []byte) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Для продолжения необходимо авторизоваться"})
			ctx.Abort()
			return
		}

		// Попытка парсинга JWT-токена с использованием секретного ключа
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtSecretKey, nil
		})
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Недействительный токен"})
			ctx.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный формат токена"})
			ctx.Abort()
			return
		}

		// Получение идентификатора пользователя из токена
		subClaim, ok := claims["userID"].(float64)
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный формат токена"})
			ctx.Abort()
			return
		}

		userID := int(subClaim)

		// Проверка, существует ли токен в Redis
		tokenValue, err := redisClient.Get(strconv.Itoa(userID)).Result()
		if err != nil || tokenValue != tokenString {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Токен не действителен"})
			ctx.Abort()
			return
		}

		ctx.Set("userID", userID) // Сохранение идентификатора пользователя в контексте Gin
		ctx.Next()
	}
}
