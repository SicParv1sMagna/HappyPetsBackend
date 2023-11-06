package middleware

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/golang-jwt/jwt"
)

func Authenticated(redisClient *redis.Client, jwtSecretKey []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(redisClient)
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Для продолжения необходимо авторизоваться"})
			c.Abort()
			return
		}

		// Попытка парсинга JWT-токена с использованием секретного ключа
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtSecretKey, nil
		})
		fmt.Println(token)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Недействительный токен"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		fmt.Println(claims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный формат токена"})
			c.Abort()
			return
		}

		// Получение идентификатора пользователя из токена
		subClaim, ok := claims["userID"].(float64)
		fmt.Println(subClaim)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный формат токена"})
			c.Abort()
			return
		}

		userID := int(subClaim)

		// Проверка, существует ли токен в Redis
		tokenValue, err := redisClient.Get(strconv.Itoa(userID)).Result()
		if err != nil || tokenValue != tokenString {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Токен не действителен"})
			c.Abort()
			return
		}

		c.Set("userID", userID) // Сохранение идентификатора пользователя в контексте Gin
	}
}
