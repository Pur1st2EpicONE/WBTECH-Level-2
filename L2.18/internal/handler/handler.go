package handler

import (
	"net/http"

	v1 "L2.18/internal/handler/v1"
	"L2.18/internal/service"
	"L2.18/pkg/logger"
	"github.com/gin-gonic/gin"
)

func NewHandler(service service.Service, logger logger.Logger) http.Handler {

	router := gin.New()

	apiV1 := router.Group("/api/v1")
	handlerV1 := v1.NewHandler(service, logger)

	apiV1.POST("/create_event", handlerV1.CreateEvent)
	// apiV1.POST("/update_event", handlerV1.UpdateEvent)
	// apiV1.POST("/delete_event", handlerV1.DeleteEvent)

	// apiV1.GET("/events_for_day/:day", handlerV1.GetEventsForDay)
	// apiV1.GET("/events_for_week/:week", handlerV1.GetEventsForWeek)
	// apiV1.GET("/events_for_month/:month", handlerV1.GetEventsForMonth)

	return router

}
