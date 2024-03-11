package handler

import (
	"barTrApp/pkg/service"

	_ "barTrApp/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		courses := api.Group("/courses")
		{
			courses.GET("", h.getCourses)
		}

		lessons := api.Group("/lessons")
		{
			lessons.GET("", h.getLessons)
		}

		lection := api.Group("/lections")
		{
			lection.GET("", h.getLections)
		}

		information := api.Group("/information")
		{
			information.GET("", h.getInformations)
		}

		tests := api.Group("/tests")
		{
			tests.GET("", h.getTests)
		}

		questions := api.Group("/questions")
		{
			questions.GET("", h.getQuestions)
		}

		answers := api.Group("/answers")
		{
			answers.GET("", h.getAnswers)
		}
	}

	return router
}
