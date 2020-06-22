package api

//L2ConnectionResponse get connection by uuid response
type L2ConnectionResponse struct {
	UUID                string   `json:"uuid,omitempty"`
	Name                string   `json:"name,omitempty"`
	SellerServiceUUID   string   `json:"sellerServiceUUID,omitempty"`
	Speed               int      `json:"speed,omitempty"`
	SpeedUnit           string   `json:"speedUnit,omitempty"`
	Status              string   `json:"status,omitempty"`
	Notifications       []string `json:"notifications"`
	PurchaseOrderNumber string   `json:"purchaseOrderNumber"`
	PortUUID            string   `json:"portUUID,omitempty"`
	VlanSTag            int      `json:"vlanSTag,omitempty"`
	VlanCTag            int      `json:"vlanCTag,omitempty"`
	ZSidePortUUID       string   `json:"zSidePortUUID,omitempty"`
	ZSideVlanCTag       int      `json:"zSideVlanCTag,omitempty"`
	ZSideVlanSTag       int      `json:"zSideVlanSTag,omitempty"`
	SellerRegion        string   `json:"sellerRegion,omitempty"`
	SellerMetroCode     string   `json:"sellerMetroCode,omitempty"`
	AuthorizationKey    string   `json:"authorizationKey,omitempty"`
	RedundantUUID       string   `json:"redundantUUID,omitempty"`
}
