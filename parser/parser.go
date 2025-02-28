package parser

import (
	"context"
	"database/sql"
	"fmt"

	db "example.com/m/v2/db/sqlc"
	"github.com/xuri/excelize/v2"
)

func ParseXLSXToJSON(filePath string) ([]map[string]string, error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	sheets := f.GetSheetList()
	if len(sheets) == 0 {
		return nil, fmt.Errorf("no sheets found")
	}
	sheetName := sheets[0]

	rows, err := f.GetRows(sheetName)
	if err != nil {
		return nil, err
	}

	if len(rows) == 0 {
		return nil, fmt.Errorf("no data in sheet")
	}

	headers := rows[0]
	var jsonData []map[string]string

	for _, row := range rows[1:] {
		entry := make(map[string]string)
		for i, cell := range row {
			if i < len(headers) {
				entry[headers[i]] = cell
			}
		}

		jsonData = append(jsonData, entry)
	}
	return jsonData, nil
}

func AddRecordsToDatabase(ctx context.Context, pool *sql.DB, records []map[string]string) error {
	tx, err := pool.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return fmt.Errorf("failed to start transaction: %v", err)
	}
	defer tx.Rollback()

	queries := db.New(tx)

	for _, record := range records {
		swiftData := db.CreateSwiftDataParams{
			CountryIso2Code: record["COUNTRY ISO2 CODE"],
			SwiftCode:       record["SWIFT CODE"],
			CodeType:        record["CODE TYPE"],
			BankName:        record["NAME"],
			Address:         record["ADDRESS"],
			TownName:        record["TOWN NAME"],
			CountryName:     record["COUNTRY NAME"],
			TimeZone:        record["TIME ZONE"],
		}

		err := queries.CreateSwiftData(ctx, swiftData)
		if err != nil {
			return fmt.Errorf("failed to insert data: %v", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}
