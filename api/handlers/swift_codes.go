package api

import (
	"net/http"

	models "example.com/m/v2/api/models"
	db "example.com/m/v2/db/sqlc"
	"github.com/gin-gonic/gin"
)

type GetDefailsSwiftCodeResponse struct {
	Address       string          `json:"address"`
	BankName      string          `json:"bankName"`
	CountryISO2   string          `json:"countryISO2"`
	CountryName   string          `json:"countryName"`
	IsHeadquarter bool            `json:"isHeadquarter"`
	SwiftCode     string          `json:"swiftCode"`
	Branches      []models.Branch `json:"branches"`
}

func convertToResponse(rows []db.GetDetailsSwiftRow) (*GetDefailsSwiftCodeResponse, error) {
	var response GetDefailsSwiftCodeResponse

	if len(rows) == 1 {
		row := rows[0]
		response.Address = row.Address
		response.BankName = row.BankName
		response.CountryISO2 = row.CountryIso2Code
		response.CountryName = row.CountryName
		response.IsHeadquarter = false
		response.SwiftCode = row.SwiftCode
		return &response, nil
	}

	for _, row := range rows {
		if row.Parent == "PARENT" {
			response.Address = row.Address
			response.BankName = row.BankName
			response.CountryISO2 = row.CountryIso2Code
			response.CountryName = row.CountryName
			response.IsHeadquarter = true
			response.SwiftCode = row.SwiftCode
		} else {
			branch := models.Branch{
				Address:       row.Address,
				BankName:      row.BankName,
				CountryISO2:   row.CountryIso2Code,
				IsHeadquarter: false,
				SwiftCode:     row.SwiftCode,
			}
			response.Branches = append(response.Branches, branch)
		}
	}

	return &response, nil
}

func (handler *Handler) GetDetailsSwiftCode(c *gin.Context) {
	swiftCode := c.Param("swift-code")
	if len(swiftCode) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "no swift-code provided",
		})
		return
	}

	res, err := handler.Store.GetDetailsSwift(c, swiftCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error in query",
		})
		return
	}

	formattedResponse, err := convertToResponse(res)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "cannot format response",
		})
		return
	}
	c.JSON(http.StatusOK, formattedResponse)

}
