package routing

import (
	"fmt"

	"pulsr/internal/handler"
	"pulsr/internal/midlewere"
	"pulsr/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func Ping() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("PONG")
		status := "up"
		ctx.JSON(200, gin.H{
			"status": status + "dasdasdasdsad",
		})
	}
}

func SetupRouting(db *gorm.DB, route *gin.Engine, rdb *redis.Client) {
	supabasePath := "v1/supabase/"
	// vercelPath := "v1/vercel/"
	route.GET("/ping", Ping())

	//Vercel

	// MongoDb

	// Redis

	// Domain

	//Netlfy

	//CloudFlare

	// Supabase (belum di isi function)
		route.GET("/v1/supabase/projects",midlewere.CheckingAuthorization(), handler.HandlerSupabaseGetAllProject(db, rdb))

		route.GET("/v1/supabase/org", midlewere.CheckingAuthorization(), handler.HandlerOrg())

		route.GET("/v1/supabase/org/:id", midlewere.CheckingAuthorization(), handler.HandlerOrgDetail())

		route.GET(fmt.Sprintf("%sprojects/:id", supabasePath), midlewere.CheckingAuthorization(), handler.HandlerSupabaseGetAllProjectId(db))
		
		route.GET(fmt.Sprintf("%sanalytics/usage/:id", supabasePath),
		midlewere.CheckingAuthorization(), handler.HandlerSupabaseAnalytics())

		route.GET(fmt.Sprintf("%sanalytics/logs/:id", supabasePath), midlewere.CheckingAuthorization(), 
		handler.HandlerSupabaseAnalyticsLogs()) //NEED query  ( iso_timestamp_start, iso_timestamp_end )

		// Belum perlu
		route.GET(fmt.Sprintf("%sedge/status", supabasePath))

		route.POST(fmt.Sprintf("%sproject/pause/:id", supabasePath), midlewere.CheckingAuthorization(), handler.HandlerSupabasePauseProject())

		route.POST(fmt.Sprintf("%sproject/start/:id", supabasePath), midlewere.CheckingAuthorization(), handler.HandlerSupabaseStartProject())
	// Koyeb

		route.GET("v1/redis/dat", database.GetValueRedis(rdb))
}