package errorwrapper

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApiError struct {
	ErrorKey string       `json:"errorKey"`
	Title    string       `json:"title"`
	Status   int          `json:"status"`
	Detail   string       `json:"detail,omitempty"`
	Fields   []FieldError `json:"fields,omitempty"`
}

type FieldError struct {
	Field string `json:"field"`
	Msg   string `json:"msg"`
}

func New(errorKey string, title string) *ApiError {
	return &ApiError{
		ErrorKey: errorKey,
		Title:    title,
		Status:   http.StatusInternalServerError,
	}
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("%s - %s", e.ErrorKey, e.Title)
}

func (e *ApiError) WithField(field string, msg string) *ApiError {
	if e.Fields == nil {
		e.Fields = []FieldError{}
	}
	e.Fields = append(e.Fields, FieldError{field, msg})
	return e
}

func (e *ApiError) WithDetail(detail string) *ApiError {
	e.Detail = detail
	return e
}

func (e *ApiError) WithStatus(status int) *ApiError {
	e.Status = status
	return e
}

func SendError(c *gin.Context, err error) {
	var apiErr *ApiError
	if ok := errors.As(err, &apiErr); !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	c.JSON(apiErr.Status, apiErr)
}
