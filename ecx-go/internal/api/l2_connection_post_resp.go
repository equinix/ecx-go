package api

//CreateL2ConnectionResponse post l2 connection response
type CreateL2ConnectionResponse struct {
	Message               string `json:"message,omitempty"`
	PrimaryConnectionID   string `json:"primaryConnectionId,omitempty"`
	SecondaryConnectionID string `json:"secondaryConnectionId,omitempty"`
	Status                string `json:"status,omitempty"`
}
