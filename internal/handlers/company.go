package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"

	"job-port-api/internal/auth"
	"job-port-api/internal/middleware"
	"job-port-api/internal/models"
)

// AddCompany handles the addition of a company
func (h *handler) AddCompany(c *gin.Context) {
	ctx := c.Request.Context()
	traceID, traceIDExists := ctx.Value(middleware.TraceIdKey).(string)
	if !traceIDExists {
		log.Error().Msg("traceid missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return
	}
	_, traceIDExists = ctx.Value(auth.Ctxkey).(jwt.RegisteredClaims)
	if !traceIDExists {
		log.Error().Str("Trace Id", traceID).Msg("login first")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
		return
	}
	var companyData models.Company

	err := json.NewDecoder(c.Request.Body).Decode(&companyData)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceID)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "please provide valid name and location",
		})
		return
	}
	validate := validator.New()
	err = validate.Struct(companyData)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceID)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "please provide valid name and location",
		})
		return
	}

	companyData, err = h.service.AddCompany(ctx, companyData)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceID)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, companyData)

}

// ViewCompany handles viewing details of a specific company.
func (h *handler) FetchCompanyById(c *gin.Context) {
	ctx := c.Request.Context()
	traceID, traceIDExists := ctx.Value(middleware.TraceIdKey).(string)
	if !traceIDExists {
		log.Error().Msg("traceid missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return
	}
	_, traceIDExists = ctx.Value(auth.Ctxkey).(jwt.RegisteredClaims)
	if !traceIDExists {
		log.Error().Str("Trace Id", traceID).Msg("login first")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
		return
	}
	id := c.Param("id")

	cid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	companyData, err := h.service.FetchCompByid(cid)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceID)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, companyData)
}

// ViewAllCompanies lists all available companies.
func (h *handler) FetchAllCompanies(c *gin.Context) {
	ctx := c.Request.Context()
	traceID, traceIDExists := ctx.Value(middleware.TraceIdKey).(string)
	if !traceIDExists {
		log.Error().Msg("traceid missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return
	}
	_, traceIDExists = ctx.Value(auth.Ctxkey).(jwt.RegisteredClaims)
	if !traceIDExists {
		log.Error().Str("Trace Id", traceID).Msg("login first")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
		return
	}

	companyDetails, err := h.service.FetchAllCompanies()
	if err != nil {
		log.Error().Err(err).Str("trace id", traceID)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, companyDetails)
}

// func (h *handler) FetchCompanyById(c *gin.Context) {
// 	ctx := c.Request.Context()
// 	traceId, ok := ctx.Value(middleware.TraceIdKey).(string)
// 	if !ok {
// 		log.Error().Msg("traceId missing from context")
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
// 		return
// 	}
// 	stringCmpnyId := c.Param("id")
// 	cid, err := strconv.ParseUint(stringCmpnyId, 10, 64)
// 	if err != nil {

// 		log.Print("conversion string to int error", err)
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "error found at conversion.."})
// 		return

// 	}
// 	companyData, err := h.service.FetchCompByid(cid)
// 	if err != nil {
// 		log.Error().Err(err).Str("Trace Id", traceId).Msg("problem in fetching company by id")
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "failed to get company details by id"})
// 		return
// 	}
// 	// If everything goes right, respond with the created user
// 	c.JSON(http.StatusOK, companyData)
// }
