package handler

import (
	"errors"
	"net/http"

	"github.com/datti-api/pkg/domain/model"
	"github.com/datti-api/pkg/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler interface {
	HandlerCreate(c *gin.Context)
	HandlerGet(c *gin.Context)
	HandlerUpdate(c *gin.Context)
}

type userHandler struct {
	useCase usecase.UserUseCase
}

func NewUserHandler(userUseCase usecase.UserUseCase) UserHandler {
	return &userHandler{
		useCase: userUseCase,
	}
}

// HandlerCreate implements UserHandler.
func (uh *userHandler) HandlerCreate(c *gin.Context) {
	user := new(model.User)

	// リクエストボディから構造体へバインディング
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	}

	// ユーザー情報の新規登録
	newUser, err := uh.useCase.CreateUser(c, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	} else {
		c.JSON(http.StatusCreated, newUser)
	}
}

// HandlerGet implements UserHandler.
func (uh *userHandler) HandlerGet(c *gin.Context) {
	user := new(model.User)
	name, exsist := c.Get("name")
	if exsist {
		user.Name = name.(string)
	}
	email, exsist := c.Get("email")
	if exsist {
		user.Email = email.(string)
	}

	findUser, err := uh.useCase.GetUserByEmail(c, user)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"err": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		}
	} else {
		c.JSON(http.StatusOK, findUser)
	}
}

// HandlerUpdate implements UserHandler.
func (uh *userHandler) HandlerUpdate(c *gin.Context) {
	user := new(model.User)

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	}

	email, exsist := c.Get("email")
	if exsist {
		user.Email = email.(string)
	}

	updateUser, err := uh.useCase.UpdateUser(c, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	} else {
		c.JSON(http.StatusOK, updateUser)
	}
}
