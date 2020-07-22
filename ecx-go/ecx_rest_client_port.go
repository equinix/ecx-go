package ecx

import (
	"ecx-go/v3/internal/api"
	"fmt"

	"github.com/go-resty/resty/v2"
)

//GetUserPort operation retrieves ECXF user port by its name
func (c RestClient) GetUserPort(name string) (*Port, error) {
	url := fmt.Sprintf("%s/ecx/v3/port/userport", c.baseURL)
	respBody := []api.Port{}
	req := c.R().SetResult(&respBody)
	if err := c.execute(req, resty.MethodGet, url); err != nil {
		return nil, err
	}
	for _, v := range respBody {
		if v.Name == name {
			return mapPortAPIToDomain(v), nil
		}
	}
	return nil, fmt.Errorf("port with name '%s' was not found", name)
}

func mapPortAPIToDomain(apiPort api.Port) *Port {
	return &Port{
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
