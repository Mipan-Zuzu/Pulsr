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

func generateKeyRandom(length int) (string, error) {
	key := make([]byte, length)
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
		token, err := ctx.Cookie("auth_key_supabase")
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

		keyrand, errRand := generateKeyRandom(10)
		var authkey Authkey

		if errRand != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message" : "cannot create random key",
			})
		}

		if err := ctx.ShouldBindJSON(&authkey.Key); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message" : err,
			})
		}

		err := rdb.Set(ctx , fmt.Sprintf("token_%s", keyrand), authkey, 5*time.Minute).Err()
		if err != nil {
			ctx.JSON(402, gin.H{"message" : "failed to save redis key"})
			return
		}

		ctx.SetSameSite(http.SameSiteLaxMode)
		ctx.SetCookie(
			"auth_key_supabase",
			authkey.Key,
			300,
			"/",
			"",
			false,
			true,
		)

		ctx.JSON(200, gin.H{"message" : "succsesfull to save redis key"})
	}
}