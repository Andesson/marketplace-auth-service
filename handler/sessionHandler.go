package handler

import (
	"net/http"
	"time"

	"github.com/Andesson/marketplace-auth-service/config"
	"github.com/Andesson/marketplace-auth-service/model"
	"github.com/Andesson/marketplace-auth-service/schemas"
	"github.com/Andesson/marketplace-auth-service/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateSession(userID uuid.UUID, tokenString string) {
	session := model.Session{
		UserID:    userID,
		Token:     tokenString,
		ExpiresAt: time.Now().Add(time.Hour * 24),
	}
	db.Callback().Create().Before("gorm:create").Register("before_create", func(db *gorm.DB) {
		if err := utils.BeforeCreate(db); err != nil {
			logger.Errorf("❌ Erro no hook BeforeCreate: %v", err)
		}
	})
	if err := db.Create(&session).Error; err != nil {
		logger.Errorf("❌ Erro ao criar credenciais: %v", err)
		sendError(nil, http.StatusInternalServerError, "Erro ao criar credenciais no banco")
		return
	}
	logger.Infof("🎉 Session Created")
}

func GetSession(ctx *gin.Context, tokenString string) {
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	userID, ok := claims["sub"].(string)
	if !ok || userID == "" {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		ctx.Abort()
		return
	}

	var session schemas.Sessions
	err := config.GetPostgres().Where("user_id = ? AND token = ? AND expires_at > ?", userID, tokenString, time.Now()).First(&session).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired session"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error verifying session"})
		}
		ctx.Abort()
		return
	}
}
