package api

import (
	"github.com/go-playground/validator/v10"
	"unico/internal/unico/database"
)

type controlerHandler struct {
	db       database.Database
	validate *validator.Validate
}
