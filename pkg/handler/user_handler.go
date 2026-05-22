package handler

import (
	"app/pkg/app"
	"app/pkg/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	App *app.App
}

func NewUserHandler(app *app.App) *UserHandler {
	return &UserHandler{
		App: app,
	}
}

func (h *UserHandler) AddUser(c *gin.Context) {
	user, err := h.App.Queries.AddUser(c.Request.Context(), sqlc.AddUserParams{Name: "arpit", Email: "arpit@gmail.com"}) // need to use sqlc and not h.App.Queries since its a object of *sqlc.Queries types and does not contain the type  AddUserParam
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	user, err := h.App.Queries.GetUser(c.Request.Context(), "arpit@gmail.com")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}
