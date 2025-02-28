package main

import (
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// func main() {
// 	conn, err := sql.Open("postgres", "postgres://admin:secretpassword@localhost:5432/swift_db?sslmode=disable")
// 	if err != nil {
// 		log.Fatalf("unable to connect to database: %v\n", err)
// 	}
// 	defer conn.Close()

// 	filePath := "data/swift_codes.xlsx"

// 	records, err := parser.ParseXLSXToJSON(filePath)
// 	if err != nil {
// 		log.Fatalf("failed to parse XLSX file: %v\n", err)
// 	}
// 	err = parser.AddRecordsToDatabase(context.Background(), conn, records)
// 	if err != nil {
// 		log.Fatalf("failed to add records to database: %v\n", err)
// 	}

// 	fmt.Println("Records added successfully!")
// }
