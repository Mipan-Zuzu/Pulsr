package handler

import (
	"fmt"
	"os"
	// "os"
	"pulsr/internal/struct"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"gorm.io/gorm"
)

func HandlerSupabaseGetAllProject(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authtoken := ctx.GetHeader("Authorization")
		client := resty.New()
		var result []Struct.Project
		res, err := client.R().
		SetAuthToken(authtoken).
		SetResult(&result).
		Get(fmt.Sprintf("%sprojects", os.Getenv("SUPABASE_URL")))

		if err != nil {
			ctx.JSON(400, gin.H{
				"message" : err,
			})
			return
		}

		if result == nil {
			ctx.JSON(401, gin.H{
				"message" : "invalid authorization token",
			})
			return
		}

		fmt.Println(res)
		ctx.JSON(200, gin.H{
			"status": 200,
			"data" : result,
		})
	}
}

func HandlerSupabaseGetAllProjectId (db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authtoken := ctx.GetHeader("Authorization")
		id := ctx.Param("id")
		client := resty.New()
		var result Struct.Project
		_, err := client.R().
		SetAuthToken(authtoken).
		SetResult(&result).
		Get(fmt.Sprintf("%sproject/%s", os.Getenv("SUPABASE_URL"),id))

		if id == "" {
			ctx.JSON(404, gin.H{
				"message" : "cannot find project with this id",
			})
		}

		if err != nil {
			ctx.JSON(401, gin.H{
				"message" : err,
			})
			return
		}

		// if result == nil {
		// 	ctx.JSON(401, gin.H{
		// 		"message" : "invalid Authorization token",
		// 	})
		// 	return 
		// }

		ctx.JSON(200, gin.H{
			"data" : result,
			"status" : 200,
		})
	}
}