package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	api "example.com/m/v2/api/handlers"
	db "example.com/m/v2/db/sqlc"
	"example.com/m/v2/parser"
	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func ExtractXLSXDataToDB(conn *sql.DB) error {
	filePath := "data/swift_codes.xlsx"

	records, err := parser.ParseXLSXToJSON(filePath)
	if err != nil {
		return fmt.Errorf("failed to parse XLSX file: %v", err)
	}
	err = parser.AddRecordsToDatabase(context.Background(), conn, records)
	if err != nil {
		return fmt.Errorf("failed to add records to database: %v", err)
	}

	return nil
}

func main() {
	conn, err := sql.Open("postgres", "postgres://admin:secretpassword@localhost:5432/swift_db?sslmode=disable")
	if err != nil {
		log.Fatalf("unable to connect to database: %v\n", err)
	}
	defer conn.Close()
	err = ExtractXLSXDataToDB(conn)
	if err != nil {
		log.Fatalf("unable to parse .xlsx data: %v", err)
	}
	log.Print("success: imported .xlsx data")

	store := db.NewStore(conn)
	handler := &api.Handler{Store: store}

	router := gin.Default()
	router.GET("/v1/swift-codes/:swift-code", handler.GetDetailsSwiftCode)
	router.GET("/v1/swift-codes/country/:countryISO2code", handler.GetCountryIsoDetails)
	router.POST("/v1/swift-codes", handler.AddSwiftCode)
	router.DELETE("/v1/swift-codes/:swift-code", handler.DeleteSwiftCode)
	router.Run(":8080")
}
