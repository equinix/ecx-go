package api

//DeleteL2ConnectionResponse l2 connection delete response
type DeleteL2ConnectionResponse struct {
	Message             string `json:"message,omitempty"`
	PrimaryConnectionID string `json:"primaryConnectionId,omitempty"`
}
