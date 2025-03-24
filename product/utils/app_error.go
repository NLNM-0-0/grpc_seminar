package utils

import (
	"errors"
	"fmt"
	"net/http"
)

type AppError struct {
	StatusCode int    `json:"statusCode" example:"400"`
	RootErr    error  `json:"-"`
	Message    string `json:"message" example:"error message"`
	Log        string `json:"log" example:"error log"`
	Key        string `json:"errorKey" example:"ErrKey"`
}

func NewErrorResponse(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewCustomError(root error, msg, key string) *AppError {
	if root != nil {
		return NewErrorResponse(root, msg, root.Error(), key)
	}

	return NewErrorResponse(errors.New(msg), msg, msg, key)
}

func (e *AppError) RootError() error {
	var err *AppError
	if errors.As(e.RootErr, &err) {
		return err.RootError()
	}
	return e.RootErr
}

func (e *AppError) Error() string {
	return e.RootError().Error()
}

func ErrDB(err error) *AppError {
	return NewErrorResponse(
		err,
		"An error occurred with the database, please try again later",
		err.Error(),
		DB_ERROR_CODE,
	)
}

func ErrDuplicateDB(err error, field string) *AppError {
	return NewErrorResponse(
		err,
		fmt.Sprintf("A duplicate entry error occurred in the database for field '%s'.", field),
		err.Error(),
		DB_DUPLICATE_ERROR_CODE,
	)
}

func ErrRecordNotFound(err error, field string) *AppError {
	return NewErrorResponse(
		err,
		fmt.Sprintf("Could not find a record with field '%s'", field),
		err.Error(),
		DB_RECORD_NOT_FOUND_CODE,
	)
}

var (
	DB_RECORD_NOT_FOUND_CODE = "DB_RECORD_NOT_FOUND"
	DB_DUPLICATE_ERROR_CODE  = "DB_DUPLICATE_ERROR"
	DB_ERROR_CODE            = "DB_ERROR"
)
