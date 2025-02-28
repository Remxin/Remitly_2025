package parser

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func ParseXLSXToJSON(path string) ([]map[string]string, error) {
	f, err := excelize.OpenFile(path)
	if err != nil {
		return nil, fmt.Errorf("cannot open .xlsx file")
	}
	defer f.Close()

	sheets := f.GetSheetList()
	if len(sheets) == 0 {
		return nil, fmt.Errorf("no sheets found in .xlsx file")
	}
	swiftSheet := sheets[0]
	rows, err := f.GetRows(swiftSheet)
	if err != nil {
		return nil, err
	}

	if len(rows) == 0 {
		return nil, fmt.Errorf("no data found in sheet")
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
