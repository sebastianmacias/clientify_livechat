package httperrors

import (
	"fmt"
	"runtime"

	"github.com/go-openapi/swag"
)

func ErrXRef(fileCode string, httpError *HTTPError, code *int64, err error) *HTTPError {
	returnError := *httpError

	_, _, fileLine, ok := runtime.Caller(1)

	if ok {
		returnError.XRef = *swag.String(fmt.Sprintf("%s_%d", fileCode, fileLine))
	} else {
		returnError.XRef = *swag.String(fileCode)
	}

	if code != nil {
		returnError.Code = code
	} else {
		returnError.Code = swag.Int64(400)
	}

	if err != nil {
		errorMsg := err.Error()
		returnError.Detail = errorMsg
	}

	return &returnError
}
