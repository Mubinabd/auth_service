package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "project/api/docs"
	"project/api/handlers"
)

func NewGin(h *handlers.HTTPHandler) *gin.Engine {
	r := gin.Default()
	r.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	user := r.Group("/user")
	user.GET("/:id", h.GetUser)
	user.POST("/", h.CreateUser)
	user.PUT("/", h.UpdateUser)
	user.DELETE("/:id", h.DeleteUser)
	user.GET("/:id/my-resumes", h.GetAllUserResumes)
	user.GET("/:id/my-interviews", h.GetAllUserInterviews)
	r.GET("/users", h.GetAllUsers)

	resume := r.Group("/resume")
	resume.GET("/:id", h.GetResume)
	resume.POST("/", h.CreateResume)
	resume.PUT("/", h.UpdateResume)
	resume.DELETE("/:id", h.DeleteResume)
	r.GET("/resumes", h.GetAllResumes)

	company := r.Group("/company")
	company.GET("/:id", h.GetCompany)
	//company.POST("/", h.CreateCompany)
	company.PUT("/", h.UpdateCompany)
	//company.DELETE("/:id", h.DeleteCompany)
	//r.GET("/companies", h.GetAllCompanies)

	//recruiter := r.Group("/recruiter")
	//recruiter.GET("/:id", h.GetRecruiter)
	//recruiter.POST("/", h.CreateRecruiter)
	//recruiter.PUT("/", h.UpdateRecruiter)
	//recruiter.DELETE("/:id", h.DeleteRecruiter)
	//r.GET("/recruiters", h.GetAllRecruiters)
	//
	//vacancy := r.Group("/vacancy")
	//vacancy.GET("/:id", h.GetVacancy)
	//vacancy.POST("/", h.CreateVacancy)
	//vacancy.PUT("/", h.UpdateVacancy)
	//vacancy.DELETE("/:id", h.DeleteVacancy)
	//r.GET("/vacancies", h.GetAllVacancies)
	//
	//interview := r.Group("/interview")
	//interview.GET("/:id", h.GetInterview)
	//interview.POST("/", h.CreateInterview)
	//interview.PUT("/", h.UpdateInterview)
	//interview.DELETE("/:id", h.DeleteInterview)
	//r.GET("/interviews", h.GetAllInterviews)

	// The end
	return r
}
