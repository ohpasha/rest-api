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
	userId, ok := getUserId(c)

	if ok != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "no userId in context")

		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	item, err := h.services.TodoItem.GetById(userId, itemId)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}
	c.JSON(http.StatusOK, item)

}

func (h *Handler) updateItem(c *gin.Context) {
	userId, ok := getUserId(c)

	if ok != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "no userId in context")

		return
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "id not integer")

		return
	}

	var input todo.UpdateItemInput

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	err = h.services.TodoItem.Update(userId, id, input)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})

}

func (h *Handler) deleteItem(c *gin.Context) {
	userId, ok := getUserId(c)

	if ok != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "no userId in context")

		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	err = h.services.TodoItem.Delete(userId, itemId)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}
	c.JSON(http.StatusOK, statusResponse{
		"ok",
	})
}
