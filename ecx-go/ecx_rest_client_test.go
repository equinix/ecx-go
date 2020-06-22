package ecx

import (
	"context"
	"ecx-go/v3/internal/api"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestSingleError(t *testing.T) {
	//given
	resp := api.ErrorResponse{}
	if err := readJSONData("./test-fixtures/ecx_httpapi_error_generic.json", &resp); err != nil {
		assert.Fail(t, "Cannont read test response")
	}
	testURL := "http://localhost:8888"
	testHc := &http.Client{}
	httpmock.ActivateNonDefault(testHc)
	httpmock.RegisterResponder("GET", testURL,
		func(r *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(500, resp)
			return resp, nil
		},
	)
	defer httpmock.DeactivateAndReset()

	//when
	cli := NewClient(context.Background(), baseURL, testHc)
	err := cli.execute(cli.R(), resty.MethodGet, testURL)

	//then
	assert.NotNil(t, err, "Error should be returned")
	assert.IsType(t, RestError{}, err, "Error should be RestError type")
	restErr := err.(RestError)
	assert.Equal(t, 500, restErr.HTTPCode, "RestError should have valid httpCode")
	assert.Equal(t, 1, len(restErr.Errors), "RestError should have one domain error")
	neError := restErr.Errors[0]
	assert.Equal(t, resp.ErrorCode, neError.ErrorCode, "RestError domain error code matches")
	assert.Equal(t, resp.ErrorMessage, neError.ErrorMessage, "RestError domain error message matches")
}

func TestMultipleError(t *testing.T) {
	//given
	resp := api.ErrorResponses{}
	if err := readJSONData("./test-fixtures/ecx_httpapi_multierror_generic.json", &resp); err != nil {
		assert.Fail(t, "Cannont read test response")
	}
	testURL := "http://localhost:8888"
	testHc := &http.Client{}
	httpmock.ActivateNonDefault(testHc)
	httpmock.RegisterResponder("GET", testURL,
		func(r *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(500, resp)
			return resp, nil
		},
	)
	defer httpmock.DeactivateAndReset()

	//when
	cli := NewClient(context.Background(), testURL, testHc)
	err := cli.execute(cli.R(), resty.MethodGet, testURL)

	//then
	assert.NotNil(t, err, "Error should be returned")
	assert.IsType(t, RestError{}, err, "Error should be RestError type")
	restErr := err.(RestError)
	assert.Equal(t, 500, restErr.HTTPCode, "RestError should have valid httpCode")
	assert.Equal(t, len(resp), len(restErr.Errors), "RestError should have valid number of domain errors")
}

func readJSONData(filePath string, target interface{}) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, target); err != nil {
		return err
	}
	return nil
}
