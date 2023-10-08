# \SatKeysAPI

All URIs are relative to *https://www.facturapi.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CatalogsProductsGet**](SatKeysAPI.md#CatalogsProductsGet) | **Get** /catalogs/products | Producto/Servicio
[**CatalogsUnitsGet**](SatKeysAPI.md#CatalogsUnitsGet) | **Get** /catalogs/units | Unidades de medida



## CatalogsProductsGet

> ProductCatalogSearchResult CatalogsProductsGet(ctx).Q(q).Page(page).Limit(limit).Execute()

Producto/Servicio



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
    q := "q_example" // string | Consulta. Texto a buscar en la descripción de la categoría. (optional)
    page := TODO // interface{} | Página de resultados a regresar, empezando desde la página 1. (optional)
    limit := TODO // interface{} | Número del 1 al 100 que representa la cantidad máxima de resultados a regresar con motivos de paginación. (optional) (default to 50)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SatKeysAPI.CatalogsProductsGet(context.Background()).Q(q).Page(page).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SatKeysAPI.CatalogsProductsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CatalogsProductsGet`: ProductCatalogSearchResult
    fmt.Fprintf(os.Stdout, "Response from `SatKeysAPI.CatalogsProductsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCatalogsProductsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Consulta. Texto a buscar en la descripción de la categoría. | 
 **page** | [**interface{}**](interface{}.md) | Página de resultados a regresar, empezando desde la página 1. | 
 **limit** | [**interface{}**](interface{}.md) | Número del 1 al 100 que representa la cantidad máxima de resultados a regresar con motivos de paginación. | [default to 50]

### Return type

[**ProductCatalogSearchResult**](ProductCatalogSearchResult.md)

### Authorization

[Secret Test Key](../README.md#Secret Test Key), [Secret Live Key](../README.md#Secret Live Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CatalogsUnitsGet

> UnitCatalogSearchResult CatalogsUnitsGet(ctx).Q(q).Page(page).Limit(limit).Execute()

Unidades de medida



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
    q := "q_example" // string | Consulta. Texto a buscar en la descripción de la unidad de medida. (optional)
    page := TODO // interface{} | Página de resultados a regresar, empezando desde la página 1. (optional)
    limit := TODO // interface{} | Número del 1 al 100 que representa la cantidad máxima de resultados a regresar con motivos de paginación. (optional) (default to 50)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SatKeysAPI.CatalogsUnitsGet(context.Background()).Q(q).Page(page).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SatKeysAPI.CatalogsUnitsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CatalogsUnitsGet`: UnitCatalogSearchResult
    fmt.Fprintf(os.Stdout, "Response from `SatKeysAPI.CatalogsUnitsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCatalogsUnitsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Consulta. Texto a buscar en la descripción de la unidad de medida. | 
 **page** | [**interface{}**](interface{}.md) | Página de resultados a regresar, empezando desde la página 1. | 
 **limit** | [**interface{}**](interface{}.md) | Número del 1 al 100 que representa la cantidad máxima de resultados a regresar con motivos de paginación. | [default to 50]

### Return type

[**UnitCatalogSearchResult**](UnitCatalogSearchResult.md)

### Authorization

[Secret Test Key](../README.md#Secret Test Key), [Secret Live Key](../README.md#Secret Live Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

