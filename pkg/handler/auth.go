package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	todo "github.com/ohpasha/rest-api"
	"github.com/sirupsen/logrus"
)

// @Summary      signUp
// @Description  create account
// @ID			 signup
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input   body      todo.User  true  "account info"
// @Success      200  {integer}  integer 1
// @Failure      400  {object}  errorResponse
// @Failure      404  {object}  errorResponse
// @Failure      500  {object}  errorResponse
// @Router       /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var input todo.User

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid input body")

		return
	}

	id, err := h.services.Authorization.CreateUser(input)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary      signIn
// @Description  login
// @ID			 signin
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input   body      signInInput  true  "credentials"
// @Success      200  {integer}  string "token"
// @Failure      400  {object}  errorResponse
// @Failure      404  {object}  errorResponse
// @Failure      500  {object}  errorResponse
// @Router       /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		logrus.Errorf(err.Error())

		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})

}
