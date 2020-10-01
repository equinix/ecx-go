package ecx

import (
	"context"
	"net/http"

	"github.com/equinix/rest-go"
)

//RestClient describes ECX Fabric client that uses REST API
type RestClient struct {
	*rest.Client
}

//NewClient creates new ECX REST API client with a given baseURL and http.Client
func NewClient(ctx context.Context, baseURL string, httpClient *http.Client) *RestClient {
	rest := rest.NewClient(ctx, baseURL, httpClient)
	rest.SetHeader("User-agent", "equinix/ecx-go")
	return &RestClient{rest}
}
