package handler

import (
	"crypto/rand"
	"encoding/hex"
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

func SetValueRedis ( rdb *redis.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := rdb.Set(ctx , "suki", "value", 5*time.Minute).Err()
		if err != nil {
			ctx.JSON(402, gin.H{"message" : "failed to save redis key"})
			return
		}

		ctx.JSON(200, gin.H{"message" : "succsesfull to save redis key"})
	}
}