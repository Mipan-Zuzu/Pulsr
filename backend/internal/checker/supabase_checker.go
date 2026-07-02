package checker

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SupabaseProjectAll(authToken string, db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp, err := http.Get(os.Getenv("SUPABASE_URL"))
		if err != nil {
		
	}
	}
}