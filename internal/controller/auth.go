package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/koalitz/backend/internal/controller/dao"
	"github.com/koalitz/backend/internal/controller/dto"
	"github.com/koalitz/backend/pkg/middleware/errs"
	"net/http"
)

func (h *Handler) sendCodeToEmail(c *gin.Context, to dto.Email) error {

	code := h.sess.GenerateSecretCode(5)
	if err := h.auth.SetCodes(to.Email, code); err != nil {
		return err
	}

	if err := h.mail.DialAndSend("Verify email for koalitz account", code, to.Email); err != nil {
		return errs.EmailError.AddErr(err)
	}

	c.Status(http.StatusOK)
	return nil
}

func (h *Handler) signInByEmail(c *gin.Context, auth dto.EmailWithCode) error {
	if oki, err := h.auth.EqualsPopCode(auth.Email, auth.Code); err != nil {
		return err
	} else if !oki {
		return errs.MailCodeError.AddErr(err)
	}

	customer, err := h.auth.AuthUserByEmail(auth.Email)

	if err != nil {
		return err
	}

	err = h.sess.SetNewCookie(customer.ID, c)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, dao.TransformToMe(customer))
	return nil
}

func (h *Handler) signOut(c *gin.Context) error {
	session, _ := c.Cookie(h.cfg.Session.CookieName)
	info, _, err := h.sess.ValidateSession(session)
	if err != nil {
		return errs.UnAuthorized.AddErr(err)
	}

	h.auth.DelKeys(session)

	user, err := h.user.FindUserByID(info.ID)
	if err != nil {
		return err
	}

	for i, v := range user.Sessions {
		if v == session {
			if err = user.Update().SetSessions(append(user.Sessions[:i], user.Sessions[i+1:]...)).
				Exec(context.Background()); err != nil {
				return err
			}
		}
	}

	c.Status(http.StatusOK)
	return nil
}
