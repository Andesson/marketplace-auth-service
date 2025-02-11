package handler

import (
	"net/http"

	"github.com/Andesson/marketplace-auth-service/model"
	"github.com/Andesson/marketplace-auth-service/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Logon(ctx *gin.Context) {
	logger.Infof("📩 Recebendo requisição de login...")

	var request CreateLoginRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		logger.Errorf("❌ Erro ao decodificar JSON: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	var user model.User
	if err := db.Where("email = ?", request.Email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			logger.Warningf("⚠️ Usuário não encontrado: %s", request.Email)
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciais inválidas"})
			return
		}
		logger.Errorf("❌ Erro ao buscar usuário: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro interno"})
		return
	}

	var authCredential model.AuthCredential
	if err := db.Where("user_id = ?", user.ID).First(&authCredential).Error; err != nil {
		logger.Errorf("❌ Erro ao buscar usuário: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro interno"})
		return
	}
	if !utils.ComparePasswordHash(request.Password, authCredential.PassHash, authCredential.Salt) {
		logger.Warningf("⚠️ Senha incorreta para o usuário: %s", request.Email)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciais inválidas"})
		return
	}

	// Gerar Token JWT
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		logger.Errorf("❌ Erro ao gerar token JWT: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar token"})
		return
	}
	logger.Infof("📩 Token gerado")
	CreateSession(user.ID, token)
	ctx.JSON(http.StatusOK, LoginResponse{Token: token})
}
