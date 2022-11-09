package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	todo "github.com/ohpasha/rest-api"
)

func (h *Handler) createItem(c *gin.Context) {
	userId, ok := getUserId(c)

	if ok != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "no userId in context")

		return
	}

	listId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	var input todo.TodoItem

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	id, err := h.services.TodoItem.Create(userId, listId, input)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "can't create list item")

		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

func (h *Handler) getAllItems(c *gin.Context) {
	userId, ok := getUserId(c)

	if ok != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "no userId in context")

		return
	}

	listId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "no listId")
	}

	items, err := h.services.TodoItem.GetAll(userId, listId)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) getItemById(c *gin.Context) {

}

func (h *Handler) updateItem(c *gin.Context) {

}

func (h *Handler) deleteItem(c *gin.Context) {

}
