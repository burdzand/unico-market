package api

import (
	"unico/internal/unico/services"
	"unico/pkg/rest"

	"github.com/gin-gonic/gin"
)

// ImportData godoc
// @Summary Import CSV Data
// @Description Import CSV Data
// @Accept  multipart/form-data
// @Produce  json
// @Param file formData file true "File Path"
// @Success 200
// @Failure 500 {object} rest.Response
// @Router /import-data [post]
func (ch *controlerHandler) ImportData(c *gin.Context) {
	// Upload Single CSV file
	file, err := c.FormFile("file")
	if err != nil {
		rest.ResponseBadRequest(c, err)
		return
	}

	log.WithField("file", file.Filename).
		Info("Load CSV")

	// Market ImportCSV
	err = services.ImportDataCSV(c, ch.db, file)
	if err != nil {
		rest.ResponseBadRequest(c, err)
		return
	}
	rest.ResponseSuccess(c)

}
