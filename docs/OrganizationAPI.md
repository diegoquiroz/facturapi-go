# \OrganizationAPI

All URIs are relative to *https://www.facturapi.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckDomainAvailability**](OrganizationAPI.md#CheckDomainAvailability) | **Get** /organizations/domain-check | Revisar dominio disponible
[**CreateOrganization**](OrganizationAPI.md#CreateOrganization) | **Post** /organizations | Crear organización
[**DeleteOrganization**](OrganizationAPI.md#DeleteOrganization) | **Delete** /organizations/{organization_id} | Eliminar organización
[**EditOrganizationCustomization**](OrganizationAPI.md#EditOrganizationCustomization) | **Put** /organizations/{organization_id}/customization | Editar personalización
[**EditOrganizationDomain**](OrganizationAPI.md#EditOrganizationDomain) | **Put** /organizations/{organization_id}/domain | Elegir dominio de autofactura
[**EditOrganizationLegal**](OrganizationAPI.md#EditOrganizationLegal) | **Put** /organizations/{organization_id}/legal | Editar datos fiscales
[**EditOrganizationReceiptsSettings**](OrganizationAPI.md#EditOrganizationReceiptsSettings) | **Put** /organizations/{organization_id}/receipts | Editar config. recibos
[**GetOrganization**](OrganizationAPI.md#GetOrganization) | **Get** /organizations/{organization_id} | Obtener organización por ID
[**GetTestApiKey**](OrganizationAPI.md#GetTestApiKey) | **Get** /organizations/{organization_id}/apikeys/test | Obtener Test Api Key
[**ListOrganizations**](OrganizationAPI.md#ListOrganizations) | **Get** /organizations | Listar organizaciones
[**RenewLiveApiKey**](OrganizationAPI.md#RenewLiveApiKey) | **Put** /organizations/{organization_id}/apikeys/live | Renovar Live Api Key
[**RenewTestApiKey**](OrganizationAPI.md#RenewTestApiKey) | **Put** /organizations/{organization_id}/apikeys/test | Renovar Test Api Key
[**UploadOrganizationCertificate**](OrganizationAPI.md#UploadOrganizationCertificate) | **Put** /organizations/{organization_id}/certificate | Subir certificados (CSD)
[**UploadOrganizationLogo**](OrganizationAPI.md#UploadOrganizationLogo) | **Put** /organizations/{organization_id}/logo | Subir logotipo



## CheckDomainAvailability

> CheckDomainAvailability200Response CheckDomainAvailability(ctx).Domain(domain).Execute()

Revisar dominio disponible



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
    domain := "domain_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationAPI.CheckDomainAvailability(context.Background()).Domain(domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.CheckDomainAvailability``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckDomainAvailability`: CheckDomainAvailability200Response
    fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.CheckDomainAvailability`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckDomainAvailabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domain** | **string** |  | 

### Return type

[**CheckDomainAvailability200Response**](CheckDomainAvailability200Response.md)

### Authorization

[secretUserKey](../README.md#secretUserKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganization

> Organization CreateOrganization(ctx).OrganizationCreateInput(organizationCreateInput).Execute()

Crear organización



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
    organizationCreateInput := *openapiclient.NewOrganizationCreateInput("Name_example") // OrganizationCreateInput |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationAPI.CreateOrganization(context.Background()).OrganizationCreateInput(organizationCreateInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.CreateOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganization`: Organization
    fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.CreateOrganization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationCreateInput** | [**OrganizationCreateInput**](OrganizationCreateInput.md) |  | 

### Return type

[**Organization**](Organization.md)

### Authorization

[secretUserKey](../README.md#secretUserKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganization

> Organization DeleteOrganization(ctx, organizationId).Execute()

Eliminar organización



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
    organizationId := "organizationId_example" // string | ID del objeto a eliminar

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationAPI.DeleteOrganization(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.DeleteOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteOrganization`: Organization
    fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.DeleteOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | ID del objeto a eliminar | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Organization**](Organization.md)

### Authorization

[secretUserKey](../README.md#secretUserKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditOrganizationCustomization

> Organization EditOrganizationCustomization(ctx, organizationId).OrganizationCustomizationInput(organizationCustomizationInput).Execute()

Editar personalización



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
    organizationId := "organizationId_example" // string | ID de la organización
    organizationCustomizationInput := TODO // OrganizationCustomizationInput |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationAPI.EditOrganizationCustomization(context.Background(), organizationId).OrganizationCustomizationInput(organizationCustomizationInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.EditOrganizationCustomization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditOrganizationCustomization`: Organization
    fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.EditOrganizationCustomization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | ID de la organización | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditOrganizationCustomizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organizationCustomizationInput** | [**OrganizationCustomizationInput**](OrganizationCustomizationInput.md) |  | 

### Return type

[**Organization**](Organization.md)

### Authorization

[secretUserKey](../README.md#secretUserKey)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditOrganizationDomain

> Organization EditOrganizationDomain(ctx, organizationId).OrganizationDomainInput(organizationDomainInput).Execute()

Elegir dominio de autofactura



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
    organizationId := "organizationId_example" // string | ID de la organización
    organizationDomainInput := TODO // OrganizationDomainInput |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationAPI.EditOrganizationDomain(context.Background(), organizationId).OrganizationDomainInput(organizationDomainInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.EditOrganizationDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditOrganizationDomain`: Organization
    fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.EditOrganizationDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | ID de la organización | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditOrganizationDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organizationDomainInput** | [**OrganizationDomainInput**](OrganizationDomainInput.md) |  | 

### Return type

[**Organization**](Organization.md)

### Authorization

[secretUserKey](../README.md#secretUserKey)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditOrganizationLegal

> Organization EditOrganizationLegal(ctx, organizationId).OrganizationLegalInput(organizationLegalInput).Execute()

Editar datos fiscales



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
    organizationId := "organizationId_example" // string | ID de la organización
    organizationLegalInput := *openapiclient.NewOrganizationLegalInput("Name_example", "LegalName_example", "601", *openapiclient.NewOrganizationLegalInputAddress("Blvd. Atardecer", "142", "86500")) // OrganizationLegalInput |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationAPI.EditOrganizationLegal(context.Background(), organizationId).OrganizationLegalInput(organizationLegalInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.EditOrganizationLegal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditOrganizationLegal`: Organization
    fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.EditOrganizationLegal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | ID de la organización | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditOrganizationLegalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organizationLegalInput** | [**OrganizationLegalInput**](OrganizationLegalInput.md) |  | 

### Return type

[**Organization**](Organization.md)

### Authorization

[secretUserKey](../README.md#secretUserKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditOrganizationReceiptsSettings

> Organization EditOrganizationReceiptsSettings(ctx, organizationId).OrganizationReceiptsInput(organizationReceiptsInput).Execute()

Editar config. recibos



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
    organizationId := "organizationId_example" // string | ID de la organización
    organizationReceiptsInput := TODO // OrganizationReceiptsInput |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationAPI.EditOrganizationReceiptsSettings(context.Background(), organizationId).OrganizationReceiptsInput(organizationReceiptsInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.EditOrganizationReceiptsSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditOrganizationReceiptsSettings`: Organization
    fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.EditOrganizationReceiptsSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | ID de la organización | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditOrganizationReceiptsSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organizationReceiptsInput** | [**OrganizationReceiptsInput**](OrganizationReceiptsInput.md) |  | 

### Return type

[**Organization**](Organization.md)

### Authorization

[secretUserKey](../README.md#secretUserKey)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganization

> Organization GetOrganization(ctx, organizationId).Execute()

Obtener organización por ID



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
    organizationId := "organizationId_example" // string | ID de la organización

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationAPI.GetOrganization(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.GetOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganization`: Organization
    fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.GetOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | ID de la organización | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Organization**](Organization.md)

### Authorization

[secretUserKey](../README.md#secretUserKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTestApiKey

> string GetTestApiKey(ctx, organizationId).Execute()

Obtener Test Api Key



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
    organizationId := "organizationId_example" // string | ID de la organización

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationAPI.GetTestApiKey(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.GetTestApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTestApiKey`: string
    fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.GetTestApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | ID de la organización | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTestApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[secretUserKey](../README.md#secretUserKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrganizations

> OrganizationSearchResult ListOrganizations(ctx).Q(q).Date(date).Page(page).Limit(limit).Execute()

Listar organizaciones



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
    q := "q_example" // string | Consulta. Texto a buscar en `name` (nombre comercial), `legal_name` (nombre fiscal) o en `tax_id` (RFC). (optional)
    date := *openapiclient.NewDateRange() // DateRange | Objeto con rango de fechas solicitado. (optional)
    page := TODO // interface{} | Página de resultados a regresar, empezando desde la página 1. (optional)
    limit := TODO // interface{} | Número del 1 al 100 que representa la cantidad máxima de resultados a regresar con motivos de paginación. (optional) (default to 50)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationAPI.ListOrganizations(context.Background()).Q(q).Date(date).Page(page).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.ListOrganizations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrganizations`: OrganizationSearchResult
    fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.ListOrganizations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Consulta. Texto a buscar en &#x60;name&#x60; (nombre comercial), &#x60;legal_name&#x60; (nombre fiscal) o en &#x60;tax_id&#x60; (RFC). | 
 **date** | [**DateRange**](DateRange.md) | Objeto con rango de fechas solicitado. | 
 **page** | [**interface{}**](interface{}.md) | Página de resultados a regresar, empezando desde la página 1. | 
 **limit** | [**interface{}**](interface{}.md) | Número del 1 al 100 que representa la cantidad máxima de resultados a regresar con motivos de paginación. | [default to 50]

### Return type

[**OrganizationSearchResult**](OrganizationSearchResult.md)

### Authorization

[secretUserKey](../README.md#secretUserKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenewLiveApiKey

> string RenewLiveApiKey(ctx, organizationId).Execute()

Renovar Live Api Key



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
    organizationId := "organizationId_example" // string | ID de la organización

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationAPI.RenewLiveApiKey(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.RenewLiveApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RenewLiveApiKey`: string
    fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.RenewLiveApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | ID de la organización | 

### Other Parameters

Other parameters are passed through a pointer to a apiRenewLiveApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[secretUserKey](../README.md#secretUserKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenewTestApiKey

> string RenewTestApiKey(ctx, organizationId).Execute()

Renovar Test Api Key



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
    organizationId := "organizationId_example" // string | ID de la organización

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationAPI.RenewTestApiKey(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.RenewTestApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RenewTestApiKey`: string
    fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.RenewTestApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | ID de la organización | 

### Other Parameters

Other parameters are passed through a pointer to a apiRenewTestApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[secretUserKey](../README.md#secretUserKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadOrganizationCertificate

> Organization UploadOrganizationCertificate(ctx, organizationId).OrganizationCertsInput(organizationCertsInput).Execute()

Subir certificados (CSD)



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
    organizationId := "organizationId_example" // string | ID de la organización
    organizationCertsInput := TODO // OrganizationCertsInput |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationAPI.UploadOrganizationCertificate(context.Background(), organizationId).OrganizationCertsInput(organizationCertsInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.UploadOrganizationCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadOrganizationCertificate`: Organization
    fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.UploadOrganizationCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | ID de la organización | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadOrganizationCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organizationCertsInput** | [**OrganizationCertsInput**](OrganizationCertsInput.md) |  | 

### Return type

[**Organization**](Organization.md)

### Authorization

[secretUserKey](../README.md#secretUserKey)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadOrganizationLogo

> Organization UploadOrganizationLogo(ctx, organizationId).OrganizationLogoInput(organizationLogoInput).Execute()

Subir logotipo



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
    organizationId := "organizationId_example" // string | ID de la organización
    organizationLogoInput := TODO // OrganizationLogoInput |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationAPI.UploadOrganizationLogo(context.Background(), organizationId).OrganizationLogoInput(organizationLogoInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.UploadOrganizationLogo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadOrganizationLogo`: Organization
    fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.UploadOrganizationLogo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | ID de la organización | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadOrganizationLogoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organizationLogoInput** | [**OrganizationLogoInput**](OrganizationLogoInput.md) |  | 

### Return type

[**Organization**](Organization.md)

### Authorization

[secretUserKey](../README.md#secretUserKey)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

