package routes

import (
	"account-service/pkg/encryption"
	"account-service/pkg/httperrors"
	"account-service/pkg/models"
	"account-service/pkg/repositories"

	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type signupReq struct {
	Email    string `json:"email" binding:"required,email"`
	FullName string `json:"fullName" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
}

type signupRes struct {
	StatusCode int    `json:"statusCode"`
	Id         string `json:"id"`
}

func (routes *Routes) SignUp(ctx *gin.Context) {
	req := &signupReq{}

	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, httperrors.HttpBindingJSONError(err, http.StatusBadRequest))
		return
	}

	pass := encryption.Encrypt(
		[]byte(req.Password),
		routes.cfg.EnvVars.EncryptsionKey,
	)

	usrModel := models.NewUser(req.Email, req.FullName, req.Username, string(pass))
	userRepository := repositories.NewUserRepo(routes.Db)
	c, cancel := context.WithTimeout(context.TODO(), time.Second*10)
	defer cancel()
	id, err := userRepository.Create(c, usrModel)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusCreated, signupRes{StatusCode: http.StatusCreated, Id: id})
}
