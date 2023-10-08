# \ReceiptAPI

All URIs are relative to *https://www.facturapi.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelReceipt**](ReceiptAPI.md#CancelReceipt) | **Delete** /receipts/{receipt_id} | Cancelar recibo
[**CreateGlobalInvoice**](ReceiptAPI.md#CreateGlobalInvoice) | **Post** /receipts/global-invoice | Crear factura global
[**CreateReceipt**](ReceiptAPI.md#CreateReceipt) | **Post** /receipts | Crear recibo
[**DownloadReceiptPdf**](ReceiptAPI.md#DownloadReceiptPdf) | **Get** /receipts/{receipt_id}/pdf | Descargar PDF
[**GetReceipt**](ReceiptAPI.md#GetReceipt) | **Get** /receipts/{receipt_id} | Obtener recibo por ID
[**InvoiceReceipt**](ReceiptAPI.md#InvoiceReceipt) | **Post** /receipts/{receipt_id}/invoice | Facturar recibo
[**ListReceipts**](ReceiptAPI.md#ListReceipts) | **Get** /receipts | Listar recibos
[**SendReceiptByEmail**](ReceiptAPI.md#SendReceiptByEmail) | **Post** /receipts/{receipt_id}/email | Enviar recibo por correo electrónico



## CancelReceipt

> Receipt CancelReceipt(ctx, receiptId).Execute()

Cancelar recibo



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
    receiptId := "receiptId_example" // string | ID del recibo a cancelar

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReceiptAPI.CancelReceipt(context.Background(), receiptId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReceiptAPI.CancelReceipt``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelReceipt`: Receipt
    fmt.Fprintf(os.Stdout, "Response from `ReceiptAPI.CancelReceipt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**receiptId** | **string** | ID del recibo a cancelar | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelReceiptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Receipt**](Receipt.md)

### Authorization

[secretTestKey](../README.md#secretTestKey), [secretLiveKey](../README.md#secretLiveKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGlobalInvoice

> Invoice CreateGlobalInvoice(ctx).GlobalInvoiceInput(globalInvoiceInput).Execute()

Crear factura global



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
    globalInvoiceInput := *openapiclient.NewGlobalInvoiceInput("Periodicity_example") // GlobalInvoiceInput |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReceiptAPI.CreateGlobalInvoice(context.Background()).GlobalInvoiceInput(globalInvoiceInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReceiptAPI.CreateGlobalInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGlobalInvoice`: Invoice
    fmt.Fprintf(os.Stdout, "Response from `ReceiptAPI.CreateGlobalInvoice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGlobalInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **globalInvoiceInput** | [**GlobalInvoiceInput**](GlobalInvoiceInput.md) |  | 

### Return type

[**Invoice**](Invoice.md)

### Authorization

[secretTestKey](../README.md#secretTestKey), [secretLiveKey](../README.md#secretLiveKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateReceipt

> Receipt CreateReceipt(ctx).ReceiptInput(receiptInput).Execute()

Crear recibo



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
    receiptInput := *openapiclient.NewReceiptInput([]openapiclient.LineItemInput{*openapiclient.NewLineItemInput(openapiclient.LineItemInput_product{LineItemProductInput: openapiclient.NewLineItemProductInput()})}) // ReceiptInput |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReceiptAPI.CreateReceipt(context.Background()).ReceiptInput(receiptInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReceiptAPI.CreateReceipt``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateReceipt`: Receipt
    fmt.Fprintf(os.Stdout, "Response from `ReceiptAPI.CreateReceipt`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateReceiptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **receiptInput** | [**ReceiptInput**](ReceiptInput.md) |  | 

### Return type

[**Receipt**](Receipt.md)

### Authorization

[secretUserKey](../README.md#secretUserKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadReceiptPdf

> *os.File DownloadReceiptPdf(ctx, receiptId).Execute()

Descargar PDF



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
    receiptId := "receiptId_example" // string | ID del objeto a descargar

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReceiptAPI.DownloadReceiptPdf(context.Background(), receiptId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReceiptAPI.DownloadReceiptPdf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadReceiptPdf`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `ReceiptAPI.DownloadReceiptPdf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**receiptId** | **string** | ID del objeto a descargar | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadReceiptPdfRequest struct via the builder pattern


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


## GetReceipt

> Receipt GetReceipt(ctx, receiptId).Execute()

Obtener recibo por ID



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
    receiptId := "receiptId_example" // string | ID del objeto a obtener

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReceiptAPI.GetReceipt(context.Background(), receiptId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReceiptAPI.GetReceipt``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReceipt`: Receipt
    fmt.Fprintf(os.Stdout, "Response from `ReceiptAPI.GetReceipt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**receiptId** | **string** | ID del objeto a obtener | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReceiptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Receipt**](Receipt.md)

### Authorization

[secretTestKey](../README.md#secretTestKey), [secretLiveKey](../README.md#secretLiveKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoiceReceipt

> Invoice InvoiceReceipt(ctx, receiptId).InvoiceReceiptInput(invoiceReceiptInput).Execute()

Facturar recibo



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
    receiptId := "receiptId_example" // string | ID del objeto a obtener
    invoiceReceiptInput := *openapiclient.NewInvoiceReceiptInput(openapiclient.InvoiceCommonInputProperties_allOf_customer{CustomerCreateInput: openapiclient.NewCustomerCreateInput(*openapiclient.NewCustomerPropertiesAllOfAddress())}) // InvoiceReceiptInput |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReceiptAPI.InvoiceReceipt(context.Background(), receiptId).InvoiceReceiptInput(invoiceReceiptInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReceiptAPI.InvoiceReceipt``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvoiceReceipt`: Invoice
    fmt.Fprintf(os.Stdout, "Response from `ReceiptAPI.InvoiceReceipt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**receiptId** | **string** | ID del objeto a obtener | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceReceiptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **invoiceReceiptInput** | [**InvoiceReceiptInput**](InvoiceReceiptInput.md) |  | 

### Return type

[**Invoice**](Invoice.md)

### Authorization

[secretTestKey](../README.md#secretTestKey), [secretLiveKey](../README.md#secretLiveKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListReceipts

> ReceiptSearchResult ListReceipts(ctx).Q(q).PaymentForm(paymentForm).Date(date).Page(page).Limit(limit).Execute()

Listar recibos



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
    q := "q_example" // string | Consulta. Texto a buscar en la descripción de los conceptos del recibo o el SKU.  (optional)
    paymentForm := "02" // string | Código que representa la forma de pago, de acuerdo al [catálogo del SAT](#forma-de-pago). Si se incluye, los recibos se agruparán y se listarán de acuerdo a la forma de pago.  (optional)
    date := *openapiclient.NewDateRange() // DateRange | Objeto con rango de fechas solicitado. (optional)
    page := TODO // interface{} | Página de resultados a regresar, empezando desde la página 1. (optional)
    limit := TODO // interface{} | Número del 1 al 100 que representa la cantidad máxima de resultados a regresar con motivos de paginación. (optional) (default to 50)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReceiptAPI.ListReceipts(context.Background()).Q(q).PaymentForm(paymentForm).Date(date).Page(page).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReceiptAPI.ListReceipts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListReceipts`: ReceiptSearchResult
    fmt.Fprintf(os.Stdout, "Response from `ReceiptAPI.ListReceipts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListReceiptsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Consulta. Texto a buscar en la descripción de los conceptos del recibo o el SKU.  | 
 **paymentForm** | **string** | Código que representa la forma de pago, de acuerdo al [catálogo del SAT](#forma-de-pago). Si se incluye, los recibos se agruparán y se listarán de acuerdo a la forma de pago.  | 
 **date** | [**DateRange**](DateRange.md) | Objeto con rango de fechas solicitado. | 
 **page** | [**interface{}**](interface{}.md) | Página de resultados a regresar, empezando desde la página 1. | 
 **limit** | [**interface{}**](interface{}.md) | Número del 1 al 100 que representa la cantidad máxima de resultados a regresar con motivos de paginación. | [default to 50]

### Return type

[**ReceiptSearchResult**](ReceiptSearchResult.md)

### Authorization

[secretTestKey](../README.md#secretTestKey), [secretLiveKey](../README.md#secretLiveKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendReceiptByEmail

> *os.File SendReceiptByEmail(ctx, receiptId).SendReceiptByEmailRequest(sendReceiptByEmailRequest).Execute()

Enviar recibo por correo electrónico



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
    receiptId := "receiptId_example" // string | ID del objeto a obtener
    sendReceiptByEmailRequest := *openapiclient.NewSendReceiptByEmailRequest("otro@correo.com") // SendReceiptByEmailRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReceiptAPI.SendReceiptByEmail(context.Background(), receiptId).SendReceiptByEmailRequest(sendReceiptByEmailRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReceiptAPI.SendReceiptByEmail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendReceiptByEmail`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `ReceiptAPI.SendReceiptByEmail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**receiptId** | **string** | ID del objeto a obtener | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendReceiptByEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sendReceiptByEmailRequest** | [**SendReceiptByEmailRequest**](SendReceiptByEmailRequest.md) |  | 

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

