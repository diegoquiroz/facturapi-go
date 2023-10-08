# \InvoiceAPI

All URIs are relative to *https://www.facturapi.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelInvoice**](InvoiceAPI.md#CancelInvoice) | **Delete** /invoices/{invoice_id} | Cancelar factura
[**CreateInvoice**](InvoiceAPI.md#CreateInvoice) | **Post** /invoices | Crear factura (CFDI 4.0)
[**DownloadCancellationReceiptXml**](InvoiceAPI.md#DownloadCancellationReceiptXml) | **Get** /invoices/{invoice_id}/cancellation_receipt/xml | Descargar acuse de cancelación
[**DownloadInvoice**](InvoiceAPI.md#DownloadInvoice) | **Get** /invoices/{invoice_id}/{format} | Descargar factura
[**GetInvoice**](InvoiceAPI.md#GetInvoice) | **Get** /invoices/{invoice_id} | Obtener factura por ID
[**ListInvoices**](InvoiceAPI.md#ListInvoices) | **Get** /invoices | Listar facturas
[**SendInvoiceByEmail**](InvoiceAPI.md#SendInvoiceByEmail) | **Post** /invoices/{invoice_id}/email | Enviar factura por correo electrónico



## CancelInvoice

> Invoice CancelInvoice(ctx, invoiceId).Motive(motive).Substitution(substitution).Execute()

Cancelar factura



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
    invoiceId := "invoiceId_example" // string | ID de la factura a cancelar
    motive := "motive_example" // string | Clave que representa el motivo de la cancelación de la factura.  - `01`: **Comprobante emitido con errores con relación**. Cuando la   factura contiene algún error en las cantidades, claves o cualquier otro dato y ya   se ha emitido el comprobante que la sustituye, el cual deberá indicarse por medio   del atributo `substitution`. - `02`: **Comprobante emitido con errores sin relación**. Cuando la   factura contiene algún error en las cantidades, claves o cualquier otro dato y no   se requiere relacionar con otra factura. - `03`: **No se llevó a cabo la operación**. Cuando la venta o transacción no se concretó. - `04`: **Operación nominativa relacionada en la factura global**. Cuando se requiere cancelar   una factura al público en general porque el cliente solicita su comprobante. 
    substitution := "substitution_example" // string | ID de la factura que sustituye a la factura que se está cancelando.  Puedes usar el ID de Facturapi o el folio fiscal (UUID).  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoiceAPI.CancelInvoice(context.Background(), invoiceId).Motive(motive).Substitution(substitution).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoiceAPI.CancelInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelInvoice`: Invoice
    fmt.Fprintf(os.Stdout, "Response from `InvoiceAPI.CancelInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** | ID de la factura a cancelar | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **motive** | **string** | Clave que representa el motivo de la cancelación de la factura.  - &#x60;01&#x60;: **Comprobante emitido con errores con relación**. Cuando la   factura contiene algún error en las cantidades, claves o cualquier otro dato y ya   se ha emitido el comprobante que la sustituye, el cual deberá indicarse por medio   del atributo &#x60;substitution&#x60;. - &#x60;02&#x60;: **Comprobante emitido con errores sin relación**. Cuando la   factura contiene algún error en las cantidades, claves o cualquier otro dato y no   se requiere relacionar con otra factura. - &#x60;03&#x60;: **No se llevó a cabo la operación**. Cuando la venta o transacción no se concretó. - &#x60;04&#x60;: **Operación nominativa relacionada en la factura global**. Cuando se requiere cancelar   una factura al público en general porque el cliente solicita su comprobante.  | 
 **substitution** | **string** | ID de la factura que sustituye a la factura que se está cancelando.  Puedes usar el ID de Facturapi o el folio fiscal (UUID).  | 

### Return type

[**Invoice**](Invoice.md)

### Authorization

[secretTestKey](../README.md#secretTestKey), [secretLiveKey](../README.md#secretLiveKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInvoice

> Invoice CreateInvoice(ctx).CreateInvoiceRequest(createInvoiceRequest).Execute()

Crear factura (CFDI 4.0)



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
    createInvoiceRequest := openapiclient.createInvoice_request{InvoiceEgresoInput: openapiclient.NewInvoiceEgresoInput("Type_example", "03", []openapiclient.LineItemEgresoInput{*openapiclient.NewLineItemEgresoInput(openapiclient.LineItemEgresoInput_product{LineItemProductEgresoInput: openapiclient.NewLineItemProductEgresoInput()})})} // CreateInvoiceRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoiceAPI.CreateInvoice(context.Background()).CreateInvoiceRequest(createInvoiceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoiceAPI.CreateInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateInvoice`: Invoice
    fmt.Fprintf(os.Stdout, "Response from `InvoiceAPI.CreateInvoice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createInvoiceRequest** | [**CreateInvoiceRequest**](CreateInvoiceRequest.md) |  | 

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


## DownloadCancellationReceiptXml

> *os.File DownloadCancellationReceiptXml(ctx, invoiceId).Execute()

Descargar acuse de cancelación



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
    invoiceId := "invoiceId_example" // string | ID del objeto a obtener

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoiceAPI.DownloadCancellationReceiptXml(context.Background(), invoiceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoiceAPI.DownloadCancellationReceiptXml``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadCancellationReceiptXml`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `InvoiceAPI.DownloadCancellationReceiptXml`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** | ID del objeto a obtener | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadCancellationReceiptXmlRequest struct via the builder pattern


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


## DownloadInvoice

> *os.File DownloadInvoice(ctx, invoiceId, format).Execute()

Descargar factura



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
    invoiceId := "invoiceId_example" // string | ID del objeto a descargar
    format := "format_example" // string | Formato del archivo de descarga

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoiceAPI.DownloadInvoice(context.Background(), invoiceId, format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoiceAPI.DownloadInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadInvoice`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `InvoiceAPI.DownloadInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** | ID del objeto a descargar | 
**format** | **string** | Formato del archivo de descarga | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadInvoiceRequest struct via the builder pattern


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


## GetInvoice

> Invoice GetInvoice(ctx, invoiceId).Execute()

Obtener factura por ID



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
    invoiceId := "invoiceId_example" // string | ID del objeto a obtener

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoiceAPI.GetInvoice(context.Background(), invoiceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoiceAPI.GetInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInvoice`: Invoice
    fmt.Fprintf(os.Stdout, "Response from `InvoiceAPI.GetInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** | ID del objeto a obtener | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Invoice**](Invoice.md)

### Authorization

[secretTestKey](../README.md#secretTestKey), [secretLiveKey](../README.md#secretLiveKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInvoices

> InvoiceSearchResult ListInvoices(ctx).Q(q).Customer(customer).Type_(type_).PaymentMethod(paymentMethod).Date(date).Page(page).Limit(limit).Execute()

Listar facturas



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
    q := "q_example" // string | Consulta. Texto a buscar en la factura.  La búsqueda se realizará por coincidencias **parciales** en los campos:  - `items[].product.description` - `customer.legal_name`  Y por coincidencias **exactas** en los campos:  - `id` - `uuid` - `customer.tax_id` - `folio_number` - `total`  (optional)
    customer := "customer_example" // string | Identificador del cliente. Útil para obtener las facturas emitidas a un sólo cliente. (optional)
    type_ := "type__example" // string | Tipo de factura. Búsqueda por tipo de factura con las claves exactas. (optional)
    paymentMethod := "paymentMethod_example" // string | Método de pago. Búsqueda exacta por método de pago. (optional)
    date := *openapiclient.NewDateRange() // DateRange | Objeto con rango de fechas solicitado. (optional)
    page := TODO // interface{} | Página de resultados a regresar, empezando desde la página 1. (optional)
    limit := TODO // interface{} | Número del 1 al 100 que representa la cantidad máxima de resultados a regresar con motivos de paginación. (optional) (default to 50)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoiceAPI.ListInvoices(context.Background()).Q(q).Customer(customer).Type_(type_).PaymentMethod(paymentMethod).Date(date).Page(page).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoiceAPI.ListInvoices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInvoices`: InvoiceSearchResult
    fmt.Fprintf(os.Stdout, "Response from `InvoiceAPI.ListInvoices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListInvoicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Consulta. Texto a buscar en la factura.  La búsqueda se realizará por coincidencias **parciales** en los campos:  - &#x60;items[].product.description&#x60; - &#x60;customer.legal_name&#x60;  Y por coincidencias **exactas** en los campos:  - &#x60;id&#x60; - &#x60;uuid&#x60; - &#x60;customer.tax_id&#x60; - &#x60;folio_number&#x60; - &#x60;total&#x60;  | 
 **customer** | **string** | Identificador del cliente. Útil para obtener las facturas emitidas a un sólo cliente. | 
 **type_** | **string** | Tipo de factura. Búsqueda por tipo de factura con las claves exactas. | 
 **paymentMethod** | **string** | Método de pago. Búsqueda exacta por método de pago. | 
 **date** | [**DateRange**](DateRange.md) | Objeto con rango de fechas solicitado. | 
 **page** | [**interface{}**](interface{}.md) | Página de resultados a regresar, empezando desde la página 1. | 
 **limit** | [**interface{}**](interface{}.md) | Número del 1 al 100 que representa la cantidad máxima de resultados a regresar con motivos de paginación. | [default to 50]

### Return type

[**InvoiceSearchResult**](InvoiceSearchResult.md)

### Authorization

[secretTestKey](../README.md#secretTestKey), [secretLiveKey](../README.md#secretLiveKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendInvoiceByEmail

> *os.File SendInvoiceByEmail(ctx, invoiceId).SendInvoiceByEmailRequest(sendInvoiceByEmailRequest).Execute()

Enviar factura por correo electrónico



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
    invoiceId := "invoiceId_example" // string | ID del objeto a obtener
    sendInvoiceByEmailRequest := *openapiclient.NewSendInvoiceByEmailRequest() // SendInvoiceByEmailRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoiceAPI.SendInvoiceByEmail(context.Background(), invoiceId).SendInvoiceByEmailRequest(sendInvoiceByEmailRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoiceAPI.SendInvoiceByEmail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendInvoiceByEmail`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `InvoiceAPI.SendInvoiceByEmail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** | ID del objeto a obtener | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendInvoiceByEmailRequest struct via the builder pattern


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

