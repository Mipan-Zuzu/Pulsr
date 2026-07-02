package handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"pulsr/internal/checker"
)

func HandlerSupabaseGetAllProject(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authToken := ctx.GetHeader("Authorization")
		
		checker.SupabaseProjectAll(authToken, db)
	}
}