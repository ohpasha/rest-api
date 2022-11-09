package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	todo "github.com/ohpasha/rest-api"
)

// @Summary      create todo list
// @Description  create todo list
// @ID			 create-list
// @Tags         lists
// @Security 	 ApiKeyAuth
// @Accept       json
// @Produce      json
// @Param        input   body      todo.TodoList  true  "list info"
// @Success      200  {integer}  string "token"
// @Failure      400  {object}  errorResponse
// @Failure      404  {object}  errorResponse
// @Failure      500  {object}  errorResponse
// @Router       /api/lists [post]
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

	list, err := h.services.TodoList.GetById(userId, id)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, list)
}

func (h *Handler) updateList(c *gin.Context) {
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

func (h *Handler) deleteList(c *gin.Context) {
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

	err = h.services.TodoList.Delete(userId, id)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
