package errors

// This package implements the RFC 7807 standard for problem details in HTTP APIs.
// More information can be found here: https://datatracker.ietf.org/doc/html/rfc7807
// Helper functions are provided to create common HTTP error responses.

import (
	"net/http"
)

// ErrorDTO represents a standardized error response according to RFC 7807 (problem details).
// To create extensions on the error DTO, embed this struct in another struct.
// Example:
//
//	type CustomErrorDTO struct {
//	   *ErrorDTO
//	   CustomField string `json:"customField"`
//	}
type ErrorDTO struct {
	Type     string `json:"type" binding:"required"`
	Title    string `json:"title" binding:"required"`
	Status   int    `json:"status" binding:"required"`
	Detail   string `json:"detail" binding:"required"`
	Instance string `json:"instance"`
	Code     string `json:"code"`
}

func NewError(errType string, title string, status int, detail, instance string, code string) *ErrorDTO {
	return &ErrorDTO{
		Type:     errType,
		Title:    title,
		Status:   status,
		Detail:   detail,
		Instance: instance,
		Code: 	  code,
	}
}

func NewErrorWithoutType(title string, status int, detail, instance string, code string) *ErrorDTO {
	return &ErrorDTO{
		Type:     "about:blank",
		Title:    title,
		Status:   status,
		Detail:   detail,
		Instance: instance,
		Code: 	  code,
	}
}

// 400 — Bad Request
func NewBadRequestError(title string, detail string, code string) *ErrorDTO {
	return &ErrorDTO{
		Type:   UrlBadRequest,
		Title:  title,
		Status: http.StatusBadRequest,
		Detail: detail,
		Code:   code,
	}
}

// 401 — Unauthorized
func NewUnauthorizedError(title string, detail string, code string) *ErrorDTO {
	return &ErrorDTO{
		Type:   UrlUnauthorized,
		Title:  title,
		Status: http.StatusUnauthorized,
		Detail: detail,
		Code:   code,
	}
}

// 403 — Forbidden
func NewForbiddenError(title string, detail string, code string) *ErrorDTO {
	return &ErrorDTO{
		Type:   UrlForbidden,
		Title:  title,
		Status: http.StatusForbidden,
		Detail: detail,
		Code:   code,
	}
}

// 404 — Not Found
func NewNotFoundError(title string, detail string, code string) *ErrorDTO {
	return &ErrorDTO{
		Type:   UrlNotFound,
		Title:  title,
		Status: http.StatusNotFound,
		Detail: detail,
		Code:   code,
	}
}

// 500 — Internal Server Error
func NewStandardInternalServerError(detail string) *ErrorDTO {
	return &ErrorDTO{
		Type:   UrlInternalServerError,
		Title:  TitleInternalServerError,
		Status: http.StatusInternalServerError,
		Detail: detail,
		Code:  "500",
	}
}
func NewInternalServerError(title string, detail string) *ErrorDTO {
	return &ErrorDTO{
		Type:   UrlInternalServerError,
		Title:  title,
		Status: http.StatusInternalServerError,
		Detail: detail,
		Code:   "500",
	}
}

func NewErrorWithInstance(err *ErrorDTO, instance string) *ErrorDTO {
	err.Instance = instance
	return err
}

const (
	UrlBadRequest          = "https://datatracker.ietf.org/doc/html/rfc7231#section-6.5.1"
	UrlUnauthorized        = "https://datatracker.ietf.org/doc/html/rfc7235#section-3.1"
	UrlForbidden           = "https://datatracker.ietf.org/doc/html/rfc7231#section-6.5.3"
	UrlNotFound            = "https://datatracker.ietf.org/doc/html/rfc7231#section-6.5.4"
	UrlInternalServerError = "https://datatracker.ietf.org/doc/html/rfc7231#section-6.6.1"
)

const (
	TitleBadRequest          = "Bad Request"
	TitleUnauthorized        = "Unauthorized"
	TitleForbidden           = "Forbidden"
	TitleNotFound            = "Not Found"
	TitleInternalServerError = "Internal Server Error"
)
