package api

import (
	"net/http"

	models "example.com/m/v2/api/models"
	db "example.com/m/v2/db/sqlc"
	"github.com/gin-gonic/gin"
)

type GetDetailsCountryCodeResponse struct {
	CountryISO2 string          `json:"countryISO2"`
	CountryName string          `json:"countryName"`
	Branches    []models.Branch `json:"branches"`
}

func convertToGetDetailsCountryResponse(rows []db.GetDetailsCountryRow) (*GetDetailsCountryCodeResponse, error) {
	var response GetDetailsCountryCodeResponse
	response.CountryISO2 = rows[0].CountryIso2Code
	response.CountryName = rows[0].CountryName
	for _, row := range rows {
		branch := models.Branch{
			Address:       row.Address,
			BankName:      row.BankName,
			CountryISO2:   row.CountryIso2Code,
			IsHeadquarter: row.Parent == "PARENT",
			SwiftCode:     row.SwiftCode,
		}
		response.Branches = append(response.Branches, branch)
	}
	return &response, nil
}

func (handler *Handler) GetCountryIsoDetails(c *gin.Context) {
	countryISO2Code := c.Param("countryISO2code")
	if len(countryISO2Code) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "no swift-code provided",
		})
		return
	}
	response, err := handler.Store.GetDetailsCountry(c, countryISO2Code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error in query",
		})
		return
	}

	formattedResponse, err := convertToGetDetailsCountryResponse(response)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "cannot format response",
		})
		return
	}
	c.JSON(http.StatusOK, formattedResponse)
}
