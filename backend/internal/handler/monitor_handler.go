package handler

import (
	"fmt"
	"net/http"
	"os"

	"time"

	"pulsr/internal/database"
	"pulsr/internal/struct"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)



func HandlerSupabaseGetAllProject(db *gorm.DB, rdb *redis.Client) gin.HandlerFunc {
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

		errRedisSet, keyrand := database.SetValueRedis(ctx ,rdb, authtoken);

		if  errRedisSet != nil {
			ctx.JSON(http.StatusBadRequest, gin.H {
				"message" : errRedisSet,
			})
		}

		ctx.SetSameSite(http.SameSiteLaxMode)
		ctx.SetCookie(
			"supabase_key",
			keyrand,
			300,
			"/",
			"",
			false,
			true,
		)

		fmt.Println(res)
		ctx.JSON(200, gin.H{
			"status": 200,
			"data":   result,
		})
	}
}

func HandlerOrg() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authtoken := ctx.GetHeader("Authorization")

		client := resty.New()
		var org []Struct.Org
		res, err := client.R().
			SetAuthToken(authtoken). 
			SetResult(&org). 
			Get("https://api.supabase.com/v1/organizations")

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message" : err,
			})
			return
		}

		if res.StatusCode() == http.StatusUnauthorized {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message" : "invalid authorize token",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data" : org,
			"status" : 200,
		})
	}
}



func HandlerOrgDetail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authtoken := ctx.GetHeader("Authorization")
		id := ctx.Param("id")

		client := resty.New()
		var result Struct.OrgDetail
		res, err := client.R().
			SetAuthToken(authtoken).
			SetResult(&result). 
			Get(fmt.Sprintf("https://api.supabase.com/v1/organizations/%s", id))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message" : err,
			})
			return 
		}

		if res.StatusCode() == http.StatusUnauthorized {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message" : "request not auhtorized",
			})
			return 
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data" : result,
			"status" : 200,
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
		fmt.Println(result)

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

func HandlerMetrics (rdb *redis.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authtoken := ctx.GetHeader("Authorization")
		id := ctx.Param("id")
		if id == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message" : "bad request missing data id",
			})
			return
		}

		client := resty.New()

		res, err := client.R().
		SetAuthToken(authtoken).
		Get(fmt.Sprintf("https://api.supabase.com/v1/projects/%s/analytics/endpoints/metrics", id))

		if err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{
				"message" : "upstream request failed",
				"error" : err.Error(),
			})
		}

		if res.IsError() {
			ctx.JSON(http.StatusBadGateway, gin.H{
				"message" : "upstream returned an error",
				"status" : res.StatusCode(),
				"data" : res.String(),
			})
		}

		ctx.Data(http.StatusOK, "text/plain; version=0.0.4", res.Body())
	}
}