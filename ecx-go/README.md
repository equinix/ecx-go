ECX Fabric Go client
==================

Equinix Cloud Exchange (ECX) Fabric client library written in Go.

Purpose
------------------
ECXF client library was written in Go for purpose of managing ECX resources from Terraform provider plugin.

Library gives possibility to create L2 connections and service profiles on Equinix platform to any Cloud Service Provider, other Enterprise or between own devices.

Features
------------------
Client library consumes ECX Fabric's REST API version 3 and allows to:
- authenticate and obtain bearer tokens from Equinix's oAuth token endpoint
- manage ECX L2 connections
  - retrieve L2 connection details
  - create non redundant L2 connection
  - create redundant L2 connection
  - delete L2 connection
- manage ECX L2 service profiles

**NOTE**: scope of this library is limited to needs of Terraform provider plugin and it is not providing full capabilities of ECXF API

Usage
------------------
### Project setup
**NOTE**: this project may be moved to Github in future for easier usage. For now, use below instruction for local setup

1. Checkout equinix-terraform-sdk
2. In your project's `go.mod` use `replace` directive and point it out to  `ecx-go` directory from checked out repository
   ```
   require ecx-go/v3 v3.0.0
   replace ecx-go/v3 v3.0.0 => ../ecx-go
   ```

### Code
1. Add ecx-go modules to import statement
   ```
   import (
	  "oauth2-go"
	  "ecx-go/v3"
   )
   ```

2. Define baseURL that will be used in all REST API requests
    ```
    baseURL := "https://sandboxapi.equinix.com"
    ```
3. Create oAuth configuration and oAuth enabled `http.Client`
    ```
    authConfig := oauth2.Config{
      ClientID:     "someClientId",
      ClientSecret: "someSecret",
      BaseURL:      baseURL}
    ctx := context.Background()
    authClient := authConfig.New(ctx)
    ```
4. Create ECX REST client with a given `baseURL` and oauth's `http.Client`
    ```
    var ecxClient ecx.Client = ecx.NewClient(ctx, baseURL, authClient)
    ```
5. Use ECX client to perform some operation `i.e. fetch`
    ```
    l2conn, err := ecxClient.GetL2Connection("myUUID")
    if err != nil {
      log.Printf("Error while fetching connection - %v", err)
    } else {
      log.Printf("Retrieved connection - %+v", l2conn)
    }
    ```
