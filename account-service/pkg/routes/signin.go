package routes

import (
	"account-service/pkg/encryption"
	"account-service/pkg/httperrors"
	"account-service/pkg/models"
	"account-service/pkg/repositories"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type signInReq struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type signinRes struct {
	StatusCode int          `json:"statusCode"`
	User       *models.User `json:"user"`
}

func (routes *Routes) Signin(ctx *gin.Context) {
	req := &signInReq{}

	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, httperrors.HttpBindingJSONError(err, http.StatusBadRequest))
		return
	}

	userRepository := repositories.NewUserRepo(routes.Db)

	c, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	user, err := userRepository.FindByUserNameAndPassword(c, req.UserName, string(encryption.Encrypt([]byte(req.Password), routes.cfg.EnvVars.EncryptsionKey)))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError})
		return
	}

	ctx.JSON(http.StatusOK, signinRes{StatusCode: http.StatusOK, User: user})
}
