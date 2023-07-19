package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/koalitz/backend/internal/controller/dao"
	"github.com/koalitz/backend/internal/controller/dto"
	"net/http"
	"os"
)

func (h *Handler) getMe(c *gin.Context, info *dao.Session) error {

	user, err := h.user.FindUserByID(info.ID)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, user)
	return nil
}

func (h *Handler) getImages(c *gin.Context, l dto.Limit) error {
	entries, err := os.ReadDir("./" + h.cfg.Files.Path)
	if err != nil {
		return err
	}

	if l.Limit != 0 {
		entries = entries[:l.Limit]
	}

	c.JSON(http.StatusOK, entries)
	return nil
}
