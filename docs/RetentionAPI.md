# \RetentionAPI

All URIs are relative to *https://www.facturapi.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelRetention**](RetentionAPI.md#CancelRetention) | **Delete** /retentions/{retention_id} | Cancelar retención
[**CreateRetention**](RetentionAPI.md#CreateRetention) | **Post** /retentions | Crear retención
[**DownloadRetention**](RetentionAPI.md#DownloadRetention) | **Get** /retentions/{retention_id}/{format} | Descargar retención
[**GetRetention**](RetentionAPI.md#GetRetention) | **Get** /retentions/{retention_id} | Obtener retención por ID
[**ListRetentions**](RetentionAPI.md#ListRetentions) | **Get** /retentions | Listar retenciones
[**SendRetentionByEmail**](RetentionAPI.md#SendRetentionByEmail) | **Post** /retentions/{retention_id}/email | Enviar retención por correo electrónico



## CancelRetention

> Retention CancelRetention(ctx, retentionId).Execute()

Cancelar retención



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
    retentionId := "retentionId_example" // string | ID de la retención a cancelar

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RetentionAPI.CancelRetention(context.Background(), retentionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RetentionAPI.CancelRetention``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelRetention`: Retention
    fmt.Fprintf(os.Stdout, "Response from `RetentionAPI.CancelRetention`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**retentionId** | **string** | ID de la retención a cancelar | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelRetentionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Retention**](Retention.md)

### Authorization

[secretTestKey](../README.md#secretTestKey), [secretLiveKey](../README.md#secretLiveKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRetention

> Retention CreateRetention(ctx).RetentionInput(retentionInput).Execute()

Crear retención



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
    retentionInput := *openapiclient.NewRetentionInput(openapiclient.InvoiceCommonInputProperties_allOf_customer{CustomerCreateInput: openapiclient.NewCustomerCreateInput(*openapiclient.NewCustomerPropertiesAllOfAddress())}, "26", *openapiclient.NewRetentionInputPeriodo(int32(9), int32(9), int32(2021)), *openapiclient.NewRetentionInputTotales(float32(123), float32(123), []openapiclient.RetentionInputTotalesImpRetenidosInner{*openapiclient.NewRetentionInputTotalesImpRetenidosInner(float32(123), "TipoPagoRet_example")})) // RetentionInput |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RetentionAPI.CreateRetention(context.Background()).RetentionInput(retentionInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RetentionAPI.CreateRetention``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRetention`: Retention
    fmt.Fprintf(os.Stdout, "Response from `RetentionAPI.CreateRetention`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRetentionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **retentionInput** | [**RetentionInput**](RetentionInput.md) |  | 

### Return type

[**Retention**](Retention.md)

### Authorization

[secretTestKey](../README.md#secretTestKey), [secretLiveKey](../README.md#secretLiveKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadRetention

> *os.File DownloadRetention(ctx, retentionId, format).Execute()

Descargar retención



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
    retentionId := "retentionId_example" // string | ID del objeto a descargar
    format := "format_example" // string | Formato del archivo de descarga

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RetentionAPI.DownloadRetention(context.Background(), retentionId, format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RetentionAPI.DownloadRetention``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadRetention`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `RetentionAPI.DownloadRetention`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**retentionId** | **string** | ID del objeto a descargar | 
**format** | **string** | Formato del archivo de descarga | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadRetentionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

[secretTestKey](../README.md#secretTestKey), [secretLiveKey](../README.md#secretLiveKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRetention

> Retention GetRetention(ctx, retentionId).Execute()

Obtener retención por ID



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
    retentionId := "retentionId_example" // string | ID del objeto a obtener

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RetentionAPI.GetRetention(context.Background(), retentionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RetentionAPI.GetRetention``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRetention`: Retention
    fmt.Fprintf(os.Stdout, "Response from `RetentionAPI.GetRetention`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**retentionId** | **string** | ID del objeto a obtener | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRetentionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Retention**](Retention.md)

### Authorization

[secretTestKey](../README.md#secretTestKey), [secretLiveKey](../README.md#secretLiveKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRetentions

> RetentionSearchResult ListRetentions(ctx).Q(q).Customer(customer).Date(date).Page(page).Limit(limit).Execute()

Listar retenciones



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
    q := "q_example" // string | Consulta. Texto a buscar en el nombre fiscal del cliente o su RFC. (optional)
    customer := "customer_example" // string | Identificador del cliente. Útil para obtener las retenciones emitidas a un sólo cliente. (optional)
    date := *openapiclient.NewDateRange() // DateRange | Objeto con rango de fechas solicitado. (optional)
    page := TODO // interface{} | Página de resultados a regresar, empezando desde la página 1. (optional)
    limit := TODO // interface{} | Número del 1 al 100 que representa la cantidad máxima de resultados a regresar con motivos de paginación. (optional) (default to 50)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RetentionAPI.ListRetentions(context.Background()).Q(q).Customer(customer).Date(date).Page(page).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RetentionAPI.ListRetentions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRetentions`: RetentionSearchResult
    fmt.Fprintf(os.Stdout, "Response from `RetentionAPI.ListRetentions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRetentionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Consulta. Texto a buscar en el nombre fiscal del cliente o su RFC. | 
 **customer** | **string** | Identificador del cliente. Útil para obtener las retenciones emitidas a un sólo cliente. | 
 **date** | [**DateRange**](DateRange.md) | Objeto con rango de fechas solicitado. | 
 **page** | [**interface{}**](interface{}.md) | Página de resultados a regresar, empezando desde la página 1. | 
 **limit** | [**interface{}**](interface{}.md) | Número del 1 al 100 que representa la cantidad máxima de resultados a regresar con motivos de paginación. | [default to 50]

### Return type

[**RetentionSearchResult**](RetentionSearchResult.md)

### Authorization

[secretTestKey](../README.md#secretTestKey), [secretLiveKey](../README.md#secretLiveKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendRetentionByEmail

> *os.File SendRetentionByEmail(ctx, retentionId).SendInvoiceByEmailRequest(sendInvoiceByEmailRequest).Execute()

Enviar retención por correo electrónico



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
    retentionId := "retentionId_example" // string | ID del objeto a obtener
    sendInvoiceByEmailRequest := *openapiclient.NewSendInvoiceByEmailRequest() // SendInvoiceByEmailRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RetentionAPI.SendRetentionByEmail(context.Background(), retentionId).SendInvoiceByEmailRequest(sendInvoiceByEmailRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RetentionAPI.SendRetentionByEmail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendRetentionByEmail`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `RetentionAPI.SendRetentionByEmail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**retentionId** | **string** | ID del objeto a obtener | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendRetentionByEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sendInvoiceByEmailRequest** | [**SendInvoiceByEmailRequest**](SendInvoiceByEmailRequest.md) |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[secretTestKey](../README.md#secretTestKey), [secretLiveKey](../README.md#secretLiveKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

