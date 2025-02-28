package api

import (
	db "example.com/m/v2/db/sqlc"
)

type Handler struct {
	Store db.Store
}
