# \CustomerAPI

All URIs are relative to *https://www.facturapi.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomer**](CustomerAPI.md#CreateCustomer) | **Post** /customers | Crear cliente
[**DeleteCustomer**](CustomerAPI.md#DeleteCustomer) | **Delete** /customers/{customer_id} | Eliminar cliente
[**EditCustomer**](CustomerAPI.md#EditCustomer) | **Put** /customers/{customer_id} | Editar cliente
[**GetCustomer**](CustomerAPI.md#GetCustomer) | **Get** /customers/{customer_id} | Obtener cliente por ID
[**ListCustomers**](CustomerAPI.md#ListCustomers) | **Get** /customers | Listar clientes
[**ValidateCustomerTaxInfo**](CustomerAPI.md#ValidateCustomerTaxInfo) | **Get** /customers/{customer_id}/tax-info-validation | Validar información fiscal



## CreateCustomer

> Customer CreateCustomer(ctx).CustomerCreateInput(customerCreateInput).Execute()

Crear cliente



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
    customerCreateInput := *openapiclient.NewCustomerCreateInput(*openapiclient.NewCustomerPropertiesAllOfAddress()) // CustomerCreateInput |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerAPI.CreateCustomer(context.Background()).CustomerCreateInput(customerCreateInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerAPI.CreateCustomer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCustomer`: Customer
    fmt.Fprintf(os.Stdout, "Response from `CustomerAPI.CreateCustomer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerCreateInput** | [**CustomerCreateInput**](CustomerCreateInput.md) |  | 

### Return type

[**Customer**](Customer.md)

### Authorization

[secretTestKey](../README.md#secretTestKey), [secretLiveKey](../README.md#secretLiveKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomer

> Customer DeleteCustomer(ctx, customerId).Execute()

Eliminar cliente



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
    customerId := "customerId_example" // string | ID del objeto a eliminar

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerAPI.DeleteCustomer(context.Background(), customerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerAPI.DeleteCustomer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCustomer`: Customer
    fmt.Fprintf(os.Stdout, "Response from `CustomerAPI.DeleteCustomer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID del objeto a eliminar | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Customer**](Customer.md)

### Authorization

[secretTestKey](../README.md#secretTestKey), [secretLiveKey](../README.md#secretLiveKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditCustomer

> Customer EditCustomer(ctx, customerId).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

Editar cliente



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
    customerId := "customerId_example" // string | ID del objeto a editar
    uNKNOWNBASETYPE := TODO // UNKNOWN_BASE_TYPE |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerAPI.EditCustomer(context.Background(), customerId).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerAPI.EditCustomer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditCustomer`: Customer
    fmt.Fprintf(os.Stdout, "Response from `CustomerAPI.EditCustomer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID del objeto a editar | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md) |  | 

### Return type

[**Customer**](Customer.md)

### Authorization

[secretTestKey](../README.md#secretTestKey), [secretLiveKey](../README.md#secretLiveKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomer

> Customer GetCustomer(ctx, customerId).Execute()

Obtener cliente por ID



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
    customerId := "customerId_example" // string | ID del objeto a obtener

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerAPI.GetCustomer(context.Background(), customerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerAPI.GetCustomer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomer`: Customer
    fmt.Fprintf(os.Stdout, "Response from `CustomerAPI.GetCustomer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID del objeto a obtener | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Customer**](Customer.md)

### Authorization

[secretTestKey](../README.md#secretTestKey), [secretLiveKey](../README.md#secretLiveKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCustomers

> CustomerSearchResult ListCustomers(ctx).Q(q).Date(date).Page(page).Limit(limit).Execute()

Listar clientes



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
    q := "q_example" // string | Consulta. Texto a buscar en `legal_name` (nombre fiscal) o en `tax_id` (RFC). (optional)
    date := *openapiclient.NewDateRange() // DateRange | Objeto con rango de fechas solicitado. (optional)
    page := TODO // interface{} | Página de resultados a regresar, empezando desde la página 1. (optional)
    limit := TODO // interface{} | Número del 1 al 100 que representa la cantidad máxima de resultados a regresar con motivos de paginación. (optional) (default to 50)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerAPI.ListCustomers(context.Background()).Q(q).Date(date).Page(page).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerAPI.ListCustomers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCustomers`: CustomerSearchResult
    fmt.Fprintf(os.Stdout, "Response from `CustomerAPI.ListCustomers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCustomersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Consulta. Texto a buscar en &#x60;legal_name&#x60; (nombre fiscal) o en &#x60;tax_id&#x60; (RFC). | 
 **date** | [**DateRange**](DateRange.md) | Objeto con rango de fechas solicitado. | 
 **page** | [**interface{}**](interface{}.md) | Página de resultados a regresar, empezando desde la página 1. | 
 **limit** | [**interface{}**](interface{}.md) | Número del 1 al 100 que representa la cantidad máxima de resultados a regresar con motivos de paginación. | [default to 50]

### Return type

[**CustomerSearchResult**](CustomerSearchResult.md)

### Authorization

[secretTestKey](../README.md#secretTestKey), [secretLiveKey](../README.md#secretLiveKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateCustomerTaxInfo

> ValidateCustomerTaxInfo200Response ValidateCustomerTaxInfo(ctx, customerId).Execute()

Validar información fiscal



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
    customerId := "customerId_example" // string | ID del objeto `Customer` a validar

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerAPI.ValidateCustomerTaxInfo(context.Background(), customerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerAPI.ValidateCustomerTaxInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ValidateCustomerTaxInfo`: ValidateCustomerTaxInfo200Response
    fmt.Fprintf(os.Stdout, "Response from `CustomerAPI.ValidateCustomerTaxInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID del objeto &#x60;Customer&#x60; a validar | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateCustomerTaxInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ValidateCustomerTaxInfo200Response**](ValidateCustomerTaxInfo200Response.md)

### Authorization

[secretLiveKey](../README.md#secretLiveKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

