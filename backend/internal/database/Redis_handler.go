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
		token, err := ctx.Cookie("auth_key_supabase")
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "cookie not found",
			})
			return
		}

		token, errdb := rdb.Get(ctx, token).Result()

		if errdb != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "key not found",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": token,
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
