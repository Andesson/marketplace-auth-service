package handler

import (
	"net/http"

	"github.com/Andesson/marketplace-auth-service/hook"
	"github.com/Andesson/marketplace-auth-service/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateUserHandler(ctx *gin.Context) {
	logger.Infof("📩 Recebendo requisição para criar usuário...")

	request := CreateUserRequest{}
	if err := ctx.BindJSON(&request); err != nil {
		logger.Errorf("❌ Erro ao decodificar JSON: %v", err)
		sendError(ctx, http.StatusBadRequest, "Erro ao processar requisição")
		return
	}
	logger.Infof("✅ Requisição validada. Email: %s | Nome: %s | Senha: %s", request.Email, request.FullName, request.Password)

	if err := request.Validate(); err != nil {
		logger.Errorf("⚠️ Erro de validação: %v", err)
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	logger.Infof("🔐 Gerando hash da senha...")
	passwordHash, salt, err := HashPassword(request.Password)
	if err != nil {
		logger.Errorf("❌ Erro ao gerar hash da senha: %v", err)
		sendError(ctx, http.StatusInternalServerError, "Erro ao processar senha")
		return
	}
	logger.Infof("✅ Hash gerado com sucesso.")

	user := model.User{
		Email:    request.Email,
		FullName: request.FullName,
	}

	db.Callback().Create().Before("gorm:create").Register("before_create", func(db *gorm.DB) {
		if err := hook.BeforeCreate(db); err != nil {
			logger.Errorf("❌ Erro no hook BeforeCreate: %v", err)
		}
	})

	logger.Infof("🛠️ Criando usuário no banco de dados...")
	if err := db.Create(&user).Error; err != nil {
		logger.Errorf("❌ Erro ao criar usuário: %v", err)
		sendError(ctx, http.StatusInternalServerError, "Erro ao criar usuário no banco")
		return
	}
	logger.Infof("✅ Usuário criado com sucesso. ID: %v", user.ID)

	authCredential := model.AuthCredential{
		UserID:        user.ID,
		PassHash:      passwordHash,
		HashAlgorithm: "bcrypt",
		Salt:          salt,
	}
	logger.Infof("🔑 Criando credenciais de autenticação para o usuário...")
	if err := db.Create(&authCredential).Error; err != nil {
		logger.Errorf("❌ Erro ao criar credenciais: %v", err)
		sendError(ctx, http.StatusInternalServerError, "Erro ao criar credenciais no banco")
		return
	}
	logger.Infof("✅ Credenciais criadas com sucesso.")

	sendCreatedSucess(ctx, "create-user", user)
	logger.Infof("🎉 Usuário criado com sucesso e retornado para o cliente.")
}

func HashPassword(password string) (string, string, error) {
	salt := "random-salt-value"
	saltedPassword := password + salt

	hash, err := bcrypt.GenerateFromPassword([]byte(saltedPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", "", err
	}

	return string(hash), salt, nil
}
