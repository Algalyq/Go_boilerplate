package handler

import (

	"github.com/Algalyq/Go_boilerplate/pkg/service"
	"github.com/gin-gonic/gin"

)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
    return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	//  gin.SetMode(gin.ReleaseMode)
	 router := gin.New()
	 auth := router.Group("/auth")
	 {
		auth.GET("/signup", h.signup			)
	 }
	 return router
}