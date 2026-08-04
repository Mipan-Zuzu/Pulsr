package handler

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func generateKeyRandom() (string, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(key), nil
}

type Authkey struct {
	Key string `json:"key" binding:"required"`
}

func GetValueRedis ( rdb *redis.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := ctx.Cookie("auth_key")
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message" : "cookie not found",
			})
			return 
		}

		token, errdb := rdb.Get(ctx, token).Result()

		if errdb != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message" : "key not found",
			})
			return 
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message" : token,
		})
	}
}

func SetValueRedis ( rdb *redis.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var authkey Authkey

		if err := ctx.ShouldBindJSON(&authkey); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message" : err,
			})
		}

		err := rdb.Set(ctx , "auth_key", authkey, 5*time.Minute).Err()
		if err != nil {
			ctx.JSON(402, gin.H{"message" : "failed to save redis key"})
			return
		}

		ctx.JSON(200, gin.H{"message" : "succsesfull to save redis key"})
	}
}