package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	todo "github.com/ohpasha/rest-api"
)

func (h *Handler) createList(c *gin.Context) {
	id, ok := getUserId(c)

	if ok != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "no userId in context")

		return
	}

	var input todo.TodoList

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	id, err := h.services.TodoList.Create(id, input)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllListResponse struct {
	Data []todo.TodoList `json:"data"`
}

func (h *Handler) getAllLists(c *gin.Context) {
	userId, ok := getUserId(c)

	if ok != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "no userId in context")

		return
	}

	lists, err := h.services.TodoList.GetAll(userId)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, getAllListResponse{
		Data: lists,
	})

}

func (h *Handler) getListById(c *gin.Context) {

}

func (h *Handler) updateList(c *gin.Context) {

}

func (h *Handler) deleteList(c *gin.Context) {

}
