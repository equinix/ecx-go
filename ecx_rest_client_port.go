package ecx

import (
	"fmt"

	"github.com/equinix/ecx-go/internal/api"
	"github.com/go-resty/resty/v2"
)

//GetUserPorts operation retrieves ECXF user ports
func (c RestClient) GetUserPorts() ([]Port, error) {
	url := fmt.Sprintf("%s/ecx/v3/port/userport", c.baseURL)
	respBody := []api.Port{}
	req := c.R().SetResult(&respBody)
	if err := c.execute(req, resty.MethodGet, url); err != nil {
		return nil, err
	}
	mapped := make([]Port, len(respBody))
	for i := range respBody {
		mapped[i] = mapPortAPIToDomain(respBody[i])
	}
	return mapped, nil
}

func mapPortAPIToDomain(apiPort api.Port) Port {
	return Port{
		UUID:          apiPort.UUID,
		Name:          apiPort.Name,
		Region:        apiPort.Region,
		IBX:           apiPort.IBX,
		MetroCode:     apiPort.MetroCode,
		Priority:      apiPort.DevicePriority,
		Encapsulation: apiPort.Encapsulation,
		Buyout:        apiPort.Buyout,
		Bandwidth:     fmt.Sprintf("%d", apiPort.TotalBandwidth),
		Status:        apiPort.ProvisionStatus,
	}
}
