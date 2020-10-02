package ecx

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/equinix/ecx-go/internal/api"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

var testPrimaryConnection = L2Connection{
	Name:                "name",
	ProfileUUID:         "profileUUID",
	Speed:               666,
	SpeedUnit:           "MB",
	Notifications:       []string{"janek@equinix.com", "marek@equinix.com"},
	PurchaseOrderNumber: "orderNumber",
	PortUUID:            "primaryPortUUID",
	VlanSTag:            100,
	VlanCTag:            101,
	NamedTag:            "Private",
	AdditionalInfo:      []L2ConnectionAdditionalInfo{{Name: "asn", Value: "1543"}, {Name: "global", Value: "false"}},
	ZSidePortUUID:       "primaryZSidePortUUID",
	ZSideVlanSTag:       200,
	ZSideVlanCTag:       201,
	SellerRegion:        "EMEA",
	SellerMetroCode:     "AM",
	AuthorizationKey:    "authorizationKey"}

func TestGetL2OutgoingConnections(t *testing.T) {
	//Given
	respBody := api.L2BuyerConnectionsResponse{}
	if err := readJSONData("./test-fixtures/ecx_l2connections_get_resp.json", &respBody); err != nil {
		assert.Failf(t, "Cannot read test response due to %s", err.Error())
	}
	pageSize := respBody.PageSize
	testHc := &http.Client{}
	httpmock.ActivateNonDefault(testHc)
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/ecx/v3/l2/buyer/connections?pageSize=%d&status=%s", baseURL, pageSize, url.QueryEscape("PROVISIONED,PROVISIONING")),
		func(r *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, respBody)
			return resp, nil
		},
	)
	defer httpmock.DeactivateAndReset()

	//When
	ecxClient := NewClient(context.Background(), baseURL, testHc)
	ecxClient.SetPageSize(pageSize)
	conns, err := ecxClient.GetL2OutgoingConnections([]string{ConnectionStatusProvisioned, ConnectionStatusProvisioning})

	//Then
	assert.Nil(t, err, "Client should not return an error")
	assert.NotNil(t, conns, "Client should return a response")
	assert.Equal(t, len(respBody.Content), len(conns), "Number of connections matches")
	for i := range respBody.Content {
		verifyL2Connection(t, conns[i], respBody.Content[i])
	}
}

func TestGetL2Connection(t *testing.T) {
	//Given
	respBody := api.L2ConnectionResponse{}
	if err := readJSONData("./test-fixtures/ecx_l2connection_get_resp.json", &respBody); err != nil {
		assert.Failf(t, "Cannot read test response due to %s", err.Error())
	}
	connID := "connId"
	testHc := &http.Client{}
	httpmock.ActivateNonDefault(testHc)
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/ecx/v3/l2/connections/%s", baseURL, connID),
		func(r *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, respBody)
			return resp, nil
		},
	)
	defer httpmock.DeactivateAndReset()

	//When
	ecxClient := NewClient(context.Background(), baseURL, testHc)
	conn, err := ecxClient.GetL2Connection(connID)

	//Then
	assert.Nil(t, err, "Client should not return an error")
	assert.NotNil(t, conn, "Client should return a response")
	verifyL2Connection(t, *conn, respBody)
}

func TestCreateL2Connection(t *testing.T) {
	//Given
	respBody := api.CreateL2ConnectionResponse{}
	if err := readJSONData("./test-fixtures/ecx_l2connection_post_resp.json", &respBody); err != nil {
		assert.Failf(t, "Cannot read test response due to %s", err.Error())
	}
	reqBody := api.L2ConnectionRequest{}
	testHc := &http.Client{}
	httpmock.ActivateNonDefault(testHc)
	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/ecx/v3/l2/connections", baseURL),
		func(r *http.Request) (*http.Response, error) {
			if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
				return httpmock.NewStringResponse(400, ""), nil
			}
			resp, _ := httpmock.NewJsonResponse(200, respBody)
			return resp, nil
		},
	)
	defer httpmock.DeactivateAndReset()
	newConnection := testPrimaryConnection

	//When
	ecxClient := NewClient(context.Background(), baseURL, testHc)
	conn, err := ecxClient.CreateL2Connection(newConnection)

	//Then
	assert.Nil(t, err, "Client should not return an error")
	assert.NotNil(t, conn, "Client should return a response")
	verifyL2ConnectionRequest(t, *conn, reqBody)
	assert.Equal(t, conn.UUID, respBody.PrimaryConnectionID, "UUID matches")
}

func TestCreateDeviceL2Connection(t *testing.T) {
	//Given
	respBody := api.CreateL2ConnectionResponse{}
	if err := readJSONData("./test-fixtures/ecx_l2connection_post_resp.json", &respBody); err != nil {
		assert.Failf(t, "Cannot read test response due to %s", err.Error())
	}
	reqBody := api.L2ConnectionRequest{}
	testHc := &http.Client{}
	httpmock.ActivateNonDefault(testHc)
	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/ne/v1/l2/connections", baseURL),
		func(r *http.Request) (*http.Response, error) {
			if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
				return httpmock.NewStringResponse(400, ""), nil
			}
			resp, _ := httpmock.NewJsonResponse(200, respBody)
			return resp, nil
		},
	)
	defer httpmock.DeactivateAndReset()
	newConnection := testPrimaryConnection
	newConnection.DeviceUUID = "deviceUUID"

	//When
	ecxClient := NewClient(context.Background(), baseURL, testHc)
	conn, err := ecxClient.CreateL2Connection(newConnection)

	//Then
	assert.Nil(t, err, "Client should not return an error")
	assert.NotNil(t, conn, "Client should return a response")
	verifyL2ConnectionRequest(t, *conn, reqBody)
	assert.Equal(t, conn.UUID, respBody.PrimaryConnectionID, "UUID matches")
}

func TestCreateRedundantL2Connection(t *testing.T) {
	//Given
	respBody := api.CreateL2ConnectionResponse{}
	if err := readJSONData("./test-fixtures/ecx_l2connection_post_resp.json", &respBody); err != nil {
		assert.Failf(t, "Cannot read test response due to %s", err.Error())
	}
	reqBody := api.L2ConnectionRequest{}
	testHc := &http.Client{}
	httpmock.ActivateNonDefault(testHc)
	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/ecx/v3/l2/connections", baseURL),
		func(r *http.Request) (*http.Response, error) {
			if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
				return httpmock.NewStringResponse(400, ""), nil
			}
			resp, _ := httpmock.NewJsonResponse(200, respBody)
			return resp, nil
		},
	)
	defer httpmock.DeactivateAndReset()
	newPriConn := testPrimaryConnection
	newSecConn := L2Connection{
		Name:          "secName",
		PortUUID:      "secondaryPortUUID",
		DeviceUUID:    "secondaryDeviceUUID",
		VlanSTag:      690,
		VlanCTag:      691,
		ZSidePortUUID: "secondaryZSidePortUUID",
		ZSideVlanSTag: 717,
		ZSideVlanCTag: 718}

	//When
	ecxClient := NewClient(context.Background(), baseURL, testHc)
	conn, err := ecxClient.CreateL2RedundantConnection(newPriConn, newSecConn)

	//Then
	assert.Nil(t, err, "Client should not return an error")
	assert.NotNil(t, conn, "Client should return a response")
	verifyRedundantL2ConnectionRequest(t, newPriConn, newSecConn, reqBody)
	assert.Equal(t, conn.UUID, respBody.PrimaryConnectionID, "UUID matches")
	assert.Equal(t, conn.RedundantUUID, respBody.SecondaryConnectionID, "RedundantUUID matches")
}

func TestDeleteL2Connection(t *testing.T) {
	//Given
	respBody := api.DeleteL2ConnectionResponse{}
	if err := readJSONData("./test-fixtures/ecx_l2connection_delete_resp.json", &respBody); err != nil {
		assert.Failf(t, "Cannot read test response due to %s", err.Error())
	}
	connID := "connId"
	testHc := &http.Client{}
	httpmock.ActivateNonDefault(testHc)
	httpmock.RegisterResponder("DELETE", fmt.Sprintf("%s/ecx/v3/l2/connections/%s", baseURL, connID),
		func(r *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, respBody)
			return resp, nil
		})
	defer httpmock.DeactivateAndReset()

	//When
	ecxClient := NewClient(context.Background(), baseURL, testHc)
	err := ecxClient.DeleteL2Connection(connID)

	//Then
	assert.Nil(t, err, "Client should not return an error")
}

func TestUpdateL2Connection(t *testing.T) {
	//Given
	respBody := api.L2ConnectionUpdateResponse{}
	if err := readJSONData("./test-fixtures/ecx_l2connection_update_resp.json", &respBody); err != nil {
		assert.Failf(t, "Cannot read test response due to %s", err.Error())
	}
	connID := "connId"
	newName := "newConnName"
	newSpeed := 500
	newSpeedUnit := "MB"
	reqBody := api.L2ConnectionUpdateRequest{}
	testHc := &http.Client{}
	httpmock.ActivateNonDefault(testHc)
	httpmock.RegisterResponder("PATCH", fmt.Sprintf("%s/ecx/v3/l2/connections/%s?action=update", baseURL, connID),
		func(r *http.Request) (*http.Response, error) {
			if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
				return httpmock.NewStringResponse(400, ""), nil
			}
			resp, _ := httpmock.NewJsonResponse(200, respBody)
			return resp, nil
		},
	)
	defer httpmock.DeactivateAndReset()

	//When
	c := NewClient(context.Background(), baseURL, testHc)
	err := c.NewL2ConnectionUpdateRequest(connID).
		WithName(newName).
		WithBandwidth(newSpeed, newSpeedUnit).
		Execute()

	//Then
	assert.Nil(t, err, "Client should not return an error")
	assert.Equal(t, newName, reqBody.Name, "Name matches")
	assert.Equal(t, newSpeed, reqBody.Speed, "Speed matches")
	assert.Equal(t, newSpeedUnit, reqBody.SpeedUnit, "SpeedUnit matches")
}

func verifyL2Connection(t *testing.T, conn L2Connection, resp api.L2ConnectionResponse) {
	assert.Equal(t, resp.UUID, conn.UUID, "UUID matches")
	assert.Equal(t, resp.Name, conn.Name, "Name matches")
	assert.Equal(t, resp.SellerServiceUUID, conn.ProfileUUID, "Name matches")
	assert.Equal(t, resp.Speed, conn.Speed, "Speed matches")
	assert.Equal(t, resp.SpeedUnit, conn.SpeedUnit, "SpeedUnit matches")
	assert.Equal(t, resp.Status, conn.Status, "Status matches")
	assert.Equal(t, resp.ProviderStatus, conn.ProviderStatus, "ProviderStatus matches")
	assert.ElementsMatch(t, resp.Notifications, conn.Notifications, "Notifications match")
	assert.Equal(t, resp.PurchaseOrderNumber, conn.PurchaseOrderNumber, "PurchaseOrderNumber match")
	assert.Equal(t, resp.PortUUID, conn.PortUUID, "PrimaryPortUUID matches")
	assert.Equal(t, resp.VirtualDeviceUUID, conn.DeviceUUID, "VirtualDeviceUUID matches")
	assert.Equal(t, resp.VlanSTag, conn.VlanSTag, "PrimaryVlanSTag matches")
	assert.Equal(t, resp.VlanCTag, conn.VlanCTag, "PrimaryVlanCTag matches")
	assert.Equal(t, resp.NamedTag, conn.NamedTag, "NamedTag matches")
	assert.Equal(t, resp.ZSidePortUUID, conn.ZSidePortUUID, "PrimaryZSidePortUUID matches")
	assert.Equal(t, resp.ZSideVlanSTag, conn.ZSideVlanSTag, "PrimaryZSideVlanSTag matches")
	assert.Equal(t, resp.ZSideVlanCTag, conn.ZSideVlanCTag, "PrimaryZSideVlanCTag matches")
	assert.Equal(t, resp.SellerMetroCode, conn.SellerMetroCode, "SellerMetroCode matches")
	assert.Equal(t, resp.AuthorizationKey, conn.AuthorizationKey, "AuthorizationKey matches")
	assert.Equal(t, resp.RedundantUUID, conn.RedundantUUID, "RedundantUUID key matches")
	assert.Equal(t, resp.RedundancyType, conn.RedundancyType, "RedundancyType matches")
	assert.Equal(t, len(resp.AdditionalInfo), len(conn.AdditionalInfo), "AdditionalInfo array size matches")
	for i := range resp.AdditionalInfo {
		verifyL2ConnectionAdditionalInfo(t, conn.AdditionalInfo[i], resp.AdditionalInfo[i])
	}
}

func verifyL2ConnectionRequest(t *testing.T, conn L2Connection, req api.L2ConnectionRequest) {
	assert.Equal(t, conn.Name, req.PrimaryName, "Name matches")
	assert.Equal(t, conn.ProfileUUID, req.ProfileUUID, "ProfileUUID matches")
	assert.Equal(t, conn.Speed, req.Speed, "Speed matches")
	assert.Equal(t, conn.SpeedUnit, req.SpeedUnit, "SpeedUnit matches")
	assert.ElementsMatch(t, conn.Notifications, req.Notifications, "Notifications match")
	assert.Equal(t, conn.PurchaseOrderNumber, req.PurchaseOrderNumber, "PurchaseOrderNumber matches")
	assert.Equal(t, conn.PortUUID, req.PrimaryPortUUID, "PrimaryPortUUID matches")
	assert.Equal(t, conn.DeviceUUID, req.VirtualDeviceUUID, "VirtualDeviceUUID matches")
	assert.Equal(t, conn.VlanSTag, req.PrimaryVlanSTag, "PrimaryVlanSTag matches")
	assert.Equal(t, conn.VlanCTag, req.PrimaryVlanCTag, "PrimaryVlanCTag matches")
	assert.Equal(t, conn.NamedTag, req.NamedTag, "NamedTag matches")
	assert.Equal(t, conn.ZSidePortUUID, req.PrimaryZSidePortUUID, "PrimaryZSidePortUUID matches")
	assert.Equal(t, conn.ZSideVlanSTag, req.PrimaryZSideVlanSTag, "PrimaryZSideVlanSTag matches")
	assert.Equal(t, conn.ZSideVlanCTag, req.PrimaryZSideVlanCTag, "PrimaryZSideVlanCTag matches")
	assert.Equal(t, conn.SellerRegion, req.SellerRegion, "SellerRegion matches")
	assert.Equal(t, conn.SellerMetroCode, req.SellerMetroCode, "SellerMetroCode matches")
	assert.Equal(t, conn.AuthorizationKey, req.AuthorizationKey, "Authorization key matches")

	assert.Equal(t, len(conn.AdditionalInfo), len(req.AdditionalInfo), "AdditionalInfo array size matches")
	for i := range conn.AdditionalInfo {
		verifyL2ConnectionAdditionalInfo(t, conn.AdditionalInfo[i], req.AdditionalInfo[i])
	}
}

func verifyRedundantL2ConnectionRequest(t *testing.T, primary L2Connection, secondary L2Connection, req api.L2ConnectionRequest) {
	verifyL2ConnectionRequest(t, primary, req)
	assert.Equal(t, secondary.Name, req.SecondaryName, "SecondaryName matches")
	assert.Equal(t, secondary.PortUUID, req.SecondaryPortUUID, "SecondaryPortUUID matches")
	assert.Equal(t, secondary.DeviceUUID, req.SecondaryVirtualDeviceUUID, "SecondaryVirtualDeviceUUID matches")
	assert.Equal(t, secondary.VlanSTag, req.SecondaryVlanSTag, "SecondaryVlanSTag matches")
	assert.Equal(t, secondary.VlanCTag, req.SecondaryVlanCTag, "SecondaryVlanCTag matches")
	assert.Equal(t, secondary.ZSidePortUUID, req.SecondaryZSidePortUUID, "SecondaryZSidePortUUID matches")
	assert.Equal(t, secondary.ZSideVlanSTag, req.SecondaryZSideVlanSTag, "SecondaryZSideVlanSTag matches")
	assert.Equal(t, secondary.ZSideVlanCTag, req.SecondaryZSideVlanCTag, "SecondaryZSideVlanCTag matches")
}

func verifyL2ConnectionAdditionalInfo(t *testing.T, info L2ConnectionAdditionalInfo, apiInfo api.L2ConnectionAdditionalInfo) {
	assert.Equal(t, info.Name, apiInfo.Name, "Name matches")
	assert.Equal(t, info.Value, apiInfo.Value, "Value matches")
}
