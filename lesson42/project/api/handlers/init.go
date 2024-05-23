package handlers

import (
	"project/config/logger"
	"project/postgresql/managers"
)

type HTTPHandler struct {
	UM *managers.UserManager
	RM *managers.ResumeManager
	//CM     *managers.CompanyManager
	//RCM    *managers.RecruiterManager
	//VM     *managers.VacancyManager
	//IM     *managers.InterviewManager
	Logger logger.Logger
}

func NewHTTPHandler(um *managers.UserManager, rm *managers.ResumeManager, // cm *managers.CompanyManager, rcm *managers.RecruiterManager, vm *managers.VacancyManager, im *managers.InterviewManager
	logger logger.Logger) *HTTPHandler {
	return &HTTPHandler{UM: um, RM: rm,
		//CM: cm, RCM: rcm, VM: vm, IM: im,
		Logger: logger}
}
