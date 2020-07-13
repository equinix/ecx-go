package ecx

import (
	"ecx-go/v3/internal/api"
	"fmt"
	"net/url"

	"github.com/go-resty/resty/v2"
)

type restL2ConnectionUpdateRequest struct {
	uuid      string
	name      string
	speed     int
	speedUnit string
	c         RestClient
}

//GetL2Connection operation retrieves layer 2 connection with a given UUID
func (c RestClient) GetL2Connection(uuid string) (*L2Connection, error) {
	url := fmt.Sprintf("%s/ecx/v3/l2/connections/%s", c.baseURL, url.PathEscape(uuid))
	respBody := api.L2ConnectionResponse{}
	req := c.R().SetResult(&respBody)
	if err := c.execute(req, resty.MethodGet, url); err != nil {
		return nil, err
	}
	return mapGETToL2Connection(respBody), nil
}

//CreateL2Connection operation creates non-redundant layer 2 connection with a given connection structure.
//Upon successful creation, connection structure, enriched with assigned UUID, will be returned
func (c RestClient) CreateL2Connection(l2connection L2Connection) (*L2Connection, error) {
	url := fmt.Sprintf("%s/ecx/v3/l2/connections", c.baseURL)
	reqBody := createL2ConnectionRequest(l2connection)
	respBody := api.CreateL2ConnectionResponse{}
	req := c.R().SetBody(&reqBody).SetResult(&respBody)
	if err := c.execute(req, resty.MethodPost, url); err != nil {
		return nil, err
	}
	l2connection.UUID = respBody.PrimaryConnectionID
	return &l2connection, nil
}

//CreateL2RedundantConnection operation creates redundant layer2 connection with given connection structures.
//Primary connection structure is used as a baseline for underlaying API call, whereas secondary connection strucutre provices
//supplementary information only.
//Upon successful creation, primary connection structure, enriched with assigned UUID and redundant connection UUID, will be returned
func (c RestClient) CreateL2RedundantConnection(primary L2Connection, secondary L2Connection) (*L2Connection, error) {
	url := fmt.Sprintf("%s/ecx/v3/l2/connections", c.baseURL)
	reqBody := createL2RedundantConnectionRequest(primary, secondary)
	respBody := api.CreateL2ConnectionResponse{}
	req := c.R().SetBody(&reqBody).SetResult(&respBody)
	if err := c.execute(req, resty.MethodPost, url); err != nil {
		return nil, err
	}
	primary.UUID = respBody.PrimaryConnectionID
	primary.RedundantUUID = respBody.SecondaryConnectionID
	return &primary, nil
}

//DeleteL2Connection deletes layer 2 connection with a given UUID
func (c RestClient) DeleteL2Connection(uuid string) error {
	url := fmt.Sprintf("%s/ecx/v3/l2/connections/%s", c.baseURL, url.PathEscape(uuid))
	respBody := api.DeleteL2ConnectionResponse{}
	req := c.R().SetResult(&respBody)
	if err := c.execute(req, resty.MethodDelete, url); err != nil {
		return err
	}
	return nil
}

//NewL2ConnectionUpdateRequest creates new composite update request for a connection with a given UUID
func (c RestClient) NewL2ConnectionUpdateRequest(uuid string) L2ConnectionUpdateRequest {
	return &restL2ConnectionUpdateRequest{
		uuid: uuid,
		c:    c,
	}
}

//WithName sets new connection name in a composite connection update request
func (req *restL2ConnectionUpdateRequest) WithName(name string) L2ConnectionUpdateRequest {
	req.name = name
	return req
}

//WithBandwidth sets new connection bandwidth in a composite connection update request
func (req *restL2ConnectionUpdateRequest) WithBandwidth(speed int, speedUnit string) L2ConnectionUpdateRequest {
	req.speed = speed
	req.speedUnit = speedUnit
	return req
}

//Execute attempts to update connection according new data set in composite update request.
//This is not atomic operation and if any update will fail, other changes won't be reverted.
//UpdateError will be returned if any of requested data failed to update
func (req *restL2ConnectionUpdateRequest) Execute() error {
	url := fmt.Sprintf("%s/ecx/v3/l2/connections/%s", req.c.baseURL, url.PathEscape(req.uuid))
	reqBody := api.L2ConnectionUpdateRequest{
		Name:      req.name,
		Speed:     req.speed,
		SpeedUnit: req.speedUnit,
	}
	if req.name != "" || (req.speed > 0 && req.speedUnit != "") {
		restReq := req.c.R().SetQueryParam("action", "update").SetBody(&reqBody)
		if err := req.c.execute(restReq, resty.MethodPatch, url); err != nil {
			return err
		}
	}
	return nil
}

func mapGETToL2Connection(getResponse api.L2ConnectionResponse) *L2Connection {
	return &L2Connection{
		UUID:                getResponse.UUID,
		Name:                getResponse.Name,
		ProfileUUID:         getResponse.SellerServiceUUID,
		Speed:               getResponse.Speed,
		SpeedUnit:           getResponse.SpeedUnit,
		Status:              getResponse.Status,
		Notifications:       getResponse.Notifications,
		PurchaseOrderNumber: getResponse.PurchaseOrderNumber,
		PortUUID:            getResponse.PortUUID,
		VlanSTag:            getResponse.VlanSTag,
		VlanCTag:            getResponse.VlanCTag,
		NamedTag:            getResponse.NamedTag,
		AdditionalInfo:      mapAdditionalInfoAPIToDomain(getResponse.AdditionalInfo),
		ZSidePortUUID:       getResponse.ZSidePortUUID,
		ZSideVlanSTag:       getResponse.ZSideVlanSTag,
		ZSideVlanCTag:       getResponse.ZSideVlanCTag,
		SellerRegion:        getResponse.SellerRegion,
		SellerMetroCode:     getResponse.SellerMetroCode,
		AuthorizationKey:    getResponse.AuthorizationKey,
		RedundantUUID:       getResponse.RedundantUUID}
}

func createL2ConnectionRequest(l2connection L2Connection) api.L2ConnectionRequest {
	return api.L2ConnectionRequest{
		PrimaryName:          l2connection.Name,
		ProfileUUID:          l2connection.ProfileUUID,
		Speed:                l2connection.Speed,
		SpeedUnit:            l2connection.SpeedUnit,
		Notifications:        l2connection.Notifications,
		PurchaseOrderNumber:  l2connection.PurchaseOrderNumber,
		PrimaryPortUUID:      l2connection.PortUUID,
		PrimaryVlanSTag:      l2connection.VlanSTag,
		PrimaryVlanCTag:      l2connection.VlanCTag,
		NamedTag:             l2connection.NamedTag,
		AdditionalInfo:       mapAdditionalInfoDomainToAPI(l2connection.AdditionalInfo),
		PrimaryZSidePortUUID: l2connection.ZSidePortUUID,
		PrimaryZSideVlanSTag: l2connection.ZSideVlanSTag,
		PrimaryZSideVlanCTag: l2connection.ZSideVlanCTag,
		SellerRegion:         l2connection.SellerRegion,
		SellerMetroCode:      l2connection.SellerMetroCode,
		AuthorizationKey:     l2connection.AuthorizationKey}
}

func createL2RedundantConnectionRequest(primary L2Connection, secondary L2Connection) api.L2ConnectionRequest {
	connReq := createL2ConnectionRequest(primary)
	connReq.SecondaryName = secondary.Name
	connReq.SecondaryPortUUID = secondary.PortUUID
	connReq.SecondaryVlanSTag = secondary.VlanSTag
	connReq.SecondaryVlanCTag = secondary.VlanCTag
	connReq.SecondaryZSidePortUUID = secondary.ZSidePortUUID
	connReq.SecondaryZSideVlanSTag = secondary.ZSideVlanSTag
	connReq.SecondaryZSideVlanCTag = secondary.ZSideVlanCTag
	return connReq
}

func mapAdditionalInfoDomainToAPI(info []L2ConnectionAdditionalInfo) []api.L2ConnectionAdditionalInfo {
	apiInfo := make([]api.L2ConnectionAdditionalInfo, len(info))
	for i, v := range info {
		apiInfo[i] = api.L2ConnectionAdditionalInfo{
			Name:  v.Name,
			Value: v.Value,
		}
	}
	return apiInfo
}

func mapAdditionalInfoAPIToDomain(apiInfo []api.L2ConnectionAdditionalInfo) []L2ConnectionAdditionalInfo {
	info := make([]L2ConnectionAdditionalInfo, len(apiInfo))
	for i, v := range apiInfo {
		info[i] = L2ConnectionAdditionalInfo{
			Name:  v.Name,
			Value: v.Value,
		}
	}
	return info
}
