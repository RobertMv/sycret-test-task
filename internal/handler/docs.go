package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sycret-test-task/internal/model"
)

func (h *Handlers) generateDoc(c *gin.Context) {
	var input model.Input

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.services.DocsGenerator.Generate(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}
