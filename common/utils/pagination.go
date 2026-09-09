package utils

import (
	"math"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

// PaginationMeta metadata for enterprise paginated queries
type PaginationMeta struct {
	CurrentPage int   `json:"current_page"`
	PerPage     int   `json:"per_page"`
	TotalItems  int64 `json:"total_items"`
	TotalPages  int   `json:"total_pages"`
	HasNext     bool  `json:"has_next"`
	HasPrev     bool  `json:"has_prev"`
}

// PaginatedResponse wrapper for paginated endpoints
type PaginatedResponse struct {
	Message    string         `json:"message"`
	Data       interface{}    `json:"data"`
	Pagination PaginationMeta `json:"pagination"`
}

// GetPaginationQuery extracts and sanitizes page, limit, offset, and search from echo context
func GetPaginationQuery(c echo.Context) (page int, limit int, offset int, search string) {
	pageStr := c.QueryParam("page")
	limitStr := c.QueryParam("limit")
	search = strings.TrimSpace(c.QueryParam("search"))

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err = strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	} else if limit > 100 {
		limit = 100 // Hard cap to protect server memory
	}

	offset = (page - 1) * limit
	return page, limit, offset, search
}

// BuildPaginationMeta constructs pagination metadata
func BuildPaginationMeta(totalItems int64, page int, limit int) PaginationMeta {
	totalPages := int(math.Ceil(float64(totalItems) / float64(limit)))
	if totalPages < 1 {
		totalPages = 1
	}

	return PaginationMeta{
		CurrentPage: page,
		PerPage:     limit,
		TotalItems:  totalItems,
		TotalPages:  totalPages,
		HasNext:     page < totalPages,
		HasPrev:     page > 1,
	}
}

// SendPaginatedSuccess sends an enterprise JSON response with pagination meta
func SendPaginatedSuccess(c echo.Context, statusCode int, message string, data interface{}, meta PaginationMeta) error {
	return c.JSON(statusCode, PaginatedResponse{
		Message:    message,
		Data:       data,
		Pagination: meta,
	})
}
