package model

import "math"

type CustomerCreateResponse struct {
	ID       int `json:"id"`
	WalletID int `json:"wallet_id"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

// generic error for checking in presentation layer
// in the case of returning this error, http should be
// in 400 range
type ServiceError string

func (se ServiceError) Error() string {
	return string(se)
}

func (se ServiceError) Response() *ErrorResponse {
	return &ErrorResponse{Message: se.Error()}
}

type PaginationResponse[T any] struct {
	Data        []T   `json:"data"`
	Total       int64 `json:"total"`
	LastPage    int   `json:"last_page"`
	PageSize    int   `json:"page_size"`
	CurrentPage int   `json:"current_page"`
}

func CreatePagination[T any](data []T, total int64, page, pageSize int) *PaginationResponse[T] {
	return &PaginationResponse[T]{
		Total:       total,
		LastPage:    int(math.Ceil(float64(total) / float64(pageSize))),
		PageSize:    pageSize,
		CurrentPage: page,
		Data:        data,
	}
}
