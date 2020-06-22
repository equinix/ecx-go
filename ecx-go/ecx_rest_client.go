package ecx

import (
	"context"
	"ecx-go/v3/internal/api"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

//RestClient describes ECX Fabric client that uses REST API
type RestClient struct {
	baseURL string
	ctx     context.Context
	*resty.Client
}

//RestError describe REST client error that occurs during operation processing
type RestError struct {
	HTTPCode int
	Message  string
	Errors   []Error
}

func (e RestError) Error() string {
	return fmt.Sprintf("ecx rest error: httpCode: %v, message: %v", e.HTTPCode, e.Message)
}

//NewClient creates new ECX REST API client with a given baseURL and http.Client
func NewClient(ctx context.Context, baseURL string, httpClient *http.Client) *RestClient {
	resty := resty.NewWithClient(httpClient)
	resty.SetHeader("User-agent", "equinix/ecx-go/v3")
	resty.SetHeader("Accept", "application/json")
	return &RestClient{
		baseURL,
		ctx,
		resty}
}

func (c RestClient) execute(req *resty.Request, method string, url string) error {
	resp, err := req.SetContext(c.ctx).Execute(method, url)
	if err != nil {
		restErr := RestError{Message: fmt.Sprintf("operation failed: %s", err)}
		if resp != nil {
			restErr.HTTPCode = resp.StatusCode()
		}
		return restErr
	}
	if resp.IsError() {
		err := transformErrorBody(resp.Body())
		err.HTTPCode = resp.StatusCode()
		return err
	}
	return nil
}

func transformErrorBody(body []byte) RestError {
	apiError := api.ErrorResponse{}
	if err := json.Unmarshal(body, &apiError); err == nil {
		return mapErrorAPIToDomain(apiError)
	}
	apiErrors := api.ErrorResponses{}
	if err := json.Unmarshal(body, &apiErrors); err == nil {
		return mapErrorsAPIToDomain(apiErrors)
	}
	return RestError{
		Message: string(body)}
}

func mapErrorAPIToDomain(apiError api.ErrorResponse) RestError {
	return RestError{
		Message: apiError.ErrorMessage,
		Errors:  []Error{{apiError.ErrorCode, apiError.ErrorMessage}},
	}
}

func mapErrorsAPIToDomain(apiErrors api.ErrorResponses) RestError {
	errors := make([]Error, len(apiErrors))
	for i, v := range apiErrors {
		errors[i] = Error{v.ErrorCode, v.ErrorMessage}
	}
	return RestError{
		Message: "Multiple errors occured",
		Errors:  errors,
	}
}
