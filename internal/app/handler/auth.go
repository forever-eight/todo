package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/forever-eight/todo.git/internal/app/ds"
)

func (h *Handler) signUp(c *gin.Context) {
	var input = new(ds.User)
	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}

func (h *Handler) signIn(c *gin.Context) {

}
