package handler

import (
	"barTrApp/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	courses := router.Group("/courses", h.employeeIdentity)
	{
		courses.GET("", h.getCourses)
	}

	lessons := router.Group("/lessons", h.employeeIdentity)
	{
		lessons.GET("", h.getLessons)
	}

	lection := router.Group("/lections", h.employeeIdentity)
	{
		lection.GET("", h.getLections)
	}

	information := router.Group("/informations", h.employeeIdentity)
	{
		information.GET("", h.getInformations)
	}

	tests := router.Group("/tests", h.employeeIdentity)
	{
		tests.GET("", h.getTests)
	}

	questions := router.Group("/questions", h.employeeIdentity)
	{
		questions.GET("", h.getQuestions)
	}

	answers := router.Group("/answers", h.employeeIdentity)
	{
		answers.GET("", h.getAnswers)
	}

	return router
}
