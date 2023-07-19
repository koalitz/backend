package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/koalitz/backend/internal/controller/dao"
	"net/http"
)

func (h *Handler) getMe(c *gin.Context, info *dao.Session) error {

	user, err := h.user.FindMe(info)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, user)
	return nil
}
