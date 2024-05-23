package api

import (
	"project/api/handlers"

	"github.com/gin-gonic/gin"
)

func NewGin(h *handlers.HTTPHandler) *gin.Engine {
	r := gin.Default()

	company := r.Group("/company")
	company.GET("/:id", h.GetCompany)
	company.POST("/", h.CreateCompany)
	company.PUT("/", h.UpdateCompany)
	company.DELETE("/:id", h.DeleteCompany)
	r.GET("/companies", h.GetAllCompanies)

	recruiter := r.Group("/recruiter")
	recruiter.GET("/:id", h.GetRecruiter)
	recruiter.POST("/", h.CreateRecruiter)
	recruiter.PUT("/", h.UpdateRecruiter)
	recruiter.DELETE("/:id", h.DeleteRecruiter)
	r.GET("/recruiters", h.GetAllRecruiters)

	vacancy := r.Group("/vacancy")
	vacancy.GET("/:id", h.GetVacancy)
	vacancy.POST("/", h.CreateVacancy)
	vacancy.PUT("/", h.UpdateVacancy)
	vacancy.DELETE("/:id", h.DeleteVacancy)
	r.GET("/vacancies", h.GetAllVacancies)

	interview := r.Group("/interview")
	interview.GET("/:id", h.GetInterview)
	interview.POST("/", h.CreateInterview)
	interview.PUT("/", h.UpdateInterview)
	interview.DELETE("/:id", h.DeleteInterview)
	r.GET("/interviews", h.GetAllInterviews)

	// The end
	return r
}
