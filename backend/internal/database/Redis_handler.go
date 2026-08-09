package database

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
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

func GetValueRedis(rdb *redis.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, _ := ctx.Cookie("supabase_key")
		if token == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message" : "invalid cookie",
			})
		}
		dat, err := rdb.Get(ctx, token).Result()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message" : "invalid get redis",
			})
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message": dat,
		})
	}
}

func SetValueRedis(ctx context.Context ,rdb *redis.Client, key string) (error, string)  {

		var err error

		if key == "" {
			err = errors.New("Key was empty")
			return err , ""
		}

		keyrand, errorRand := generateKeyRandom(10)

		if errorRand != nil {
			return errorRand, ""
		}


		if rdbErr := rdb.Set(ctx, fmt.Sprintf("token_%s", keyrand), key, 5*time.Minute).Err(); rdbErr != nil {
			return rdbErr, ""
		}


		return nil , fmt.Sprintf("token_%s", keyrand)

	}
