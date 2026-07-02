package routing

import (
	"fmt"

	"pulsr/internal/handler"
	"pulsr/internal/midlewere"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Ping() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("PONG")
		status := "up"
		ctx.JSON(200, gin.H{
			"status": status,
		})
	}
}

func SetupRouting(db *gorm.DB, route *gin.Engine) {
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
		route.GET(fmt.Sprintf("%sprojects", supabasePath), midlewere.CheckingAuthorization(), handler.HandlerSupabaseGetAllProject(db))

		route.GET(fmt.Sprintf("%sprojects/:id", supabasePath))
		route.GET(fmt.Sprintf("%sanalytics/usage/:id", supabasePath))
		route.GET(fmt.Sprintf("%sanalytics/logs/:id", supabasePath)) //NEED query  ( iso_timestamp_start, iso_timestamp_end )
		route.GET(fmt.Sprintf("%sedge/status", supabasePath))

		route.POST(fmt.Sprintf("%sproject/start/:id", supabasePath))
		route.POST(fmt.Sprintf("%sproject/pause/:id", supabasePath))
	// Koyeb
}
