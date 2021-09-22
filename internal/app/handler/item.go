package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/forever-eight/todo.git/internal/app/ds"
)

func (h *Handler) createItem(c *gin.Context) {
	userId, err := h.GetIntId(c)
	if err != nil {
		return
	}
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "invalid list id param")
		return
	}
	var input ds.TodoItem
	err = c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "invalid list id param")
		return
	}
	id, err := h.services.TodoItem.Create(userId, listId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "invalid list id param")
		return
	}
	c.JSON(http.StatusOK, map[string]string{
		"id": id,
	})
}

func (h *Handler) getAllItems(c *gin.Context) {

}

func (h *Handler) getItemById(c *gin.Context) {

}

func (h *Handler) updateItem(c *gin.Context) {

}

func (h *Handler) deleteItem(c *gin.Context) {

}
