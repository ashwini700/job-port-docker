package handlers

import (
	"encoding/json"
	"job-port-api/internal/middleware"
	"job-port-api/internal/models"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func (h *handler) AcceptApplicants(c *gin.Context) {
	ctx := c.Request.Context()
	trackerId, ok := ctx.Value(middleware.TraceIdKey).(string)
	if !ok {
		log.Error().Msg("TrackerId missing from context in Filter Applicant")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}

	var newApplicant []models.ApplicantsRequest

	err := json.NewDecoder(c.Request.Body).Decode(&newApplicant)
	if err != nil {
		log.Error().Err(err).Str("tracker Id", trackerId).Send()
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}
	validate := validator.New()
	for _, data := range newApplicant {
		err = validate.Struct(data)
		if err != nil {
			log.Error().Err(err).Str("tracker Id", trackerId).Interface("body", newApplicant).Send()
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": "All fields are mandatory"})
			return
		}
	}

	filteredData, err := h.service.ApplicantsFilter(ctx, newApplicant)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}

	c.JSON(http.StatusOK, filteredData)
}
