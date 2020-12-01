package controllers

import (
	"context"
	"strconv"

	"github.com/Piichet/app/ent"
	"github.com/Piichet/app/ent/title"
	"github.com/gin-gonic/gin"
)

// TitleController defines the struct for the title controller
type TitleController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateTitle handles POST requests for adding Title entities
// @Summary Create title
// @Description Create title
// @ID create-title
// @Accept   json
// @Produce  json
// @Param title body ent.User true "title entity"
// @Success 200 {object} ent.Title
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /titles [post]
func (ctl *TitleController) CreateTitle(c *gin.Context) {
	obj := ent.Title{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "title binding failed",
		})
		return
	}
	t, err := ctl.client.Title.
		Create().
		SetTitle(obj.Title).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, t)
}

// GetTitle handles GET requests to retrieve a title entity
// @Summary Get a title entity by ID
// @Description get title by ID
// @ID get-title
// @Produce  json
// @Param id path int true "title ID"
// @Success 200 {object} ent.Title
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /titles/{id} [get]
func (ctl *TitleController) GetTitle(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	t, err := ctl.client.Title.
		Query().
		Where(title.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, t)
}

// ListTitle handles request to get a list of title entities
// @Summary List title entities
// @Description list title entities
// @ID list-title
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Title
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /titles [get]
func (ctl *TitleController) ListTitle(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	titles, err := ctl.client.Title.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, titles)
}

// NewTitleController creates and registers handles for the title controller
func NewTitleController(router gin.IRouter, client *ent.Client) *TitleController {
	tc := &TitleController{
		client: client,
		router: router,
	}
	tc.register()
	return tc
}

// InitTitleController registers routes to the main engine
func (ctl *TitleController) register() {
	titles := ctl.router.Group("/titles")
	titles.POST("", ctl.CreateTitle)
	titles.GET(":id", ctl.GetTitle)
	titles.GET("", ctl.ListTitle)
}
