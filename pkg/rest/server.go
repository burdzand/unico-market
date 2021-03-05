package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

//ResponseInternalServerError responds with a StatusInternalServerError, status FAILURE ad errorMessage
func ResponseInternalServerError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, Response{Status: "FAILURE", Message: err.Error()})
	c.Abort()
}

//ResponseInternalServerError responds with a StatusInternalServerError, status FAILURE ad errorMessage
func ResponseBadRequest(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, Response{Status: "FAILURE", Message: err.Error()})
	c.Abort()
}

//ResponseSuccess responds with a Status200ok, status SUCCESS ad errorMessage empty
func ResponseSuccess(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "SUCCESS"})
	c.Abort()
}

func ResponseNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, Response{Status: "FAILURE", Message: "Not Found"})
	c.Abort()
}
