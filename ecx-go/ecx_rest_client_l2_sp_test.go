package ecx

import (
	"context"
	"ecx-go/v3/internal/api"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

var testProfile = L2ServiceProfile{
	AlertPercentage:              30.2,
	AllowCustomSpeed:             true,
	AllowOverSubscription:        false,
	APIAvailable:                 true,
	AuthKeyLabel:                 "authKeyLabel",
	ConnectionNameLabel:          "connectionNameLabel",
	CTagLabel:                    "cTagLabel",
	EnableAutoGenerateServiceKey: false,
	EquinixManagedPortAndVlan:    false,
	Features: L2ServiceProfileFeatures{
		CloudReach:  true,
		TestProfile: true,
	},
	IntegrationID:                       "integrationID",
	Name:                                "name",
	OnBandwidthThresholdNotification:    []string{"miro@equinix.com", "jane@equinix.com"},
	OnProfileApprovalRejectNotification: []string{"miro@equinix.com", "jane@equinix.com"},
	OnVcApprovalRejectionNotification:   []string{"miro@equinix.com", "jane@equinix.com"},
	OverSubscription:                    "2x",
	Ports: []L2ServiceProfilePort{
		{
			ID:        "port-id1",
			MetroCode: "FR",
		}, {
			ID:        "port-id2",
			MetroCode: "AM",
		},
	},
	Private:            true,
	PrivateUserEmails:  []string{"miro@equinix.com", "jane@equinix.com"},
	RequiredRedundancy: false,
	SpeedBands: []L2ServiceProfileSpeedBand{
		{
			Speed:     100,
			SpeedUnit: "MB",
		}, {
			Speed:     1000,
			SpeedUnit: "MB",
		},
	},
	SpeedFromAPI:      false,
	TagType:           "tagType",
	VlanSameAsPrimary: false,
}

func TestGetL2ServiceProfile(t *testing.T) {
	//Given
	respBody := api.L2ServiceProfile{}
	if err := readJSONData("./test-fixtures/ecx_l2serviceprofile_get_resp.json", &respBody); err != nil {
		assert.Failf(t, "Cannont read test response due to %s", err.Error())
	}
	profileID := "spId"
	testHc := &http.Client{}
	httpmock.ActivateNonDefault(testHc)
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/ecx/v3/l2/serviceprofiles/%s", baseURL, profileID),
		func(r *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, respBody)
			return resp, nil
		},
	)
	//When
	ecxClient := NewClient(context.Background(), baseURL, testHc)
	prof, err := ecxClient.GetL2ServiceProfile(profileID)

	//Then
	assert.Nil(t, err, "Client should not return an error")
	assert.NotNil(t, prof, "Client should return a response")
	//verifyL2ServiceProfileResponse(t, serverResp, *resp)
	verifyL2ServiceProfile(t, *prof, respBody)
}

func TestCreateL2ServiceProfile(t *testing.T) {
	//Given
	respBody := api.CreateL2ServiceProfileResponse{}
	if err := readJSONData("./test-fixtures/ecx_l2serviceprofile_post_resp.json", &respBody); err != nil {
		assert.Failf(t, "Cannont read test response due to %s", err.Error())
	}
	reqBody := api.L2ServiceProfile{}
	testHc := &http.Client{}
	httpmock.ActivateNonDefault(testHc)
	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/ecx/v3/l2/serviceprofiles", baseURL),
		func(r *http.Request) (*http.Response, error) {
			if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
				return httpmock.NewStringResponse(400, ""), nil
			}
			resp, _ := httpmock.NewJsonResponse(200, respBody)
			return resp, nil
		},
	)
	defer httpmock.DeactivateAndReset()
	newProfile := testProfile

	//When
	ecxClient := NewClient(context.Background(), baseURL, testHc)
	prof, err := ecxClient.CreateL2ServiceProfile(newProfile)

	//Then
	assert.Nil(t, err, "Client should not return an error")
	assert.NotNil(t, prof, "Client should return a response")
	verifyL2ServiceProfile(t, *prof, reqBody)
}

func TestUpdateL2ServiceProfile(t *testing.T) {
	//Given
	respBody := api.CreateL2ServiceProfileResponse{}
	if err := readJSONData("./test-fixtures/ecx_l2serviceprofile_post_resp.json", &respBody); err != nil {
		assert.Failf(t, "Cannont read test response due to %s", err.Error())
	}
	reqBody := api.L2ServiceProfile{}
	testHc := &http.Client{}
	httpmock.ActivateNonDefault(testHc)
	httpmock.RegisterResponder("PUT", fmt.Sprintf("%s/ecx/v3/l2/serviceprofiles", baseURL),
		func(r *http.Request) (*http.Response, error) {
			if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
				return httpmock.NewStringResponse(400, ""), nil
			}
			resp, _ := httpmock.NewJsonResponse(200, respBody)
			return resp, nil
		},
	)
	defer httpmock.DeactivateAndReset()
	newProfile := testProfile
	newProfile.UUID = "someUUID"
	newProfile.State = "APPROVED"

	//When
	ecxClient := NewClient(context.Background(), baseURL, testHc)
	prof, err := ecxClient.UpdateL2ServiceProfile(newProfile)

	//Then
	assert.Nil(t, err, "Client should not return an error")
	assert.NotNil(t, prof, "Client should return a response")
	verifyL2ServiceProfileUpdate(t, *prof, reqBody)
}

func TestDeleteServiceProfile(t *testing.T) {
	//Given
	respBody := api.CreateL2ServiceProfileResponse{}
	if err := readJSONData("./test-fixtures/ecx_l2serviceprofile_delete_resp.json", &respBody); err != nil {
		assert.Failf(t, "Cannont read test response due to %s", err.Error())
	}
	profileID := "existingId"
	testHc := &http.Client{}
	httpmock.ActivateNonDefault(testHc)
	httpmock.RegisterResponder("DELETE", fmt.Sprintf("%s/ecx/v3/l2/serviceprofiles/%s", baseURL, profileID),
		func(r *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, respBody)
			return resp, nil
		})
	defer httpmock.DeactivateAndReset()

	//When
	ecxClient := NewClient(context.Background(), baseURL, testHc)
	err := ecxClient.DeleteL2ServiceProfile(profileID)

	//Then
	assert.Nil(t, err, "Client should not return an error")
}

func verifyL2ServiceProfile(t *testing.T, prof L2ServiceProfile, resp api.L2ServiceProfile) {
	assert.Equal(t, resp.AlertPercentage, prof.AlertPercentage, "AlertPercentage matches")
	assert.Equal(t, resp.AllowCustomSpeed, prof.AllowCustomSpeed, "AllowCustomSpeed matches")
	assert.Equal(t, resp.AllowOverSubscription, prof.AllowOverSubscription, "AllowOverSubscription matches")
	assert.Equal(t, resp.APIAvailable, prof.APIAvailable, "APIAvailable matches")
	assert.Equal(t, resp.AuthKeyLabel, prof.AuthKeyLabel, "AuthKeyLabel matches")
	assert.Equal(t, resp.ConnectionNameLabel, prof.ConnectionNameLabel, "ConnectionNameLabel matches")
	assert.Equal(t, resp.CTagLabel, prof.CTagLabel, "CTagLabel matches")
	assert.Equal(t, resp.EnableAutoGenerateServiceKey, prof.EnableAutoGenerateServiceKey, "EnableAutoGenerateServiceKey matches")
	assert.Equal(t, resp.EquinixManagedPortAndVlan, prof.EquinixManagedPortAndVlan, "EquinixManagedPortAndVlan matches")
	assert.Equal(t, resp.IntegrationID, prof.IntegrationID, "IntegrationID matches")
	assert.Equal(t, resp.Name, prof.Name, "Name matches")
	assert.ElementsMatch(t, resp.OnBandwidthThresholdNotification, prof.OnBandwidthThresholdNotification, "OnBandwidthThresholdNotification match")
	assert.ElementsMatch(t, resp.OnProfileApprovalRejectNotification, prof.OnProfileApprovalRejectNotification, "OnProfileApprovalRejectNotification match")
	assert.ElementsMatch(t, resp.OnVcApprovalRejectionNotification, prof.OnVcApprovalRejectionNotification, "OnVcApprovalRejectionNotification match")
	assert.Equal(t, resp.OverSubscription, prof.OverSubscription, "OverSubscription matches")
	assert.Equal(t, resp.Private, prof.Private, "Private matches")
	assert.ElementsMatch(t, resp.PrivateUserEmails, prof.PrivateUserEmails, "PrivateUserEmails match")
	assert.Equal(t, resp.RequiredRedundancy, prof.RequiredRedundancy, "RequiredRedundancy matches")
	assert.Equal(t, resp.SpeedFromAPI, prof.SpeedFromAPI, "SpeedFromAPI matches")
	assert.Equal(t, resp.TagType, prof.TagType, "TagType matches")
	assert.Equal(t, resp.VlanSameAsPrimary, prof.VlanSameAsPrimary, "VlanSameAsPrimary matches")

	assert.Equal(t, resp.Features.CloudReach, prof.Features.CloudReach, "Features.CloudReach matches")
	assert.Equal(t, resp.Features.TestProfile, prof.Features.TestProfile, "Features.TestProfile matches")

	for i := range prof.Ports {
		assert.Equal(t, resp.Ports[i].ID, prof.Ports[i].ID, fmt.Sprintf("Ports[%v].id matches", i))
		assert.Equal(t, resp.Ports[i].MetroCode, prof.Ports[i].MetroCode, fmt.Sprintf("Ports[%v].metroCode matches", i))
	}
	for i := range prof.SpeedBands {
		assert.Equal(t, resp.SpeedBands[i].Speed, prof.SpeedBands[i].Speed, fmt.Sprintf("SpeedBands[%v].Speed matches", i))
		assert.Equal(t, resp.SpeedBands[i].SpeedUnit, prof.SpeedBands[i].SpeedUnit, fmt.Sprintf("SpeedBands[%v].SpeedUnit matches", i))
	}
}

func verifyL2ServiceProfileResponse(t *testing.T, prof L2ServiceProfile, resp api.L2ServiceProfile) {
	assert.Equal(t, resp.UUID, prof.UUID, "UUID matches")
	assert.Equal(t, resp.UUID, prof.State, "State matches")
	verifyL2ServiceProfile(t, prof, resp)
}

func verifyL2ServiceProfileUpdate(t *testing.T, prof L2ServiceProfile, req api.L2ServiceProfile) {
	assert.Equal(t, prof.UUID, req.UUID, "UUID matches")
	assert.Equal(t, prof.State, req.State, "State matches")
	verifyL2ServiceProfile(t, prof, req)
}
