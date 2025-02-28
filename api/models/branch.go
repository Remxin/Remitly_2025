package api

type Branch struct {
	Address       string
	BankName      string
	CountryISO2   string
	isHeadquarter bool
	swiftCode     string
}
