package handler

import (
	"fmt"
	"net/http"
	"os"

	"time"

	// "os"
	"pulsr/internal/struct"
	// "reflect"
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
				"message": err,
			})
			return
		}

		if result == nil {
			ctx.JSON(401, gin.H{
				"message": "invalid authorization token",
			})
			return
		}

		fmt.Println(res)
		ctx.JSON(200, gin.H{
			"status": 200,
			"data":   result,
		})
	}
}

func HandlerSupabaseGetAllProjectId(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authtoken := ctx.GetHeader("Authorization")
		id := ctx.Param("id")
		client := resty.New()
		var result Struct.Project
		res, err := client.R().
			SetAuthToken(authtoken).
			SetResult(&result).
			Get(fmt.Sprintf("%sprojects/%s", os.Getenv("SUPABASE_URL"), id))

		if id == "" {
			ctx.JSON(404, gin.H{
				"message": "cannot find spesific project",
			})
			return
		}

		if err != nil {
			ctx.JSON(401, gin.H{
				"message": err,
			})
			return
		}

		if res.StatusCode() == http.StatusUnauthorized {
			ctx.JSON(401, gin.H{
				"message": "invalid Authorization token",
			})
			return
		}

		ctx.JSON(200, gin.H{
			"data":   result,
			"status": 200,
		})
	}
}

func HandlerSupabaseAnalytics() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authtoken := ctx.GetHeader("Authorization")
		id := ctx.Param("id")
		client := resty.New()
		var result []Struct.Analytics
		res, err := client.R().
			SetAuthToken(authtoken).
			SetResult(&result).
			Get(fmt.Sprintf("%sprojects/%s/analytics/endpoints/usage.api-counts", os.Getenv("SUPABASE_URL"), id))
		fmt.Println(result, "dasdasdasdsad")

		fmt.Println(res)
		if id == "" {
			ctx.JSON(404, gin.H{
				"message": "cannot find spesific project",
			})
		}
		if err != nil {
			ctx.JSON(401, gin.H{
				"message": err,
			})
		}

		ctx.JSON(200, gin.H{
			"message": result,
		})
	}
}

func HandlerSupabaseAnalyticsLogs() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		timeNow := time.Now().UTC()
		isoTimestampEnd := timeNow.Format(time.RFC3339)
		isoTimestampStart := timeNow.Add(-1 * time.Hour).Format(time.RFC3339)
		authtoken := ctx.GetHeader("Authorization")
		id := ctx.Param("id")
		client := resty.New()
		var result Struct.Logs
		res, err := client.R().
			SetAuthToken(authtoken).
			SetResult(&result).
			Get(fmt.Sprintf("https://api.supabase.com/v1/projects/%s/analytics/endpoints/logs.all?iso_timestamp_start=%s&iso_timestamp_end=%s", id, isoTimestampStart, isoTimestampEnd))

		
		fmt.Println(res)

		if id == "" {
			ctx.JSON(404, gin.H{
				"message": "cannot find spesific id",
			})
			return
		}

		if err != nil {
			ctx.JSON(401, gin.H{
				"message": err,
			})
			return
		}

		if res.StatusCode() == http.StatusUnauthorized {
			ctx.JSON(401, gin.H{
				"message": "invalid authorize token",
			})
			return
		}
		fmt.Println(isoTimestampStart)
		ctx.JSON(200, gin.H{
			"result": result,
			"err":    err,
		})
	}
}

func HandlerSupabasePauseProject () gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authtoken := ctx.GetHeader("Authorization")
		id := ctx.Param("id")
		var Message Struct.Pause
		client := resty.New()

		res,err := client.R(). 
		SetAuthToken(authtoken).
		SetResult(&Message).
		SetError(&Message). 
		Post(fmt.Sprintf("%sprojects/%s/pause", os.Getenv("SUPABASE_URL"), id))

		if id == "" {
			ctx.JSON(404, gin.H{
				"message" : "cannot find spesific project",
			})
			return 
		}

		if err != nil {
			ctx.JSON(401, gin.H{
				"message" : err,
			})
			return
		}

		if res.StatusCode() == http.StatusUnauthorized {
			ctx.JSON(401, gin.H{
				"message" : Message.Message,
			})
			return 
		}

		if res.StatusCode() == http.StatusBadGateway {
			ctx.JSON(400, gin.H{
				"message" : Message.Message,
			})
			return
		}

		ctx.JSON(200, gin.H{
			"message" : "succsesfull PAUSE service",
			"id" : id,
		})
	}
}

func HandlerSupabaseStartProject () gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authtoken := ctx.GetHeader("Authorization")
		id := ctx.Param("id")
		var Result Struct.Start
		client := resty.New()
		if id == "" {
			ctx.JSON(404, gin.H{
				"message" : "cannot find spesific project",
			})
			return
		}
		res,err := client.R(). 
		SetAuthToken(authtoken). 
		SetResult(&Result). 
		SetError(&Result).
		Post(fmt.Sprintf("%sprojects/%s/restore",os.Getenv("SUPABASE_URL"), id ))

		if err != nil {
			ctx.JSON(401, gin.H{
				"message" : err,
			})
			return 
		}
		
		if res.StatusCode() == http.StatusUnauthorized {
			ctx.JSON(401, gin.H{
				"message" : "invalid authorization",
			})
			return 
		}

		if res.StatusCode() == http.StatusBadRequest {
			ctx.JSON(400, gin.H{
				"message" : Result.Message,
			})
			return 
		}

		fmt.Println(res)
		ctx.JSON(200, gin.H{
			"message" : "succsesfull START the service",
		})
	}
}