package ecx

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"github.com/equinix/ecx-go/internal/api"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGetUserPort(t *testing.T) {
	//Given
	respBody := []api.Port{}
	if err := readJSONData("./test-fixtures/ecx_ports_get.json", &respBody); err != nil {
		assert.Failf(t, "Cannont read test response due to %s", err.Error())
	}
	testHc := &http.Client{}
	httpmock.ActivateNonDefault(testHc)
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/ecx/v3/port/userport", baseURL),
		func(r *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, respBody)
			return resp, nil
		},
	)
	//When
	ecxClient := NewClient(context.Background(), baseURL, testHc)
	ports, err := ecxClient.GetUserPorts()

	//Then
	assert.Nil(t, err, "Client should not return an error")
	assert.NotNil(t, ports, "Client should return a response")
	assert.Equal(t, len(respBody), len(ports), "Client returned valid number of ports")
	for i := range ports {
		verifyPort(t, ports[i], respBody[i])
	}
}

func verifyPort(t *testing.T, port Port, apiPort api.Port) {
	assert.Equal(t, apiPort.UUID, port.UUID, "UUID matches")
	assert.Equal(t, apiPort.Name, port.Name, "Name matches")
	assert.Equal(t, apiPort.Region, port.Region, "Region matches")
	assert.Equal(t, apiPort.IBX, port.IBX, "IBX matches")
	assert.Equal(t, apiPort.MetroCode, port.MetroCode, "MetroCode matches")
	assert.Equal(t, apiPort.DevicePriority, port.Priority, "DevicePriority matches")
	assert.Equal(t, apiPort.Encapsulation, port.Encapsulation, "Encapsulation matches")
	assert.Equal(t, apiPort.Buyout, port.Buyout, "Buyout matches")
	assert.Equal(t, apiPort.ProvisionStatus, port.Status, "ProvisionStatus matches")
	convBandwidth, err := strconv.ParseInt(port.Bandwidth, 10, 64)
	assert.Nil(t, err, "Bandwidth string should be convertable to int64")
	assert.Equal(t, apiPort.TotalBandwidth, convBandwidth, "TotalBandwidth matches")

}
