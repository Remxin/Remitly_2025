package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	api "example.com/m/v2/api"
	db "example.com/m/v2/db/sqlc"
	"example.com/m/v2/parser"
	"example.com/m/v2/utils"
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
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatalf("unable to load confi file: %v\n", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
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
	api.SetupRouter(store)
}
