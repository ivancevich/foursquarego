package foursquarego

import (
	"fmt"
	"net/http"
	"net/url"
)

type ApiError struct {
	StatusCode int
	Meta       Meta
	Url        *url.URL
}

func newApiError(resp *http.Response, meta Meta) *ApiError {
	return &ApiError{
		StatusCode: resp.StatusCode,
		Meta:       meta,
		Url:        resp.Request.URL,
	}
}

func (a ApiError) Error() string {
	return fmt.Sprintf(
		"Get %s returned status %d. Code: %d, %s: %s",
		a.Url,
		a.StatusCode,
		a.Meta.Code,
		a.Meta.ErrorDetail,
		a.Meta.ErrorType)
}
