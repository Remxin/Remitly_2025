package api

import (
	api "example.com/m/v2/api/handlers"
	db "example.com/m/v2/db/sqlc"
	"github.com/gin-gonic/gin"
)

func SetupRouter(store db.Store) *gin.Engine {
	handler := &api.Handler{Store: store}

	router := gin.Default()
	router.GET("/v1/swift-codes/:swift-code", handler.GetDetailsSwiftCode)
	router.GET("/v1/swift-codes/country/:countryISO2code", handler.GetCountryIsoDetails)
	router.POST("/v1/swift-codes", handler.AddSwiftCode)
	router.DELETE("/v1/swift-codes/:swift-code", handler.DeleteSwiftCode)
	router.Run(":8080")
	return router
}
