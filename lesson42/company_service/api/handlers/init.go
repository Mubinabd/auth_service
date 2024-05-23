package handlers

import (
	"project/config/logger"
	"project/postgresql/managers"
)

type HTTPHandler struct {
	CM     *managers.CompanyManager
	RCM    *managers.RecruiterManager
	VM     *managers.VacancyManager
	IM     *managers.InterviewManager
	Logger logger.Logger
}

func NewHTTPHandler(cm *managers.CompanyManager, rcm *managers.RecruiterManager, vm *managers.VacancyManager, im *managers.InterviewManager, logger logger.Logger) *HTTPHandler {
	return &HTTPHandler{CM: cm, RCM: rcm, VM: vm, IM: im, Logger: logger}
}
