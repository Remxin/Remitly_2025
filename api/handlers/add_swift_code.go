package api

import (
	"net/http"

	models "example.com/m/v2/api/models"
	db "example.com/m/v2/db/sqlc"
	"github.com/gin-gonic/gin"
)

type AddSwiftCodeRequest struct {
	Address       string `json:"address"`
	BankName      string `json:"bankName"`
	CountryISO2   string `json:"countryISO2"`
	CountryName   string `json:"countryName"`
	IsHeadquarter bool   `json:"isHeadquarter"`
	SwiftCode     string `json:"swiftCode"`
}

func validateAddSwiftCodeFields(body *AddSwiftCodeRequest) []models.FieldViolation {
	var fieldViolations []models.FieldViolation

	if len(body.CountryISO2) != 2 {
		fieldViolations = append(fieldViolations, models.FieldViolation{
			Field:   "countryISO2",
			Message: "countryISO2 code should have lenght of 2",
		})
	}
	if len(body.SwiftCode) != 11 {
		fieldViolations = append(fieldViolations, models.FieldViolation{
			Field:   "swiftCode",
			Message: "swiftCode should have lenght of 11",
		})
	}

	return fieldViolations
}

func (handler *Handler) AddSwiftCode(c *gin.Context) {
	var body AddSwiftCodeRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fieldViolations := validateAddSwiftCodeFields(&body)
	if len(fieldViolations) != 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":          "failed",
			"fieldViolations": fieldViolations,
		})
		return
	}

	_, err := handler.Store.AddNewSwiftCode(c, db.AddNewSwiftCodeParams{
		Address:         body.Address,
		BankName:        body.BankName,
		CountryIso2Code: body.CountryISO2,
		CountryName:     body.CountryName,
		SwiftCode:       body.SwiftCode,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
