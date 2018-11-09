package utils

import (
	"net/http"
	"strconv"

	"github.com/pkg/errors"
)

// GetURLQueryValueInt64 ...
func GetURLQueryValueInt64(queryKey string, defaultValue int64, r *http.Request) (int64, error) {
	retVal := defaultValue
	queryValueStr := r.URL.Query().Get(queryKey)
	if queryValueStr != "" {
		queryConvertedValue, err := strconv.ParseInt(queryValueStr, 10, 64)
		if err != nil {
			return 0, errors.Wrap(err, "Failed to convert query value")
		}
		retVal = queryConvertedValue
	}
	return retVal, nil
}

// GetURLQueryValueUint64 ...
func GetURLQueryValueUint64(queryKey string, defaultValue uint64, r *http.Request) (uint64, error) {
	retVal := defaultValue
	queryValueStr := r.URL.Query().Get(queryKey)
	if queryValueStr != "" {
		queryConvertedValue, err := strconv.ParseUint(queryValueStr, 10, 64)
		if err != nil {
			return 0, errors.Wrap(err, "Failed to convert query value")
		}
		retVal = queryConvertedValue
	}
	return retVal, nil
}
