package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/koalitz/backend/internal/controller/dao"
	"github.com/koalitz/backend/internal/controller/dto"
	"github.com/koalitz/backend/pkg/middleware/errs"
	"image"
	"net/http"
)

func (h *Handler) getPostById(c *gin.Context, id dto.ID) error {
	p, err := h.post.FindPostByID(id.ID)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, p)
	return nil
}

func (h *Handler) getPostByTitle(c *gin.Context, t dto.Title) error {
	p, err := h.post.FindPostByTitle(t.Title)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, p)
	return nil
}

func (h *Handler) createPost(c *gin.Context, post dto.PostInfo, info *dao.Session) error {
	var imageName *string
	if post.Image != nil {
		img, err := post.Image.Open()
		if err != nil {
			return err
		}

		var imgConfig image.Config
		imgConfig, _, err = image.DecodeConfig(img)
		if err != nil {
			return errs.UnsupportedImageType.AddErr(err)
		}

		if imgConfig.Width > 3000 || imgConfig.Height > 3000 {
			return errs.ImageTooLarge
		}

		name, err := h.sess.GenerateFileName(c, post.Image)
		if err != nil {
			return err
		}
		imageName = &name
	}
	p, err := h.post.CreatePost(imageName, post, info.ID)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, p)
	return nil
}
