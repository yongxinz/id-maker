package v1

import (
	"id-maker/internal/entity"
	"id-maker/internal/usecase"
	"id-maker/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type segmentRoutes struct {
	s usecase.Segment
	l logger.Interface
}

func newSegmentRoutes(handler *gin.RouterGroup, s usecase.Segment, l logger.Interface) {
	r := &segmentRoutes{s, l}

	h := handler.Group("/")
	{
		h.GET("/ping", r.pong)
		h.GET("/id/:tag", r.GetId)
		h.GET("/snowid", r.GetSnowId)
		h.POST("/tag", r.CreateTag)
	}
}

func (r *segmentRoutes) pong(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}

// @Summary     GetId
// @Description get id
// @Tags  	    segments
// @Accept      json
// @Produce     json
// @Success     200 {object} response
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /id/:tag [get]
func (r *segmentRoutes) GetId(c *gin.Context) {
	var (
		tag string
		id  int64
		err error
	)

	tag = c.Param("tag")
	if tag == "" {
		r.l.Error(err, "http - v1 - GetId")
		errorResponse(c, http.StatusBadRequest, "tag cannot empty")

		return
	}

	if id, err = r.s.GetId(tag); err != nil {
		r.l.Error(err, "http - v1 - GetId")
		errorResponse(c, http.StatusInternalServerError, "service problems")

		return
	}

	c.JSON(http.StatusOK, id)
}

func (r *segmentRoutes) GetSnowId(c *gin.Context) {
	id := r.s.SnowFlakeGetId()

	c.JSON(http.StatusOK, id)
}

func (r *segmentRoutes) CreateTag(c *gin.Context) {
	var request entity.Segments
	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - v1 - CreateTag")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	if err := r.s.CreateTag(&request); err != nil {
		r.l.Error(err, "http - v1 - CreateTag")
		errorResponse(c, http.StatusBadRequest, "service problems")

		return
	}

	c.JSON(http.StatusOK, nil)
}
