package constants

type StatusCode int

const (
	StatusOK                  StatusCode = 200
	StatusCreated             StatusCode = 201
	StatusBadRequest          StatusCode = 400
	StatusNotFound            StatusCode = 404
	StatusInternalServerError StatusCode = 500
	StatusConflict            StatusCode = 409
	StatusUnauthorized        StatusCode = 401
	StatusForbidden           StatusCode = 403
)
