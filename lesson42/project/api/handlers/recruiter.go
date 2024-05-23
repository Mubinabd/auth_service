package handlers

//
//func (h *HTTPHandler) GetRecruiter(c *gin.Context) {
//	id := c.Param("id")
//	recruiter, err := h.RCM.GetRecruiterByID(id)
//	if err != nil {
//		c.JSON(http.StatusNotFound, gin.H{"Error getting recruiter": err.Error()})
//		return
//	}
//	c.JSON(http.StatusOK, recruiter)
//}
//
//func (h *HTTPHandler) GetAllRecruiters(c *gin.Context) {
//	companyID := c.Query("company_id")
//	gender := c.Query("gender")
//	age := c.Query("age")
//	recruiters, err := h.RCM.GetAllRecruiters(companyID, gender, age)
//	if err != nil {
//		c.JSON(http.StatusNotFound, gin.H{"Error getting all recruiters": err.Error()})
//		return
//	}
//	c.JSON(http.StatusOK, recruiters)
//}
//
//func (h *HTTPHandler) CreateRecruiter(c *gin.Context) {
//	var recruiter models.RecruiterCreated
//	c.BindJSON(&recruiter)
//	err := h.RCM.CreateRecruiter(&recruiter)
//	if err != nil {
//		c.JSON(500, gin.H{"Error creating recruiter": err.Error()})
//		h.Logger.ERROR.Println("Recruiter not created: ", err.Error())
//		return
//	}
//	c.String(200, "Recruiter created")
//	h.Logger.INFO.Println("Recruiter created")
//}
//
//func (h *HTTPHandler) UpdateRecruiter(c *gin.Context) {
//	var recruiter models.RecruiterUpdated
//	err := c.BindJSON(&recruiter)
//	if err != nil {
//		c.JSON(500, gin.H{"Error updating recruiter": err.Error()})
//		h.Logger.ERROR.Println("Recruiter not updated: ", err.Error())
//		return
//	}
//	err = h.RCM.UpdateRecruiter(&recruiter)
//	if err != nil {
//		c.JSON(500, gin.H{"Error updating recruiter": err.Error()})
//		h.Logger.ERROR.Println("Recruiter not updated: ", err.Error())
//		return
//	}
//	c.String(200, "Recruiter updated")
//	h.Logger.INFO.Println("Recruiter updated")
//}
//
//func (h *HTTPHandler) DeleteRecruiter(c *gin.Context) {
//	id := c.Param("id")
//	err := h.RCM.DeleteRecruiter(id)
//	if err != nil {
//		c.JSON(500, gin.H{"Error deleting recruiter": err.Error()})
//		h.Logger.ERROR.Println("Recruiter not deleted: ", err.Error())
//		return
//	}
//	c.String(200, "Recruiter deleted")
//	h.Logger.INFO.Println("Recruiter deleted")
//}
