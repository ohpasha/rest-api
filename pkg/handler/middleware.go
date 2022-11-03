package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	AuthHeader  = "Authorization"
	userContext = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(AuthHeader)

	if header == "" {
		NewErrorResponse(c, http.StatusUnauthorized, fmt.Sprintf("empty %s header", AuthHeader))

		return
	}

	headerParts := strings.Split(header, " ")

	if len(headerParts) != 2 {
		NewErrorResponse(c, http.StatusUnauthorized, fmt.Sprintf("wrong %s header", AuthHeader))

		return
	}

	userId, err := h.services.Authorization.ParseToken(headerParts[1])

	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userContext, userId)
}
