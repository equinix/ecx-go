package ecx

import (
	"ecx-go/v3/internal/api"
	"fmt"
	"net/url"
	"github.com/go-resty/resty/v2"
)

//ConfirmL2Connection operation accepts a hosted connection
func (c RestClient) ConfirmL2Connection(uuid string, connToConfirm L2ConnectionToConfirm) (*L2ConnectionConfirmation, error) {
	url := fmt.Sprintf("%s/ecx/v3/l2/connections/%s", c.baseURL, url.PathEscape(uuid))
	reqBody := confirmL2ConnectionRequest(connToConfirm)
	respBody := api.ConfirmL2ConnectionResponse{}
	req := c.R().
		SetQueryParam("action", "Approve").
		SetBody(&reqBody).
		SetResult(&respBody)
	if err := c.execute(req, resty.MethodPatch, url); err != nil {
		return nil, err
	}

	confirmation := L2ConnectionConfirmation{}
	confirmation.PrimaryConnectionID = respBody.PrimaryConnectionID
	confirmation.Message = respBody.Message
	return &confirmation, nil
}

func confirmL2ConnectionRequest(connToConfirm L2ConnectionToConfirm) api.ConfirmL2ConnectionRequest {
	return api.ConfirmL2ConnectionRequest{
		AccessKey: connToConfirm.AccessKey,
		SecretKey: connToConfirm.SecretKey}
}