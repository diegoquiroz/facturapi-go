# \ToolsAPI

All URIs are relative to *https://www.facturapi.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckGet**](ToolsAPI.md#CheckGet) | **Get** /check | Health check (Pulso)
[**ToolsTaxIdValidationGet**](ToolsAPI.md#ToolsTaxIdValidationGet) | **Get** /tools/tax_id_validation | Validar RFC



## CheckGet

> CheckGet200Response CheckGet(ctx).Execute()

Health check (Pulso)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/diegoquiroz/facturapi-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ToolsAPI.CheckGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.CheckGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckGet`: CheckGet200Response
    fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.CheckGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCheckGetRequest struct via the builder pattern


### Return type

[**CheckGet200Response**](CheckGet200Response.md)

### Authorization

[Secret Test Key](../README.md#Secret Test Key), [Secret Live Key](../README.md#Secret Live Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToolsTaxIdValidationGet

> TaxIdValidationResult ToolsTaxIdValidationGet(ctx).TaxId(taxId).Execute()

Validar RFC



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/diegoquiroz/facturapi-go"
)

func main() {
    taxId := "BBA830831LJ2" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ToolsAPI.ToolsTaxIdValidationGet(context.Background()).TaxId(taxId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.ToolsTaxIdValidationGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ToolsTaxIdValidationGet`: TaxIdValidationResult
    fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.ToolsTaxIdValidationGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiToolsTaxIdValidationGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taxId** | **string** |  | 

### Return type

[**TaxIdValidationResult**](TaxIdValidationResult.md)

### Authorization

[Secret Test Key](../README.md#Secret Test Key), [Secret Live Key](../README.md#Secret Live Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

