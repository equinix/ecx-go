ECX Fabric Go client
==================

Equinix Cloud Exchange (ECX) Fabric client library written in Go.

[![Build Status](https://travis-ci.com/equinix/ecx-go.svg?branch=master)](https://travis-ci.com/github/equinix/ecx-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/equinix/ecx-go)](https://goreportcard.com/report/github.com/equinix/ecx-go)
[![GoDoc](https://godoc.org/github.com/equinix/ecx-go?status.svg)](https://godoc.org/github.com/equinix/ecx-go)
![GitHub](https://img.shields.io/github/license/equinix/ecx-go)

---

Purpose
------------------
ECXF client library was written in Go for purpose of managing ECX resources from Terraform provider plugin.

Library gives possibility to create L2 connections and service profiles on Equinix platform to any Cloud Service Provider, other Enterprise or between own devices.

Features
------------------
Client library consumes ECX Fabric's REST API and allows to:
- manage ECXF L2 connections
  - retrieve L2 connection details
  - create non redundant L2 connection
  - create redundant L2 connection
  - delete L2 connection
  - update L2 connection (name and speed)
- manage ECXF L2 service profiles
- retrieve list of ECXF user ports
- retrieve list of ECXF L2 seller profiles

**NOTE**: scope of this library is limited to needs of Terraform provider plugin and it is not providing full capabilities of ECXF API

Usage
------------------
### Code
1. Add ecx-go module to import statement.
   In below example, Equinix `oauth2-go` module is imported as well
   ```
   import (
	  "github.com/equinix/oauth2-go"
	  "github.com/equinix/ecx-go"
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
