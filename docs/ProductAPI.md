# \ProductAPI

All URIs are relative to *https://www.facturapi.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProduct**](ProductAPI.md#CreateProduct) | **Post** /products | Crear producto
[**DeleteProduct**](ProductAPI.md#DeleteProduct) | **Delete** /products/{product_id} | Eliminar producto
[**EditProduct**](ProductAPI.md#EditProduct) | **Put** /products/{product_id} | Editar producto
[**GetProduct**](ProductAPI.md#GetProduct) | **Get** /products/{product_id} | Obtener producto por ID
[**ListProducts**](ProductAPI.md#ListProducts) | **Get** /products | Listar productos



## CreateProduct

> Product CreateProduct(ctx).CreateProductRequest(createProductRequest).Execute()

Crear producto



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
    createProductRequest := *openapiclient.NewCreateProductRequest() // CreateProductRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductAPI.CreateProduct(context.Background()).CreateProductRequest(createProductRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.CreateProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProduct`: Product
    fmt.Fprintf(os.Stdout, "Response from `ProductAPI.CreateProduct`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createProductRequest** | [**CreateProductRequest**](CreateProductRequest.md) |  | 

### Return type

[**Product**](Product.md)

### Authorization

[secretTestKey](../README.md#secretTestKey), [secretLiveKey](../README.md#secretLiveKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProduct

> Product DeleteProduct(ctx, productId).Execute()

Eliminar producto



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
    productId := "productId_example" // string | ID del objeto a eliminar

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductAPI.DeleteProduct(context.Background(), productId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.DeleteProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteProduct`: Product
    fmt.Fprintf(os.Stdout, "Response from `ProductAPI.DeleteProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** | ID del objeto a eliminar | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Product**](Product.md)

### Authorization

[secretTestKey](../README.md#secretTestKey), [secretLiveKey](../README.md#secretLiveKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditProduct

> Product EditProduct(ctx, productId).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

Editar producto



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
    productId := "productId_example" // string | ID del objeto a editar
    uNKNOWNBASETYPE := TODO // UNKNOWN_BASE_TYPE |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductAPI.EditProduct(context.Background(), productId).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.EditProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditProduct`: Product
    fmt.Fprintf(os.Stdout, "Response from `ProductAPI.EditProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** | ID del objeto a editar | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md) |  | 

### Return type

[**Product**](Product.md)

### Authorization

[secretTestKey](../README.md#secretTestKey), [secretLiveKey](../README.md#secretLiveKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProduct

> Product GetProduct(ctx, productId).Execute()

Obtener producto por ID



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
    productId := "productId_example" // string | ID del objeto a obtener

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductAPI.GetProduct(context.Background(), productId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.GetProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProduct`: Product
    fmt.Fprintf(os.Stdout, "Response from `ProductAPI.GetProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** | ID del objeto a obtener | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Product**](Product.md)

### Authorization

[secretTestKey](../README.md#secretTestKey), [secretLiveKey](../README.md#secretLiveKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProducts

> ProductSearchResult ListProducts(ctx).Q(q).Page(page).Limit(limit).Execute()

Listar productos



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
    q := "q_example" // string | Consulta.  Texto a buscar en la descripción del producto o SKU. (optional)
    page := TODO // interface{} | Página de resultados a regresar, empezando desde la página 1. (optional)
    limit := TODO // interface{} | Número del 1 al 100 que representa la cantidad máxima de resultados a regresar con motivos de paginación. (optional) (default to 50)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductAPI.ListProducts(context.Background()).Q(q).Page(page).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ListProducts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListProducts`: ProductSearchResult
    fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ListProducts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Consulta.  Texto a buscar en la descripción del producto o SKU. | 
 **page** | [**interface{}**](interface{}.md) | Página de resultados a regresar, empezando desde la página 1. | 
 **limit** | [**interface{}**](interface{}.md) | Número del 1 al 100 que representa la cantidad máxima de resultados a regresar con motivos de paginación. | [default to 50]

### Return type

[**ProductSearchResult**](ProductSearchResult.md)

### Authorization

[secretTestKey](../README.md#secretTestKey), [secretLiveKey](../README.md#secretLiveKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

