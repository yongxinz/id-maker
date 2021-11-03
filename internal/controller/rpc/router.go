package rpc

import (
	"id-maker/internal/usecase"
	"id-maker/pkg/logger"
)

func NewRouter(s usecase.Segment, l logger.Interface) {
	newSegmentRoutes(s, l)
}
