/*package handler

import (
	"encoding/json"
	"github.com/amrchnk/auth_service/pkg/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var input models.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	input.Id = id
	userSession, err := json.Marshal(input)
	if err != nil {
		return
	}
	session := sessions.Default(c)
	session.Set("UserSession",userSession)
	session.Save()

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInRequest

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userSession, err := h.services.Authorization.CheckUser(input.Login, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	session := sessions.Default(c)
	session.Set("UserSession", userSession)
	session.Save()

	c.JSON(http.StatusOK, map[string]interface{}{
		"session": userSession,
	})
}*/
