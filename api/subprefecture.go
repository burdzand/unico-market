package api

import (
	"net/http"
	"strconv"
	"unico/pkg/rest"

	"github.com/gin-gonic/gin"
)

// GetAllSubprefecture godoc
// @Summary List existing Subprefecture
// @Description Get all the existing Subprefecture
// @Accept  json
// @Produce  json
// @Success 200 {array} entities.SubPrefecture
// @Failure 500 {object} rest.Response
// @Router /subprefecture [get]
func (ch *controlerHandler) GetAllSubprefecture(c *gin.Context) {
	records, err := ch.db.GetAllSubPrefecture()
	if err != nil {
		log.Error(err)
		rest.ResponseBadRequest(c, err)
		return
	}
	c.JSON(http.StatusOK, records)
}

// GetAllDistrict godoc
// @Summary List existing District by Subprefecture
// @Description Get all the existing District by Subprefecture
// @Accept  json
// @Produce  json
// @Param  subprefecture path string true "Subprefecture ID"
// @Success 200 {array} entities.District
// @Failure 500 {object} rest.Response
// @Router /district/{subprefecture} [get]
func (ch *controlerHandler) GetAllDistrict(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("subprefecture"), 10, 32)
	if err != nil {
		rest.ResponseBadRequest(c, err)
		return
	}
	records, err := ch.db.GetAllDistrict(id)
	c.JSON(http.StatusOK, records)
}
