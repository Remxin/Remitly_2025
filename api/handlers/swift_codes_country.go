package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *Handler) GetCountryIsoDetails(c *gin.Context) {
	countryISO2Code := c.Param("countryISO2code")
	if len(countryISO2Code) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "no swift-code provided",
		})
	}

}
