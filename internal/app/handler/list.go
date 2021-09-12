package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/forever-eight/todo.git/internal/app/ds"
)

func (h *Handler) createList(c *gin.Context) {
	userId, err := h.GetIntId(c)
	if err != nil {
		return
	}
	var input ds.TodoList
	err = c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// call service method
	id, err := h.services.TodoList.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK,
		map[string]interface{}{
			"id": id,
		})
}

func (h *Handler) getAllLists(c *gin.Context) {

}

func (h *Handler) getListById(c *gin.Context) {

}

func (h *Handler) updateList(c *gin.Context) {

}

func (h *Handler) deleteList(c *gin.Context) {

}
