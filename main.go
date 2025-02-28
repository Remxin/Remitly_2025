package main

import (
	"database/sql"
	"log"

	api "example.com/m/v2/api/handlers"
	db "example.com/m/v2/db/sqlc"
	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	conn, err := sql.Open("postgres", "postgres://admin:secretpassword@localhost:5432/swift_db?sslmode=disable")
	if err != nil {
		log.Fatalf("unable to connect to database: %v\n", err)
	}
	defer conn.Close()
	store := db.NewStore(conn)
	handler := &api.Handler{Store: store}

	router := gin.Default()
	router.GET("/v1/swift-codes/:swift-code", handler.GetDetailsSwiftCode)
	router.GET("/v1/swift-codes/country/:countryISO2code", handler.GetCountryIsoDetails)
	router.POST("/v1/swift-codes", handler.AddSwiftCode)
	router.DELETE("/v1/swift-codes/:swift-code", handler.DeleteSwiftCode)
	router.Run(":8080")
}
