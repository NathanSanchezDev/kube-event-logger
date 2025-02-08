package utils

import (
	"net/http"
	"strconv"
)

// ParsePagination extracts pagination parameters from the request
func ParsePagination(r *http.Request) (page int, limit int) {
	query := r.URL.Query()

	page, err := strconv.Atoi(query.Get("page"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err = strconv.Atoi(query.Get("limit"))
	if err != nil || limit < 1 {
		limit = 100
	}

	return page, limit
}
