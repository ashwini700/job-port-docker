package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"job-port-api/internal/auth"
	"job-port-api/internal/middleware"
	"job-port-api/internal/service"
)

func API(a auth.TokenAuth, sc service.UserService) *gin.Engine {
	r := gin.New()
	m, err := middleware.NewMid(a)
	if err != nil {
		log.Panic("middleware not setup")
		return nil
	}
	h, err := NewHandlerFunc(sc)
	if err != nil {
		log.Panic("handler not setup")
		return nil
	}
	r.Use(middleware.Log(), gin.Recovery())

	r.GET("/check", m.Authenticate((check)))
	r.POST("/register", h.SignUp)
	r.POST("/login", h.Login)
	r.POST("/Addcomp", m.Authenticate(h.AddCompany))
	r.GET("/FetchCompByid/:id", m.Authenticate(h.FetchCompanyById))
	r.GET("/Allcomp", m.Authenticate(h.FetchAllCompanies))
	//
	r.POST("/Addjobs/:id/jobs", m.Authenticate(h.AddJob))
	r.GET("/FetchJobsBycompid/:id", m.Authenticate(h.FetchJobByCompanyId))
	r.GET("/jobs/Alljobs", m.Authenticate(h.FetchAllJobs))
	// r.GET("/jobs/FetchjobsByid/:id", m.Authenticate(h.FetchJobByCompanyId))
	r.POST("/applicant",m.Authenticate(h.AcceptApplicants))

	return r
}
func check(c *gin.Context) {
	c.JSON(http.StatusOK, "Msg :ok")

}
