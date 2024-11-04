package validation

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"gobarber/pkg/errorwrapper"
	"net/http"
)

func CustomValidationError(err error) *errorwrapper.ApiError {
	var ve validator.ValidationErrors
	apiErr := errorwrapper.New(
		"validation-error",
		"Erro ao validar parâmetros da requisição",
	).WithStatus(http.StatusUnprocessableEntity)

	if errors.As(err, &ve) {
		for _, fe := range ve {
			_ = apiErr.WithField(fe.Field(), msgForTag(fe.Tag(), fe.Field()))
		}
	}

	return apiErr
}

func msgForTag(tag string, field string) string {
	switch tag {
	case "required":
		return fmt.Sprintf("o campo %s é obrigatório.", field)
	case "email":
		return "email inválido."
	}
	return ""
}
