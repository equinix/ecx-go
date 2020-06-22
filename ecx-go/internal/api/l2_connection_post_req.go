package api

//L2ConnectionRequest post l2 connections request
type L2ConnectionRequest struct {
	PrimaryName            string   `json:"primaryName,omitempty"`
	ProfileUUID            string   `json:"profileUUID,omitempty"`
	Speed                  int      `json:"speed,omitempty"`
	SpeedUnit              string   `json:"speedUnit,omitempty"`
	Notifications          []string `json:"notifications"`
	PurchaseOrderNumber    string   `json:"purchaseOrderNumber"`
	PrimaryPortUUID        string   `json:"primaryPortUUID,omitempty"`
	PrimaryVlanSTag        int      `json:"primaryVlanSTag,omitempty"`
	PrimaryVlanCTag        int      `json:"primaryVlanCTag,omitempty"`
	PrimaryZSidePortUUID   string   `json:"primaryZSidePortUUID,omitempty"`
	PrimaryZSideVlanSTag   int      `json:"primaryZSideVlanSTag,omitempty"`
	PrimaryZSideVlanCTag   int      `json:"primaryZSideVlanCTag,omitempty"`
	SecondaryName          string   `json:"secondaryName,omitempty"`
	SecondaryPortUUID      string   `json:"secondaryPortUUID,omitempty"`
	SecondaryVlanSTag      int      `json:"secondaryVlanSTag,omitempty"`
	SecondaryVlanCTag      int      `json:"secondaryVlanCTag,omitempty"`
	SecondaryZSidePortUUID string   `json:"secondaryZSidePortUUID,omitempty"`
	SecondaryZSideVlanSTag int      `json:"secondaryZSideVlanSTag,omitempty"`
	SecondaryZSideVlanCTag int      `json:"secondaryZSideVlanCTag,omitempty"`
	SellerRegion           string   `json:"sellerRegion,omitempty"`
	SellerMetroCode        string   `json:"sellerMetroCode,omitempty"`
	AuthorizationKey       string   `json:"authorizationKey,omitempty"`
}
