package ecx

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/equinix/ecx-go/v2/internal/api"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

var testConnectionToConfirm = L2ConnectionToConfirm{
	AccessKey: String("accessKey"),
	SecretKey: String("secretKey"),
}

func TestConfirmL2Connection(t *testing.T) {
	//Given
	respBody := api.ConfirmL2ConnectionResponse{}
	if err := readJSONData("./test-fixtures/ecx_l2connection_patch_resp.json", &respBody); err != nil {
		assert.Failf(t, "Cannot read test response due to %s", err.Error())
	}
	connID := "connId"
	reqBody := api.ConfirmL2ConnectionRequest{}
	testHc := &http.Client{}
	httpmock.ActivateNonDefault(testHc)
	httpmock.RegisterResponder("PATCH", fmt.Sprintf("%s/ecx/v3/l2/connections/%s?action=Approve", baseURL, connID),
		func(r *http.Request) (*http.Response, error) {
			if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
				return httpmock.NewStringResponse(400, ""), nil
			}
			resp, _ := httpmock.NewJsonResponse(200, respBody)
			return resp, nil
		},
	)
	defer httpmock.DeactivateAndReset()
	connectionToConfirm := testConnectionToConfirm

	//When
	ecxClient := NewClient(context.Background(), baseURL, testHc)
	confirmation, err := ecxClient.ConfirmL2Connection(connID, connectionToConfirm)

	//Then
	assert.Nil(t, err, "Client should not return an error")
	assert.NotNil(t, confirmation, "Client should return a response")
	assert.Equal(t, confirmation.PrimaryConnectionID, respBody.PrimaryConnectionID, "UUID matches")
}
